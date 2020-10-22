package cmd

import (
	"github.com/clindet/bget/api/fetch"
	"github.com/clindet/bget/api/types"
	"github.com/spf13/cobra"
)

var covid19Endp types.Covid19Endpoints
var Covid19Cmd = &cobra.Command{
	Use:   "covid19",
	Short: "Query covid19api.com website APIs.",
	Long:  `Query covid19api.com website APIs. Detail see https://covid19api.com`,
	Run: func(cmd *cobra.Command, args []string) {
		Covid19CmdRunOptions(cmd, args)
	},
}

func Covid19CmdRunOptions(cmd *cobra.Command, args []string) {
	if fetch.Covid19(&covid19Endp, &bapiClis, func() { initCmd(cmd, args) }, nil) {
		bapiClis.HelpFlags = false
	}
	if bapiClis.HelpFlags {
		cmd.Help()
	}
}

func init() {
	setGlobalFlag(Covid19Cmd, &bapiClis)
	Covid19Cmd.Flags().BoolVarP(&covid19Endp.AllRoute, "all", "", false, `Get all Data.`)
	Covid19Cmd.Flags().BoolVarP(&covid19Endp.CountriesRoute, "cts", "", false, `Get list of countries.`)
	Covid19Cmd.Flags().BoolVarP(&covid19Endp.CountryDayOneRoute, "ct-d-one", "", false, `Get list of cases per country per province by case type from the first recorded case.`)
	Covid19Cmd.Flags().BoolVarP(&covid19Endp.CountryDayOneTotalRoute, "ct-d-one-total", "", false, `Get list of cases per country by case type from the first recorded case.`)
	Covid19Cmd.Flags().BoolVarP(&covid19Endp.CountryRoute, "ct", "", false, `Get list of cases per country per province by case type.`)
	Covid19Cmd.Flags().BoolVarP(&covid19Endp.CountryStatusDayOneLiveRoute, "ct-st-d-one-live", "", false, `Get list of cases per country per province by case type from the first recorded case,  updated with latest live count.`)
	Covid19Cmd.Flags().BoolVarP(&covid19Endp.CountryStatusDayOneRoute, "ct-st-d-one", "", false, `Get list of cases per country per province by case type from the first recorded case".`)
	Covid19Cmd.Flags().BoolVarP(&covid19Endp.CountryStatusDayOneTotalRoute, "ct-st-d-one-total", "", false, `Get list of cases per country by case type from the first recorded case.`)
	Covid19Cmd.Flags().BoolVarP(&covid19Endp.CountryStatusLiveRoute, "ct-st-live", "", false, `Daily list of cases per country per province by case type, updated with latest live count.`)
	Covid19Cmd.Flags().BoolVarP(&covid19Endp.CountryStatusRoute, "ct-st", "", false, `Get list of cases per country per province by case type.`)
	Covid19Cmd.Flags().BoolVarP(&covid19Endp.CountryStatusTotalRoute, "ct-st-total", "", false, `Get list of cases per country by case type.`)
	Covid19Cmd.Flags().BoolVarP(&covid19Endp.CountryTotalRoute, "ct-total", "", false, `Get list of cases per country by case type.`)
	Covid19Cmd.Flags().BoolVarP(&covid19Endp.ExportRoute, "export", "", false, `Get all data as a zip file.`)
	Covid19Cmd.Flags().BoolVarP(&covid19Endp.LiveCountryRoute, "live-ct", "", false, `Get live list of cases per country per province by case type.`)
	Covid19Cmd.Flags().BoolVarP(&covid19Endp.LiveCountryStatusAfterDateRoute, "live-ct-st-date", "", false, `Get a time series of cases per country per province by case type after a date.`)
	Covid19Cmd.Flags().BoolVarP(&covid19Endp.LiveCountryStatusRoute, "live-ct-st", "", false, `Get a time series of cases per country per province by case type.`)
	Covid19Cmd.Flags().BoolVarP(&covid19Endp.SummaryRoute, "summary", "", false, `Summary of new and total cases per country.`)
	Covid19Cmd.Flags().StringVarP(&covid19Endp.WebhookRoute, "webhook", "", "", `Add a webhook to be notified when new data becomes available.`)

	Covid19Cmd.Flags().StringVarP(&covid19Endp.Country, "name", "", "", `Country name.`)
	Covid19Cmd.Flags().StringVarP(&covid19Endp.Status, "status", "", "", `Patient status [confirmed, recovered, deaths].`)
	Covid19Cmd.Flags().StringVarP(&covid19Endp.Date, "date", "", "", `After a given date.`)

	Covid19Cmd.Example = `  # returns all data in the system. Warning: this request returns 8MB+ and takes 5+ seconds
  bget api covid19 --all
  # returns all countries and associated provinces
  bget api covid19 --cts
  # returns all cases by case type for a country from the first recorded case.
  bget api covid19 --ct --name China
  # returns all cases by case type for a country from the first recorded case.
  bget api covid19 --ct-d-one --name China
  # returns all cases by case type for a country.
  bget api covid19 --ct-d-one-total --name China
  # returns all cases by case type for a country from the first recorded case with the latest record being the live count.
  bget api covid19 --ct-st-d-one --name China --status confirmed
  # returns all cases by case type for a country from the first recorded case.
  bget api covid19 --ct-st-d-one-live --name China --status confirmed
  # returns all cases by case type for a country from the first recorded case
  bget api covid19 --ct-st-d-one-total --name China --status confirmed
  # returns all cases by case type for a country with the latest record being the live count.
  bget api covid19 --ct-st-live --name China --status confirmed
  # returns all cases by case type for a country.
  bget api covid19 --ct-st --name China --status confirmed
  # returns all cases by case type for a country.
  bget api covid19 --ct-st-total --name China --status confirmed
  # returns all cases of a country.
  bget api covid19 --ct-total --name China
  # returns all live cases by case type for a country. 
  bget api covid19 --live-ct --name China
  # returns all live cases by case type for a country after a given date.
  bget api covid19 --live-ct-st-date --name China --status confirmed --date 2020-04-20T06:20:47Z
  # returns all live cases by case type for a country.
  bget api covid19 --live-ct-st --name China --status confirmed
  # a summary of new and total cases per country
  bget api covid19 --summary
  bget api covid19 --export
  bget api covid19 --webhook https://your_webhook.com`
}
