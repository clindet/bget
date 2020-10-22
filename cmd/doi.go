package cmd

import (
	//"bufio"
	//"bytes"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	neturl "net/url"
	"path"
	"strings"
	"time"

	"github.com/clindet/bget/api/fetch"
	"github.com/clindet/bget/api/types"
	"github.com/clindet/bget/spider"
	cio "github.com/openbiox/ligo/io"
	cnet "github.com/openbiox/ligo/net"
	"github.com/openbiox/ligo/slice"
	stringo "github.com/openbiox/ligo/stringo"
	"github.com/spf13/cobra"
	"github.com/tidwall/pretty"
)

var fullText string
var suppl bool
var pmc bool
var universeSpider bool
var scihub bool
var printSiteMeta bool
var printCrossRefMeta bool

var DoiCmd = &cobra.Command{
	Use:   "doi [doi1 doi2 doi3...]",
	Short: "Can be used to access files via DOI.",
	Long:  `Can be used to access files via DOI. More see here https://github.com/clindet/bget.`,
	Run: func(cmd *cobra.Command, args []string) {
		doiCmdRunOptions(cmd, args)
	},
}

func doiCmdRunOptions(cmd *cobra.Command, args []string) {
	initCmd(cmd, args)
	checkArgs(cmd, "doi")
	checkDownloadDir(bgetClis.Doi != "")
	if bgetClis.Doi != "" || bgetClis.ListFile != "" {
		downloadDoi(cmd, args)
		bgetClis.HelpFlags = false
	}
	if bgetClis.HelpFlags {
		cmd.Help()
	}
}

func downloadDoi(cmd *cobra.Command, args []string) {
	var sem = make(chan bool, bgetClis.Thread)
	var urls = []string{}
	var destDirArray []string
	var doi = parseArgsDoi()
	for _, v := range doi {
		sem <- true
		go func(v string) {
			defer func() {
				<-sem
			}()
			urlsTmp, opt := doiSpiders(v)
			*urlsTmp = slice.DropSliceDup(*urlsTmp)
			var outDir = path.Join(bgetClis.DownloadDir, v)
			if opt != nil && opt.PrintSiteMeta {
				outputSiteMetaData(outDir, opt)
			}
			if opt != nil && opt.PrintCrossRefMeta {
				outputCrossRefData(outDir, opt, cmd, args)
			}
			for range *urlsTmp {
				destDirArray = append(destDirArray, outDir)
			}
			urls = append(urls, *urlsTmp...)
		}(v)
	}
	for i := 0; i < cap(sem); i++ {
		sem <- true
	}
	netOpt := setNetParams(&bgetClis)
	cnet.HTTPGetURLs(urls, destDirArray, netOpt)
}

func parseArgsDoi() (doi []string) {
	if bgetClis.Doi != "" && strings.Contains(bgetClis.Doi, bgetClis.Seperator) {
		doi = strings.Split(bgetClis.Doi, bgetClis.Seperator)
	} else if bgetClis.Doi != "" && strings.Contains(bgetClis.Doi, "\n") {
		doi = strings.Split(bgetClis.Doi, "\n")
	} else if bgetClis.Doi != "" {
		doi = []string{bgetClis.Doi}
	} else if bgetClis.ListFile != "" {
		doi = cio.ReadLines(bgetClis.ListFile)
	}
	return doi
}

func outputSiteMetaData(outDir string, opt *spider.DoiSpiderOpt) {
	var err error
	var of io.Writer
	var buf *bytes.Buffer
	metaJSON, _ := json.Marshal(*opt.CitationMeta)
	log.Infof("Website metadata of %s: %s", opt.Doi, string(metaJSON))
	outfn := path.Join(outDir, "website.meta.json")
	log.Infof("Saving website metadata of %s => %s", opt.Doi, outfn)
	cio.CreateFileParDir(outfn)
	if cio.CheckFilesOverwrite([]string{outfn}, bgetClis.Overwrite) == -1 {
		of, err = cio.Open(outfn)
		if err != nil {
			log.Warnln(err)
		} else {
			indent := ""
			for i := 0; i < 2; i++ {
				indent = indent + " "
			}
			opt := pretty.Options{
				Indent:   indent,
				SortKeys: false,
			}
			buf = bytes.NewBuffer(pretty.PrettyOptions(metaJSON, &opt))
			io.Copy(of, buf)
		}
	}
}

func outputCrossRefData(outDir string, opt *spider.DoiSpiderOpt,
	cmd *cobra.Command, args []string) {
	var err error
	var crossRefEndp = types.CrossRefEndpoints{}
	crossRefEndp.Doi.Doi = opt.Doi
	var bapiClis = types.BapiClisT{
		TaskID:       bgetClis.TaskID,
		Verbose:      bgetClis.Verbose,
		SaveLog:      bgetClis.SaveLog,
		LogDir:       bgetClis.LogDir,
		HelpFlags:    false,
		Proxy:        bgetClis.Proxy,
		Timeout:      bgetClis.Timeout,
		RetSleepTime: bgetClis.RetSleepTime,
		Retries:      bgetClis.Retries,
		PrettyJSON:   false,
		XML2json:     true,
	}
	var buf bytes.Buffer
	var of = io.Writer(&buf)
	var outfn string
	fetch.CrossRef(&crossRefEndp, &bapiClis, func() { initCmd(cmd, args) }, of)
	log.Infof("Crossref metadata of %s: %s", opt.Doi, string(buf.Bytes()))
	outfn = path.Join(outDir, "crossref.citation.json")
	log.Infof("Saving crossref metadata of %s => %s", opt.Doi, outfn)
	cio.CreateFileParDir(outfn)
	if cio.CheckFilesOverwrite([]string{outfn}, bgetClis.Overwrite) == -1 {
		of, err = cio.Open(outfn)
		if err != nil {
			log.Warnln(err)
		} else {
			indent := ""
			for i := 0; i < 2; i++ {
				indent = indent + " "
			}
			opt := pretty.Options{
				Indent:   indent,
				SortKeys: false,
			}
			bufTmp := bytes.NewBuffer(pretty.PrettyOptions(buf.Bytes(), &opt))
			io.Copy(of, bufTmp)
		}
	}
}

