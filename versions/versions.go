package versions

import (
	"context"
	"encoding/json"
	"fmt"
	neturl "net/url"
	"os"
	"strings"
	"sync"

	"github.com/Miachol/bget/urlpool"
	"github.com/google/go-github/v27/github"
	"github.com/olekukonko/tablewriter"
	log "github.com/openbiox/butils/log"
	"github.com/openbiox/butils/stringo"
	"golang.org/x/oauth2"
)

// DefaultVersions set default key string versions
func DefaultVersions(key string, env *map[string]string) {
	if (*env)["version"] == "" && strings.Contains(key, "miniconda") {
		(*env)["version"] = "latest"
	}
	if key == "reffa" && (*env)["site"] == "" {
		(*env)["site"] = "genecode"
	}
	if key == "reffa" && (*env)["version"] == "" &&
		((*env)["site"] == "genecode" || (*env)["site"] == "ensemble" ||
			(*env)["site"] == "defuse") {
		(*env)["version"] = "GRCh38"
	} else if key == "reffa" && (*env)["version"] == "" {
		(*env)["version"] = "hg38"
	}
	if key == "reffa" && (*env)["site"] == "genecode" && (*env)["release"] == "" {
		(*env)["release"] = "31"
	} else if key == "reffa" && (*env)["site"] == "ensemble" && (*env)["release"] == "" {
		(*env)["release"] = "97"
	}
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
func QueryKeysInfo(keys []string, env *map[string]string) (urls, postShellCmd, vers map[string][]string) {
	urls = make(map[string][]string)
	postShellCmd = make(map[string][]string)
	vers = make(map[string][]string)
	for k := range keys {
		key, version, site, release := ParseMeta(keys[k])
		if (*env)["version"] == "" {
			(*env)["version"] = version
		}
		if (*env)["site"] == "" {
			(*env)["site"] = site
		}
		if (*env)["release"] == "" {
			(*env)["release"] = release
		}
		DefaultVersions(key, env)
		tmp, tmp2, tmp3 := urlpool.QueryBgetTools(key, env)
		if len(tmp) > 0 {
			if (*env)["version"] == "" {
				versions := GitHubVersionSpider(tmp[0])
				if len(versions) > 1 {
					(*env)["version"] = versions[0]
				}
			}
			urls[key] = append(urls[key], tmp...)
			postShellCmd[key] = append(postShellCmd[key], tmp2...)
			vers[key] = append(vers[key], tmp3...)
		}
		tmp, tmp2, tmp3 = urlpool.QueryBgetFiles(key, env)
		if len(tmp) > 0 {
			if (*env)["version"] == "" {
				versions := GitHubVersionSpider(tmp[0])
				if len(versions) > 1 {
					(*env)["version"] = versions[0]
				}
			}
			urls[key] = append(urls[key], tmp...)
			postShellCmd[key] = append(postShellCmd[key], tmp2...)
			vers[key] = append(vers[key], tmp3...)
		}
	}
	return urls, postShellCmd, vers
}

// QueryKeysVersions get keys versions
func QueryKeysVersions(keys []string, env *map[string]string) map[string][]string {
	versions := make(map[string][]string)
	table := tablewriter.NewWriter(os.Stdout)
	if (*env)["printFormat"] == "table" {
		table.SetHeader([]string{"Key", "Versions"})
		table.SetRowLine(true)
		table.SetRowSeparator("-")
		table.SetAlignment(tablewriter.ALIGN_LEFT)
	}
	wg := sync.WaitGroup{}
	for i := range keys {
		wg.Add(1)
		urls, _, vers := QueryKeysInfo([]string{keys[i]}, env)
		key, _, _, _ := ParseMeta(keys[i])
		if len(vers[key]) > 0 {
			versions[key] = vers[key]
			wg.Done()
		} else {
			url := urls[key][0]
			go func(url string) {
				if tmp := GitHubVersionSpider(url); len(tmp) > 0 {
					versions[key] = tmp
				} else if tmp := KeyFixedVersions(key); len(tmp) > 0 {
					versions[key] = tmp
				}
				wg.Done()
			}(url)
		}
	}
	wg.Wait()
	for k := range versions {
		if len(versions[k]) > 0 {
			if (*env)["printFormat"] == "table" {
				table.Append([]string{k, strings.Join(versions[k], ", ")})
			} else if (*env)["printFormat"] == "txt" {
				fmt.Println(fmt.Sprintf("key> %s\nversions> %s\n-----------", k, strings.Join(versions[k], ", ")))
			}
		}
	}
	if (*env)["printFormat"] == "table" {
		table.Render()
	} else if (*env)["printFormat"] == "json" {
		mapVersions, _ := json.Marshal(versions)
		fmt.Println(string(mapVersions))
	}
	return versions
}

// ParseMeta parse @version %site #release format string
func ParseMeta(key string) (keyNew string, version string, site string, release string) {
	info := stringo.StrSplit(key, "@|%|#", 4)
	info1 := stringo.StrSplit(key, "@", 2)
	info2 := stringo.StrSplit(key, "%", 2)
	info3 := stringo.StrSplit(key, "#", 2)
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
