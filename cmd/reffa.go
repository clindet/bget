package cmd

import (
	"fmt"
	"path"
	"strings"

	"github.com/JhuangLab/bget/log"
)

func installReffa() {
	reffaArray := []string{}
	if strings.Contains(downloadClis.installReffa, downloadClis.separator) {
		reffaArray = strings.Split(downloadClis.installReffa, downloadClis.separator)
	} else {
		reffaArray = []string{downloadClis.installReffa}
	}
	urlsPool := []string{}
	destDirPool := []string{}
	defuse := ""
	for reffaI := range reffaArray {
		if !strings.Contains(reffaArray[reffaI], "_") {
			log.Fatal(fmt.Sprintf("Invaild parameters (--get-reffa): %s. Optional (GRCh38_genecode_31, GRCh37_genecode_31, hg38_ucsc, and hg19_ucsc). Multiple usage: seperate items with comma (e.g. hg38_ucsc,hg19_ucsc)", reffaArray[reffaI]))
		}
		reffa := strings.Split(reffaArray[reffaI], "_")
		version := reffa[0]
		site := reffa[1]
		release := ""
		if len(reffa) == 3 {
			release = reffa[2]
		}
		destDir := ""
		if downloadClis.autoPath {
			destDir = path.Join(downloadClis.downloadDir, "reffa", site, version, release)
		} else {
			destDir = downloadClis.downloadDir
		}
		if site == "defuse" {
			defuse = destDir
		}
		urls := getEnvFilesURL("reffa", site, version, release)
		urlsPool = append(urlsPool, urls...)
		for range urls {
			destDirPool = append(destDirPool, destDir)
		}
	}

	HTTPGetURLs(urlsPool, destDirPool, downloadClis.engine, taskID, downloadClis.mirror,
		downloadClis.concurrency, downloadClis.axelThread,
		overwrite, downloadClis.ignore, quiet, saveLog)
	if defuse != "" {
		gunzipDefuseReffa(defuse)
	}
}
