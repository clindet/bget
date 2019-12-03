package fetch

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/openbiox/bget/bapi/types"
	cio "github.com/openbiox/butils/io"
	"github.com/openbiox/butils/log"
	cnet "github.com/openbiox/butils/net"
	"github.com/tidwall/pretty"
)

const Dataset2toolsHost = "http://amp.pharm.mssm.edu/datasets2tools/api/search?"

// Dataset2tools access http://amp.pharm.mssm.edu/datasets2tools/ API
func Dataset2tools(endpoints *types.Datasets2toolsEndpoints, bapiClis *types.BapiClisT) {
	var netopt = &cnet.BnetParams{
		Pbar: pg,
	}
	netopt.Quiet = bapiClis.Quiet
	netopt.Retries = bapiClis.Retries
	netopt.Timeout = bapiClis.Timeout
	netopt.RetSleepTime = bapiClis.RetSleepTime
	netopt.Proxy = bapiClis.Proxy
	url := Dataset2toolsHost + setDatasets2toolsQuerySuffix(endpoints)
	client := cnet.NewHTTPClient(bapiClis.Timeout, bapiClis.Proxy)
	method := "GET"
	req, err := http.NewRequest(method, url, nil)
	cnet.SetDefaultReqHeader(req)
	if err != nil {
		log.Warn(err)
	}
	log.Infof("Query datasets2tools API: %s.", url)
	resp, err := cnet.RetriesClient(client, req, netopt)
	if err != nil {
		return
	}
	if resp != nil {
		defer resp.Body.Close()
	} else {
		return
	}
	of := cio.NewOutStream(bapiClis.Outfn, req.URL.String())
	buf, _ := ioutil.ReadAll(resp.Body)
	if bapiClis.PrettyJSON {
		indent := ""
		for i := 0; i < bapiClis.Indent; i++ {
			indent = indent + " "
		}
		opt := pretty.Options{
			Indent:   indent,
			SortKeys: bapiClis.SortKeys,
		}
		buf = pretty.PrettyOptions(buf, &opt)
	}
	_, err = io.Copy(of, bytes.NewBuffer(buf))
	if err != nil {
		log.Warn(err)
	}
	defer resp.Body.Close()
	defer of.Close()

	return
}

func setDatasets2toolsQuerySuffix(endpoints *types.Datasets2toolsEndpoints) (suffix string) {
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
	if len(suffixList) > 0 {
		suffix = strings.Join(suffixList, "&")
	}
	return suffix
}
