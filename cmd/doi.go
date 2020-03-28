package cmd

import (
	"fmt"
	"net/http"
	neturl "net/url"
	"path"
	"strings"
	"time"

	"github.com/openbiox/bget/spider"
	cio "github.com/openbiox/ligo/io"
	cnet "github.com/openbiox/ligo/net"
	"github.com/openbiox/ligo/slice"
	stringo "github.com/openbiox/ligo/stringo"
	"github.com/spf13/cobra"
)

var fullText string
var suppl bool
var pmc bool
var universeSpider bool
var scihub bool

var DoiCmd = &cobra.Command{
	Use:   "doi [doi1 doi2 doi3...]",
	Short: "Can be used to access files via DOI.",
	Long:  `Can be used to access files via DOI. More see here https://github.com/openbiox/bget.`,
	Run: func(cmd *cobra.Command, args []string) {
		doiCmdRunOptions(cmd, args)
	},
}

func doiCmdRunOptions(cmd *cobra.Command, args []string) {
	initCmd(cmd, args)
	checkArgs(cmd, "doi")
	checkDownloadDir(bgetClis.Doi != "")
	if bgetClis.Doi != "" || bgetClis.ListFile != "" {
		downloadDoi()
		bgetClis.HelpFlags = false
	}
	if bgetClis.HelpFlags {
		cmd.Help()
	}
}

func downloadDoi() {
	sem := make(chan bool, bgetClis.Thread)
	doi := []string{}
	urls := []string{}
	var destDirArray []string
	if bgetClis.Doi != "" && strings.Contains(bgetClis.Doi, bgetClis.Seperator) {
		doi = strings.Split(bgetClis.Doi, bgetClis.Seperator)
	} else if bgetClis.Doi != "" && strings.Contains(bgetClis.Doi, "\n") {
		doi = strings.Split(bgetClis.Doi, "\n")
	} else if bgetClis.Doi != "" {
		doi = []string{bgetClis.Doi}
	} else if bgetClis.ListFile != "" {
		doi = cio.ReadLines(bgetClis.ListFile)
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
				destDirArray = append(destDirArray, path.Join(bgetClis.DownloadDir, v))
			}
			urls = append(urls, urlsTmp...)
		}(v)
	}
	for i := 0; i < cap(sem); i++ {
		sem <- true
	}
	netOpt := setNetParams(&bgetClis)
	cnet.HTTPGetURLs(urls, destDirArray, netOpt)
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
		Proxy:         bgetClis.Proxy,
		Timeout:       bgetClis.Timeout,
		FullText:      fullText == "true",
		Supplementary: suppl,
	}
	if pmc {
		urls = spider.PmcSpider(&opt)
		if len(urls) > 0 && !suppl {
			return urls
		} else if len(urls) == 0 {
			return
		}
	}
	if len(urls) == 0 && scihub {
		urls = spider.ScihupSpider(&opt)
		if len(urls) > 0 && !suppl {
			return urls
		} else if len(urls) == 0 {
			return
		}
	}
	if len(urls) > 0 {
		opt.FullText = false
	}
	client := cnet.NewHTTPClient(opt.Timeout, opt.Proxy)
	req, _ := http.NewRequest("HEAD", "https://doi.org/"+opt.Doi, nil)
	resp, err := client.Do(req)
	if err != nil && strings.Contains(err.Error(), "http") {
		link := stringo.StrExtract(err.Error(), `".*"`, 1)[0]
		link = stringo.StrReplaceAll(link, `"`, "")
		u, _ := neturl.Parse(link)
		opt.URL = u
	} else {
		defer resp.Body.Close()
		u, _ := neturl.Parse(resp.Request.URL.String())
		opt.URL = u
	}
	for k := range spider.DoiSpidersPool {
		if k == doiOrg {
			for t = 0; t < bgetClis.Retries; t++ {
				urls = append(urls, spider.DoiSpidersPool[doiOrg](&opt)...)
				if len(urls) == 0 {
					log.Warnf("%s returns empty, on attempt %d... retrying after %d seconds.", doi, t+1, bgetClis.RetSleepTime)
					time.Sleep(time.Duration(bgetClis.RetSleepTime) * time.Second)
					continue
				} else {
					break
				}
			}
		}
	}
	if len(urls) == 0 && universeSpider {
		urls = append(urls, spider.UniVersalDoiSpider(&opt)...)
	}
	return urls
}

func init() {
	universeSpider = true
	DoiCmd.Flags().BoolVarP(&scihub, "enable-scihub", "", false, "enable try scihub spider.")
	DoiCmd.Flags().BoolVarP(&pmc, "enable-pmc", "", false, "enable try PMC database.")
	DoiCmd.Flags().StringVarP(&fullText, "full-text", "", "true", "access full text.")
	DoiCmd.Flags().BoolVarP(&suppl, "suppl", "", false, "access supplementary files.")
	setGlobalFlag(DoiCmd, &bgetClis)
	setKeyListFlag(DoiCmd, &bgetClis, "dois")
	exampleXML2Json := "`bioctl cvrt --xml2json pubmed titleSearch.XML | bioextr --mode pubmed - | grep Doi | tr -d ' ,(Doi:)\"'`"
	DoiCmd.Example = fmt.Sprintf(`  bget doi 10.5281/zenodo.3363060 10.5281/zenodo.3357455 10.5281/zenodo.3351812 -t 3
  bget doi 10.1016/j.devcel.2017.03.001 10.1016/j.stem.2019.07.009 10.1016/j.celrep.2018.03.072 -t 2

  bget api ncbi -q '((The PARK10 gene USP24 is a negative regulator of autophagy and ULK1 protein stability[Title]) OR Coordinate regulation of autophagy and the ubiquitin proteasome system by MTOR[Title])' -o titleSearch.XML
	
  # github.com/openbiox/bioctl / github.com/openbiox/bioextr
  dois=%s
  echo ${dois}
  bget doi ${dois}
  bget doi 10.1080/15548627.2018.1505155 --proxy http://username:password@hostname:port
  bget doi 10.1182/blood.2019000200 --enable-scihub
  bget doi 10.1109/JPROC.2019.2905423 -n`, exampleXML2Json)
}
