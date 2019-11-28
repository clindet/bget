package urlpool

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"context"
	neturl "net/url"

	"github.com/openbiox/butils/stringo"

	"github.com/google/go-github/v27/github"
	log "github.com/openbiox/butils/log"
	"golang.org/x/oauth2"
)

type bgetToolsURLType struct {
	Name         string
	Description  string
	Versions     []string
	VersionsAPI  string
	Tags         []string
	URL          map[string][]string
	PostShellCmd []string
}

type bgetFilesURLType struct {
	Name         string
	Description  string
	URL          []string
	Versions     []string
	Tags         []string
	PostShellCmd []string
}

// BgetToolsPool an object bioinformatics tools URL
var BgetToolsPool = []bgetToolsURLType{}

// BgetFilesPool an object bioinformatics files URL
var BgetFilesPool = []bgetFilesURLType{}

func setOsStr(env *map[string]string) (ostype string) {
	if (*env)["osType"] == "linux" {
		ostype = "Linux"
	} else if (*env)["osType"] == "windows" {
		ostype = "windows"
	} else {
		ostype = "Mac"
	}
	return ostype
}
func QueryBgetTools(name string, env *map[string]string) (urls, postShellCmd, versions []string) {
	ostype := setOsStr(env)
	for i := range BgetToolsPool {
		if BgetToolsPool[i].Name == name {
			tmpUrls := []string{}
			for k, v := range *env {
				kstr := fmt.Sprintf("{{%s}}", k)
				for j, _ := range BgetToolsPool[i].URL[ostype] {
					BgetToolsPool[i].URL[ostype][j] = strings.Replace(BgetToolsPool[i].URL[ostype][j], kstr, v, 10000)
				}
				tmpUrls = BgetToolsPool[i].URL[ostype]
			}
			urls = append(urls, tmpUrls...)
			tmp := ""
			for j := range BgetToolsPool[i].PostShellCmd {
				for k, v := range *env {
					kstr := fmt.Sprintf("{{%s}}", k)
					if tmp == "" {
						tmp = strings.Replace(BgetToolsPool[i].PostShellCmd[j], kstr, v, 10000)
					} else {
						tmp = strings.Replace(tmp, kstr, v, 10000)
					}
				}
				postShellCmd = append(postShellCmd, tmp)
			}
			if BgetToolsPool[i].VersionsAPI != "" && strings.Contains(BgetToolsPool[i].VersionsAPI, "github.com") {
				versions = GitHubVersionSpider(BgetToolsPool[i].VersionsAPI)
			} else {
				versions = BgetToolsPool[i].Versions
			}
		}
	}
	return urls, postShellCmd, versions
}

func formatURL(tmp string, key string, rep string, url string) string {
	kstr := fmt.Sprintf("{{%s}}", key)
	if tmp == "" {
		tmp = strings.Replace(url,
			kstr, rep, 10000)
	} else {
		tmp = strings.Replace(tmp,
			kstr, rep, 10000)
	}
	return tmp
}

func formatURLSlice(tmpSlice []string, env *map[string]string) (urls []string) {
	chrom := []string{}
	for i := 1; i < 23; i++ {
		chrom = append(chrom, strconv.Itoa(i))
	}
	chrom = append(chrom, "X", "Y", "MT")
	for _, v := range tmpSlice {
		if stringo.StrDetect(v, "{{chrom}}") {
			raw := v
			for k := range chrom {
				v = strings.Replace(raw, "{{chrom}}", chrom[k], 10000)
				urls = append(urls, v)
			}
		}
		if !stringo.StrDetect(v, "{{chrom}}") {
			urls = append(urls, v)
		}
	}
	return urls
}

func QueryBgetFiles(name string, env *map[string]string) (urls []string, postShellCmd []string, versions []string) {
	for f := range BgetFilesPool {
		if BgetFilesPool[f].Name == name {
			for _, url := range BgetFilesPool[f].URL {
				tmp := ""
				tmpSlice := []string{}
				for k, v := range *env {
					if strings.Contains(v, ",") {
						v = stringo.StrReplaceAll(v, " ", "")
						vSlice := strings.Split(v, ",")
						for _, v2 := range vSlice {
							tmpSlice = append(tmpSlice, formatURL(tmp, k, v2, url))
						}
					} else {
						tmp = formatURL(tmp, k, v, url)
					}
					for k2, _ := range tmpSlice {
						tmpSlice[k2] = formatURL(tmpSlice[k2], k, v, url)
					}
				}
				if len(tmpSlice) == 0 {
					tmpSlice = append(tmpSlice, tmp)
				}
				urls = append(urls, formatURLSlice(tmpSlice, env)...)
			}
			for j := range BgetFilesPool[f].PostShellCmd {
				tmp := ""
				for k, v := range *env {
					kstr := fmt.Sprintf("{{%s}}", k)
					if tmp == "" {
						tmp = strings.Replace(BgetFilesPool[f].PostShellCmd[j],
							kstr, v, 10000)
					} else {
						tmp = strings.Replace(tmp,
							kstr, v, 10000)
					}
				}
				postShellCmd = append(postShellCmd, tmp)
			}
			versions = BgetFilesPool[f].Versions
		}
	}
	return urls, postShellCmd, versions
}

func genomeVersionConvertor(url string, version string) string {
	if stringo.StrDetect(url, "http://hgdownload.cse.ucsc.edu/goldenPath") {
		if strings.ToLower(version) == "grch38" {
			version = "hg38"
		} else if strings.ToLower(version) == "grch37" {
			version = "hg19"
		} else if strings.ToLower(version) == "grcm38" {
			version = "mm10"
		} else if strings.ToLower(version) == "grcm37" {
			version = "mm9"
		}
	}
	return version
}

// GitHubVersionSpider get all tags and branch
func GitHubVersionSpider(url string) (versions []string) {
	accessToken := os.Getenv("GITHUB_TOKEN")
	if accessToken == "" {
		log.Fatal("Please set GITHUB_TOKEN environment variable.")
	}
	u, err := neturl.Parse(url)
	if err != nil {
		log.Fatal(err)
	}
	if u.Host != "github.com" {
		return
	}
	pathStr := strings.Split(u.Path, "/")
	user, repo := pathStr[1], pathStr[2]
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	opt := &github.ListOptions{}
	//version, _, _ := client.Repositories.ListTags(ctx, user, repo, opt)
	vers, _, err := client.Repositories.ListTags(ctx, user, repo, opt)
	if err != nil {
		log.Fatal(err)
	}
	brchs, _, err := client.Repositories.ListBranches(ctx, user, repo, opt)
	if err != nil {
		log.Fatal(err)
	}
	for i := range vers {
		versions = append(versions, vers[i].GetName())
	}
	for i := range brchs {
		versions = append(versions, brchs[i].GetName())
	}
	return versions
}

func init() {
	BgetToolsPool = append(BgetToolsPool, toolsLinks...)
	BgetFilesPool = append(BgetFilesPool, reffaFiles...)
	BgetFilesPool = append(BgetFilesPool, githubRepos...)
	BgetFilesPool = append(BgetFilesPool, journalsMeta...)
	BgetFilesPool = append(BgetFilesPool, annovarLinks...)
	BgetFilesPool = append(BgetFilesPool, wkflFiles...)
	BgetFilesPool = append(BgetFilesPool, otherFiles...)
}
