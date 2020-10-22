package fetch

import (
	"bytes"
	"io"
	"os"
	"time"

	"github.com/biogo/ncbi"
	"github.com/biogo/ncbi/entrez"
	"github.com/clindet/bget/api/types"
	cnet "github.com/openbiox/ligo/net"
)

// Ncbi modified from https://github.com/biogo/ncbi BSD license
func Ncbi(bapiClis *types.BapiClisT, ncbiClis *types.NcbiClisT, of io.Writer) {
	setLog(bapiClis)
	ncbi.SetTimeout(time.Duration(bapiClis.Timeout) * time.Second)
	tool := "entrez.example"
	h := entrez.History{}
	parms := entrez.Parameters{
		APIKey: "193124979d2e7f360c150dadc5b1e3bfec09",
	}
	var t int
	var s *entrez.Search
	var err error
	for t = 0; t < bapiClis.Retries+1; t++ {
		s, err = entrez.DoSearch(ncbiClis.NcbiDB, bapiClis.Query, &parms, &h, tool, bapiClis.Email)
		if err != nil {
			log.Warnf("Failed to retrieve on attempt %d... error: %v ... retrying after %d seconds.", t+1, err, bapiClis.RetSleepTime)
			time.Sleep(time.Duration(bapiClis.RetSleepTime) * time.Second)
			continue
		}
		break
	}
	log.Infof("Available retrieve %d records.", s.Count)
	if s.Count == 0 {
		return
	}
	from, end := cnet.SetQueryFromEnd(bapiClis.From, bapiClis.Size, s.Count)
	if from == 0 {
		from = 1
		end = end + 1
	}
	if bapiClis.Size != -1 {
		end = end + 1
	}
	log.Infof("Will retrieve %d records, from %d to %d.", end-from, from, end-1)

	var (
		buf   = &bytes.Buffer{}
		p     = &entrez.Parameters{RetMax: ncbiClis.NcbiRetmax, RetType: bapiClis.Format, RetMode: "text"}
		bn, n int64
	)
	if p.RetMax > end-from {
		p.RetMax = end - from
	}
	for p.RetStart = from - 1; p.RetStart < end-1; p.RetStart += p.RetMax {
		log.Infof("Attempting to retrieve %d records: %d-%d with %d retries.", p.RetMax, p.RetStart+1, p.RetMax+p.RetStart, bapiClis.Retries)
		var t int
		for t = 0; t < bapiClis.Retries+1; t++ {
			buf.Reset()
			var (
				r   io.ReadCloser
				_bn int64
			)
			r, err = entrez.Fetch(ncbiClis.NcbiDB, p, tool, bapiClis.Email, &h)
			if err != nil {
				if r != nil {
					r.Close()
				}
				if bapiClis.Retries != 0 {
					log.Warnf("Failed to retrieve on attempt %d... error: %v ... retrying after %d seconds.", t+1, err, bapiClis.RetSleepTime)
					time.Sleep(time.Duration(bapiClis.RetSleepTime) * time.Second)
				} else {
					log.Warnf("Failed to retrieve on attempt %d... error: %v ...", t+1, err)
				}
				continue
			}
			_bn, err = io.Copy(buf, r)
			r.Close()
			if err == nil {
				if bapiClis.XML2json {
					xml2json(buf, bapiClis.PrettyJSON, bapiClis.Indent, bapiClis.SortKeys)
					bn += int64(buf.Len())
				} else {
					bn += _bn
				}
				break
			}
			log.Warnf("Failed to buffer on attempt %d... error: %v ... retrying after %d seconds.", t+1, err, bapiClis.RetSleepTime)
			time.Sleep(time.Duration(bapiClis.RetSleepTime) * time.Second)
		}
		if err != nil {
			os.Exit(1)
		}

		log.Infof("Retrieved records with %d retries... writing out.", t+1)
		_n, err := io.Copy(of, buf)
		n += _n
		if err != nil {
			log.Fatalf("Error: %v\n", err)
		}
	}
	if bn != n {
		log.Warnf("Writethrough mismatch: %d != %d", bn, n)
	}
}
