package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/JhuangLab/bget/log"
	"github.com/JhuangLab/bget/utils"
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
var version = "v0.1.0"

type downloadCliT struct {
	downloadDir  string
	mirror       string
	installReffa string
	autoPath     bool
	installSpack bool
	installConda string
	engine       string
	urls         string
	urlsFile     string
	separator    string
	axelThread   int
	concurrency  int
	ignore       bool
	helpFlags    bool
}

var downloadClis = downloadCliT{
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
		rootCmdRunOptions(cmd, &downloadClis)
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

func init() {
	osType = runtime.GOOS
	wd, _ := os.Getwd()
	rootCmd.Flags().BoolVar(&(downloadClis.installSpack), "get-spack", false, "Logical indicating that whether to install spack in tools directory.")
	rootCmd.Flags().StringVarP(&(downloadClis.installConda), "get-miniconda", "", "", "Install miniconda2 or miniconda3 in tools directory. Optional (2 or 3).")
	rootCmd.Flags().StringVarP(&(downloadClis.installReffa), "get-reffa", "", "", "Download reference in download directory. Format is genomeVersion_site_releaseVersion.\nOptional (GRCh38_genecode_31, GRCh37_genecode_31, hg38_ucsc, hg19_ucsc, GRCh38_ensemble_97, GRCh38_defuse_97).\nMultiple use comma to seperate (e.g. hg38_ucsc,hg19_ucsc).")
	rootCmd.Flags().StringVarP(&(downloadClis.engine), "engine", "g", "go-http", "Point the download engine: go-http, wget, curl, axel, git, and rsync.")
	rootCmd.Flags().IntVarP(&(downloadClis.axelThread), "thread-axel", "", 5, "Set the thread of axel.")
	rootCmd.Flags().StringVarP(&(downloadClis.downloadDir), "outdir", "o", filepath.Join(wd, "_download"), "Set the download dir for get-urls.")
	rootCmd.Flags().StringVarP(&(downloadClis.urls), "get-urls", "u", "", "URLs to be download.")
	rootCmd.Flags().StringVarP(&(downloadClis.urlsFile), "get-urls-list", "l", "", "A file contains URLs for download.")
	rootCmd.Flags().StringVarP(&(downloadClis.separator), "separator", "s", ",", "Separator for --get-reffa and -u.")
	rootCmd.Flags().IntVarP(&(downloadClis.concurrency), "thread", "n", 1, "Concurrency download thread.")
	rootCmd.Flags().StringVarP(&(downloadClis.mirror), "mirror", "m", "", "Set the mirror of resources.")
	rootCmd.Flags().BoolVar(&(downloadClis.ignore), "ignore", false, "Contine to download and skip the check of existed files.")
	rootCmd.Flags().BoolVar(&(downloadClis.autoPath), "autopath", true, "Logical indicating that whether to create subdir in download dir: e.g. reffa/{{site}}/{{version}}")
	rootCmd.Flags().StringVarP(&cmdExtraFromFlag, "extra-cmd", "e", "", "Extra flags and values pass to internal CMDs")
	rootCmd.Flags().BoolVarP(&overwrite, "overwrite", "f", false, "Logical indicating that whether to overwrite existing files.")
	rootCmd.Flags().StringVarP(&taskID, "task-id", "k", utils.GetRandomString(15), "Task ID.")
	rootCmd.Flags().StringVarP(&logDir, "log-dir", "", path.Join(wd, "_log"), "Log dir.")
	rootCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "No output.")
	rootCmd.Flags().BoolVarP(&saveLog, "save-log", "", true, "Save download log to local file].")
	rootCmd.Example = `  urls="https://dldir1.qq.com/weixin/Windows/WeChatSetup.exe,http://download.oray.com/pgy/windows/PgyVPN_4.1.0.21693.exe,https://dldir1.qq.com/qqfile/qq/PCQQ9.1.6/25786/QQ9.1.6.25786.exe" && echo $urls | sed 's/,/\n/g'> /tmp/urls.list

  bget -u ${urls} -n 2 -o /tmp/download
  bget -u ${urls} -n 3 -o /tmp/download -f -g wget
  bget -u ${urls} -n 3 -o /tmp/download -g wget --ignore
  bget -l /tmp/urls.list -o /tmp/download -f -n 3
  bget --get-spack
  bget --get-miniconda 3 -o /tmp/testenv
  bget --get-miniconda 3 --engine wget
  bget --get-miniconda 3 --engine axel
  bget --get-reffa GRCh38_defuse_97 -n 10
  bget --get-reffa hg38_ucsc,GRCh37_genecode_31`

	rootCmd.Version = version
}

func rootCmdRunOptions(cmd *cobra.Command, bamClis *downloadCliT) {
	if quiet {
		log.SetOutput(ioutil.Discard)
	} else {
		log.SetOutput(os.Stderr)
	}

	if hasDir, _ := utils.PathExists(downloadClis.downloadDir); !hasDir {
		if downloadClis.installSpack || downloadClis.installConda != "" ||
			downloadClis.installReffa != "" || downloadClis.urls != "" || downloadClis.urlsFile != "" {
			if err := utils.CreateDir(downloadClis.downloadDir); err != nil {
				log.FATAL(fmt.Sprintf("Could not to create %s", downloadClis.downloadDir))
			}
		}
	}
	if downloadClis.installSpack {
		installSpack()
		downloadClis.helpFlags = false
	}
	if downloadClis.installConda != "" {
		installMiniconda()
		downloadClis.helpFlags = false
	}
	if downloadClis.installReffa != "" {
		installReffa()
		downloadClis.helpFlags = false
	}
	if downloadClis.urls != "" || downloadClis.urlsFile != "" {
		downloadUrls()
		downloadClis.helpFlags = false
	}
	if downloadClis.helpFlags {
		cmd.Help()
	}
}
