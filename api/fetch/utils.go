package fetch

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/openbiox/bget/api/types"
	cio "github.com/openbiox/ligo/io"
	clog "github.com/openbiox/ligo/log"
	cnet "github.com/openbiox/ligo/net"
	"github.com/tidwall/pretty"
)

var log = clog.Logger
var logPrefix string
var logCon io.Writer

func setNetOpt(bapiClis *types.BapiClisT) *cnet.Params {
	var netopt = &cnet.Params{
		Pbar: pg,
	}
	netopt.Quiet = bapiClis.Verbose == 0
	netopt.Retries = bapiClis.Retries
	netopt.Timeout = bapiClis.Timeout
	netopt.RetSleepTime = bapiClis.RetSleepTime
	netopt.Proxy = bapiClis.Proxy
	return netopt
}

func PostReq(siteName, url string, data []byte, bapiClis *types.BapiClisT, netopt *cnet.Params) {
	method := "POST"
	client := cnet.NewHTTPClient(bapiClis.Timeout, bapiClis.Proxy)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	cnet.SetDefaultReqHeader(req)
	if err != nil {
		log.Warn(err)
	}
	log.Infof("Query %s API: %s.", siteName, url)
	log.Infof("POST Data:\n%v", string(data))
	query(client, req, bapiClis, netopt)
}

func GetReq(siteName, url string, bapiClis *types.BapiClisT, netopt *cnet.Params) {
	method := "GET"
	client := cnet.NewHTTPClient(bapiClis.Timeout, bapiClis.Proxy)
	req, err := http.NewRequest(method, url, nil)
	cnet.SetDefaultReqHeader(req)
	if err != nil {
		log.Warn(err)
	}
	log.Infof("Query %s API: %s.", siteName, url)
	query(client, req, bapiClis, netopt)
}

func query(client *http.Client, req *http.Request, bapiClis *types.BapiClisT, netopt *cnet.Params) {
	resp, err := cnet.RetriesClient(client, req, netopt)
	if err != nil {
		log.Warnln(err)
		return
	}
	if resp != nil {
		defer resp.Body.Close()
	} else {
		return
	}
	of := cio.NewOutStream(bapiClis.Outfn, req.URL.String())
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Warnln(err)
		return
	}
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

func setLog(BapiClis *types.BapiClisT) {
	if BapiClis.SaveLog {
		if BapiClis.LogDir == "" {
			BapiClis.LogDir = filepath.Join(os.TempDir(), "_log")
		}
		logPrefix = fmt.Sprintf("%s/%s", BapiClis.LogDir, BapiClis.TaskID)
		cio.CreateDir(BapiClis.LogDir)
		logCon, _ = cio.Open(logPrefix + ".log")
	}
	clog.SetLogStream(log, BapiClis.Verbose == 0, BapiClis.SaveLog, &logCon)
}
