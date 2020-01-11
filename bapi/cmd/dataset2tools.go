package cmd

import (
	"github.com/openbiox/bget/bapi/fetch"
	"github.com/openbiox/bget/bapi/types"
	"github.com/spf13/cobra"
)

var dendp types.Datasets2toolsEndpoints
var dataset2toolsCmd = &cobra.Command{
	Use:   "dta",
	Short: "Query dataset2tools website APIs: datasets (d), tools (t), and canned analysis (a).",
	Long:  `Query dataset2tools website APIs: datasets (d), tools (t), and canned analysis (a).`,
	Run: func(cmd *cobra.Command, args []string) {
		dataset2toolsCmdRunOptions(cmd)
	},
}

func dataset2toolsCmdRunOptions(cmd *cobra.Command) {
	dendp.Query = BapiClis.Query
	dendp.PageSize = BapiClis.Size
	if dendp.ObjectType != "" || dendp.DatasetAccession != "" ||
		dendp.CannedAnalysisAccession != "" ||
		dendp.ToolName != "" || dendp.DiseaseName != "" || dendp.Gneset != "" {
		fetch.Dataset2tools(&dendp, &BapiClis)
		BapiClis.HelpFlags = false
	}
	if BapiClis.HelpFlags {
		cmd.Help()
	}
}

func init() {
	setGlobalFlag(dataset2toolsCmd, &BapiClis)
	dataset2toolsCmd.Flags().StringVarP(&dendp.ObjectType, "type", "", "", "Object type [tool, dataset, canned_analysis].")
	dataset2toolsCmd.Flags().StringVarP(&dendp.ToolName, "tool", "t", "", "Tool name, e.g. bwa.")
	dataset2toolsCmd.Flags().StringVarP(&dendp.DiseaseName, "disease", "d", "", "Disease name, e.g. prostate cancer")
	dataset2toolsCmd.Flags().StringVarP(&dendp.DatasetAccession, "dataset-acc", "s", "", "Dataset accession number, e.g. GSE31106.")
	dataset2toolsCmd.Flags().StringVarP(&dendp.CannedAnalysisAccession, "analysis-acc", "a", "", "Canned analysis accession	, e.g. DCA00000060.")
	dataset2toolsCmd.Flags().StringVarP(&dendp.Gneset, "geneset", "g", "", "With dataset accession, e.g. upregulated.")
	dataset2toolsCmd.Flags().StringVarP(&BapiClis.Outfn, "outfn", "o", "", "Out specifies destination of the returned data (default to stdout).")

	dataset2toolsCmd.Example = `  # query canned analysis accession	, e.g. DCA00000060.
  bget api dta -a DCA00000060
  # query dataset accession number, e.g. GSE31106 
  bget api dta -s GSE31106 | bget fmt --json-pretty -
  # query via object type
  bget api dta --type dataset | bget fmt --json-pretty --indent 2 -
  # props of dataset accession, e.g. upregulated.
  bget api dta -g upregulated | json2csv -o out.csv`
}
