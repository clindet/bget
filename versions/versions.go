package versions

import (
	"context"
	"encoding/json"
	"fmt"
	neturl "net/url"
	"os"
	"strings"
	"sync"

	"github.com/JhuangLab/bget/urlpool"
	butils "github.com/JhuangLab/butils"
	log "github.com/JhuangLab/butils/log"
	"github.com/google/go-github/v27/github"
	"github.com/olekukonko/tablewriter"
	"golang.org/x/oauth2"
)

// DefaultVersions set default key string versions
func DefaultVersions(key string, version string, release string, site string, osType string) (string, string, string) {
	if version == "" && strings.Contains(key, "miniconda") {
		version = "latest"
	}
	if key == "reffa" && site == "" {
		site = "genecode"
	}
	if key == "reffa" && version == "" && (site == "genecode" || site == "ensemble" || site == "defuse") {
		version = "GRCh38"
	} else if key == "reffa" && version == "" {
		version = "hg38"
	}
	if key == "reffa" && site == "genecode" && release == "" {
		release = "31"
	} else if key == "reffa" && site == "ensemble" && release == "" {
		release = "97"
	}
	return version, site, release
}

// GitHubVersionSpider get all tags and branch
func GitHubVersionSpider(url string) (versions []string) {
	accessToken := "4d00c84fa5da085df3f1bb6a6a7f6dd0972e869f"
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
	vers, _, err := client.Repositories.ListReleases(ctx, user, repo, opt)
	if err != nil {
		log.Fatal(err)
	}
	brchs, _, err := client.Repositories.ListBranches(ctx, user, repo, opt)
	if err != nil {
		log.Fatal(err)
	}
	for i := range vers {
		versions = append(versions, vers[i].GetTagName())
	}
	for i := range brchs {
		versions = append(versions, brchs[i].GetName())
	}
	return versions
}

// KeyFixedVersions for fixed versions of key
func KeyFixedVersions(key string) []string {
	vers := make(map[string][]string)
	vers = map[string][]string{"reffa": []string{"GRCh38 %genecode #31", "GRCh37 %genecode #31", "hg38 %ucsc", "hg19 %ucsc", "GRCh38 %ensemble #97", "GRCh38 %defuse #97",
		"%fusioncatcher #95"}}
	for k := range vers {
		if k == key {
			return vers[key]
		}
	}
	return []string{}
}

// QueryKeysInfo get keys URL and post shell command
func QueryKeysInfo(keys []string, osType string) (urls map[string][]string, postShellCmd map[string][]string) {
	version := ""
	site := ""
	release := ""
	key := ""
	urls = make(map[string][]string)
	postShellCmd = make(map[string][]string)
	for k := range keys {
		key, version, site, release = ParseMeta(keys[k])
		version, site, release = DefaultVersions(key, version, release, site, osType)
		if tmp, tmp2 := urlpool.QueryBgetTools(key, version, release, osType); len(tmp) > 0 {
			if version == "" {
				versions := GitHubVersionSpider(tmp[0])
				if len(versions) > 1 {
					version = versions[0]
				}
				tmp, tmp2 = urlpool.QueryBgetTools(key, version, release, osType)
			}
			urls[key] = append(urls[key], tmp...)
			postShellCmd[key] = append(postShellCmd[key], tmp2...)
		}

		if tmp, tmp2 := urlpool.QueryBgetFiles(key, version, release, site); len(tmp) > 0 {
			if version == "" {
				versions := GitHubVersionSpider(tmp[0])
				if len(versions) > 1 {
					version = versions[0]
				}
				tmp, tmp2 = urlpool.QueryBgetFiles(key, version, release, site)
			}
			urls[key] = append(urls[key], tmp...)
			postShellCmd[key] = append(postShellCmd[key], tmp2...)
		}
	}
	return urls, postShellCmd
}

// QueryKeysVersions get keys versions
func QueryKeysVersions(keys []string, osType string, printFormat string) map[string][]string {
	versions := make(map[string][]string)
	wg := sync.WaitGroup{}
	table := tablewriter.NewWriter(os.Stdout)
	if printFormat == "table" {
		table.SetHeader([]string{"Key", "Versions"})
		table.SetRowLine(true)
		table.SetRowSeparator("-")
		table.SetAlignment(tablewriter.ALIGN_LEFT)
	}

	for i := range keys {
		urls, _ := QueryKeysInfo([]string{keys[i]}, osType)
		for _, url := range urls {
			for j := range url {
				key, _, _, _ := ParseMeta(keys[i])
				wg.Add(1)
				go func(i int, j int, key string) {
					if tmp := GitHubVersionSpider(url[j]); len(tmp) > 0 {
						versions[key] = tmp
					} else if tmp := KeyFixedVersions(key); len(tmp) > 0 {
						versions[key] = tmp
					}
					wg.Done()
				}(i, j, key)
			}
		}
	}
	wg.Wait()
	for k := range versions {
		if len(versions[k]) > 0 {
			if printFormat == "table" {
				table.Append([]string{k, strings.Join(versions[k], ", @")})
			} else if printFormat == "txt" {
				fmt.Println(fmt.Sprintf("key> %s\n@%s\n-----------", k, strings.Join(versions[k], ", @")))
			}
		}
	}
	if printFormat == "table" {
		table.Render()
	} else if printFormat == "json" {
		mapVersions, _ := json.Marshal(versions)
		fmt.Println(string(mapVersions))
	}
	return versions
}

// ParseMeta parse @version %site #release format string
func ParseMeta(key string) (keyNew string, version string, site string, release string) {
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
