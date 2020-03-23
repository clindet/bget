package cmd

import (
	"github.com/openbiox/bget/api/fetch"
	"github.com/openbiox/bget/api/types"
	"github.com/spf13/cobra"
)

var cligovEndp types.CligovEndpoints
var cligovCmd = &cobra.Command{
	Use:   "cligov",
	Short: "Query https://clinicaltrials.gov/ website APIs.",
	Long:  `Query https://clinicaltrials.gov/ website APIs. Detail see https://clinicaltrials.gov/api/gui/ref/api_urls`,
	Run: func(cmd *cobra.Command, args []string) {
		cligovCmdRunOptions(cmd, args)
	},
}

func cligovCmdRunOptions(cmd *cobra.Command, args []string) {
	if fetch.Cligov(&cligovEndp, &BapiClis, func() { initCmd(cmd, args) }) {
		BapiClis.HelpFlags = false
	}
	if BapiClis.HelpFlags {
		cmd.Help()
	}
}

func init() {
	setGlobalFlag(cligovCmd, &BapiClis)
	cligovCmd.Flags().BoolVarP(&cligovEndp.InfoDataVrs, "info-dat-vers", "", false, `Returns the date when the ClinicalTrials.gov dataset was posted.`)
	cligovCmd.Flags().BoolVarP(&cligovEndp.InfoAPIVrs, "info-api-vers", "", false, `Returns the current version number of the ClinicalTrials.gov API.`)
	cligovCmd.Flags().BoolVarP(&cligovEndp.InfoAPIDefs, "info-api-defs", "", false, `Returns detailed definitions.`)
	cligovCmd.Flags().BoolVarP(&cligovEndp.InfoStuStru, "info-study-struct", "", false, `Returns all available data elements for a single study record.`)
	cligovCmd.Flags().BoolVarP(&cligovEndp.InfoStuFieldsList, "info-study-fields", "", false, `Returns all study data elements.`)
	cligovCmd.Flags().BoolVarP(&cligovEndp.InfoStuStat, "info-study-stat", "", false, `Returns an annotated version of the Study Structure info URL.`)
	cligovCmd.Flags().BoolVarP(&cligovEndp.InfoSearchArea, "info-search-area", "", false, `Returns groups of weighted study fields, or "search areas".`)

	cligovCmd.Flags().BoolVarP(&cligovEndp.FullStudies, "full-studies", "", false, `Returns all content for a small set of study records.`)
	cligovCmd.Flags().BoolVarP(&cligovEndp.FieldValues, "field-values", "", false, `Returns all values found in a single API field for a set of study records.`)
	cligovCmd.Flags().BoolVarP(&cligovEndp.StuFields, "study-fields", "", false, `Returns values from selected API fields for a large set of study records.`)

	cligovCmd.Flags().StringVarP(&cligovEndp.Fields, "fields", "", "", `Specifies which fields to return results for in a Study Fields query.`)
	cligovCmd.Flags().StringVarP(&cligovEndp.Field, "field", "", "", `Specifies which field to collect values for in a Field Values query.`)
	cligovCmd.Flags().IntVarP(&BapiClis.From, "from", "", -1, "Parameters of API control the start item of retrived data.")
	cligovCmd.Flags().IntVarP(&BapiClis.Size, "size", "", -1, "Parameters of API control the lenth of retrived data. Default is auto determined.")

	cligovCmd.Example = `  # returns the date when the ClinicalTrials.gov dataset was posted.
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
