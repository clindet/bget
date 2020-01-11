package fetch

import (
	"strconv"
	"strings"

	"github.com/openbiox/bget/bapi/types"
)

const CligovHost = "https://clinicaltrials.gov/api/"

// Cligov access https://clinicaltrials.gov API
func Cligov(endpoints *types.CligovEndpoints, bapiClis *types.BapiClisT) bool {
	if bapiClis.Format == "" {
		bapiClis.Format = "json"
	}
	netopt := setNetOpt(bapiClis)
	url := CligovHost + setCligovQuerySuffix(endpoints, bapiClis)
	if url == CligovHost {
		return false
	}
	queryAPI("clinicaltrials.gov", url, bapiClis, netopt)

	return true
}

func setCligovQuerySuffix(endpoints *types.CligovEndpoints, bapiClis *types.BapiClisT) (suffix string) {
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
	if bapiClis.Query != "" {
		suffixList = append(suffixList, "expr="+bapiClis.Query)
	}
	if endpoints.Field != "" {
		suffixList = append(suffixList, "field="+endpoints.Field)
	}
	if endpoints.Fields != "" {
		suffixList = append(suffixList, "fields="+endpoints.Fields)
	}
	if bapiClis.Format != "" {
		suffixList = append(suffixList, "fmt="+bapiClis.Format)
	}
	if bapiClis.From != -1 {
		suffixList = append(suffixList, "min_rnk="+strconv.Itoa(bapiClis.From))
	}
	if bapiClis.Size != -1 && bapiClis.From != -1 {
		suffixList = append(suffixList, "max_rnk="+strconv.Itoa(bapiClis.From+bapiClis.Size))
	} else if bapiClis.Size != -1 {
		suffixList = append(suffixList, "max_rnk="+strconv.Itoa(1+bapiClis.Size))
	}
	if bapiClis.Extra != "" {
		suffixList = append(suffixList, bapiClis.Extra)
	}
	if len(suffixList) > 0 {
		suffix = suffix + "?" + strings.Join(suffixList, "&")
	}
	return suffix
}
