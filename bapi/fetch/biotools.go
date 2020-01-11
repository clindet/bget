package fetch

import (
	"strconv"
	"strings"

	"github.com/openbiox/bget/bapi/types"
)

const BioToolsHost = "https://bio.tools/api/tool/"

// BioTools access https://clinicaltrials.gov API
func BioTools(endpoints *types.BioToolsEndpoints, bapiClis *types.BapiClisT) bool {
	if bapiClis.Format == "" {
		bapiClis.Format = "json"
	}
	netopt := setNetOpt(bapiClis)
	url := BioToolsHost + setBioToolsQuerySuffix(endpoints, bapiClis)
	if url == BioToolsHost {
		return false
	}
	queryAPI("clinicaltrials.gov", url, bapiClis, netopt)

	return true
}

func setBioToolsQuerySuffix(endpoints *types.BioToolsEndpoints, bapiClis *types.BapiClisT) (suffix string) {
	suffixList := []string{}
	if endpoints.Tool != "" {
		suffix = suffix + endpoints.Tool + "/"
	}
	if endpoints.ID != "" {
		suffixList = append(suffixList, `biotoolsID="`+endpoints.ID+`"`)
	}
	if endpoints.Name != "" {
		suffixList = append(suffixList, `name=`+endpoints.ID)
	}
	if endpoints.Topic != "" {
		suffixList = append(suffixList, `Topic="`+endpoints.Topic+`"`)
	}
	if endpoints.DataType != "" {
		suffixList = append(suffixList, `dataType="`+endpoints.DataType+`"`)
	}
	if endpoints.DataFormat != "" {
		suffixList = append(suffixList, `dataFormat="`+endpoints.DataFormat+`"`)
	}
	if endpoints.OutputFormat != "" {
		suffixList = append(suffixList, `outputFormat="`+endpoints.OutputFormat+`"`)
	}
	if endpoints.Publication != "" {
		suffixList = append(suffixList, `publication=`+endpoints.Publication)
	}
	if bapiClis.Query != "" {
		suffixList = append(suffixList, "q="+bapiClis.Query)
	}
	if bapiClis.Format != "" {
		suffixList = append(suffixList, "format="+bapiClis.Format)
	}
	if bapiClis.Size != -1 {
		suffixList = append(suffixList, "page="+strconv.Itoa(bapiClis.From+bapiClis.Size))
	}
	if bapiClis.Extra != "" {
		suffixList = append(suffixList, bapiClis.Extra)
	}
	if len(suffixList) > 0 {
		suffix = suffix + "?" + strings.Join(suffixList, "&")
	}
	return suffix
}