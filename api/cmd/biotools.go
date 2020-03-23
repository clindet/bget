package cmd

import (
	"github.com/openbiox/bget/api/fetch"
	"github.com/openbiox/bget/api/types"
	"github.com/spf13/cobra"
)

var bioToolsEndp types.BioToolsEndpoints
var bioToolsCmd = &cobra.Command{
	Use:   "biots",
	Short: "Query https://bio.tools/ website APIs.",
	Long:  `Query https://bio.tools/ website APIs. Detail see https://biotools.readthedocs.io/en/latest/api_reference.html`,
	Run: func(cmd *cobra.Command, args []string) {
		bioToolsCmdRunOptions(cmd, args)
	},
}

func bioToolsCmdRunOptions(cmd *cobra.Command, args []string) {
	if fetch.BioTools(&bioToolsEndp, &BapiClis, func() { initCmd(cmd, args) }) {
		BapiClis.HelpFlags = false
	}
	if BapiClis.HelpFlags {
		cmd.Help()
	}
}

func init() {
	setGlobalFlag(bioToolsCmd, &BapiClis)
	bioToolsCmd.Flags().StringVarP(&bioToolsEndp.Tool, "tool", "", "", `Obtain information about a single tool (https://bio.tools/api/tool/:id/).`)
	bioToolsCmd.Flags().StringVarP(&bioToolsEndp.ID, "id", "", "", `Search for bio.tools tool ID e.g signalp)`)
	bioToolsCmd.Flags().StringVarP(&bioToolsEndp.Name, "name", "", "", `Search for bio.tools tool name e.g signalp)`)
	bioToolsCmd.Flags().StringVarP(&bioToolsEndp.Topic, "topic", "", "", `Search for EDAM Topic (term)`)
	bioToolsCmd.Flags().StringVarP(&bioToolsEndp.DataType, "dtype", "", "", `Fuzzy search over input and output for EDAM Data (term)`)
	bioToolsCmd.Flags().StringVarP(&bioToolsEndp.DataFormat, "dfmt", "", "", `Fuzzy search over input and output for EDAM Format (term)`)
	bioToolsCmd.Flags().StringVarP(&bioToolsEndp.OutputFormat, "ofmt", "", "", `Fuzzy search over output for EDAM Format (term)`)
	bioToolsCmd.Flags().StringVarP(&bioToolsEndp.Publication, "publication", "", "", `Fuzzy search over publication (DOI, PMID, PMCID, publication type and tool version)`)
	bioToolsCmd.Flags().IntVarP(&BapiClis.Size, "page", "", 1, "Page index.")

	bioToolsCmd.Example = `  # query item detail
  bget api biots --tool signalp
  # search item
  bget api biots --name signalp
  bget api biots --topic Proteomics
  bget api biots --dtype 'Protein sequence'
  bget api biots --dfmt FASTA
  bget api biots --ofmt 'ClustalW format'`
}
