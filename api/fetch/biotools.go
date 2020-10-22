package fetch

import (
	"io"
	"strconv"
	"strings"

	"github.com/clindet/bget/api/types"
)

const BioToolsHost = "https://bio.tools/api/tool/"

// BioTools access https://clinicaltrials.gov API
func BioTools(endpoints *types.BioToolsEndpoints, BapiClis *types.BapiClisT, f func(), of io.Writer) bool {
	setLog(BapiClis)
	if BapiClis.Format == "" {
		BapiClis.Format = "json"
	}
	netopt := setNetOpt(BapiClis)
	url := BioToolsHost + setBioToolsQuerySuffix(endpoints, BapiClis)
	if url == BioToolsHost || url == BioToolsHost+"?format=json" {
		return false
	}
	f()
	GetReq("bio.tools", url, BapiClis, netopt, of)

	return true
}

func setBioToolsQuerySuffix(endpoints *types.BioToolsEndpoints, BapiClis *types.BapiClisT) (suffix string) {
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
	if BapiClis.From != -1 && BapiClis.From > 10 {
		suffixList = append(suffixList, `page=`+strconv.Itoa(1+BapiClis.From/10))
	}
	if BapiClis.Query != "" {
		suffixList = append(suffixList, "q="+BapiClis.Query)
	}
	if BapiClis.Format != "" {
		suffixList = append(suffixList, "format="+BapiClis.Format)
	}
	if BapiClis.Extra != "" {
		suffixList = append(suffixList, BapiClis.Extra)
	}
	if len(suffixList) > 0 {
		suffix = suffix + "?" + strings.Join(suffixList, "&")
	}
	return suffix
}
