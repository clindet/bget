package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/clindet/bget/api/types"
	cvrt "github.com/openbiox/ligo/convert"
	cio "github.com/openbiox/ligo/io"
	clog "github.com/openbiox/ligo/log"

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
	if bapiClis.Verbose == 2 {
		logEnv.Infof("Prog: %s", cmd.CommandPath())
		logEnv.Infof("TaskID: %s", bapiClis.TaskID)
		if bapiClis.SaveLog && logPrefix != "" {
			logEnv.Infof("Log: %s.log", logPrefix)
		}
		if len(args) > 0 {
			logEnv.Infof("Args: %s", strings.Join(args, " "))
		}
		logEnv.Infof("Global: %v", cvrt.Struct2Map(bapiClis))
	}
}

func setLog() {
	if bapiClis.SaveLog {
		if bapiClis.LogDir == "" {
			bapiClis.LogDir = filepath.Join(os.TempDir(), "_log")
		}
		logPrefix = fmt.Sprintf("%s/%s", bapiClis.LogDir, bapiClis.TaskID)
		cio.CreateDir(bapiClis.LogDir)
		logCon, _ = cio.Open(logPrefix + ".log")
	}
	clog.SetLogStream(log, bapiClis.Verbose == 0, bapiClis.SaveLog, &logCon)
}

func setGlobalFlag(cmd *cobra.Command, bapiClis *types.BapiClisT) {
	cmd.Flags().StringVarP(&bapiClis.Format, "format", "", "", "Rettype specifies the format of the returned data (CSV, TSV, JSON for gdc; XML/TEXT for ncbi).")
	cmd.Flags().BoolVarP(&bapiClis.XML2json, "xml2json", "", false, "convert returned XML data to json format.")
	cmd.PersistentFlags().BoolVarP(&bapiClis.PrettyJSON, "json-pretty", "", false, "Pretty json files.")
	cmd.PersistentFlags().IntVarP(&bapiClis.Indent, "indent", "", 4, "Control the indent of output json files.")
	cmd.PersistentFlags().BoolVarP(&bapiClis.SortKeys, "sort-keys", "", false, "Control wheather to sort JSON key.")
	cmd.PersistentFlags().IntVarP(&bapiClis.Retries, "retries", "r", 2, "Retry specifies the number of attempts to retrieve the data.")
	cmd.PersistentFlags().IntVarP(&bapiClis.Timeout, "timeout", "", 35, "Set the timeout of per request.")
	cmd.PersistentFlags().StringVarP(&(bapiClis.Proxy), "proxy", "", "", "HTTP proxy to query.")
	cmd.PersistentFlags().IntVarP(&bapiClis.RetSleepTime, "retries-sleep-time", "", 5, "Sleep time after one retry.")
	cmd.PersistentFlags().StringVarP(&bapiClis.Extra, "extra", "", "", "Extra query parameters.")
}
