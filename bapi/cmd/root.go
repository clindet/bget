package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/openbiox/bget/bapi/types"
	cio "github.com/openbiox/butils/io"
	"github.com/openbiox/butils/log"
	"github.com/spf13/cobra"
)

var BapiClis = types.BapiClisT{}

var BapiCmd = &cobra.Command{
	Use:   "api",
	Short: "Query bioinformatics website APIs.",
	Long:  `Query bioinformatics website APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		BapiCmdRunOptions(cmd)
	},
}

// Execute main interface of butils
func Execute() {
	if err := BapiCmd.Execute(); err != nil {
		if !BapiCmd.HasFlags() && !BapiCmd.HasSubCommands() {
			BapiCmd.Help()
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func init() {
	BapiClis.Quiet = false
	BapiClis.HelpFlags = true
	BapiClis.Version = "v0.1.0"
	BapiCmd.AddCommand(ncbiCmd)
	BapiCmd.AddCommand(gdcCmd)
	BapiCmd.AddCommand(dataset2toolsCmd)
	BapiCmd.PersistentFlags().StringVarP(&BapiClis.Query, "query", "q", "", "Query specifies the search query for record retrieval (required).")
	BapiCmd.PersistentFlags().StringVarP(&BapiClis.Format, "format", "", "", "Rettype specifies the format of the returned data (CSV, TSV, JSON for gdc; XML/TEXT for ncbi).")
	BapiCmd.PersistentFlags().BoolVarP(&BapiClis.Quiet, "quiet", "", false, "No log output.")
	BapiCmd.PersistentFlags().IntVarP(&BapiClis.From, "from", "", -1, "Parameters of API control the start item of retrived data.")
	BapiCmd.PersistentFlags().IntVarP(&BapiClis.Size, "size", "", -1, "Parameters of API control the lenth of retrived data. Default is auto determined.")
	BapiCmd.PersistentFlags().StringVarP(&BapiClis.Email, "email", "e", "your_email@domain.com", "Email specifies the email address to be sent to the server (NCBI website is required).")
	BapiCmd.PersistentFlags().IntVarP(&BapiClis.Retries, "retries", "r", 5, "Retry specifies the number of attempts to retrieve the data.")
	BapiCmd.PersistentFlags().IntVarP(&BapiClis.Timeout, "timeout", "", 35, "Set the timeout of per request.")
	BapiCmd.PersistentFlags().IntVarP(&BapiClis.RetSleepTime, "retries-sleep-time", "", 5, "Sleep time after one retry.")
	BapiCmd.PersistentFlags().BoolVarP(&BapiClis.PrettyJSON, "json-pretty", "", false, "Pretty json files.")
	BapiCmd.PersistentFlags().IntVarP(&BapiClis.Indent, "indent", "", 4, "Control the indent of output json files.")
	BapiCmd.PersistentFlags().BoolVarP(&BapiClis.SortKeys, "sort-keys", "", false, "Control wheather to sort JSON key.")
	BapiCmd.Version = BapiClis.Version
}

func BapiCmdRunOptions(cmd *cobra.Command) {
	if BapiClis.Quiet {
		log.SetOutput(ioutil.Discard)
	} else {
		log.SetOutput(os.Stderr)
	}
	if hasDir, _ := cio.PathExists(BapiClis.Outfn); BapiClis.Outfn != "" && !hasDir {
		if err := cio.CreateDir(path.Dir(BapiClis.Outfn)); err != nil {
			log.FATAL(fmt.Sprintf("Could not to create %s", path.Dir(BapiClis.Outfn)))
		}
	}
	if BapiClis.HelpFlags {
		cmd.Help()
	}
}
