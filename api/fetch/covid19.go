package fetch

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/openbiox/ligo/stringo"

	"github.com/clindet/bget/api/types"
)

const Covid19Host = "https://api.covid19api.com/"

// Covid19 access api.covid19api.com API
func Covid19(endpoints *types.Covid19Endpoints, bapiClis *types.BapiClisT, f func(), of io.Writer) bool {
	setLog(bapiClis)
	if bapiClis.Format == "" {
		bapiClis.Format = "json"
	}
	netopt := setNetOpt(bapiClis)
	url := Covid19Host + setCovid19QuerySuffix(endpoints, bapiClis)
	if url == Covid19Host || url == Covid19Host+"?fmt=json" {
		return false
	}
	f()
	if endpoints.WebhookRoute != "" {
		val := make(map[string]string)
		val["URL"] = endpoints.WebhookRoute
		postData, _ := json.MarshalIndent(val, "", " ")

		PostReq("api.mg-rast.org", url, postData, bapiClis, netopt, of)
	} else {
		GetReq("api.covid19api.com", url, bapiClis, netopt, of)
	}

	return true
}

func setCovid19QuerySuffix(endpoints *types.Covid19Endpoints, BapiClis *types.BapiClisT) (suffix string) {
	suffixList := []string{}
	if endpoints.AllRoute {
		suffix = "all"
	} else if endpoints.CountriesRoute {
		suffix = "countries"
	} else if endpoints.CountryDayOneRoute {
		suffix = "dayone/country/:country"
	} else if endpoints.CountryDayOneTotalRoute {
		suffix = "total/dayone/country/:country"
	} else if endpoints.CountryRoute {
		suffix = "country/:country"
	} else if endpoints.CountryStatusDayOneLiveRoute {
		suffix = "dayone/country/:country/status/:status/live"
	} else if endpoints.CountryStatusDayOneRoute {
		suffix = "dayone/country/:country/status/:status"
	} else if endpoints.CountryStatusDayOneTotalRoute {
		suffix = "total/dayone/country/:country/status/:status"
	} else if endpoints.CountryStatusLiveRoute {
		suffix = "country/:country/status/:status/live"
	} else if endpoints.CountryStatusRoute {
		suffix = "country/:country/status/:status"
	} else if endpoints.CountryStatusTotalRoute {
		suffix = "total/country/:country/status/:status"
	} else if endpoints.CountryTotalRoute {
		suffix = "total/country/:country"
	} else if endpoints.ExportRoute {
		suffix = "export"
	} else if endpoints.LiveCountryRoute {
		suffix = "live/country/:country"
	} else if endpoints.LiveCountryStatusAfterDateRoute {
		suffix = "live/country/:country/status/:status/date/:date"
	} else if endpoints.LiveCountryStatusRoute {
		suffix = "live/country/:country/status/:status"
	} else if endpoints.SummaryRoute {
		suffix = "summary"
	} else if endpoints.WebhookRoute != "" {
		suffix = "webhook"
	}
	if strings.Contains(suffix, ":country") {
		suffix = stringo.StrReplaceAll(suffix, ":country", endpoints.Country)
	}
	if strings.Contains(suffix, ":status") {
		suffix = stringo.StrReplaceAll(suffix, ":status", endpoints.Status)
	}
	if strings.Contains(suffix, ":date") {
		suffix = stringo.StrReplaceAll(suffix, ":date", endpoints.Date)
	}
	if BapiClis.Extra != "" {
		suffixList = append(suffixList, BapiClis.Extra)
	}
	if len(suffixList) > 0 {
		suffix = suffix + "?" + strings.Join(suffixList, "&")
	}
	return suffix
}
