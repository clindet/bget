package cmd

import (
	"fmt"
	"path"
	"strings"

	"github.com/Miachol/bget/spider"
	butils "github.com/openbiox/butils"
	"github.com/spf13/cobra"
)

var doiCmd = &cobra.Command{
	Use:   "doi [doi1 doi2 doi3...]",
	Short: "Can be used to access files via DOI.",
	Long:  `Can be used to access files via DOI. More see here https://github.com/Miachol/bget.`,
	Run: func(cmd *cobra.Command, args []string) {
		doiCmdRunOptions(cmd)
	},
}

func downloadDoi() {
	sem := make(chan bool, bgetClis.concurrency)
	doi := []string{}
	urls := []string{}
	var destDirArray []string
	if bgetClis.doi != "" && strings.Contains(bgetClis.doi, bgetClis.separator) {
		doi = strings.Split(bgetClis.doi, bgetClis.separator)
	} else if bgetClis.doi != "" {
		doi = []string{bgetClis.doi}
	} else if bgetClis.listFile != "" {
		doi = butils.ReadLines(bgetClis.listFile)
	}
	for _, v := range doi {
		sem <- true
		go func(v string) {
			defer func() {
				<-sem
			}()
			urlsTmp := doiSpiders(v)
			urlsTmp = butils.RemoveRepeatEle(urlsTmp)
			for range urlsTmp {
				destDirArray = append(destDirArray, path.Join(bgetClis.downloadDir, v))
			}
			urls = append(urls, urlsTmp...)
		}(v)
	}
	for i := 0; i < cap(sem); i++ {
		sem <- true
	}
	HTTPGetURLs(urls, destDirArray, bgetClis.engine, cmdExtraFromFlag, taskID, bgetClis.mirror,
		bgetClis.concurrency, bgetClis.axelThread, overwrite, ignore, quiet, saveLog, bgetClis.retries, bgetClis.timeout, bgetClis.retSleepTime, bgetClis.remoteName)
}

func doiSpiders(doi string) (urls []string) {
	if !strings.Contains(doi, "/") {
		return urls
	}
	doiTmp := strings.Split(doi, "/")
	doiOrg := doiTmp[0]
	runFlag := false
	for k := range spider.DoiSpidersPool {
		if k == doiOrg {
			urls = spider.DoiSpidersPool[doiOrg](doi)
			runFlag = true
		}
	}
	if !runFlag || len(urls) == 0 {
		//urls = spider.ScihubSpider(doi)
	}
	return urls
}

func doiCmdRunOptions(cmd *cobra.Command) {
	checkQuiet()
	checkDownloadDir(bgetClis.doi != "")
	items := []string{}
	if len(cmd.Flags().Args()) >= 1 {
		items = append(items, cmd.Flags().Args()...)
		bgetClis.doi = strings.Join(items, bgetClis.separator)
	}
	if bgetClis.doi != "" || bgetClis.listFile != "" {
		downloadDoi()
		bgetClis.helpFlags = false
	}
	if bgetClis.helpFlags {
		cmd.Help()
	}
}

func init() {
	doiCmd.Flags().StringVarP(&(bgetClis.engine), "engine", "g", "go-http", "Point the download engine: go-http, wget, curl, axel, git, and rsync.")
	doiCmd.Flags().IntVarP(&(bgetClis.axelThread), "thread-axel", "", 5, "Set the thread of axel.")
	doiCmd.Flags().StringVarP(&(bgetClis.mirror), "mirror", "m", "", "Set the mirror of resources.")
	doiCmd.Flags().StringVarP(&(bgetClis.listFile), "list-file", "l", "", "A file contains dois for download.")
	doiCmd.Example = fmt.Sprintf(`  bget doi 10.5281/zenodo.3363060 10.5281/zenodo.3357455 10.5281/zenodo.3351812 -t 3
  bget doi 10.1016/j.devcel.2017.03.001 10.1016/j.stem.2019.07.009 10.1016/j.celrep.2018.03.072 -t 2

  bapi ncbi -q '((The PARK10 gene USP24 is a negative regulator of autophagy and ULK1 protein stability[Title]) OR Coordinate regulation of autophagy and the ubiquitin proteasome system by MTOR[Title])' -o titleSearch.XML
  dois=%s
  echo ${dois}
  bget doi ${dois}`, "`bapi ncbi --xml2json pubmed titleSearch.XML |grep Doi| tr -d ' ,(Doi:)\"'`")
}
