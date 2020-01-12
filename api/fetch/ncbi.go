package fetch

import (
	"bytes"
	"fmt"
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
func Ncbi(bapiClis *types.BapiClisT, ncbiClis *types.NcbiClisT) {
	SetLogStream(bapiClis.Quiet == "true", bapiClis.SaveLog == "true", fmt.Sprintf("%s/%s.log", bapiClis.LogDir, bapiClis.TaskID))
	ncbi.SetTimeout(time.Duration(bapiClis.Timeout) * time.Second)
	tool := "entrez.example"
	h := entrez.History{}
	parms := entrez.Parameters{
		APIKey: "193124979d2e7f360c150dadc5b1e3bfec09",
	}
	s, err := entrez.DoSearch(ncbiClis.NcbiDB, bapiClis.Query, &parms, &h, tool, bapiClis.Email)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	log.Infof("Available retrieve %d records.", s.Count)
	from, end := cnet.SetQueryFromEnd(bapiClis.From, bapiClis.Size, s.Count)
	log.Infof("Will retrieve %d records, from %d to %d.", end-from, from+1, end)

	of := cio.NewOutStream(bapiClis.Outfn, "")
	defer of.Close()
	var (
		buf   = &bytes.Buffer{}
		p     = &entrez.Parameters{RetMax: ncbiClis.NcbiRetmax, RetType: bapiClis.Format, RetMode: "text"}
		bn, n int64
	)
	if p.RetMax > end-from {
		p.RetMax = end - from
	}
	for p.RetStart = from; p.RetStart < end; p.RetStart += p.RetMax {
		log.Infof("Attempting to retrieve %d records: %d-%d with %d retries.", p.RetMax, p.RetStart+1, p.RetMax+p.RetStart, bapiClis.Retries)
		var t int
		for t = 0; t < bapiClis.Retries; t++ {
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
				log.Warnf("Failed to retrieve on attempt %d... error: %v ... retrying after %d seconds.", t+1, err, bapiClis.RetSleepTime)
				time.Sleep(time.Duration(bapiClis.RetSleepTime) * time.Second)
				continue
			}
			_bn, err = io.Copy(buf, r)
			r.Close()
			if err == nil {
				bn += _bn
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
