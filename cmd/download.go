package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/JhuangLab/bioget/log"
	"github.com/JhuangLab/bioget/utils"
	"github.com/spf13/cobra"
)

var taskID string
var overwrite bool
var osType string
var cmdExtraFromFlag string

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
	Use:   "bioget",
	Short: "Lightweight downloader for bioinformatics data, databases and files.",
	Long:  "Lightweight downloader for bioinformatics data, databases and files.",
	Run: func(cmd *cobra.Command, args []string) {
		rootCmdRunOptions(cmd, &downloadClis)
	},
}

// Execute main interface of bioget
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
	rootCmd.Flags().BoolVar(&(downloadClis.installSpack), "get-spack", false, "Logical indicating that whether to install spack in tools directory [false].")
	rootCmd.Flags().StringVarP(&(downloadClis.installConda), "get-miniconda", "", "", "Install miniconda2 or miniconda3 in tools directory. Optional (2 or 3).")
	rootCmd.Flags().StringVarP(&(downloadClis.installReffa), "get-reffa", "", "", "Download reference in download directory. Format is genomeVersion_site_releaseVersion.\nOptional (GRCh38_genecode_31, GRCh37_genecode_31, hg38_ucsc, hg19_ucsc, GRCh38_ensemble_97, GRCh38_defuse_97).\nMultiple use comma to seperate (e.g. hg38_ucsc,hg19_ucsc).")
	rootCmd.Flags().StringVarP(&(downloadClis.engine), "engine", "g", "go-http", "Point the download engine: go-http, wget, curl, axel, and rsync.")
	rootCmd.Flags().IntVarP(&(downloadClis.axelThread), "thread-axel", "", 5, "Set the thread of axel.")
	rootCmd.Flags().StringVarP(&(downloadClis.downloadDir), "outdir", "o", filepath.Join(wd, "_download"), "Set the download dir for get-urls.")
	rootCmd.Flags().StringVarP(&(downloadClis.urls), "urls", "u", "", "URLs for download.")
	rootCmd.Flags().StringVarP(&(downloadClis.urlsFile), "urls-list", "l", "", "A file contains URLs for download.")
	rootCmd.Flags().StringVarP(&(downloadClis.separator), "separator", "s", ",", "Separator for --get-reffa and --urls.")
	rootCmd.Flags().IntVarP(&(downloadClis.concurrency), "thread", "n", 1, "Concurrency download thread.")
	rootCmd.Flags().StringVarP(&(downloadClis.mirror), "mirror", "m", "", "Set the mirror of resources.")
	rootCmd.Flags().BoolVar(&(downloadClis.ignore), "ignore", false, "Contine to download and skip the check existed files. [false]")
	rootCmd.Flags().BoolVar(&(downloadClis.autoPath), "autopath", false, "Logical indicating that whether to create subdir in download dir: e.g. reffa/{{site}}/{{version}} [false]")
	rootCmd.Flags().StringVarP(&cmdExtraFromFlag, "extra-cmd", "e", "", "Extra flags and values pass to internal CMDs")
	rootCmd.Flags().BoolVarP(&overwrite, "overwrite", "f", false, "Logical indicating that whether to overwrite existing files [false].")
	rootCmd.Flags().StringVarP(&taskID, "task-id", "k", utils.GetRandomString(15), "Task ID.")

	rootCmd.Example = `  urls="https://dldir1.qq.com/weixin/Windows/WeChatSetup.exe,http://download.oray.com/pgy/windows/PgyVPN_4.1.0.21693.exe,https://dldir1.qq.com/qqfile/qq/PCQQ9.1.6/25786/QQ9.1.6.25786.exe" && echo $urls | sed 's/,/\n/g'> /tmp/urls.list

  bioget --urls ${urls} -n 2 -o /tmp/download
  bioget --urls ${urls} -n 3 -o /tmp/download -f -g wget
  bioget --urls ${urls} -n 3 -o /tmp/download -g wget --ignore
  bioget -l /tmp/urls.list -o /tmp/download -f -n 3
  bioget --get-spack
  bioget --get-miniconda 3 -o /tmp/testenv
  bioget --get-miniconda 3 --engine wget
  bioget --get-miniconda 3 --engine axel
  bioget --get-reffa hg38_ucsc
  bioget --get-reffa hg38_ucsc,GRCh37_genecode_31`
}

