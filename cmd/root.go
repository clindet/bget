package cmd

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/openbiox/butils/log"
	cnet "github.com/openbiox/butils/net"
	"github.com/openbiox/butils/stringo"
	"github.com/spf13/cobra"
	"github.com/vbauerster/mpb/v4"
)

var version = "v0.1.3-4"

type bgetCliT struct {
	downloadDir      string
	mirror           string
	autoPath         bool
	engine           string
	doi              string
	urls             string
	listFile         string
	seperator        string
	keys             string
	seqs             string
	gdcToken         string
	uncompress       bool
	keysAll          bool
	clean            bool
	printFormat      string
	axelThread       int
	thread           int
	timeout          int
	retSleepTime     int
	retries          int
	proxy            string
	cmdExtraFromFlag string
	remoteName       bool
	ignore           bool
	taskID           string
	overwrite        bool
	saveLog          bool
	logDir           string
	quiet            bool
	env              map[string]string
	showVersions     bool
	egaCredFile      string
	outjson          bool
	outxt            bool
	geoGPL           bool
	helpFlags        bool
}

var bgetClis bgetCliT

var rootCmd = &cobra.Command{
	Use:   "bget",
	Short: "Lightweight downloader for bioinformatics data, databases and files.",
	Long:  `Lightweight downloader for bioinformatics data, databases and files (under development). It will provides a simple and parallelized method to access various bioinformatics resoures. More see here https://github.com/openbiox/bget.`,
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

func checkArgs(cmd *cobra.Command, subcmd string) {
	items := []string{}
	for _, v := range cmd.Flags().Args() {
		if strings.Contains(v, "=") && subcmd == "key" {
			kvs := strings.Split(v, "=")
			bgetClis.env[kvs[0]] = strings.TrimSpace(kvs[1])
		} else {
			items = append(items, v)
		}
	}
	if len(items) == 0 {
		return
	}
	if subcmd == "url" {
		bgetClis.urls = strings.Join(items, bgetClis.seperator)
	} else if subcmd == "key" {
		bgetClis.keys = strings.Join(items, bgetClis.seperator)
	} else if subcmd == "doi" {
		bgetClis.doi = strings.Join(items, bgetClis.seperator)
	} else if subcmd == "seq" {
		bgetClis.seqs = strings.Join(items, bgetClis.seperator)
	}
}

func rootCmdRunOptions(cmd *cobra.Command) {
	checkQuiet()
	if bgetClis.clean {
		if err := os.RemoveAll("_download"); err != nil {
			log.Warn(err)
		}
		if err := os.RemoveAll("_log"); err != nil {
			log.Warn(err)
		}
		bgetClis.helpFlags = false
	}
	if bgetClis.helpFlags {
		cmd.Help()
	}
}

func setNetParams(bgetClis *bgetCliT) (netOpt *cnet.BnetParams) {
	pbar := mpb.New(
		mpb.WithWidth(45),
		mpb.WithRefreshRate(180*time.Millisecond),
	)
	return &cnet.BnetParams{
		Proxy:          bgetClis.proxy,
		Engine:         bgetClis.engine,
		ExtraArgs:      bgetClis.cmdExtraFromFlag,
		TaskID:         bgetClis.taskID,
		Mirror:         bgetClis.mirror,
		Thread:         bgetClis.thread,
		AxelThread:     bgetClis.axelThread,
		Overwrite:      bgetClis.overwrite,
		Ignore:         bgetClis.ignore,
		Quiet:          bgetClis.quiet,
		Retries:        bgetClis.retries,
		SaveLog:        bgetClis.saveLog,
		Timeout:        bgetClis.timeout,
		RetSleepTime:   bgetClis.retSleepTime,
		RemoteName:     bgetClis.remoteName,
		LogDir:         bgetClis.logDir,
		EgaCredentials: bgetClis.egaCredFile,
		Pbar:           pbar,
	}
}

func init() {
	bgetClis.helpFlags = true
	wd, _ := os.Getwd()
	rootCmd.AddCommand(urlCmd)
	rootCmd.AddCommand(doiCmd)
	rootCmd.AddCommand(keyCmd)
	rootCmd.AddCommand(seqCmd)

	rootCmd.PersistentFlags().StringVarP(&(bgetClis.proxy), "proxy", "", "", "HTTP proxy to download.")
	rootCmd.Flags().BoolVarP(&(bgetClis.clean), "clean", "", false, "Remove _download and _log in current dir.")
	rootCmd.PersistentFlags().IntVarP(&(bgetClis.thread), "thread", "t", 1, "Concurrency download thread.")
	rootCmd.Flags().IntVarP(&(bgetClis.axelThread), "thread-axel", "", 5, "Set the thread of axel.")
	rootCmd.PersistentFlags().StringVarP(&(bgetClis.downloadDir), "outdir", "o", wd, "Set the download dir.")
	rootCmd.PersistentFlags().StringVarP(&(bgetClis.seperator), "seperator", "s", ",", "Optional 'url1{seperator}url2' for multiple keys, urls, or seqs.")
	rootCmd.PersistentFlags().BoolVar(&(bgetClis.ignore), "ignore", false, "Contine to download and skip the check of existed files.")
	rootCmd.PersistentFlags().StringVarP(&(bgetClis.cmdExtraFromFlag), "extra-cmd", "e", "", "Extra flags and values pass to internal CMDs")
	rootCmd.PersistentFlags().BoolVarP(&(bgetClis.overwrite), "overwrite", "f", false, "Logical indicating that whether to overwrite existing files.")
	rootCmd.PersistentFlags().StringVarP(&(bgetClis.taskID), "task-id", "", stringo.GetRandomString(15), "Task ID (random).")
	rootCmd.PersistentFlags().StringVarP(&(bgetClis.logDir), "log-dir", "", path.Join(wd, "_log"), "Log dir.")
	rootCmd.PersistentFlags().BoolVarP(&(bgetClis.quiet), "quiet", "q", false, "No output.")
	rootCmd.PersistentFlags().BoolVarP(&(bgetClis.saveLog), "save-log", "", true, "Save download log to local file].")
	rootCmd.PersistentFlags().IntVarP(&bgetClis.retries, "retries", "r", 5, "Retry specifies the number of attempts to retrieve the data.")
	rootCmd.PersistentFlags().IntVarP(&bgetClis.timeout, "timeout", "", 35, "Set the timeout of per request.")
	rootCmd.PersistentFlags().IntVarP(&bgetClis.retSleepTime, "retries-sleep-time", "", 5, "Sleep time after one retry.")
	rootCmd.PersistentFlags().BoolVarP(&bgetClis.remoteName, "remote-name", "n", false, "Use remote defined filename.")
	rootCmd.PersistentFlags().BoolVarP(&(bgetClis.uncompress), "uncompress", "u", false, "Uncompress download files for .zip, .tar.gz, and .gz suffix files.")
	rootCmd.PersistentFlags().StringVarP(&(bgetClis.mirror), "mirror", "m", "", "Set the mirror of resources.")
	rootCmd.PersistentFlags().StringVarP(&(bgetClis.engine), "engine", "g", "go-http", "Point the download engine: go-http, wget, curl, axel, git, and rsync.")
	rootCmd.PersistentFlags().StringVarP(&(bgetClis.listFile), "list-file", "l", "", "A file contains dois for download.")
	bgetClis.env = make(map[string]string)
	bgetClis.env["osType"] = runtime.GOOS
	bgetClis.env["wd"], _ = os.Getwd()
	rootCmd.Version = version
}
