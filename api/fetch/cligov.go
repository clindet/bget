package fetch

import (
	"io"
	"strconv"
	"strings"

	"github.com/clindet/bget/api/types"
)

const CligovHost = "https://clinicaltrials.gov/api/"

// Cligov access https://clinicaltrials.gov API
func Cligov(endpoints *types.CligovEndpoints, BapiClis *types.BapiClisT, f func(), of io.Writer) bool {
	setLog(BapiClis)
	if BapiClis.Format == "" {
		BapiClis.Format = "json"
	}
	netopt := setNetOpt(BapiClis)
	url := CligovHost + setCligovQuerySuffix(endpoints, BapiClis)
	if url == CligovHost || url == CligovHost+"?fmt=json" {
		return false
	}
	f()
	GetReq("clinicaltrials.gov", url, BapiClis, netopt, of)

	return true
}

func setCligovQuerySuffix(endpoints *types.CligovEndpoints, BapiClis *types.BapiClisT) (suffix string) {
	suffixList := []string{}
	if endpoints.InfoDataVrs {
		suffix = "info/data_vrs"
	} else if endpoints.InfoAPIVrs {
		suffix = "info/api_vrs"
	} else if endpoints.InfoAPIDefs {
		suffix = "info/api_defs"
	} else if endpoints.InfoStuStru {
		suffix = "info/study_structure"
	} else if endpoints.InfoStuFieldsList {
		suffix = "info/study_fields_list"
	} else if endpoints.InfoStuStat {
		suffix = "info/study_statistics"
	} else if endpoints.InfoSearchArea {
		suffix = "info/search_areas"
	} else if endpoints.StuFields {
		suffix = "query/study_fields"
	} else if endpoints.FullStudies {
		suffix = "query/full_studies"
	} else if endpoints.FieldValues {
		suffix = "query/field_values"
	}
	if BapiClis.Query != "" {
		suffixList = append(suffixList, "expr="+BapiClis.Query)
	}
	if endpoints.Field != "" {
		suffixList = append(suffixList, "field="+endpoints.Field)
	}
	if endpoints.Fields != "" {
		suffixList = append(suffixList, "fields="+endpoints.Fields)
	}
	if BapiClis.Format != "" {
		suffixList = append(suffixList, "fmt="+BapiClis.Format)
	}
	if BapiClis.From != -1 {
		suffixList = append(suffixList, "min_rnk="+strconv.Itoa(BapiClis.From))
	}
	if BapiClis.Size != -1 && BapiClis.From != -1 {
		suffixList = append(suffixList, "max_rnk="+strconv.Itoa(BapiClis.From+BapiClis.Size))
	} else if BapiClis.Size != -1 {
		suffixList = append(suffixList, "max_rnk="+strconv.Itoa(1+BapiClis.Size))
	}
	if BapiClis.Extra != "" {
		suffixList = append(suffixList, BapiClis.Extra)
	}
	if len(suffixList) > 0 {
		suffix = suffix + "?" + strings.Join(suffixList, "&")
	}
	return suffix
}
