package cmd

import (
	"github.com/clindet/bget/api/fetch"
	"github.com/clindet/bget/api/types"
	"github.com/spf13/cobra"
)

var crossRefEndp types.CrossRefEndpoints
var CrossRefCmd = &cobra.Command{
	Use:   "crf",
	Short: "Query crossref guestquery APIs.",
	Long:  `Query crossref guestquery APIs. Detail see https://www.crossref.org/guestquery website.`,
	Run: func(cmd *cobra.Command, args []string) {
		CrossRefCmdRunOptions(cmd, args)
	},
}

func CrossRefCmdRunOptions(cmd *cobra.Command, args []string) {
	if fetch.CrossRef(&crossRefEndp, &bapiClis, func() { initCmd(cmd, args) }, nil) {
		bapiClis.HelpFlags = false
	}
	if bapiClis.HelpFlags {
		cmd.Help()
	}
}

func init() {
	setGlobalFlag(CrossRefCmd, &bapiClis)
	CrossRefCmd.Flags().StringVarP(&crossRefEndp.Doi.Doi, "doi", "", "", `unique doi.`)

	CrossRefCmd.Example = `  bget api crf --doi 10.1073/pnas.1814397115
  bget api crf --doi 10.1073/pnas.1814397115 --xml2json --json-pretty --indent 1`
}
