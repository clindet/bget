package cmd

import (
	"github.com/clindet/bget/api/fetch"
	"github.com/clindet/bget/api/types"
	"github.com/spf13/cobra"
)

var dendp types.Datasets2toolsEndpoints
var Dataset2toolsCmd = &cobra.Command{
	Use:   "dta",
	Short: "Query dataset2tools website APIs: datasets (d), tools (t), and canned analysis (a).",
	Long:  `Query dataset2tools website APIs: datasets (d), tools (t), and canned analysis (a).`,
	Run: func(cmd *cobra.Command, args []string) {
		Dataset2toolsCmdRunOptions(cmd, args)
	},
}

func Dataset2toolsCmdRunOptions(cmd *cobra.Command, args []string) {
	dendp.Query = bapiClis.Query
	dendp.PageSize = bapiClis.Size
	if dendp.ObjectType != "" || dendp.DatasetAccession != "" ||
		dendp.CannedAnalysisAccession != "" ||
		dendp.ToolName != "" || dendp.DiseaseName != "" || dendp.Gneset != "" {

		initCmd(cmd, args)
		fetch.Dataset2tools(&dendp, &bapiClis, nil)
		bapiClis.HelpFlags = false
	}
	if bapiClis.HelpFlags {
		cmd.Help()
	}
}

func init() {
	setGlobalFlag(Dataset2toolsCmd, &bapiClis)
	Dataset2toolsCmd.Flags().StringVarP(&dendp.ObjectType, "type", "", "", "Object type [tool, dataset, canned_analysis].")
	Dataset2toolsCmd.Flags().StringVarP(&dendp.ToolName, "tool", "t", "", "Tool name, e.g. bwa.")
	Dataset2toolsCmd.Flags().StringVarP(&dendp.DiseaseName, "disease", "d", "", "Disease name, e.g. prostate cancer")
	Dataset2toolsCmd.Flags().StringVarP(&dendp.DatasetAccession, "dataset-acc", "s", "", "Dataset accession number, e.g. GSE31106.")
	Dataset2toolsCmd.Flags().StringVarP(&dendp.CannedAnalysisAccession, "analysis-acc", "a", "", "Canned analysis accession	, e.g. DCA00000060.")
	Dataset2toolsCmd.Flags().StringVarP(&dendp.Gneset, "geneset", "g", "", "With dataset accession, e.g. upregulated.")
	Dataset2toolsCmd.Flags().IntVarP(&bapiClis.From, "from", "", -1, "Parameters of API control the start item of retrived data.")
	Dataset2toolsCmd.Flags().IntVarP(&bapiClis.Size, "size", "", -1, "Parameters of API control the lenth of retrived data. Default is auto determined.")
	Dataset2toolsCmd.Flags().StringVarP(&bapiClis.Query, "query", "q", "", "Query specifies the search query for record retrieval (required).")

	Dataset2toolsCmd.Example = `  # query canned analysis accession	, e.g. DCA00000060.
  bget api dta -a DCA00000060
  # query dataset accession number, e.g. GSE31106 
  bget api dta -s GSE31106 | bioctl fmt --json-pretty -
  # query via object type
  bget api dta --type dataset | bioctl fmt --json-pretty --indent 2 -
  # props of dataset accession, e.g. upregulated.
  bget api dta -g upregulated | json2csv -o out.csv`
}
