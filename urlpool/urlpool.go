package urlpool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"

	"context"
	"net/http"
	neturl "net/url"

	"github.com/openbiox/ligo/stringo"

	"github.com/google/go-github/v27/github"
	glog "github.com/openbiox/ligo/log"
	"golang.org/x/oauth2"
)

var log = glog.Logger

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
	VersionsAPI  string
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
		if strings.ReplaceAll(strings.ToLower(BgetToolsPool[i].Name), "_", "-") == name {
			if BgetToolsPool[i].VersionsAPI != "" && strings.Contains(BgetToolsPool[i].VersionsAPI, "://github.com") {
				versions = GitHubVersionSpider(BgetToolsPool[i].VersionsAPI, true)
			} else if BgetToolsPool[i].VersionsAPI != "" && strings.Contains(BgetToolsPool[i].VersionsAPI, "://bitbucket.org") {
				versions = BitbucketVersionSpider(BgetToolsPool[i].VersionsAPI)
			} else {
				versions = BgetToolsPool[i].Versions
			}
			if (*env)["version"] == "" && len(versions) > 0 {
				(*env)["version"] = versions[0]
			}
			tmpUrls := []string{}
			for k, v := range *env {
				kstr := fmt.Sprintf("{{%s}}", k)
				for j := range BgetToolsPool[i].URL[ostype] {
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
		}
		if len(urls) > 0 {
			break
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
		if strings.ReplaceAll(strings.ToLower(BgetFilesPool[f].Name), "_", "-") == name {
			if BgetFilesPool[f].VersionsAPI != "" && strings.Contains(BgetFilesPool[f].VersionsAPI, "github.com") {
				versions = GitHubVersionSpider(BgetToolsPool[f].VersionsAPI, true)
			} else if BgetFilesPool[f].VersionsAPI != "" && strings.Contains(BgetFilesPool[f].VersionsAPI, "bitbucket.org") {
				versions = BitbucketVersionSpider(BgetFilesPool[f].VersionsAPI)
			} else if strings.Contains(BgetFilesPool[f].URL[0], "github.com") {
				versions = GitHubVersionSpider(BgetFilesPool[f].URL[0], true)
			} else if strings.Contains(BgetFilesPool[f].URL[0], "bitbucket.org") {
				versions = BitbucketVersionSpider(BgetFilesPool[f].URL[0])
			} else {
				versions = BgetFilesPool[f].Versions
			}
			if (*env)["version"] == "" && len(versions) > 0 {
				(*env)["version"] = versions[0]
			}
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
					for k2 := range tmpSlice {
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
		}
		if len(urls) > 0 {
			break
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

func setGitHubCtx(url string) (user, repo string, ctx context.Context, client *github.Client, opt *github.ListOptions) {
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
	user, repo = pathStr[1], pathStr[2]
	ctx = context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client = github.NewClient(tc)
	opt = &github.ListOptions{}
	return user, repo, ctx, client, opt
}

// GitHubVersionSpider get all tags and branch
func GitHubVersionSpider(url string, includeBranches bool) (versions []string) {
	user, repo, ctx, client, opt := setGitHubCtx(url)
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
	if includeBranches {
		for i := range brchs {
			versions = append(versions, brchs[i].GetName())
		}
	}
	return versions
}

func GitHubAssetsSpider(url, version string) (urls []string) {
	user, repo, ctx, client, opt := setGitHubCtx(url)
	rels, _, err := client.Repositories.ListReleases(ctx, user, repo, opt)
	if err != nil {
		log.Fatal(err)
	}
	for i := range rels {
		if rels[i].GetTagName() == version {
			assets, _, err := client.Repositories.ListReleaseAssets(ctx, user, repo, rels[i].GetID(), opt)
			if err != nil {
				log.Fatal(err)
			}
			for j := range assets {
				urls = append(urls, assets[j].GetBrowserDownloadURL())
			}
			return urls
		}
	}
	return urls
}

// BitbucketVersionSpider query Bitbucket versions
func BitbucketVersionSpider(url string) (versions []string) {
	u, err := neturl.Parse(url)
	if err != nil {
		log.Fatal(err)
	}
	if u.Host != "bitbucket.org" {
		return
	}
	pathStr := strings.Split(u.Path, "/")
	user, repo := pathStr[1], pathStr[2]
	tagsBody := bitbucketRepoAPI(user, repo, "tags")
	var s BitbucketObj
	json.Unmarshal(tagsBody, &s)
	for i := range s.Values {
		versions = append(versions, s.Values[i].Name)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(versions)))
	brcsBody := bitbucketRepoAPI(user, repo, "branches")
	json.Unmarshal(brcsBody, &s)
	for i := range s.Values {
		versions = append(versions, s.Values[i].Name)
	}
	return versions
}

func bitbucketRepoAPI(user, repo, entry string) []byte {
	url := fmt.Sprintf("https://api.bitbucket.org/2.0/repositories/%s/%s/refs/%s", user, repo, entry)
	resp, err := http.Get(url)
	if err != nil {
		log.Warn(err)
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Warn(err)
		return nil
	}
	return []byte(string(body))
}

func bitbucketBranches(user, repo string) []byte {
	url := fmt.Sprintf("https://api.bitbucket.org/2.0/repositories/%s/%s/refs/tags", user, repo)
	resp, err := http.Get(url)
	if err != nil {
		log.Warn(err)
		return nil
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Warn(err)
		return nil
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Warn(err)
		return nil
	}
	return body
}

func init() {
	BgetToolsPool = append(BgetToolsPool, toolsLinks...)
	BgetFilesPool = append(BgetFilesPool, reffaFiles...)
	BgetFilesPool = append(BgetFilesPool, githubRepos...)
	BgetFilesPool = append(BgetFilesPool, otherGitRepos...)
	BgetFilesPool = append(BgetFilesPool, journalsMeta...)
	BgetFilesPool = append(BgetFilesPool, annovarLinks...)
	BgetFilesPool = append(BgetFilesPool, wkflFiles...)
	BgetFilesPool = append(BgetFilesPool, otherFiles...)
}
