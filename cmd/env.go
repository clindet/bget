package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	butils "github.com/JhuangLab/butils"
	"github.com/JhuangLab/butils/log"
	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Can be used to install package manager, dependences, and preset analysis environment.",
	Long:  `Can be used to install package manager, dependences, and preset analysis environment. More see here https://github.com/JhuangLab/bget.`,
	Run: func(cmd *cobra.Command, args []string) {
		envCmdRunOptions(cmd)
	},
}

func installSpack() {
	if overwrite {
		os.RemoveAll(path.Join(bgetClis.downloadDir, "spack"))
	}
	AsyncURL(getEnvToolsURL("spack", "latest", ""), path.Join(bgetClis.downloadDir, "spack"),
		"git", taskID, bgetClis.mirror, 1, quiet, saveLog)
}

func installMiniconda() {
	url := getEnvToolsURL("miniconda"+bgetClis.installConda, "latest", bgetClis.installConda)
	destFn := path.Join(os.TempDir(), path.Base(url))
	destDir := ""
	destDir = path.Join(bgetClis.downloadDir, "miniconda"+bgetClis.installConda)

	if overwrite {
		os.RemoveAll(destDir)
		os.RemoveAll(destFn)
	}
	logPath := path.Join(logDir, fmt.Sprintf("%s_installMiniconda.log", taskID))
	if saveLog {
		butils.CreateFileParDir(logPath)
	}
	if status, _ := butils.PathExists(destFn); !status {
		log.Infof("Trying %s => %s", url, destFn)
		AsyncURL3(url, destFn, bgetClis.engine, taskID, bgetClis.mirror, bgetClis.axelThread, quiet, saveLog)
	} else {
		log.Infof("%s existed.", destFn)
	}

	if osType != "windows" {
		args := []string{destFn, "-b", "-p", destDir, cmdExtraFromFlag}
		cmd := exec.Command("sh", args...)
		butils.RunExecCmdConsole(logPath, cmd, quiet, saveLog)
	}
}

func installReffa() {
	reffaArray := []string{}
	if strings.Contains(bgetClis.installReffa, bgetClis.separator) {
		reffaArray = strings.Split(bgetClis.installReffa, bgetClis.separator)
	} else {
		reffaArray = []string{bgetClis.installReffa}
	}
	urlsPool := []string{}
	destDirPool := []string{}
	defuse := ""
	for reffaI := range reffaArray {
		key, version, site, release := parseMeta("reffa@" + reffaArray[reffaI])
		destDir := ""
		if bgetClis.autoPath {
			destDir = path.Join(bgetClis.downloadDir, "reffa", site, version, release)
		} else {
			destDir = bgetClis.downloadDir
		}
		if site == "defuse" {
			defuse = destDir
		}
		urls := getEnvFilesURL(key, site, version, release)
		urlsPool = append(urlsPool, urls...)
		for range urls {
			destDirPool = append(destDirPool, destDir)
		}
	}

	HTTPGetURLs(urlsPool, destDirPool, bgetClis.engine, taskID, bgetClis.mirror,
		bgetClis.concurrency, bgetClis.axelThread,
		overwrite, bgetClis.ignore, quiet, saveLog)
	if defuse != "" {
		butils.GunzipDefuseReffa(defuse)
	}
}

func envCmdRunOptions(cmd *cobra.Command) {
	checkQuiet()
	checkArgs(cmd)
	checkDownloadDir(bgetClis.installSpack || bgetClis.installConda != "" || bgetClis.installReffa != "")

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
	if bgetClis.helpFlags {
		cmd.Help()
	}
}

func init() {
	envCmd.Flags().BoolVar(&(bgetClis.installSpack), "spack", false, "Logical indicating that whether to install spack in tools directory.")
	envCmd.Flags().StringVarP(&(bgetClis.installConda), "miniconda", "", "", "Install miniconda2 or miniconda3 in tools directory. Optional (2 or 3).")
	envCmd.Flags().StringVarP(&(bgetClis.installReffa), "reffa", "", "", "Download reference in download directory. Format is genomeVersion %site #releaseVersion.\nOptional (GRCh38 %genecode #31, GRCh37 %genecode #31, hg38 %ucsc, hg19 %ucsc, GRCh38 %ensemble #97, GRCh38 %defuse #97).\nMultiple use comma to seperate (e.g. GRCh38 %genecode #31,GRCh37 %genecode #31, %fusioncatcher #95).")

	envCmd.Example = `
  bget env --spack
  bget env --miniconda 3 -o /tmp/testenv
  bget env --miniconda 3 --engine wget
  bget env --miniconda 3 --engine axel
  bget env --reffa "GRCh38 %defuse #97" -t 10
  bget env --reffa "hg38 %ucsc, GRCh37 %genecode #31"`
}
