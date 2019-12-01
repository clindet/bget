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

var gitHubCmd = &cobra.Command{
	Use:   "github [user1/repo1 user1/repo2 user2/repo1...]",
	Short: "Can be used to access GitHub repo and assets.",
	Long:  `Can be used to access GitHub repo and assets. More see here https://github.com/openbiox/bget.`,
	Run: func(cmd *cobra.Command, args []string) {
		gitHubCmdRunOptions(cmd)
	},
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
		key := urls[i]
		urls[i] = "https://github.com/" + urls[i]
		if bgetClis.withAssets {
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
	cnet.HttpGetURLs(urls, destDirArray, netOpt)
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
func gitHubCmdRunOptions(cmd *cobra.Command) {
	checkArgs(cmd, "github")
	checkDownloadDir(bgetClis.github != "" || bgetClis.listFile != "")
	if bgetClis.github != "" || bgetClis.listFile != "" {
		downloadGitHubRepos()
		bgetClis.helpFlags = false
	}
	if bgetClis.helpFlags {
		cmd.Help()
	}
}
func init() {
	gitHubCmd.Flags().BoolVarP(&(bgetClis.withAssets), "with-assets", "", false, "Logical indicating that whether to download associated assets files.")
	gitHubCmd.Flags().StringVarP(&(bgetClis.withAssetsVersions), "with-assets-versions", "", "", "Required to get specific tagname of github assets (e.g. v2.7.1,v1.0.0).")
	gitHubCmd.Example = `  bget github Miachol/github_demo
  bget github PapenfussLab/gridss openbiox/bget --with-assets -t 5
  bget github PapenfussLab/gridss openbiox/bget --with-assets --with-assets-versions v2.7.2,v0.1.3 -t 5 `
}
