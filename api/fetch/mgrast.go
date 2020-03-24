package fetch

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/openbiox/ligo/stringo"

	"github.com/openbiox/bget/api/types"
)

const MgRastHost = "http://api.mg-rast.org/"

// MgRast access https://clinicaltrials.gov API
func MgRast(endpoints *types.MgRastEndpoints, BapiClis *types.BapiClisT, f func()) bool {
	setLog(BapiClis)
	netopt := setNetOpt(BapiClis)
	cvrtMgRastList(endpoints)
	url := MgRastHost + setMgRastQuerySuffix(endpoints, BapiClis)
	if emptyURL(url, endpoints) {
		return false
	}
	f()
	if mgRastGetMode(endpoints) {
		GetReq("api.mg-rast.org", url, BapiClis, netopt)
	} else {
		postData := setMgRastQueryData(endpoints, BapiClis)
		PostReq("api.mg-rast.org", url, postData, BapiClis, netopt)
	}
	return true
}

func mgRastGetMode(endpoints *types.MgRastEndpoints) bool {
	return (len(endpoints.ParamsAnno.Md5s) == 0 && endpoints.Annotation) || (len(endpoints.ParamsCompute.Data) == 0 && len(endpoints.ParamsCompute.Rows) == 0 && len(endpoints.ParamsCompute.Columns) == 0 && endpoints.Compute) || endpoints.DarkMatter || endpoints.Download
}

func emptyURL(url string, endpoints *types.MgRastEndpoints) bool {
	if (endpoints.Annotation || endpoints.Compute || endpoints.DarkMatter || endpoints.Download) && (endpoints.ID != "" || endpoints.Sequence != "" || endpoints.Info || len(endpoints.ParamsCompute.Columns) > 0 || len(endpoints.ParamsCompute.Rows) > 0 || len(endpoints.ParamsCompute.Data) > 0) {
		return false
	}
	return true
}

func cvrtMgRastList(endpoints *types.MgRastEndpoints) {
	if endpoints.Md5s != "" {
		endpoints.ParamsAnno.Md5s = stringo.StrSplit(endpoints.Md5s, ",", 10000000)
	}
	if endpoints.Rows != "" {
		endpoints.ParamsCompute.Rows = stringo.StrSplit(endpoints.Rows, ",", 10000000)
	}
	if endpoints.Columns != "" {
		endpoints.ParamsCompute.Columns = stringo.StrSplit(endpoints.Columns, ",", 10000000)
	}
	if endpoints.Data != "" {
		dataSlice := stringo.StrSplit(endpoints.Data, "[]],[[]", 10000000)
		for _, v := range dataSlice {
			v = stringo.StrReplaceAll(v, "[[]|[]]", "")
			tmp := stringo.StrSplit(v, ",", 10000000)
			tmp2 := []float64{}
			for _, v2 := range tmp {
				v3, _ := strconv.ParseFloat(v2, 64)
				tmp2 = append(tmp2, v3)
			}
			endpoints.ParamsCompute.Data = append(endpoints.ParamsCompute.Data, tmp2)
		}
	}
}

func setMgRastQueryData(endpoints *types.MgRastEndpoints, BapiClis *types.BapiClisT) (data []byte) {
	if endpoints.Annotation {
		data, _ = json.MarshalIndent(endpoints.ParamsAnno, "", " ")
	} else if endpoints.Compute {
		data, _ = json.MarshalIndent(endpoints.ParamsCompute, "", " ")
	}
	return data
}

func setMgRastQuerySuffix(endpoints *types.MgRastEndpoints, BapiClis *types.BapiClisT) (suffix string) {
	suffixList := []string{}
	if endpoints.Annotation {
		suffix, suffixList = setAnnoSuffix(endpoints, BapiClis)
	} else if endpoints.Compute {
		suffix, suffixList = setComputeSuffix(endpoints, BapiClis)
	} else if endpoints.Download {
		suffix, suffixList = setDownloadSuffix(endpoints, BapiClis)
	}
	suffixOther, suffixListOther := setOtherSuffix(endpoints, BapiClis)
	suffix = suffix + suffixOther
	suffixList = append(suffixList, suffixListOther...)
	if len(suffixList) > 0 {
		suffix = suffix + "?" + strings.Join(suffixList, "&")
	}
	return suffix
}

