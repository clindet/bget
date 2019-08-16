package cmd

import (
	"path"
	"strings"

	"github.com/JhuangLab/bget/spider"
	"github.com/JhuangLab/bget/utils"
)

func downloadDoi() {
	sem := make(chan bool, downloadClis.concurrency)
	doi := []string{}
	urls := []string{}
	var destDirArray []string
	if downloadClis.doi != "" && strings.Contains(downloadClis.doi, downloadClis.separator) {
		doi = strings.Split(downloadClis.doi, downloadClis.separator)
	} else if downloadClis.doi != "" {
		doi = []string{downloadClis.doi}
	}
	for _, v := range doi {
		sem <- true
		go func(v string) {
			defer func() {
				<-sem
			}()
			urlsTmp := doiSpiders(v)
			urlsTmp = utils.RemoveRepeatEle(urlsTmp)
			for range urlsTmp {
				destDirArray = append(destDirArray, path.Join(downloadClis.downloadDir, v))
			}
			urls = append(urls, urlsTmp...)
		}(v)
	}
	for i := 0; i < cap(sem); i++ {
		sem <- true
	}
	HTTPGetURLs(urls, destDirArray, downloadClis.engine, taskID, downloadClis.mirror,
		downloadClis.concurrency, downloadClis.axelThread, overwrite, downloadClis.ignore, quiet, saveLog)
}

func doiSpiders(doi string) (urls []string) {
	doiOrg := ""
	doiTmp := strings.Split(doi, "/")
	doiOrg = doiTmp[0]
	for k := range spider.DoiSpidersPool {
		if k == doiOrg {
			urls = spider.DoiSpidersPool[doiOrg](doi)
		}
	}
	return urls
}
