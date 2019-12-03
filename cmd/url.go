package cmd

import (
	"path"
	"strings"

	"github.com/openbiox/bget/urlpool"
	"github.com/openbiox/butils/archive"
	cio "github.com/openbiox/butils/io"
	log "github.com/openbiox/butils/log"
	cnet "github.com/openbiox/butils/net"
	"github.com/spf13/cobra"
)

var URLCmd = &cobra.Command{
	Use:   "url [url1 url2 url3...]",
	Short: "Can be used to access URLs via Golang http, wget, curl, axel and git, and rsync.",
	Long:  `Can be used to access URLs via Golang http, wget, curl, axel and git, and rsync. More see here https://github.com/openbiox/bget.`,
	Run: func(cmd *cobra.Command, args []string) {
		URLCmdRunOptions(cmd)
	},
}

func downloadUrls() {
	urls := []string{}
	if bgetClis.urls != "" && strings.Contains(bgetClis.urls, bgetClis.seperator) {
		urls = strings.Split(bgetClis.urls, bgetClis.seperator)
	} else if bgetClis.urls != "" {
		urls = []string{bgetClis.urls}
	} else if bgetClis.listFile != "" {
		urls = cio.ReadLines(bgetClis.listFile)
	}
	var destDirArray []string
	for i := range urls {
		urls[i] = strings.TrimSpace(urls[i])
		if !strings.Contains(urls[i], "://") && !strings.Contains(urls[i], "git@") {
			urls[i] = "https://" + urls[i]
		}
		destDirArray = append(destDirArray, bgetClis.downloadDir)
	}

	netOpt := setNetParams(&bgetClis)
	done := cnet.HttpGetURLs(urls, destDirArray, netOpt)
	for _, dest := range done {
		if bgetClis.uncompress {
			if err := archive.UnarchiveLog(dest, path.Dir(dest)); err != nil {
				log.Warn(err)
			}
		}
	}
}

func downloadGitHubRepos() {
	urls := []string{}
	if bgetClis.github != "" && strings.Contains(bgetClis.github, bgetClis.seperator) {
		urls = strings.Split(bgetClis.github, bgetClis.seperator)
	} else if bgetClis.github != "" {
		urls = []string{bgetClis.github}
	} else if bgetClis.listFile != "" {
		urls = cio.ReadLines(bgetClis.listFile)
	}
	var destDirArray []string
	assetsUrls := make(map[string][]string)
	versIdx := strings.Split(bgetClis.withAssetsVersions, bgetClis.seperator)
	for i := range urls {
		urls[i] = strings.TrimSpace(urls[i])
		urls[i] = strings.ReplaceAll(urls[i], "git@github.com:", "")
		urls[i] = strings.ReplaceAll(urls[i], "https://github.com/", "")
		urls[i] = strings.ReplaceAll(urls[i], "http://github.com/", "")
		key := urls[i]
		urls[i] = "https://github.com/" + urls[i]

		if bgetClis.withAssets || bgetClis.onlyAssets {
			vers := urlpool.GitHubVersionSpider(urls[i], false)
			if len(vers) != 0 {
				log.Infof("Availabe tags of github.com/%s: %s", key, strings.Join(vers, ", "))
				idx := vers[0]
				if len(versIdx) == len(urls) && versIdx[i] != "" {
					idx = versIdx[i]
				}
				assetsUrls[key+"/"+idx] = urlpool.GitHubAssetsSpider(urls[i], idx)
			} else {
				log.Infoln("assets file not found......")
			}
		}
		destDirArray = append(destDirArray, bgetClis.downloadDir)
	}
	netOpt := setNetParams(&bgetClis)
	if !bgetClis.onlyAssets {
		cnet.HttpGetURLs(urls, destDirArray, netOpt)
	}
	urls2 := []string{}
	destDirArray2 := []string{}
	for k, v := range assetsUrls {
		urls2 = append(urls2, v...)
		for range v {
			destDirArray2 = append(destDirArray2, path.Join(bgetClis.downloadDir, "github-assets", strings.ReplaceAll(k, "https://github.com/", "")))
		}
	}
	netOpt = setNetParams(&bgetClis)
	done := cnet.HttpGetURLs(urls2, destDirArray2, netOpt)
	for _, dest := range done {
		if bgetClis.uncompress {
			if err := archive.UnarchiveLog(dest, path.Dir(dest)); err != nil {
				log.Warn(err)
			}
		}
	}
}

func URLCmdRunOptions(cmd *cobra.Command) {
	checkArgs(cmd, "url")
	checkDownloadDir(bgetClis.urls != "" || bgetClis.listFile != "")
	if (bgetClis.urls != "" || bgetClis.listFile != "") && !bgetClis.githubMode {
		downloadUrls()
		bgetClis.helpFlags = false
	} else if (bgetClis.urls != "" || bgetClis.listFile != "") && bgetClis.githubMode {
		bgetClis.github = bgetClis.urls
		downloadGitHubRepos()
		bgetClis.helpFlags = false
	}
	if bgetClis.helpFlags {
		cmd.Help()
	}
}
func init() {
	setGlobalFlag(URLCmd, &bgetClis)
	setKeyListFlag(URLCmd, &bgetClis, "urls")
	setUncompressFlag(URLCmd, &bgetClis)
	URLCmd.Flags().BoolVarP(&(bgetClis.githubMode), "github", "", false, "GitHub mode.")
	URLCmd.Flags().BoolVarP(&(bgetClis.onlyAssets), "only-github-assets", "", false, "Logical indicating that whether to only download github repo assets files.")
	URLCmd.Flags().BoolVarP(&(bgetClis.withAssets), "with-github-assets", "", false, "Logical indicating that whether to download associated assets files of github repo.")
	URLCmd.Flags().StringVarP(&(bgetClis.withAssetsVersions), "github-assets-versions", "", "", "Required to get specific tagname of github assets (e.g. v2.7.1,v1.0.0).")
	URLCmd.Example = `  urls="https://dldir1.qq.com/weixin/Windows/WeChatSetup.exe,http://download.oray.com/pgy/windows/PgyVPN_4.1.0.21693.exe,https://dldir1.qq.com/qqfile/qq/PCQQ9.1.6/25786/QQ9.1.6.25786.exe" && echo $urls | tr "," "\n"> /tmp/urls.list

  bget url ${urls}
  bget url https://dldir1.qq.com/weixin/Windows/WeChatSetup.exe https://dldir1.qq.com/qqfile/qq/PCQQ9.1.6/25786/QQ9.1.6.25786.exe
  bget url ${urls} -t 2 -o /tmp/download
  bget url ${urls} -t 3 -o /tmp/download -f -g wget
  bget url ${urls} -t 3 -o /tmp/download -g wget --ignore
  bget url -l /tmp/urls.list -o /tmp/download -f -t 3

  bget url Miachol/github_demo --github
  bget url PapenfussLab/gridss openbiox/bget --with-github-assets -t 5 --github
  bget url PapenfussLab/gridss openbiox/bget --only-github-assets -t 5 --github
  bget url PapenfussLab/gridss openbiox/bget --with-github-assets --with-assets-versions v2.7.2,v0.1.3 -t 5 --github`
}
