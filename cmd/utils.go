package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	cvrt "github.com/openbiox/ligo/convert"
	cio "github.com/openbiox/ligo/io"
	clog "github.com/openbiox/ligo/log"
	cnet "github.com/openbiox/ligo/net"
	"github.com/openbiox/ligo/par"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var log = clog.Logger
var logBash = clog.LoggerBash
var logEnv = log.WithFields(logrus.Fields{
	"prefix": "Env"})
var logPrefix string
var wd string
var logCon io.Writer

func initCmd(cmd *cobra.Command, args []string) {
	setLog()
	if bgetClis.Clean {
		clearLogDownload()
	}
	if bgetClis.Verbose == 2 {
		logEnv.Infof("Prog: %s", cmd.CommandPath())
		logEnv.Infof("TaskID: %s", bgetClis.TaskID)
		if bgetClis.SaveLog && logPrefix != "" {
			logEnv.Infof("Log: %s.log", logPrefix)
		}
		if len(args) > 0 {
			logEnv.Infof("Args: %s", strings.Join(args, " "))
		}
		logEnv.Infof("Global: %v", cvrt.Struct2Map(bgetClis))
	}
}

func setLog() {
	if bgetClis.SaveLog {
		if bgetClis.LogDir == "" {
			bgetClis.LogDir = filepath.Join(os.TempDir(), "_log")
		}
		logPrefix = fmt.Sprintf("%s/%s", bgetClis.LogDir, bgetClis.TaskID)
		cio.CreateDir(bgetClis.LogDir)
		logCon, _ = cio.Open(logPrefix + ".log")
	}
	clog.SetLogStream(log, bgetClis.Verbose == 0, bgetClis.SaveLog, &logCon)
}

func checkDownloadDir(condtions bool) {
	if hasDir, _ := cio.PathExists(bgetClis.DownloadDir); !hasDir {
		if condtions {
			if err := cio.CreateDir(bgetClis.DownloadDir); err != nil {
				log.Fatalf("Could not to create %s", bgetClis.DownloadDir)
			}
		}
	}
}

func clearLogDownload() {
	if err := os.RemoveAll("_download"); err != nil {
		log.Warn(err)
	}
	if err := os.RemoveAll("_log"); err != nil {
		log.Warn(err)
	}
	bgetClis.HelpFlags = false
}

func setUncompressFlag(cmd *cobra.Command, bgetClis *bgetCliT) {
	cmd.PersistentFlags().BoolVarP(&(bgetClis.Uncompress), "uncompress", "u", false, "Uncompress download files for .zip, .tar.gz, and .gz suffix files.")
}
func setKeyListFlag(cmd *cobra.Command, bgetClis *bgetCliT, keyName string) {
	cmd.PersistentFlags().StringVarP(&(bgetClis.ListFile), "list-file", "l", "", fmt.Sprintf("A file contains %s for download.", keyName))
}
func setGlobalFlag(cmd *cobra.Command, bgetClis *bgetCliT) {
	cmd.PersistentFlags().StringVarP(&(bgetClis.Proxy), "proxy", "", "", "HTTP proxy to download.")
	cmd.PersistentFlags().IntVarP(&(bgetClis.Thread), "thread", "t", 1, "Concurrency download thread.")
	cmd.PersistentFlags().StringVarP(&(bgetClis.DownloadDir), "outdir", "o", filepath.Join(wd, "_download"), "Set the download dir.")
	cmd.PersistentFlags().StringVarP(&(bgetClis.Seperator), "seperator", "s", ",", "Optional 'url1{seperator}url2' for multiple keys, urls, or seqs.")
	cmd.PersistentFlags().BoolVar(&(bgetClis.Ignore), "ignore", false, "Contine to download and skip the check of existed files.")
	cmd.PersistentFlags().StringVarP(&(bgetClis.CmdExtraFromFlag), "extra-cmd", "e", "", "Extra flags and values pass to internal CMDs")
	cmd.PersistentFlags().BoolVarP(&(bgetClis.Overwrite), "overwrite", "f", false, "Logical indicating that whether to overwrite existing files.")
	cmd.PersistentFlags().IntVarP(&bgetClis.Retries, "retries", "r", 5, "Retry specifies the number of attempts to retrieve the data.")
	cmd.PersistentFlags().IntVarP(&bgetClis.Timeout, "timeout", "", 35, "Set the timeout of per request.")
	cmd.PersistentFlags().IntVarP(&bgetClis.RetSleepTime, "retries-sleep-time", "", 5, "Sleep time after one retry.")
	cmd.PersistentFlags().BoolVarP(&bgetClis.RemoteName, "remote-name", "n", false, "Use remote defined filename.")
	cmd.PersistentFlags().StringVarP(&(bgetClis.Mirror), "mirror", "m", "", "Set the mirror of resources.")
	cmd.PersistentFlags().StringVarP(&(bgetClis.Engine), "engine", "g", "go-http", "Point the download engine: go-http, wget, curl, axel, git, and rsync.")
	cmd.PersistentFlags().IntVarP(&(bgetClis.AxelThread), "thread-axel", "", 5, "Set the thread of axel.")
}

func checkArgs(cmd *cobra.Command, subcmd string) {
	items := []string{}
	for _, v := range cmd.Flags().Args() {
		if strings.Contains(v, "=") && subcmd == "key" {
			kvs := strings.Split(v, "=")
			bgetClis.Env[kvs[0]] = strings.TrimSpace(kvs[1])
		} else {
			items = append(items, v)
		}
	}
	if bgetClis.WithAssets {
		bgetClis.Env["withAssets"] = "yes"
	}
	if len(items) == 0 {
		return
	}
	if subcmd == "url" {
		bgetClis.URLs = strings.Join(items, bgetClis.Seperator)
	} else if subcmd == "key" {
		bgetClis.Keys = strings.Join(items, bgetClis.Seperator)
	} else if subcmd == "doi" {
		bgetClis.Doi = strings.Join(items, bgetClis.Seperator)
	} else if subcmd == "seq" {
		bgetClis.Seqs = strings.Join(items, bgetClis.Seperator)
	}
}

func setNetParams(bgetClis *bgetCliT) (netOpt *cnet.Params) {

	p := par.NewMpb(bgetClis.Verbose == 0, bgetClis.SaveLog, &logCon)

	return &cnet.Params{
		Proxy:          bgetClis.Proxy,
		Engine:         bgetClis.Engine,
		ExtraArgs:      bgetClis.CmdExtraFromFlag,
		TaskID:         bgetClis.TaskID,
		Mirror:         bgetClis.Mirror,
		Thread:         bgetClis.Thread,
		AxelThread:     bgetClis.AxelThread,
		Overwrite:      bgetClis.Overwrite,
		Ignore:         bgetClis.Ignore,
		Quiet:          bgetClis.Verbose == 0,
		Retries:        bgetClis.Retries,
		SaveLog:        bgetClis.SaveLog,
		Timeout:        bgetClis.Timeout,
		RetSleepTime:   bgetClis.RetSleepTime,
		RemoteName:     bgetClis.RemoteName,
		LogDir:         bgetClis.LogDir,
		EgaCredentials: bgetClis.EgaCredFile,
		Pbar:           p,
	}
}
