package cmd

import (
	"github.com/clindet/bget/api/fetch"
	"github.com/clindet/bget/api/types"
	"github.com/spf13/cobra"
)

var bioToolsEndp types.BioToolsEndpoints
var BioToolsCmd = &cobra.Command{
	Use:   "biots",
	Short: "Query bio.tools website APIs.",
	Long:  `Query bio.tools website APIs. Detail see https://biotools.readthedocs.io/en/latest/api_reference.html`,
	Run: func(cmd *cobra.Command, args []string) {
		BioToolsCmdRunOptions(cmd, args)
	},
}

func BioToolsCmdRunOptions(cmd *cobra.Command, args []string) {
	if fetch.BioTools(&bioToolsEndp, &bapiClis, func() { initCmd(cmd, args) }, nil) {
		bapiClis.HelpFlags = false
	}
	if bapiClis.HelpFlags {
		cmd.Help()
	}
}

func init() {
	setGlobalFlag(BioToolsCmd, &bapiClis)
	BioToolsCmd.Flags().StringVarP(&bioToolsEndp.Tool, "tool", "", "", `Obtain information about a single tool (https://bio.tools/api/tool/:id/).`)
	BioToolsCmd.Flags().StringVarP(&bioToolsEndp.ID, "id", "", "", `Search for bio.tools tool ID e.g signalp)`)
	BioToolsCmd.Flags().StringVarP(&bioToolsEndp.Name, "name", "", "", `Search for bio.tools tool name e.g signalp)`)
	BioToolsCmd.Flags().StringVarP(&bioToolsEndp.Topic, "topic", "", "", `Search for EDAM Topic (term)`)
	BioToolsCmd.Flags().StringVarP(&bioToolsEndp.DataType, "dtype", "", "", `Fuzzy search over input and output for EDAM Data (term)`)
	BioToolsCmd.Flags().StringVarP(&bioToolsEndp.DataFormat, "dfmt", "", "", `Fuzzy search over input and output for EDAM Format (term)`)
	BioToolsCmd.Flags().StringVarP(&bioToolsEndp.OutputFormat, "ofmt", "", "", `Fuzzy search over output for EDAM Format (term)`)
	BioToolsCmd.Flags().StringVarP(&bioToolsEndp.Publication, "publication", "", "", `Fuzzy search over publication (DOI, PMID, PMCID, publication type and tool version)`)
	BioToolsCmd.Flags().StringVarP(&bapiClis.Query, "query", "q", "", "Query specifies the search query for record retrieval (required).")
	BioToolsCmd.Flags().IntVarP(&bapiClis.From, "from", "", -1, "Parameters of API control the start item of retrived data.")
	BioToolsCmd.Flags().IntVarP(&bapiClis.Size, "size", "", -1, "Parameters of API control the lenth of retrived data. Default is auto determined.")

	BioToolsCmd.Example = `  # query item detail
  bget api biots --tool signalp
  # search item
  bget api biots --name signalp
  bget api biots --topic Proteomics
  bget api biots --dtype 'Protein sequence'
  bget api biots --dfmt FASTA
  bget api biots --ofmt 'ClustalW format'`
}
