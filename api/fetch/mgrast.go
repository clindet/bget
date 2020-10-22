package fetch

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/openbiox/ligo/stringo"

	"github.com/clindet/bget/api/types"
)

const MgRastHost = "http://api.mg-rast.org/"

// MgRast access https://clinicaltrials.gov API
func MgRast(endpoints *types.MgRastEndpoints, bapiClis *types.BapiClisT, f func(), of io.Writer) bool {
	setLog(bapiClis)
	netopt := setNetOpt(bapiClis)
	cvrtMgRastList(endpoints)
	url := MgRastHost + setMgRastQuerySuffix(endpoints, bapiClis)
	if emptyURL(url, endpoints) {
		return false
	}
	f()
	if mgRastGetMode(endpoints) {
		GetReq("api.mg-rast.org", url, bapiClis, netopt, of)
	} else {
		postData := setMgRastQueryData(endpoints, bapiClis)
		PostReq("api.mg-rast.org", url, postData, bapiClis, netopt, of)
	}
	return true
}

func mgRastGetMode(endpoints *types.MgRastEndpoints) bool {
	return (len(endpoints.ParamsAnno.Md5s) == 0 && endpoints.Annotation) || (len(endpoints.ParamsCompute.Data) == 0 && len(endpoints.ParamsCompute.Rows) == 0 && len(endpoints.ParamsCompute.Columns) == 0 && endpoints.Compute) || endpoints.DarkMatter || endpoints.Download || endpoints.Project != "" || endpoints.Library != "" || endpoints.Sample != ""
}

func emptyURL(url string, endpoints *types.MgRastEndpoints) bool {
	if (endpoints.Annotation || endpoints.Compute || endpoints.DarkMatter || endpoints.Download) && (endpoints.ID != "" || endpoints.Sequence != "" || endpoints.Info || len(endpoints.ParamsCompute.Columns) > 0 || len(endpoints.ParamsCompute.Rows) > 0 || len(endpoints.ParamsCompute.Data) > 0) || endpoints.Project != "" || endpoints.Library != "" || endpoints.Sample != "" {
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

func setMgRastQueryData(endpoints *types.MgRastEndpoints, bapiClis *types.BapiClisT) (data []byte) {
	if endpoints.Annotation {
		data, _ = json.MarshalIndent(endpoints.ParamsAnno, "", " ")
	} else if endpoints.Compute {
		data, _ = json.MarshalIndent(endpoints.ParamsCompute, "", " ")
	}
	return data
}

func setMgRastQuerySuffix(endpoints *types.MgRastEndpoints, bapiClis *types.BapiClisT) (suffix string) {
	suffixList := []string{}
	if endpoints.Annotation {
		suffix, suffixList = setAnnoSuffix(endpoints, bapiClis)
	} else if endpoints.Compute {
		suffix, suffixList = setComputeSuffix(endpoints, bapiClis)
	} else if endpoints.Download {
		suffix, suffixList = setDownloadSuffix(endpoints, bapiClis)
	} else if endpoints.Project != "" || endpoints.Library != "" || endpoints.Sample != "" {
		suffix, suffixList = setProjOrLibrarySuffix(endpoints, bapiClis)
	}
	suffixOther, suffixListOther := setOtherSuffix(endpoints, bapiClis)
	suffix = suffix + suffixOther
	suffixList = append(suffixList, suffixListOther...)
	if len(suffixList) > 0 {
		suffix = suffix + "?" + strings.Join(suffixList, "&")
	}
	return suffix
}

func setAnnoSuffix(endpoints *types.MgRastEndpoints, bapiClis *types.BapiClisT) (suffix string, suffixList []string) {
	suffix = suffix + "annotation/"
	if endpoints.Similarity {
		suffix = suffix + "similarity/" + endpoints.Sequence
	} else if endpoints.Sequence != "" {
		suffix = suffix + "sequence/" + endpoints.Sequence
	}
	if mgRastGetMode(endpoints) {
		suffixList = struct2suffixList(endpoints.ParamsAnno)
	}
	return suffix, suffixList
}

func setComputeSuffix(endpoints *types.MgRastEndpoints, bapiClis *types.BapiClisT) (suffix string, suffixList []string) {
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
		suffixList = struct2suffixList(endpoints.ParamsCompute)
	}
	return suffix, suffixList
}

func setDownloadSuffix(endpoints *types.MgRastEndpoints, bapiClis *types.BapiClisT) (suffix string, suffixList []string) {
	if endpoints.Download && endpoints.DownloadHistory {
		suffix = suffix + "download/history/" + endpoints.ID
	} else if endpoints.Download {
		suffix = suffix + "download/" + endpoints.ID
	}
	if mgRastGetMode(endpoints) {
		suffixList = struct2suffixList(endpoints.ParamsDownload)
	}
	return suffix, suffixList
}

func setProjOrLibrarySuffix(endpoints *types.MgRastEndpoints, bapiClis *types.BapiClisT) (suffix string, suffixList []string) {
	if endpoints.Project == "nil" {
		suffix = suffix + "project/"
	} else if endpoints.Project != "" && endpoints.Project != "nil" {
		suffix = suffix + "project/" + endpoints.Project
	} else if endpoints.Library == "nil" {
		suffix = suffix + "library/"
	} else if endpoints.Library != "" && endpoints.Library != "nil" {
		suffix = suffix + "library/" + endpoints.Library
	} else if endpoints.Sample == "nil" {
		suffix = suffix + "sample/"
	} else if endpoints.Sample != "" && endpoints.Sample != "nil" {
		suffix = suffix + "sample/" + endpoints.Sample
	}
	if mgRastGetMode(endpoints) {
		suffixList = struct2suffixList(endpoints.ParamsProjOrLibrary)
	}
	return suffix, suffixList
}

func struct2suffixList(dat interface{}) (suffixList []string) {
	mgrast, _ := json.MarshalIndent(dat, "", " ")
	var mgrastMap map[string]interface{}
	json.Unmarshal(mgrast, &mgrastMap)
	for k, v := range mgrastMap {
		if v != "" && v != false && v != nil {
			suffixList = append(suffixList, k+"="+fmt.Sprintf("%v", v))
		}
	}
	return suffixList
}

func setOtherSuffix(endpoints *types.MgRastEndpoints, bapiClis *types.BapiClisT) (suffix string, suffixList []string) {
	if endpoints.DarkMatter {
		suffix = suffix + "darkmatter/" + endpoints.ID
	}
	if endpoints.Inbox {
		suffix = suffix + "inbox/"
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
	if bapiClis.Format != "" {
		suffixList = append(suffixList, "format="+bapiClis.Format)
	}
	if bapiClis.Extra != "" {
		suffixList = append(suffixList, bapiClis.Extra)
	}
	return suffix, suffixList
}
