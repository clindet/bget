package cmd

import (
	"github.com/openbiox/bget/api/types"
	"github.com/spf13/cobra"
)

func setGlobalFlag(cmd *cobra.Command, bapiClis *types.BapiClisT) {
	cmd.PersistentFlags().StringVarP(&BapiClis.Format, "format", "", "", "Rettype specifies the format of the returned data (CSV, TSV, JSON for gdc; XML/TEXT for ncbi).")
	cmd.PersistentFlags().StringVarP(&BapiClis.Quiet, "quiet", "", "false", "No log output.")
	cmd.PersistentFlags().BoolVarP(&BapiClis.PrettyJSON, "json-pretty", "", false, "Pretty json files.")
	cmd.PersistentFlags().IntVarP(&BapiClis.Indent, "indent", "", 4, "Control the indent of output json files.")
	cmd.PersistentFlags().BoolVarP(&BapiClis.SortKeys, "sort-keys", "", false, "Control wheather to sort JSON key.")
	cmd.PersistentFlags().IntVarP(&BapiClis.Retries, "retries", "r", 5, "Retry specifies the number of attempts to retrieve the data.")
	cmd.PersistentFlags().IntVarP(&BapiClis.Timeout, "timeout", "", 35, "Set the timeout of per request.")
	cmd.PersistentFlags().IntVarP(&BapiClis.RetSleepTime, "retries-sleep-time", "", 5, "Sleep time after one retry.")
	cmd.PersistentFlags().StringVarP(&BapiClis.Query, "query", "q", "", "Query specifies the search query for record retrieval (required).")
	cmd.PersistentFlags().IntVarP(&BapiClis.From, "from", "", -1, "Parameters of API control the start item of retrived data.")
	cmd.PersistentFlags().IntVarP(&BapiClis.Size, "size", "", -1, "Parameters of API control the lenth of retrived data. Default is auto determined.")
	cmd.PersistentFlags().StringVarP(&BapiClis.Extra, "extra", "", "", "Extra query parameters.")
}
