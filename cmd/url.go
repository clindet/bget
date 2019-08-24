package cmd

import (
	"path"
	"strings"

	butils "github.com/JhuangLab/butils"
	log "github.com/JhuangLab/butils/log"
	"github.com/spf13/cobra"
)

var urlCmd = &cobra.Command{
	Use:   "url [url1 url2 url3...]",
	Short: "Can be used to access URLs via Golang http, wget, curl, axel and git, and rsync.",
	Long:  `Can be used to access URLs via Golang http, wget, curl, axel and git, and rsync. More see here https://github.com/JhuangLab/bget.`,
	Run: func(cmd *cobra.Command, args []string) {
		urlCmdRunOptions(cmd)
	},
}

func downloadUrls() {
	urls := []string{}
	if bgetClis.urls != "" && strings.Contains(bgetClis.urls, bgetClis.separator) {
		urls = strings.Split(bgetClis.urls, bgetClis.separator)
	} else if bgetClis.urls != "" {
		urls = []string{bgetClis.urls}
	} else if bgetClis.listFile != "" {
		urls = butils.ReadLines(bgetClis.listFile)
	}
	var destDirArray []string
	for i := range urls {
		urls[i] = strings.TrimSpace(urls[i])
		if !strings.Contains(urls[i], "://") && !strings.Contains(urls[i], "git@") {
			urls[i] = "https://" + urls[i]
		}
		destDirArray = append(destDirArray, bgetClis.downloadDir)
	}

	done := HTTPGetURLs(urls, destDirArray, bgetClis.engine, cmdExtraFromFlag, taskID, bgetClis.mirror,
		bgetClis.concurrency, bgetClis.axelThread, overwrite, ignore, quiet, saveLog)
	for _, dest := range done {
		if bgetClis.uncompress {
			if err := butils.UnarchiveLog(dest, path.Dir(dest)); err != nil {
				log.Warn(err)
			}
		}
	}
}
func urlCmdRunOptions(cmd *cobra.Command) {
	checkQuiet()
	items := []string{}
	if len(cmd.Flags().Args()) >= 1 {
		items = append(items, cmd.Flags().Args()...)
		bgetClis.urls = strings.Join(items, bgetClis.separator)
	}
	checkDownloadDir(bgetClis.urls != "" || bgetClis.listFile != "")

	if bgetClis.urls != "" || bgetClis.listFile != "" {
		downloadUrls()
		bgetClis.helpFlags = false
	}
	if bgetClis.helpFlags {
		cmd.Help()
	}
}
func init() {
	urlCmd.Flags().StringVarP(&(bgetClis.engine), "engine", "g", "go-http", "Point the download engine: go-http, wget, curl, axel, git, and rsync.")
	urlCmd.Flags().IntVarP(&(bgetClis.axelThread), "thread-axel", "", 5, "Set the thread of axel.")
	urlCmd.Flags().StringVarP(&(bgetClis.mirror), "mirror", "m", "", "Set the mirror of resources.")
	urlCmd.Flags().StringVarP(&(bgetClis.listFile), "list-file", "l", "", "A file contains URLs for download.")
	urlCmd.Flags().BoolVarP(&(bgetClis.uncompress), "uncompress", "u", false, "Uncompress download files for .zip, .tar.gz, and .gz suffix files.")
	urlCmd.Example = `  urls="https://dldir1.qq.com/weixin/Windows/WeChatSetup.exe,http://download.oray.com/pgy/windows/PgyVPN_4.1.0.21693.exe,https://dldir1.qq.com/qqfile/qq/PCQQ9.1.6/25786/QQ9.1.6.25786.exe" && echo $urls | tr "," "\n"> /tmp/urls.list

  bget url ${urls}
  bget url https://dldir1.qq.com/weixin/Windows/WeChatSetup.exe https://dldir1.qq.com/qqfile/qq/PCQQ9.1.6/25786/QQ9.1.6.25786.exe
  bget url ${urls} -t 2 -o /tmp/download
  bget url ${urls} -t 3 -o /tmp/download -f -g wget
  bget url ${urls} -t 3 -o /tmp/download -g wget --ignore
  bget url -l /tmp/urls.list -o /tmp/download -f -t 3`
}
