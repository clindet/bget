package cmd

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/JhuangLab/bget/utils"
)

func downloadKey() {
	keys := []string{}
	urls := []string{}
	if downloadClis.keys != "" && strings.Contains(downloadClis.keys, downloadClis.separator) {
		keys = strings.Split(downloadClis.keys, downloadClis.separator)
	} else if downloadClis.keys != "" {
		keys = []string{downloadClis.keys}
	}
	urls = keys2urls(keys)
	var destDirArray []string
	for i := range urls {
		u, _ := url.Parse(urls[i])
		urls[i] = strings.TrimSpace(u.String())
		destDirArray = append(destDirArray, downloadClis.downloadDir)
	}

	HTTPGetURLs(urls, destDirArray, downloadClis.engine, taskID, downloadClis.mirror,
		downloadClis.concurrency, downloadClis.axelThread, overwrite, downloadClis.ignore, quiet, saveLog)
}

func keys2urls(keys []string) (urls []string) {
	version := ""
	site := ""
	release := ""
	key := ""
	for k := range keys {
		key, version, site, release = parseMeta(keys[k])
		if version == "" {
			version = "latest"
		}
		if tmp := getEnvToolsURL(key, version, release); tmp != "" {
			urls = append(urls, tmp)
		}

		if tmp := getEnvFilesURL(key, site, version, release); len(tmp) > 0 {
			urls = append(urls, tmp...)
		}
	}
	return urls
}

func parseMeta(key string) (keyNew string, version string, site string, release string) {
	info := utils.StrSplit(key, "@|%|#", 4)
	info1 := utils.StrSplit(key, "@", 2)
	info2 := utils.StrSplit(key, "%", 2)
	info3 := utils.StrSplit(key, "#", 2)
	keyNew = strings.TrimSpace(info[0])

	if len(info) == 2 {
		if len(info1) > 1 {
			version = info[1]
		} else if len(info2) > 1 {
			site = info[1]
		} else if len(info3) > 1 {
			release = info[1]
		}
	} else if len(info) == 3 {
		if len(info1) > 1 && len(info2) > 1 {
			if strings.Contains(info1[1], "%") {
				version = info[1]
				site = info[2]
			} else {
				version = info[2]
				site = info[1]
			}
		} else if len(info1) > 1 && len(info3) > 1 {
			if strings.Contains(info1[1], "#") {
				version = info[1]
				release = info[2]
			} else {
				version = info[2]
				release = info[1]
			}
		} else if len(info2) > 1 && len(info3) > 1 {
			if strings.Contains(info2[1], "#") {
				site = info[1]
				release = info[2]
			} else {
				site = info[2]
				release = info[1]
			}
		}
	} else if len(info) == 4 {
		if strings.Contains(info1[1], "#") && strings.Contains(info1[1], "%") {
			version = info[1]
		} else if !strings.Contains(info1[1], "#") && !strings.Contains(info1[1], "%") {
			version = info[3]
		} else {
			version = info[2]
		}
		if strings.Contains(info2[1], "@") && strings.Contains(info2[1], "#") {
			site = info[1]
		} else if !strings.Contains(info2[1], "@") && !strings.Contains(info2[1], "#") {
			site = info[3]
		} else {
			site = info[2]
		}
		if strings.Contains(info3[1], "@") && strings.Contains(info3[1], "%") {
			release = info[1]
		} else if !strings.Contains(info3[1], "@") && !strings.Contains(info3[1], "%") {
			release = info[3]
		} else {
			release = info[2]
		}
	}
	keyNew = strings.TrimSpace(keyNew)
	site = strings.TrimSpace(site)
	version = strings.TrimSpace(version)
	release = strings.TrimSpace(release)
	return keyNew, version, site, release
}

func getAllKeys() (keys []string) {
	for i := range bgetFilesURLs {
		keys = append(keys, bgetFilesURLs[i].Name)
	}
	for i := range bgetToolsURLs {
		keys = append(keys, bgetToolsURLs[i].Name)
	}
	keys = utils.RemoveRepeatEle(keys)
	fmt.Printf("%s\n", strings.Join(keys, "\n"))
	return keys
}
