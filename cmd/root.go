package cmd

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	api "github.com/openbiox/bget/bapi/cmd"
	"github.com/openbiox/butils/log"
	cnet "github.com/openbiox/butils/net"
	"github.com/spf13/cobra"
	mpb "github.com/vbauerster/mpb/v4"
)

var version = "v0.2.1"

type bgetCliT struct {
	downloadDir        string
	mirror             string
	autoPath           bool
	engine             string
	doi                string
	urls               string
	listFile           string
	seperator          string
	keys               string
	seqs               string
	gdcToken           string
	uncompress         bool
	keysAll            bool
	clean              bool
	printFormat        string
	axelThread         int
	thread             int
	timeout            int
	retSleepTime       int
	retries            int
	proxy              string
	cmdExtraFromFlag   string
	remoteName         bool
	ignore             bool
	taskID             string
	overwrite          bool
	saveLog            bool
	logDir             string
	quiet              bool
	env                map[string]string
	showVersions       bool
	egaCredFile        string
	geoGPL             bool
	github             string
	githubMode         bool
	onlyAssets         bool
	withAssets         bool
	withAssetsVersions string
	helpFlags          bool
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
	if bgetClis.withAssets {
		bgetClis.env["withAssets"] = "yes"
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
	rootCmd.AddCommand(URLCmd)
	rootCmd.AddCommand(DoiCmd)
	rootCmd.AddCommand(KeyCmd)
	rootCmd.AddCommand(SeqCmd)
	rootCmd.AddCommand(FmtCmd)
	rootCmd.AddCommand(api.BapiCmd)
	rootCmd.Flags().BoolVarP(&(bgetClis.clean), "clean", "", false, "Remove _download and _log in current dir.")
	bgetClis.env = make(map[string]string)
	bgetClis.env["osType"] = runtime.GOOS
	bgetClis.env["wd"], _ = os.Getwd()
	rootCmd.Version = version
}
