package cmd

import (
	"fmt"
	"net/url"
	"strings"

	butils "github.com/JhuangLab/butils"
	"github.com/spf13/cobra"
)

var keyCmd = &cobra.Command{
	Use:   "key [key1 key2 key3...]",
	Short: "Can be used to access URLs via a key string.",
	Long:  `Can be used to access URLs via a key string. e.g. 'item' or 'item@version %site #releaseVersion', : bwa, GRCh38 %defuse #97. More see here https://github.com/JhuangLab/bget.`,
	Run: func(cmd *cobra.Command, args []string) {
		keyCmdRunOptions(cmd)
	},
}

func downloadKey() {
	keys := []string{}
	urls := []string{}
	if bgetClis.keys != "" && strings.Contains(bgetClis.keys, bgetClis.separator) {
		keys = strings.Split(bgetClis.keys, bgetClis.separator)
	} else if bgetClis.keys != "" {
		keys = []string{bgetClis.keys}
	} else if bgetClis.listFile != "" {
		keys = butils.ReadLines(bgetClis.listFile)
	}
	urls = keys2urls(keys)
	var destDirArray []string
	for i := range urls {
		u, _ := url.Parse(urls[i])
		urls[i] = strings.TrimSpace(u.String())
		destDirArray = append(destDirArray, bgetClis.downloadDir)
	}

	HTTPGetURLs(urls, destDirArray, bgetClis.engine, taskID, bgetClis.mirror,
		bgetClis.concurrency, bgetClis.axelThread, overwrite, bgetClis.ignore, quiet, saveLog)
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
	info := butils.StrSplit(key, "@|%|#", 4)
	info1 := butils.StrSplit(key, "@", 2)
	info2 := butils.StrSplit(key, "%", 2)
	info3 := butils.StrSplit(key, "#", 2)
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
	keys = butils.RemoveRepeatEle(keys)
	fmt.Printf("%s\n", strings.Join(keys, "\n"))
	return keys
}

func keyCmdRunOptions(cmd *cobra.Command) {
	checkQuiet()
	items := []string{}
	if len(cmd.Flags().Args()) >= 1 {
		items = append(items, cmd.Flags().Args()...)
		bgetClis.keys = strings.Join(items, bgetClis.separator)
	}
	checkDownloadDir(bgetClis.keys != "")
	if bgetClis.keysAll {
		getAllKeys()
		bgetClis.helpFlags = false
	}
	if bgetClis.keys != "" || bgetClis.listFile != "" {
		downloadKey()
		bgetClis.helpFlags = false
	}
	if bgetClis.helpFlags {
		cmd.Help()
	}
}

func init() {
	keyCmd.Flags().StringVarP(&(bgetClis.listFile), "list-file", "l", "", "A file contains keys for download.")
	keyCmd.Flags().BoolVarP(&(bgetClis.keysAll), "keys-all", "a", false, "Show all available string key can be download.")
	keyCmd.Example = `  bget key bwa
  bget key "reffa@GRCh38 %defuse #97" -t 10 -f`
}