func installSpack() {
	log.INFO("")
	if overwrite {
		os.RemoveAll(path.Join(downloadClis.downloadDir, "spack"))
	}
	utils.AsyncURL(getEnvToolsURL("spack", "latest", ""), path.Join(downloadClis.downloadDir, "spack"),
		"git", taskID, downloadClis.mirror, 1, true)
}

func installMiniconda() {
	log.INFO("")
	url := getEnvToolsURL("miniconda", "latest", downloadClis.installConda)
	destFn := path.Join(os.TempDir(), path.Base(url))
	destDir := path.Join(downloadClis.downloadDir, "miniconda"+downloadClis.installConda)

	if overwrite {
		os.RemoveAll(destDir)
		os.RemoveAll(destFn)
	}
	wd, _ := os.Getwd()
	logPath := path.Join(wd, "_log", fmt.Sprintf("%s_installMiniconda.log", taskID))
	utils.CreateFileParDir(logPath)
	if status, _ := utils.PathExists(destFn); !status {
		log.Infof("Trying %s => %s", url, destFn)
		utils.AsyncURL3(url, destFn, downloadClis.engine, taskID, downloadClis.mirror, downloadClis.axelThread, true)
	} else {
		log.Infof("%s existed.", destFn)
	}

	if osType != "windows" {
		args := []string{destFn, "-b", "-p", destDir, cmdExtraFromFlag}
		cmd := exec.Command("sh", args...)
		utils.RunExecCmd(logPath, cmd)
	}
}

func installReffa() {
	log.INFO("")
	reffaArray := []string{}
	if strings.Contains(downloadClis.installReffa, downloadClis.separator) {
		reffaArray = strings.Split(downloadClis.installReffa, downloadClis.separator)
	} else {
		reffaArray = []string{downloadClis.installReffa}
	}
	urlsPool := []string{}
	destDirPool := []string{}
	for reffaI := range reffaArray {
		if !strings.Contains(reffaArray[reffaI], "_") {
			log.Fatal(fmt.Sprintf("Invaild parameters (--get-reffa): %s. Optional (GRCh38_genecode_31, GRCh37_genecode_31, hg38_ucsc, and hg19_ucsc). Multiple usage: seperate items with comma (e.g. hg38_ucsc,hg19_ucsc)", reffaArray[reffaI]))
		}
		reffa := strings.Split(reffaArray[reffaI], "_")
		version := reffa[0]
		site := reffa[1]
		release := ""
		if len(reffa) == 3 {
			release = reffa[2]
		}
		destDir := ""
		if downloadClis.autoPath {
			destDir = path.Join(downloadClis.downloadDir, "reffa", site, version)
		} else {
			destDir = downloadClis.downloadDir
		}
		urls := getEnvFilesURL("reffa", site, version, release)
		urlsPool = append(urlsPool, urls...)
		for range urls {
			destDirPool = append(destDirPool, destDir)
		}
	}

	utils.HTTPGetURLs(urlsPool, destDirPool, downloadClis.engine, taskID, downloadClis.mirror,
		downloadClis.concurrency, downloadClis.axelThread,
		overwrite, downloadClis.ignore)
}

func downloadUrls() {
	urls := []string{}
	if downloadClis.urls != "" && strings.Contains(downloadClis.urls, downloadClis.separator) {
		urls = strings.Split(downloadClis.urls, downloadClis.separator)
	} else if downloadClis.urls != "" {
		urls = []string{downloadClis.urls}
	} else if downloadClis.urlsFile != "" {
		urls = utils.ReadLines(downloadClis.urlsFile)
	}
	var destDirArray []string
	for i := range urls {
		urls[i] = strings.TrimSpace(urls[i])
		destDirArray = append(destDirArray, downloadClis.downloadDir)
	}
	utils.HTTPGetURLs(urls, destDirArray, downloadClis.engine, taskID, downloadClis.mirror,
		downloadClis.concurrency, downloadClis.axelThread, overwrite, downloadClis.ignore)
}

func rootCmdRunOptions(cmd *cobra.Command, bamClis *downloadCliT) {
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
