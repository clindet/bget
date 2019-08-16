package cmd

import (
	"path"
	"strings"

	"github.com/JhuangLab/bget/utils"
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
		key, version, site, release := parseMeta("reffa@" + reffaArray[reffaI])
		destDir := ""
		if downloadClis.autoPath {
			destDir = path.Join(downloadClis.downloadDir, "reffa", site, version, release)
		} else {
			destDir = downloadClis.downloadDir
		}
		if site == "defuse" {
			defuse = destDir
		}
		urls := getEnvFilesURL(key, site, version, release)
		urlsPool = append(urlsPool, urls...)
		for range urls {
			destDirPool = append(destDirPool, destDir)
		}
	}

	HTTPGetURLs(urlsPool, destDirPool, downloadClis.engine, taskID, downloadClis.mirror,
		downloadClis.concurrency, downloadClis.axelThread,
		overwrite, downloadClis.ignore, quiet, saveLog)
	if defuse != "" {
		utils.GunzipDefuseReffa(defuse)
	}
}
