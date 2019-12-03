package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	cio "github.com/openbiox/butils/io"
	"github.com/openbiox/butils/log"
	stringo "github.com/openbiox/butils/stringo"
	"github.com/spf13/cobra"
)

var wd string

func checkQuiet() {
	if bgetClis.quiet {
		log.SetOutput(ioutil.Discard)
	} else {
		log.SetOutput(os.Stderr)
	}
}

func checkDownloadDir(condtions bool) {
	if hasDir, _ := cio.PathExists(bgetClis.downloadDir); !hasDir {
		if condtions {
			if err := cio.CreateDir(bgetClis.downloadDir); err != nil {
				log.FATAL(fmt.Sprintf("Could not to create %s", bgetClis.downloadDir))
			}
		}
	}
}

func setUncompressFlag(cmd *cobra.Command, bgetClis *bgetCliT) {
	cmd.PersistentFlags().BoolVarP(&(bgetClis.uncompress), "uncompress", "u", false, "Uncompress download files for .zip, .tar.gz, and .gz suffix files.")
}
func setKeyListFlag(cmd *cobra.Command, bgetClis *bgetCliT, keyName string) {
	cmd.PersistentFlags().StringVarP(&(bgetClis.listFile), "list-file", "l", "", fmt.Sprintf("A file contains %s for download.", keyName))
}
func setGlobalFlag(cmd *cobra.Command, bgetClis *bgetCliT) {
	cmd.PersistentFlags().StringVarP(&(bgetClis.proxy), "proxy", "", "", "HTTP proxy to download.")
	cmd.PersistentFlags().IntVarP(&(bgetClis.thread), "thread", "t", 1, "Concurrency download thread.")
	cmd.PersistentFlags().StringVarP(&(bgetClis.downloadDir), "outdir", "o", wd, "Set the download dir.")
	cmd.PersistentFlags().StringVarP(&(bgetClis.seperator), "seperator", "s", ",", "Optional 'url1{seperator}url2' for multiple keys, urls, or seqs.")
	cmd.PersistentFlags().BoolVar(&(bgetClis.ignore), "ignore", false, "Contine to download and skip the check of existed files.")
	cmd.PersistentFlags().StringVarP(&(bgetClis.cmdExtraFromFlag), "extra-cmd", "e", "", "Extra flags and values pass to internal CMDs")
	cmd.PersistentFlags().BoolVarP(&(bgetClis.overwrite), "overwrite", "f", false, "Logical indicating that whether to overwrite existing files.")
	cmd.PersistentFlags().StringVarP(&(bgetClis.taskID), "task-id", "", stringo.GetRandomString(15), "Task ID (random).")
	cmd.PersistentFlags().StringVarP(&(bgetClis.logDir), "log-dir", "", path.Join(wd, "_log"), "Log dir.")
	cmd.PersistentFlags().BoolVarP(&(bgetClis.quiet), "quiet", "q", false, "No output.")
	cmd.PersistentFlags().BoolVarP(&(bgetClis.saveLog), "save-log", "", true, "Save download log to local file].")
	cmd.PersistentFlags().IntVarP(&bgetClis.retries, "retries", "r", 5, "Retry specifies the number of attempts to retrieve the data.")
	cmd.PersistentFlags().IntVarP(&bgetClis.timeout, "timeout", "", 35, "Set the timeout of per request.")
	cmd.PersistentFlags().IntVarP(&bgetClis.retSleepTime, "retries-sleep-time", "", 5, "Sleep time after one retry.")
	cmd.PersistentFlags().BoolVarP(&bgetClis.remoteName, "remote-name", "n", false, "Use remote defined filename.")
	cmd.PersistentFlags().StringVarP(&(bgetClis.mirror), "mirror", "m", "", "Set the mirror of resources.")
	cmd.PersistentFlags().StringVarP(&(bgetClis.engine), "engine", "g", "go-http", "Point the download engine: go-http, wget, curl, axel, git, and rsync.")
	cmd.PersistentFlags().IntVarP(&(bgetClis.axelThread), "thread-axel", "", 5, "Set the thread of axel.")
}

func init() {
	wd, _ = os.Getwd()
}