func setAnnoSuffix(endpoints *types.MgRastEndpoints, BapiClis *types.BapiClisT) (suffix string, suffixList []string) {
	suffix = suffix + "annotation/"
	if endpoints.Similarity {
		suffix = suffix + "similarity/" + endpoints.Sequence
	} else if endpoints.Sequence != "" {
		suffix = suffix + "sequence/" + endpoints.Sequence
	}
	if mgRastGetMode(endpoints) {
		mgrastAnno, _ := json.MarshalIndent(endpoints.ParamsAnno, "", " ")
		var mgrastAnnoMap map[string]interface{}
		json.Unmarshal(mgrastAnno, &mgrastAnnoMap)
		for k, v := range mgrastAnnoMap {
			if k == "md5s" {
				continue
			}
			if v != "" {
				suffixList = append(suffixList, k+"="+fmt.Sprintf("%v", v))
			}
		}
	}
	return suffix, suffixList
}

func setComputeSuffix(endpoints *types.MgRastEndpoints, BapiClis *types.BapiClisT) (suffix string, suffixList []string) {
	suffix = suffix + "compute/"
	if endpoints.ComputeAlphadiversity {
		suffix = suffix + "alphadiversity/" + endpoints.Sequence
	} else if endpoints.ComputeRarefaction {
		suffix = suffix + "rarefaction/" + endpoints.Sequence
	} else if endpoints.ComputeBlast {
		suffix = suffix + "blast/" + endpoints.Sequence
	} else if endpoints.ComputeNormalize {
		suffix = suffix + "normalize" + endpoints.Sequence
	} else if endpoints.ComputeDistance {
		suffix = suffix + "distance" + endpoints.Sequence
	} else if endpoints.ComputeHeatmap {
		suffix = suffix + "heatmap" + endpoints.Sequence
	} else if endpoints.ComputePcoa {
		suffix = suffix + "pcoa" + endpoints.Sequence
	}
	if mgRastGetMode(endpoints) {
		mgrastCompute, _ := json.MarshalIndent(endpoints.ParamsCompute, "", " ")
		var mgrastComputeMap map[string]interface{}
		json.Unmarshal(mgrastCompute, &mgrastComputeMap)
		for k, v := range mgrastComputeMap {
			if k == "columns" || k == "data" || k == "rows" || k == "seq_num" && v == 0 {
				continue
			}
			if v != "" {
				suffixList = append(suffixList, k+"="+fmt.Sprintf("%v", v))
			}
		}
	}
	return suffix, suffixList
}

func setDownloadSuffix(endpoints *types.MgRastEndpoints, BapiClis *types.BapiClisT) (suffix string, suffixList []string) {
	if endpoints.Download && endpoints.DownloadHistory {
		suffix = suffix + "download/history/" + endpoints.ID
	} else if endpoints.Download {
		suffix = suffix + "download/" + endpoints.ID
	}
	if mgRastGetMode(endpoints) {
		mgrastDownload, _ := json.MarshalIndent(endpoints.ParamsDownload, "", " ")
		var mgrastDownloadMap map[string]interface{}
		json.Unmarshal(mgrastDownload, &mgrastDownloadMap)
		for k, v := range mgrastDownloadMap {
			if v != "" && v != false {
				suffixList = append(suffixList, k+"="+fmt.Sprintf("%v", v))
			}
		}
	}
	return suffix, suffixList
}

func setOtherSuffix(endpoints *types.MgRastEndpoints, BapiClis *types.BapiClisT) (suffix string, suffixList []string) {
	if endpoints.DarkMatter {
		suffix = suffix + "darkmatter/" + endpoints.ID
	}
	if endpoints.Inbox {
		suffix = suffix + "inbox/"
	}
	if endpoints.Library {
		suffix = suffix + "library/"
	}
	if endpoints.M5nr {
		suffix = suffix + "m5nr/"
	}
	if endpoints.Matrix {
		suffix = suffix + "matrix/"
	}
	if endpoints.MetaData {
		suffix = suffix + "metadata/"
	}
	if endpoints.MetaGenome {
		suffix = suffix + "metagenome/"
	}
	if endpoints.Mixs {
		suffix = suffix + "mixs/"
	}
	if endpoints.Profile {
		suffix = suffix + "profile/"
	}
	if endpoints.Project {
		suffix = suffix + "project/"
	}
	if endpoints.ResearchObject {
		suffix = suffix + "researchobject/"
	}
	if endpoints.Sample {
		suffix = suffix + "sample/"
	}
	if endpoints.Search {
		suffix = suffix + "search/"
	}
	if endpoints.Submission {
		suffix = suffix + "submission/"
	}
	if endpoints.Validation {
		suffix = suffix + "validation/"
	}

	if endpoints.Info {
		suffix = suffix + "info/"
	}
	if endpoints.Auth != "" {
		suffixList = append(suffixList, `auth=`+endpoints.Auth)
	}
	if BapiClis.Format != "" {
		suffixList = append(suffixList, "format="+BapiClis.Format)
	}
	if BapiClis.Extra != "" {
		suffixList = append(suffixList, BapiClis.Extra)
	}
	return suffix, suffixList
}
