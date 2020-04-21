package fetch

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/PuerkitoBio/goquery"
	xj "github.com/basgys/goxml2json"
	"github.com/openanno/bget/api/types"
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

func PostReq(siteName, url string, data []byte, bapiClis *types.BapiClisT, netopt *cnet.Params, of io.Writer) {
	method := "POST"
	client := cnet.NewHTTPClient(bapiClis.Timeout, bapiClis.Proxy)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	cnet.SetDefaultReqHeader(req)
	if err != nil {
		log.Warn(err)
	}
	log.Infof("Quering %s API: %s.", siteName, url)
	log.Infof("Submitting data to %s: %v", url, string(data))
	reqWithOutput(client, req, bapiClis, netopt, of)
}

func GetReq(siteName, url string, bapiClis *types.BapiClisT, netopt *cnet.Params, of io.Writer) {
	method := "GET"
	client := cnet.NewHTTPClient(bapiClis.Timeout, bapiClis.Proxy)
	req, err := http.NewRequest(method, url, nil)
	cnet.SetDefaultReqHeader(req)
	if err != nil {
		log.Warn(err)
	}
	log.Infof("Quering %s API: %s.", siteName, url)
	reqWithOutput(client, req, bapiClis, netopt, of)
}

func reqWithOutput(client *http.Client, req *http.Request, bapiClis *types.BapiClisT,
	netopt *cnet.Params, of io.Writer) {
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
	if of == nil {
		of = cio.NewOutStream(bapiClis.Outfn, req.URL.String())
	}
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

	return
}

func queryExtract(client *http.Client, req *http.Request, bapiClis *types.BapiClisT,
	netopt *cnet.Params, f func(*types.BapiClisT, *http.Request, *goquery.Document, *bytes.Buffer, io.Writer),
	of io.Writer) {
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
	buf, _ := ioutil.ReadAll(resp.Body)
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(buf))
	if err != nil {
		log.Warnln(err)
		return
	}
	f(bapiClis, req, doc, bytes.NewBuffer(buf), of)
	defer resp.Body.Close()
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

func xml2json(buf *bytes.Buffer, PrettyJSON bool, indentNum int, sortKey bool) {
	json, _ := xj.Convert(buf)
	if PrettyJSON {
		indent := ""
		for i := 0; i < indentNum; i++ {
			indent = indent + " "
		}
		opt := pretty.Options{
			Indent:   indent,
			SortKeys: sortKey,
		}
		*buf = *bytes.NewBuffer(pretty.PrettyOptions(json.Bytes(), &opt))
	} else {
		*buf = *bytes.NewBuffer(json.Bytes())
	}
}
