package fetch

import (
	"io"
	"strconv"
	"strings"

	"github.com/clindet/bget/api/types"
)

const Dataset2toolsHost = "http://amp.pharm.mssm.edu/datasets2tools/api/search?"

// Dataset2tools access http://amp.pharm.mssm.edu/datasets2tools/ API
func Dataset2tools(endpoints *types.Datasets2toolsEndpoints, BapiClis *types.BapiClisT, of io.Writer) {
	setLog(BapiClis)
	netopt := setNetOpt(BapiClis)
	url := Dataset2toolsHost + setDatasets2toolsQuerySuffix(endpoints, BapiClis)
	GetReq("datasets2tools", url, BapiClis, netopt, of)
	return
}

func setDatasets2toolsQuerySuffix(endpoints *types.Datasets2toolsEndpoints, BapiClis *types.BapiClisT) (suffix string) {
	suffixList := []string{}
	if endpoints.ObjectType != "" {
		suffixList = append(suffixList, "object_type="+endpoints.ObjectType)
	}
	if endpoints.DatasetAccession != "" {
		suffixList = append(suffixList, "dataset_accession="+endpoints.DatasetAccession)
	}
	if endpoints.CannedAnalysisAccession != "" {
		suffixList = append(suffixList, "canned_analysis_accession="+endpoints.CannedAnalysisAccession)
	}
	if endpoints.Query != "" {
		suffixList = append(suffixList, "q="+endpoints.Query)
	}
	if endpoints.ToolName != "" {
		suffixList = append(suffixList, "tool_name="+endpoints.ToolName)
	}
	if endpoints.DiseaseName != "" {
		suffixList = append(suffixList, "disease_name="+endpoints.DiseaseName)
	}
	if endpoints.Gneset != "" {
		suffixList = append(suffixList, "geneset="+endpoints.Gneset)
	}
	if endpoints.PageSize != -1 {
		suffixList = append(suffixList, "page_size="+strconv.Itoa(endpoints.PageSize))
	}
	if BapiClis.Extra != "" {
		suffixList = append(suffixList, BapiClis.Extra)
	}
	if len(suffixList) > 0 {
		suffix = strings.Join(suffixList, "&")
	}
	return suffix
}
