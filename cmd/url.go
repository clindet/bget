package cmd

import (
	"path"
	"strings"

	"github.com/clindet/bget/urlpool"
	"github.com/openbiox/ligo/archive"
	cio "github.com/openbiox/ligo/io"
	cnet "github.com/openbiox/ligo/net"
	"github.com/spf13/cobra"
)

var URLCmd = &cobra.Command{
	Use:   "url [url1 url2 url3...]",
	Short: "Can be used to access URLs via Golang http, wget, curl, axel and git, and rsync.",
	Long:  `Can be used to access URLs via Golang http, wget, curl, axel and git, and rsync. More see here https://github.com/clindet/bget.`,
	Run: func(cmd *cobra.Command, args []string) {
		URLCmdRunOptions(cmd, args)
	},
}

func URLCmdRunOptions(cmd *cobra.Command, args []string) {
	initCmd(cmd, args)
	checkArgs(cmd, "url")
	checkDownloadDir(bgetClis.URLs != "" || bgetClis.ListFile != "")
	if (bgetClis.URLs != "" || bgetClis.ListFile != "") && !bgetClis.GitHubMode {
		downloadUrls()
		bgetClis.HelpFlags = false
	} else if (bgetClis.URLs != "" || bgetClis.ListFile != "") && bgetClis.GitHubMode {
		bgetClis.GitHub = bgetClis.URLs
		downloadGitHubRepos()
		bgetClis.HelpFlags = false
	}
	if bgetClis.HelpFlags {
		cmd.Help()
	}
}

func downloadUrls() {
	urls := []string{}
	if bgetClis.URLs != "" && strings.Contains(bgetClis.URLs, bgetClis.Seperator) {
		urls = strings.Split(bgetClis.URLs, bgetClis.Seperator)
	} else if bgetClis.URLs != "" {
		urls = []string{bgetClis.URLs}
	} else if bgetClis.ListFile != "" {
		urls = cio.ReadLines(bgetClis.ListFile)
	}
	var destDirArray []string
	for i := range urls {
		urls[i] = strings.TrimSpace(urls[i])
		if !strings.Contains(urls[i], "://") && !strings.Contains(urls[i], "git@") {
			urls[i] = "https://" + urls[i]
		}
		destDirArray = append(destDirArray, bgetClis.DownloadDir)
	}

	netOpt := setNetParams(&bgetClis)
	done := cnet.HTTPGetURLs(urls, destDirArray, netOpt)
	for _, dest := range done {
		if bgetClis.Uncompress {
			if err := archive.UnarchiveLog(dest, path.Dir(dest)); err != nil {
				log.Warn(err)
			}
		}
	}
}

func downloadGitHubRepos() {
	urls := []string{}
	if bgetClis.GitHub != "" && strings.Contains(bgetClis.GitHub, bgetClis.Seperator) {
		urls = strings.Split(bgetClis.GitHub, bgetClis.Seperator)
	} else if bgetClis.GitHub != "" {
		urls = []string{bgetClis.GitHub}
	} else if bgetClis.ListFile != "" {
		urls = cio.ReadLines(bgetClis.ListFile)
	}
	var destDirArray []string
	assetsUrls := make(map[string][]string)
	versIdx := strings.Split(bgetClis.WithAssetsVersions, bgetClis.Seperator)
	for i := range urls {
		urls[i] = strings.TrimSpace(urls[i])
		urls[i] = strings.ReplaceAll(urls[i], "git@github.com:", "")
		urls[i] = strings.ReplaceAll(urls[i], "https://github.com/", "")
		urls[i] = strings.ReplaceAll(urls[i], "http://github.com/", "")
		key := urls[i]
		urls[i] = "https://github.com/" + urls[i]

		if bgetClis.WithAssets || bgetClis.OnlyAssets {
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
		destDirArray = append(destDirArray, bgetClis.DownloadDir)
	}
	netOpt := setNetParams(&bgetClis)
	if !bgetClis.OnlyAssets {
		cnet.HTTPGetURLs(urls, destDirArray, netOpt)
	}
	urls2 := []string{}
	destDirArray2 := []string{}
	for k, v := range assetsUrls {
		urls2 = append(urls2, v...)
		for range v {
			destDirArray2 = append(destDirArray2, path.Join(bgetClis.DownloadDir, "github-assets", strings.ReplaceAll(k, "https://github.com/", "")))
		}
	}
	netOpt = setNetParams(&bgetClis)
	done := cnet.HTTPGetURLs(urls2, destDirArray2, netOpt)
	for _, dest := range done {
		if bgetClis.Uncompress {
			if err := archive.UnarchiveLog(dest, path.Dir(dest)); err != nil {
				log.Warn(err)
			}
		}
	}
}

func init() {
	setGlobalFlag(URLCmd, &bgetClis)
	setKeyListFlag(URLCmd, &bgetClis, "urls")
	setUncompressFlag(URLCmd, &bgetClis)
	URLCmd.Flags().BoolVarP(&(bgetClis.GitHubMode), "github", "", false, "GitHub mode.")
	URLCmd.Flags().BoolVarP(&(bgetClis.OnlyAssets), "only-github-assets", "", false, "Logical indicating that whether to only download github repo assets files.")
	URLCmd.Flags().BoolVarP(&(bgetClis.WithAssets), "with-github-assets", "", false, "Logical indicating that whether to download associated assets files of github repo.")
	URLCmd.Flags().StringVarP(&(bgetClis.WithAssetsVersions), "github-assets-versions", "", "", "Required to get specific tagname of github assets (e.g. v2.7.1,v1.0.0).")
	URLCmd.Example = `  urls="https://dldir1.qq.com/weixin/Windows/WeChatSetup.exe,http://download.oray.com/pgy/windows/PgyVPN_4.1.0.21693.exe,https://dldir1.qq.com/qqfile/qq/PCQQ9.1.6/25786/QQ9.1.6.25786.exe" && echo $urls | tr "," "\n"> /tmp/urls.list

  bget url ${urls}
  bget url https://dldir1.qq.com/weixin/Windows/WeChatSetup.exe https://dldir1.qq.com/qqfile/qq/PCQQ9.1.6/25786/QQ9.1.6.25786.exe --save-log
  bget url ${urls} -t 3 -o /tmp/download -f -g wget --save-log --verbose 2
  bget url ${urls} -t 2 -o /tmp/download --save-log --verbose 2
	
  bget url ${urls} -t 3 -o /tmp/download -g wget --ignore
  bget url -l /tmp/urls.list -o /tmp/download -f -t 3

  bget url Miachol/github_demo --github
  bget url PapenfussLab/gridss clindet/bget --with-github-assets -t 5 --github
  bget url PapenfussLab/gridss clindet/bget --only-github-assets -t 5 --github
  bget url PapenfussLab/gridss clindet/bget --with-github-assets --with-assets-versions v2.7.2,v0.1.3 -t 5 --github`
}
