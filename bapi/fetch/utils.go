package fetch

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/openbiox/bget/bapi/types"
	cio "github.com/openbiox/butils/io"
	"github.com/openbiox/butils/log"
	cnet "github.com/openbiox/butils/net"
	"github.com/tidwall/pretty"
)

func setNetOpt(bapiClis *types.BapiClisT) *cnet.BnetParams {
	var netopt = &cnet.BnetParams{
		Pbar: pg,
	}
	netopt.Quiet = bapiClis.Quiet
	netopt.Retries = bapiClis.Retries
	netopt.Timeout = bapiClis.Timeout
	netopt.RetSleepTime = bapiClis.RetSleepTime
	netopt.Proxy = bapiClis.Proxy
	return netopt
}

func queryAPI(siteName, url string, bapiClis *types.BapiClisT, netopt *cnet.BnetParams) {
	method := "GET"
	client := cnet.NewHTTPClient(bapiClis.Timeout, bapiClis.Proxy)
	req, err := http.NewRequest(method, url, nil)
	cnet.SetDefaultReqHeader(req)
	if err != nil {
		log.Warn(err)
	}
	log.Infof("Query %s API: %s.", siteName, url)
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
