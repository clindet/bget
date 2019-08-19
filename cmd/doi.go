package cmd

import (
	"path"
	"strings"

	"github.com/JhuangLab/bget/spider"
	butils "github.com/JhuangLab/butils"
	"github.com/spf13/cobra"
)

var doiCmd = &cobra.Command{
	Use:   "doi [doi1 doi2 doi3...]",
	Short: "Can be used to access files via DOI.",
	Long:  `Can be used to access files via DOI. More see here https://github.com/JhuangLab/bget.`,
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
	HTTPGetURLs(urls, destDirArray, bgetClis.engine, taskID, bgetClis.mirror,
		bgetClis.concurrency, bgetClis.axelThread, overwrite, bgetClis.ignore, quiet, saveLog)
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
		urls = spider.ScihubSpider(doi)
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
	if bgetClis.doi != "" {
		downloadDoi()
		bgetClis.helpFlags = false
	}
	if bgetClis.helpFlags {
		cmd.Help()
	}
}

func init() {
	doiCmd.Example = `  bget doi 10.5281/zenodo.3363060 10.5281/zenodo.3357455 10.5281/zenodo.3351812 -t 3
  bget doi 10.1016/j.devcel.2017.03.001 10.1016/j.stem.2019.07.009 10.1016/j.celrep.2018.03.072 -t 2`
}
