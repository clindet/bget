package fetch

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/clindet/bget/api/types"
	cio "github.com/openbiox/ligo/io"
	cnet "github.com/openbiox/ligo/net"
	"github.com/openbiox/ligo/stringo"
)

const CrossRefHost = "https://www.crossref.org/guestquery"

// CrossRef access https://clinicaltrials.gov API
func CrossRef(endpoints *types.CrossRefEndpoints, bapiClis *types.BapiClisT, f func(), of io.Writer) bool {
	setLog(bapiClis)
	if bapiClis.Format == "" {
		bapiClis.Format = "json"
	}
	if endpoints.Doi.Doi == "" {
		return false
	}
	netopt := setNetOpt(bapiClis)
	f()
	data, bdr := setCrossRefQueryData(endpoints, bapiClis)
	method := "POST"
	client := cnet.NewHTTPClient(bapiClis.Timeout, bapiClis.Proxy)
	req, err := http.NewRequest(method, CrossRefHost,
		&data)
	//req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Content-Type", "multipart/form-data; boundary="+bdr)
	req.Header.Add("Content-Length", strconv.Itoa(len(string(data.Bytes()))))

	if err != nil {
		log.Warn(err)
	}
	log.Infof("Quering %s API: %s.", "Crossref", CrossRefHost)
	log.Infof("Submitting data to %s: %v", CrossRefHost, string(data.Bytes()))
	if of == nil {
		of = cio.NewOutStream(bapiClis.Outfn, req.URL.String())
	}
	queryExtract(client, req, bapiClis, netopt, extractCrossRef, of)
	return true
}

func extractCrossRef(bapiClis *types.BapiClisT, req *http.Request, doc *goquery.Document,
	buf *bytes.Buffer, of io.Writer) {
	doc.Find("table[name=doiresult] textarea").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		html := s.Text()
		html = strings.ReplaceAll(html, "jats:abstract", "abstract")
		html = strings.ReplaceAll(html, "<jats:p>", "")
		html = strings.ReplaceAll(html, "</jats:p>", "")
		html = stringo.StrReplaceAll(html, "/ *\n *<jats:italic> *", "/")
		html = stringo.StrReplaceAll(html, "\n *<jats:italic> *", " ")
		html = stringo.StrReplaceAll(html, "</jats:italic> *\n */ *", "/")
		html = stringo.StrReplaceAll(html, "</jats:italic> *\n *–", "–")
		html = stringo.StrReplaceAll(html, "</jats:italic> *\n *", " ")
		html = stringo.StrReplaceAll(html, "[ \n]*</abstract>", "</abstract>")
		html = stringo.StrReplaceAll(html, `<abstract xmlns:jats=".*">[ \n]*`, `<abstract>`)
		//html = strings.ReplaceAll(html, "<jats:italic>", "")
		//html = strings.ReplaceAll(html, "</jats:italic>", "")

		buf = bytes.NewBufferString(html)
		if bapiClis.XML2json {
			xml2json(buf, bapiClis.PrettyJSON, bapiClis.Indent, bapiClis.SortKeys)
		}
		_, err := io.Copy(of, buf)
		if err != nil {
			log.Warn(err)
		}
	})
}

func setCrossRefQueryData(endpoints *types.CrossRefEndpoints, bapiClis *types.BapiClisT) (bytes.Buffer, string) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	if endpoints.Doi.Doi != "" {
		fw, _ := w.CreateFormField("queryType")
		fw.Write([]byte("doi"))
		fw, _ = w.CreateFormField("resType")
		fw.Write([]byte("unixref"))
		fw, _ = w.CreateFormField("doi")
		fw.Write([]byte(endpoints.Doi.Doi))
		fw, _ = w.CreateFormField("doi_search")
		fw.Write([]byte("Search"))
	}
	w.Close()
	return b, w.Boundary()
}
