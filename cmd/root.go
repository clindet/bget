package cmd

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	butils "github.com/JhuangLab/butils"
	"github.com/spf13/cobra"
)

var taskID string
var overwrite bool
var osType string
var cmdExtraFromFlag string
var quiet bool
var saveLog bool
var wd, _ = os.Getwd()
var logDir string
var version = "v0.1.0-2"

type bgetCliT struct {
	downloadDir  string
	mirror       string
	installReffa string
	autoPath     bool
	installSpack bool
	installConda string
	engine       string
	doi          string
	urls         string
	urlsFile     string
	separator    string
	keys         string
	keysAll      bool
	axelThread   int
	concurrency  int
	ignore       bool
	helpFlags    bool
}

var bgetClis = bgetCliT{
	"",
	"",
	"",
	false,
	false,
	"",
	"go-http",
	"",
	"",
	"",
	"",
	"",
	false,
	1,
	1,
	false,
	true,
}

var rootCmd = &cobra.Command{
	Use:   "bget",
	Short: "Lightweight downloader for bioinformatics data, databases and files.",
	Long:  `Lightweight downloader for bioinformatics data, databases and files (under development). It will provides a simple and parallelized method to access various bioinformatics resoures. More see here https://github.com/JhuangLab/bget.`,
	Run: func(cmd *cobra.Command, args []string) {
		rootCmdRunOptions(cmd)
	},
}

// Execute main interface of bget
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		if !rootCmd.HasFlags() && !rootCmd.HasSubCommands() {
			rootCmd.Help()
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func checkArgs(cmd *cobra.Command) {
	items := []string{}
	if len(cmd.Flags().Args()) >= 1 && bgetClis.urls == "" && bgetClis.doi == "" {
		if bgetClis.keys != "" {
			items = []string{bgetClis.keys}
		}
		items = []string{bgetClis.keys}
		items = append(items, cmd.Flags().Args()...)
		bgetClis.keys = strings.Join(items, bgetClis.separator)
	}
}

func rootCmdRunOptions(cmd *cobra.Command) {
	checkQuiet()
	checkArgs(cmd)
	checkDownloadDir(bgetClis.installSpack || bgetClis.installConda != "" ||
		bgetClis.installReffa != "" || bgetClis.keys != "")

	if bgetClis.installSpack {
		installSpack()
		bgetClis.helpFlags = false
	}
	if bgetClis.installConda != "" {
		installMiniconda()
		bgetClis.helpFlags = false
	}
	if bgetClis.installReffa != "" {
		installReffa()
		bgetClis.helpFlags = false
	}
	if bgetClis.keysAll {
		getAllKeys()
		bgetClis.helpFlags = false
	}
	if bgetClis.keys != "" {
		downloadKey()
		bgetClis.helpFlags = false
	}
	if bgetClis.helpFlags {
		cmd.Help()
	}
}

func init() {
	osType = runtime.GOOS
	wd, _ := os.Getwd()
	rootCmd.AddCommand(urlCmd)
	rootCmd.AddCommand(doiCmd)
	rootCmd.AddCommand(keyCmd)
	rootCmd.AddCommand(envCmd)

	rootCmd.PersistentFlags().StringVarP(&(bgetClis.engine), "engine", "g", "go-http", "Point the download engine: go-http, wget, curl, axel, git, and rsync.")
	rootCmd.PersistentFlags().IntVarP(&(bgetClis.concurrency), "thread", "t", 1, "Concurrency download thread.")
	rootCmd.PersistentFlags().IntVarP(&(bgetClis.axelThread), "thread-axel", "", 5, "Set the thread of axel.")
	rootCmd.PersistentFlags().StringVarP(&(bgetClis.downloadDir), "outdir", "o", filepath.Join(wd, "_download"), "Set the download dir for get-urls.")
	rootCmd.PersistentFlags().StringVarP(&(bgetClis.separator), "separator", "s", ",", "Separator for --reffa,-k, and -u flag.")
	rootCmd.PersistentFlags().StringVarP(&(bgetClis.mirror), "mirror", "m", "", "Set the mirror of resources.")
	rootCmd.PersistentFlags().BoolVar(&(bgetClis.ignore), "ignore", false, "Contine to download and skip the check of existed files.")
	rootCmd.PersistentFlags().BoolVar(&(bgetClis.autoPath), "autopath", true, "Logical indicating that whether to create subdir in download dir (for --reffa): e.g. reffa/{{site}}/{{version}}")
	rootCmd.PersistentFlags().StringVarP(&cmdExtraFromFlag, "extra-cmd", "e", "", "Extra flags and values pass to internal CMDs")
	rootCmd.PersistentFlags().BoolVarP(&overwrite, "overwrite", "f", false, "Logical indicating that whether to overwrite existing files.")
	rootCmd.PersistentFlags().StringVarP(&taskID, "task-id", "", butils.GetRandomString(15), "Task ID (random).")
	rootCmd.PersistentFlags().StringVarP(&logDir, "log-dir", "", path.Join(wd, "_log"), "Log dir.")
	rootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "No output.")
	rootCmd.PersistentFlags().BoolVarP(&saveLog, "save-log", "", true, "Save download log to local file].")

	rootCmd.Version = version
}
