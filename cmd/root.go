package cmd

import (
	"fmt"
	"os"
	"path"
	"runtime"

	api "github.com/clindet/bget/api/cmd"
	stringo "github.com/openbiox/ligo/stringo"
	"github.com/spf13/cobra"
)

type bgetCliT struct {
	DownloadDir        string
	Mirror             string
	AutoPath           bool
	Engine             string
	Doi                string
	URLs               string
	ListFile           string
	Seperator          string
	Keys               string
	Seqs               string
	GdcToken           string
	Uncompress         bool
	KeysAll            bool
	Clean              bool
	PrintFormat        string
	ThreadQuery        int
	Thread             int
	Timeout            int
	RetSleepTime       int
	Retries            int
	Proxy              string
	CmdExtraFromFlag   string
	RemoteName         bool
	Ignore             bool
	TaskID             string
	Overwrite          bool
	SaveLog            bool
	LogDir             string
	Verbose            int
	Env                map[string]string
	ShowVersions       bool
	EgaCredFile        string
	GeoGPL             bool
	GitHub             string
	GitHubMode         bool
	OnlyAssets         bool
	WithAssets         bool
	WithAssetsVersions string
	HelpFlags          bool
}

var bgetClis bgetCliT

var rootCmd = &cobra.Command{
	Use:   "bget",
	Short: "Lightweight downloader for bioinformatics data, databases and files.",
	Long:  `Lightweight downloader for bioinformatics data, databases and files (under development). It will provides a simple and parallelized method to access various bioinformatics resoures. More see here https://github.com/clindet/bget.`,
	Run: func(cmd *cobra.Command, args []string) {
		rootCmdRunOptions(cmd, args)
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

func rootCmdRunOptions(cmd *cobra.Command, args []string) {
	initCmd(cmd, args)
	if bgetClis.HelpFlags {
		cmd.Help()
	}
}

func init() {
	wd, _ = os.Getwd()
	bgetClis.HelpFlags = true

	rootCmd.AddCommand(URLCmd)
	rootCmd.AddCommand(DoiCmd)
	rootCmd.AddCommand(KeyCmd)
	rootCmd.AddCommand(SeqCmd)
	rootCmd.Flags().BoolVarP(&(bgetClis.Clean), "clean", "", false, "remove _download and _log in current dir.")
	rootCmd.PersistentFlags().StringVarP(&(bgetClis.TaskID), "task-id", "k", stringo.RandString(15), "task ID (default is random).")
	rootCmd.PersistentFlags().StringVarP(&(bgetClis.LogDir), "log-dir", "", path.Join(wd, "_log"), "log dir.")
	rootCmd.PersistentFlags().IntVarP(&(bgetClis.Verbose), "verbose", "", 1, "verbose level (0:no output, 1: basic level, 2: with env info)")
	rootCmd.PersistentFlags().BoolVarP(&(bgetClis.SaveLog), "save-log", "", false, "Save log to file.")

	rootCmd.AddCommand(api.BapiCmd)
	bgetClis.Env = make(map[string]string)
	bgetClis.Env["osType"] = runtime.GOOS
	bgetClis.Env["wd"] = wd
	rootCmd.Version = version
}
