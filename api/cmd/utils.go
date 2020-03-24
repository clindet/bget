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

	"github.com/openbiox/bget/api/types"
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
	if BapiClis.Verbose == 2 {
		logEnv.Infof("Prog: %s", cmd.CommandPath())
		logEnv.Infof("TaskID: %s", BapiClis.TaskID)
		if BapiClis.SaveLog && logPrefix != "" {
			logEnv.Infof("Log: %s.log", logPrefix)
		}
		if len(args) > 0 {
			logEnv.Infof("Args: %s", strings.Join(args, " "))
		}
		logEnv.Infof("Global: %v", cvrt.Struct2Map(BapiClis))
	}
}

func setLog() {
	if BapiClis.SaveLog {
		if BapiClis.LogDir == "" {
			BapiClis.LogDir = filepath.Join(os.TempDir(), "_log")
		}
		logPrefix = fmt.Sprintf("%s/%s", BapiClis.LogDir, BapiClis.TaskID)
		cio.CreateDir(BapiClis.LogDir)
		logCon, _ = cio.Open(logPrefix + ".log")
	}
	clog.SetLogStream(log, BapiClis.Verbose == 0, BapiClis.SaveLog, &logCon)
}

func setGlobalFlag(cmd *cobra.Command, bapiClis *types.BapiClisT) {
	cmd.Flags().StringVarP(&BapiClis.Format, "format", "", "", "Rettype specifies the format of the returned data (CSV, TSV, JSON for gdc; XML/TEXT for ncbi).")
	cmd.PersistentFlags().BoolVarP(&BapiClis.PrettyJSON, "json-pretty", "", false, "Pretty json files.")
	cmd.PersistentFlags().IntVarP(&BapiClis.Indent, "indent", "", 4, "Control the indent of output json files.")
	cmd.PersistentFlags().BoolVarP(&BapiClis.SortKeys, "sort-keys", "", false, "Control wheather to sort JSON key.")
	cmd.PersistentFlags().IntVarP(&BapiClis.Retries, "retries", "r", 5, "Retry specifies the number of attempts to retrieve the data.")
	cmd.PersistentFlags().IntVarP(&BapiClis.Timeout, "timeout", "", 35, "Set the timeout of per request.")
	cmd.PersistentFlags().StringVarP(&(BapiClis.Proxy), "proxy", "", "", "HTTP proxy to query.")
	cmd.PersistentFlags().IntVarP(&BapiClis.RetSleepTime, "retries-sleep-time", "", 5, "Sleep time after one retry.")
	cmd.PersistentFlags().StringVarP(&BapiClis.Extra, "extra", "", "", "Extra query parameters.")
}