func doiSpiders(doi string) (urls *[]string, opt *spider.DoiSpiderOpt) {
	urls = &[]string{}
	if !strings.Contains(doi, "/") {
		return urls, opt
	}
	if stringo.StrDetect(doi, "http[s]://doi.org/") {
		doi = stringo.StrReplaceAll(doi, "http[s]://doi.org/", "")
	}
	doiTmp := strings.Split(doi, "/")
	doiOrg := doiTmp[0]
	var t int
	var citationMeta = make(map[string]string)
	opt = &spider.DoiSpiderOpt{
		Doi:               doi,
		Proxy:             bgetClis.Proxy,
		Timeout:           bgetClis.Timeout,
		FullText:          fullText == "true",
		PrintSiteMeta:     printSiteMeta,
		PrintCrossRefMeta: printCrossRefMeta,
		CitationMeta:      &citationMeta,
		Supplementary:     suppl,
	}
	if pmc {
		*urls = spider.PmcSpider(opt)
		if len(*urls) == 0 && !suppl {
			return urls, opt
		}
	}
	if len(*urls) == 0 && scihub {
		*urls = spider.ScihupSpider(opt)
		if len(*urls) > 0 && !suppl {
			return urls, opt
		}
	}
	if len(*urls) > 0 {
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
			for t = 0; t < bgetClis.Retries+1; t++ {
				*urls = append(*urls, spider.DoiSpidersPool[doiOrg](opt)...)
				if len(*urls) == 0 {
					if bgetClis.Retries != 0 {
						log.Warnf("Pool spider: %s returns empty, on attempt %d... retrying after %d seconds.", doi, t+1, bgetClis.RetSleepTime)
						time.Sleep(time.Duration(bgetClis.RetSleepTime) * time.Second)
					} else {
						log.Warnf("Pool spider: %s returns empty, on attempt %d...", doi, t+1)
					}
					continue
				} else {
					break
				}
			}
		}
	}
	if len(*urls) == 0 && universeSpider {
		for t = 0; t < bgetClis.Retries; t++ {
			*urls = append(*urls, spider.UniVersalDoiSpider(opt)...)
			if len(*urls) == 0 {
				if bgetClis.Retries != 0 {
					log.Warnf("Universal spider: %s returns empty, on attempt %d... retrying after %d seconds.", doi, t+1, bgetClis.RetSleepTime)
					time.Sleep(time.Duration(bgetClis.RetSleepTime) * time.Second)
				} else {
					log.Warnf("Pool spider: %s returns empty, on attempt %d...", doi, t+1)
				}
				continue
			} else {
				break
			}
		}
	}
	return urls, opt
}

func init() {
	universeSpider = true
	DoiCmd.Flags().BoolVarP(&scihub, "enable-scihub", "", false, "enable try scihub spider.")
	DoiCmd.Flags().BoolVarP(&pmc, "enable-pmc", "", false, "enable try PMC database.")
	DoiCmd.Flags().StringVarP(&fullText, "full-text", "", "true", "access full text.")
	DoiCmd.Flags().BoolVarP(&suppl, "suppl", "", false, "access supplementary files.")
	DoiCmd.Flags().BoolVarP(&printSiteMeta, "print-meta", "", false, "print website meta data.")
	DoiCmd.Flags().BoolVarP(&printCrossRefMeta, "print-crossref", "", false, "print crossref meta data.")

	setGlobalFlag(DoiCmd, &bgetClis)
	setKeyListFlag(DoiCmd, &bgetClis, "dois")
	exampleXML2Json := "`bget api ncbi --xml2json --json-pretty -q '30487223[pmid] or 30402350[pmid] or 29279377[pmid]' --size 3 -m 3 | grep / | grep 10. | sed 's/ .* \"//' | tr -d '\",' | sort -u`"
	DoiCmd.Example = fmt.Sprintf(`  bget doi 10.5281/zenodo.3363060 10.5281/zenodo.3357455 10.5281/zenodo.3351812 -t 3
  bget doi 10.1016/j.devcel.2017.03.001 10.1016/j.stem.2019.07.009 10.1016/j.celrep.2018.03.072 -t 2
	
  # query pdf and meta data using PubMed ID
  dois=%s && echo ${dois} && bget doi ${dois} --print-meta --print-crossref
  
  bget doi 10.1080/15548627.2018.1505155 --proxy http://username:password@hostname:port
  bget doi 10.1182/blood.2019000200 --enable-scihub
  bget doi 10.1109/JPROC.2019.2905423 -n
  bget doi 10.1073/pnas.1814397115 --print-meta --print-crossref`, exampleXML2Json)
}
