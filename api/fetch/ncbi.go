package fetch

import (
	"bytes"
	"io"
	"os"
	"time"

	"github.com/biogo/ncbi"
	"github.com/biogo/ncbi/entrez"
	"github.com/openbiox/bget/api/types"
	cio "github.com/openbiox/ligo/io"
	cnet "github.com/openbiox/ligo/net"
)

// Ncbi modified from https://github.com/biogo/ncbi BSD license
func Ncbi(BapiClis *types.BapiClisT, ncbiClis *types.NcbiClisT) {
	setLog(BapiClis)
	ncbi.SetTimeout(time.Duration(BapiClis.Timeout) * time.Second)
	tool := "entrez.example"
	h := entrez.History{}
	parms := entrez.Parameters{
		APIKey: "193124979d2e7f360c150dadc5b1e3bfec09",
	}
	s, err := entrez.DoSearch(ncbiClis.NcbiDB, BapiClis.Query, &parms, &h, tool, BapiClis.Email)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	log.Infof("Available retrieve %d records.", s.Count)
	if s.Count == 0 {
		return
	}
	from, end := cnet.SetQueryFromEnd(BapiClis.From, BapiClis.Size, s.Count)
	if from == 0 {
		from = 1
		end = end + 1
	}
	log.Infof("Will retrieve %d records, from %d to %d.", end-from, from, end-1)

	of := cio.NewOutStream(BapiClis.Outfn, "")
	defer of.Close()
	var (
		buf   = &bytes.Buffer{}
		p     = &entrez.Parameters{RetMax: ncbiClis.NcbiRetmax, RetType: BapiClis.Format, RetMode: "text"}
		bn, n int64
	)
	if p.RetMax > end-from {
		p.RetMax = end - from + 1
	}
	log.Infof("p.RetMax: %d", p.RetMax)
	for p.RetStart = from - 1; p.RetStart < end-1; p.RetStart += p.RetMax {
		log.Infof("Attempting to retrieve %d records: %d-%d with %d retries.", p.RetMax, p.RetStart, p.RetMax+p.RetStart, BapiClis.Retries)
		var t int
		for t = 0; t < BapiClis.Retries; t++ {
			buf.Reset()
			var (
				r   io.ReadCloser
				_bn int64
			)
			r, err = entrez.Fetch(ncbiClis.NcbiDB, p, tool, BapiClis.Email, &h)
			if err != nil {
				if r != nil {
					r.Close()
				}
				log.Warnf("Failed to retrieve on attempt %d... error: %v ... retrying after %d seconds.", t+1, err, BapiClis.RetSleepTime)
				time.Sleep(time.Duration(BapiClis.RetSleepTime) * time.Second)
				continue
			}
			_bn, err = io.Copy(buf, r)
			r.Close()
			if err == nil {
				bn += _bn
				break
			}
			log.Warnf("Failed to buffer on attempt %d... error: %v ... retrying after %d seconds.", t+1, err, BapiClis.RetSleepTime)
			time.Sleep(time.Duration(BapiClis.RetSleepTime) * time.Second)
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
