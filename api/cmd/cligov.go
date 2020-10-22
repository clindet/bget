package cmd

import (
	"github.com/clindet/bget/api/fetch"
	"github.com/clindet/bget/api/types"
	"github.com/spf13/cobra"
)

var cligovEndp types.CligovEndpoints
var CligovCmd = &cobra.Command{
	Use:   "cligov",
	Short: "Query clinicaltrials.gov website APIs.",
	Long:  `Query clinicaltrials.gov website APIs. Detail see https://clinicaltrials.gov/api/gui/ref/api_urls`,
	Run: func(cmd *cobra.Command, args []string) {
		CligovCmdRunOptions(cmd, args)
	},
}

func CligovCmdRunOptions(cmd *cobra.Command, args []string) {
	if fetch.Cligov(&cligovEndp, &bapiClis, func() { initCmd(cmd, args) }, nil) {
		bapiClis.HelpFlags = false
	}
	if bapiClis.HelpFlags {
		cmd.Help()
	}
}

func init() {
	setGlobalFlag(CligovCmd, &bapiClis)
	CligovCmd.Flags().BoolVarP(&cligovEndp.InfoDataVrs, "info-dat-vers", "", false, `Returns the date when the ClinicalTrials.gov dataset was posted.`)
	CligovCmd.Flags().BoolVarP(&cligovEndp.InfoAPIVrs, "info-api-vers", "", false, `Returns the current version number of the ClinicalTrials.gov API.`)
	CligovCmd.Flags().BoolVarP(&cligovEndp.InfoAPIDefs, "info-api-defs", "", false, `Returns detailed definitions.`)
	CligovCmd.Flags().BoolVarP(&cligovEndp.InfoStuStru, "info-study-struct", "", false, `Returns all available data elements for a single study record.`)
	CligovCmd.Flags().BoolVarP(&cligovEndp.InfoStuFieldsList, "info-study-fields", "", false, `Returns all study data elements.`)
	CligovCmd.Flags().BoolVarP(&cligovEndp.InfoStuStat, "info-study-stat", "", false, `Returns an annotated version of the Study Structure info URL.`)
	CligovCmd.Flags().BoolVarP(&cligovEndp.InfoSearchArea, "info-search-area", "", false, `Returns groups of weighted study fields, or "search areas".`)

	CligovCmd.Flags().BoolVarP(&cligovEndp.FullStudies, "full-studies", "", false, `Returns all content for a small set of study records.`)
	CligovCmd.Flags().BoolVarP(&cligovEndp.FieldValues, "field-values", "", false, `Returns all values found in a single API field for a set of study records.`)
	CligovCmd.Flags().BoolVarP(&cligovEndp.StuFields, "study-fields", "", false, `Returns values from selected API fields for a large set of study records.`)

	CligovCmd.Flags().StringVarP(&cligovEndp.Fields, "fields", "", "", `Specifies which fields to return results for in a Study Fields query.`)
	CligovCmd.Flags().StringVarP(&cligovEndp.Field, "field", "", "", `Specifies which field to collect values for in a Field Values query.`)
	CligovCmd.Flags().IntVarP(&bapiClis.From, "from", "", -1, "Parameters of API control the start item of retrived data.")
	CligovCmd.Flags().IntVarP(&bapiClis.Size, "size", "", -1, "Parameters of API control the lenth of retrived data. Default is auto determined.")
	CligovCmd.Flags().StringVarP(&bapiClis.Query, "query", "q", "", "Query specifies the search query for record retrieval (required).")

	CligovCmd.Example = `  # returns the date when the ClinicalTrials.gov dataset was posted.
  bget api cligov --info-dat-vers
  # returns the current version number of the ClinicalTrials.gov API
  bget api cligov --info-api-vers
  # returns detailed definitions.
  bget api cligov --info-api-defs
  # returns all available data elements for a single study record.
  bget api cligov --info-study-struct
  # returns all data elements.
  bget api cligov --info-study-fields
  # returns an annotated version of the Study Structure info URL.
  bget api cligov --info-study-stat
  # returns groups of weighted study fields, or "search areas"
  bget api cligov --info-search-area
	
  # query functions
  bget api cligov -q heart+attack --full-studies --format json
  bget api cligov -q heart+attack --fields NCTId,Condition,BriefTitle --study-fields
  bget api cligov -q heart+attack --field Condition --field-values`
}
