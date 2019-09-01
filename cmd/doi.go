package cmd

import (
	"fmt"
	"path"
	"strings"
	"time"

	"github.com/Miachol/bget/spider"
	cio "github.com/openbiox/butils/io"
	"github.com/openbiox/butils/log"
	cnet "github.com/openbiox/butils/net"
	"github.com/openbiox/butils/slice"
	stringo "github.com/openbiox/butils/stringo"
	"github.com/spf13/cobra"
)

var fullText bool
var suppl bool
var pmc bool

var doiCmd = &cobra.Command{
	Use:   "doi [doi1 doi2 doi3...]",
	Short: "Can be used to access files via DOI.",
	Long:  `Can be used to access files via DOI. More see here https://github.com/Miachol/bget.`,
	Run: func(cmd *cobra.Command, args []string) {
		doiCmdRunOptions(cmd)
	},
}

func downloadDoi() {
	sem := make(chan bool, bgetClis.thread)
	doi := []string{}
	urls := []string{}
	var destDirArray []string
	if bgetClis.doi != "" && strings.Contains(bgetClis.doi, bgetClis.separator) {
		doi = strings.Split(bgetClis.doi, bgetClis.separator)
	} else if bgetClis.doi != "" {
		doi = []string{bgetClis.doi}
	} else if bgetClis.listFile != "" {
		doi = cio.ReadLines(bgetClis.listFile)
	}
	for _, v := range doi {
		sem <- true
		go func(v string) {
			defer func() {
				<-sem
			}()
			urlsTmp := doiSpiders(v)
			urlsTmp = slice.DropSliceDup(urlsTmp)
			for range urlsTmp {
				destDirArray = append(destDirArray, path.Join(bgetClis.downloadDir, v))
			}
			urls = append(urls, urlsTmp...)
		}(v)
	}
	for i := 0; i < cap(sem); i++ {
		sem <- true
	}
	netOpt := setNetParams(&bgetClis)
	cnet.HttpGetURLs(urls, destDirArray, netOpt)
}

func doiSpiders(doi string) (urls []string) {
	if !strings.Contains(doi, "/") {
		return urls
	}
	if stringo.StrDetect(doi, "http[s]://doi.org/") {
		doi = stringo.StrReplaceAll(doi, "http[s]://doi.org/", "")
	}
	doiTmp := strings.Split(doi, "/")
	doiOrg := doiTmp[0]
	var t int
	opt := spider.DoiSpiderOpt{
		Doi:           doi,
		Proxy:         bgetClis.proxy,
		Timeout:       bgetClis.timeout,
		FullText:      fullText,
		Supplementary: suppl,
	}
	if pmc {
		urls = spider.PmcSpider(&opt)
		if len(urls) > 0 && !suppl {
			return urls
		}
	}
	if len(urls) > 0 {
		opt.FullText = false
	}
	for k := range spider.DoiSpidersPool {
		if k == doiOrg {
			for t = 0; t < bgetClis.retries; t++ {
				urls = append(urls, spider.DoiSpidersPool[doiOrg](&opt)...)
				if len(urls) == 0 {
					log.Warnf("%s returns empty, on attempt %d... retrying after %d seconds.", doi, t+1, bgetClis.retSleepTime)
					time.Sleep(time.Duration(bgetClis.retSleepTime) * time.Second)
					continue
				} else {
					break
				}
			}
		}
	}
	return urls
}

func doiCmdRunOptions(cmd *cobra.Command) {
	checkArgs(cmd, "doi")
	checkDownloadDir(bgetClis.doi != "")
	if bgetClis.doi != "" || bgetClis.listFile != "" {
		downloadDoi()
		bgetClis.helpFlags = false
	}
	if bgetClis.helpFlags {
		cmd.Help()
	}
}

func init() {
	doiCmd.Flags().BoolVarP(&pmc, "pmc", "", false, "Try PMC database.")
	doiCmd.Flags().BoolVarP(&fullText, "full-text", "", true, "Access full text.")
	doiCmd.Flags().BoolVarP(&suppl, "suppl", "", false, "Access supplementary files.")
	doiCmd.Flags().StringVarP(&(bgetClis.engine), "engine", "g", "go-http", "Point the download engine: go-http, wget, curl, axel, git, and rsync.")
	doiCmd.Flags().IntVarP(&(bgetClis.axelThread), "thread-axel", "", 5, "Set the thread of axel.")
	doiCmd.Flags().StringVarP(&(bgetClis.mirror), "mirror", "m", "", "Set the mirror of resources.")
	doiCmd.Flags().StringVarP(&(bgetClis.listFile), "list-file", "l", "", "A file contains dois for download.")
	exampleXML2Json := "`bapi ncbi --xml2json pubmed titleSearch.XML |grep Doi| tr -d ' ,(Doi:)\"'`"
	doiCmd.Example = fmt.Sprintf(`  bget doi 10.5281/zenodo.3363060 10.5281/zenodo.3357455 10.5281/zenodo.3351812 -t 3
  bget doi 10.1016/j.devcel.2017.03.001 10.1016/j.stem.2019.07.009 10.1016/j.celrep.2018.03.072 -t 2

  bapi ncbi -q '((The PARK10 gene USP24 is a negative regulator of autophagy and ULK1 protein stability[Title]) OR Coordinate regulation of autophagy and the ubiquitin proteasome system by MTOR[Title])' -o titleSearch.XML
  dois=%s
  echo ${dois}
  bget doi ${dois}
  bget doi 10.1080/15548627.2018.1505155 --proxy http://username:password@hostname:port`, exampleXML2Json)
}
