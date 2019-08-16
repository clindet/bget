package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

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
var version = "v0.1.0-2"

type downloadCliT struct {
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
	"",
	"",
	false,
	1,
	1,
	false,
	true,
}

var rootCmd = &cobra.Command{
	Use:   "bget [url1 url2... | -k key1 key2... | --doi doi1 doi2...]",
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

func init() {
	osType = runtime.GOOS
	wd, _ := os.Getwd()
	rootCmd.Flags().BoolVar(&(downloadClis.installSpack), "spack", false, "Logical indicating that whether to install spack in tools directory.")
	rootCmd.Flags().StringVarP(&(downloadClis.installConda), "miniconda", "", "", "Install miniconda2 or miniconda3 in tools directory. Optional (2 or 3).")
	rootCmd.Flags().StringVarP(&(downloadClis.installReffa), "reffa", "", "", "Download reference in download directory. Format is genomeVersion %site #releaseVersion.\nOptional (GRCh38 %genecode #31, GRCh37 %genecode #31, hg38 %ucsc, hg19 %ucsc, GRCh38 %ensemble #97, GRCh38 %defuse #97).\nMultiple use comma to seperate (e.g. GRCh38 %genecode #31,GRCh37 %genecode #31, %fusioncatcher #95).")
	rootCmd.Flags().StringVarP(&(downloadClis.doi), "doi", "", "", "Doi to be download.")
	rootCmd.Flags().StringVarP(&(downloadClis.urls), "urls", "u", "", "URLs to be download.")
	rootCmd.Flags().StringVarP(&(downloadClis.urlsFile), "urls-list", "l", "", "A file contains URLs for download.")
	rootCmd.Flags().StringVarP(&(downloadClis.keys), "keys", "k", "", "String key to be download. item@version %site #releaseVersion, e.g. bwa, GRCh38 %defuse #97")
	rootCmd.Flags().BoolVarP(&(downloadClis.keysAll), "keys-all", "a", false, "Show all available string key can be download.")
	rootCmd.Flags().StringVarP(&(downloadClis.engine), "engine", "g", "go-http", "Point the download engine: go-http, wget, curl, axel, git, and rsync.")
	rootCmd.Flags().IntVarP(&(downloadClis.concurrency), "thread", "t", 1, "Concurrency download thread.")
	rootCmd.Flags().IntVarP(&(downloadClis.axelThread), "thread-axel", "", 5, "Set the thread of axel.")
	rootCmd.Flags().StringVarP(&(downloadClis.downloadDir), "outdir", "o", filepath.Join(wd, "_download"), "Set the download dir for get-urls.")
	rootCmd.Flags().StringVarP(&(downloadClis.separator), "separator", "s", ",", "Separator for --reffa,-k, and -u flag.")
	rootCmd.Flags().StringVarP(&(downloadClis.mirror), "mirror", "m", "", "Set the mirror of resources.")
	rootCmd.Flags().BoolVar(&(downloadClis.ignore), "ignore", false, "Contine to download and skip the check of existed files.")
	rootCmd.Flags().BoolVar(&(downloadClis.autoPath), "autopath", true, "Logical indicating that whether to create subdir in download dir (for --reffa): e.g. reffa/{{site}}/{{version}}")
	rootCmd.Flags().StringVarP(&cmdExtraFromFlag, "extra-cmd", "e", "", "Extra flags and values pass to internal CMDs")
	rootCmd.Flags().BoolVarP(&overwrite, "overwrite", "f", false, "Logical indicating that whether to overwrite existing files.")
	rootCmd.Flags().StringVarP(&taskID, "task-id", "", utils.GetRandomString(15), "Task ID (random).")
	rootCmd.Flags().StringVarP(&logDir, "log-dir", "", path.Join(wd, "_log"), "Log dir.")
	rootCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "No output.")
	rootCmd.Flags().BoolVarP(&saveLog, "save-log", "", true, "Save download log to local file].")
	rootCmd.Example = `  urls="https://dldir1.qq.com/weixin/Windows/WeChatSetup.exe,http://download.oray.com/pgy/windows/PgyVPN_4.1.0.21693.exe,https://dldir1.qq.com/qqfile/qq/PCQQ9.1.6/25786/QQ9.1.6.25786.exe" && echo $urls | tr "," "\n"> /tmp/urls.list

  bget ${urls}
  bget https://dldir1.qq.com/weixin/Windows/WeChatSetup.exe https://dldir1.qq.com/qqfile/qq/PCQQ9.1.6/25786/QQ9.1.6.25786.exe
  bget -u ${urls} -t 2 -o /tmp/download
  bget -u ${urls} -t 3 -o /tmp/download -f -g wget
  bget -u ${urls} -t 3 -o /tmp/download -g wget --ignore
  bget -l /tmp/urls.list -o /tmp/download -f -t 3
  bget -k bwa
  bget --doi 10.5281/zenodo.3363060 10.5281/zenodo.3357455 10.5281/zenodo.3351812 -t 3
  bget --spack
  bget --miniconda 3 -o /tmp/testenv
  bget --miniconda 3 --engine wget
  bget --miniconda 3 --engine axel
  bget -k "reffa@GRCh38 %defuse #97" -t 10 -f
  bget --reffa "GRCh38 %defuse #97" -t 10
  bget --reffa "hg38 %ucsc, GRCh37 %genecode #31"`

	rootCmd.Version = version
}

func checkArgs(cmd *cobra.Command) {
	items := []string{}
	if len(cmd.Flags().Args()) >= 1 && downloadClis.keys == "" && downloadClis.doi == "" {
		if downloadClis.urls != "" {
			items = []string{downloadClis.urls}
		}
		items = append(items, cmd.Flags().Args()...)
		downloadClis.urls = strings.Join(items, downloadClis.separator)
	} else if len(cmd.Flags().Args()) >= 1 && downloadClis.urls == "" && downloadClis.doi == "" {
		if downloadClis.keys != "" {
			items = []string{downloadClis.keys}
		}
		items = []string{downloadClis.keys}
		items = append(items, cmd.Flags().Args()...)
		downloadClis.keys = strings.Join(items, downloadClis.separator)
	} else if len(cmd.Flags().Args()) >= 1 && downloadClis.urls == "" && downloadClis.keys == "" {
		if downloadClis.doi != "" {
			items = []string{downloadClis.doi}
		}
		items = []string{downloadClis.doi}
		items = append(items, cmd.Flags().Args()...)
		downloadClis.doi = strings.Join(items, downloadClis.separator)
	}
}

func rootCmdRunOptions(cmd *cobra.Command) {
	if quiet {
		log.SetOutput(ioutil.Discard)
	} else {
		log.SetOutput(os.Stderr)
	}
	checkArgs(cmd)
	if hasDir, _ := utils.PathExists(downloadClis.downloadDir); !hasDir {
		if downloadClis.installSpack || downloadClis.installConda != "" ||
			downloadClis.installReffa != "" || downloadClis.urls != "" || downloadClis.urlsFile != "" || downloadClis.keys != "" || downloadClis.doi != "" {
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
	if downloadClis.keysAll {
		getAllKeys()
		downloadClis.helpFlags = false
	}
	if downloadClis.keys != "" {
		downloadKey()
		downloadClis.helpFlags = false
	}
	if downloadClis.doi != "" {
		downloadDoi()
		downloadClis.helpFlags = false
	}
	if downloadClis.helpFlags {
		cmd.Help()
	}
}
