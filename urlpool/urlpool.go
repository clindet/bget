package urlpool

import (
	"fmt"
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
	Site         string
	Versions     []string
	VersionsAPI  string
	URL          map[string][]string
	PostShellCmd []string
}

type bgetFilesURLType struct {
	Name         string
	Site         string
	URL          []string
	Versions     []string
	PostShellCmd []string
}

// BgetToolsPool an object bioinformatics tools URL
var BgetToolsPool = []bgetToolsURLType{
	{Name: "miniconda2", URL: map[string][]string{
		"Linux": []string{"https://repo.anaconda.com/miniconda/Miniconda2-{{version}}-Linux-x86_64.sh"},
		"Mac":   []string{"https://repo.anaconda.com/miniconda/Miniconda2-{{version}}-MacOSX-x86_64.sh"},
		"Win":   []string{"https://repo.anaconda.com/miniconda/Miniconda2-{version}}-Windows-x86_64.exe"}},
		PostShellCmd: []string{"cd {{pdir}} && sh {{dest}} -b -p {{downloadDir}}/miniconda2"}},
	{Name: "miniconda3", URL: map[string][]string{
		"Linux": []string{"https://repo.anaconda.com/miniconda/Miniconda3-{{version}}-Linux-x86_64.sh"},
		"Mac":   []string{"https://repo.anaconda.com/miniconda/Miniconda3-{{version}}-MacOSX-x86_64.sh"},
		"Win":   []string{"https://repo.anaconda.com/miniconda/Miniconda3-{version}}-Windows-x86_64.exe"}},
		PostShellCmd: []string{"cd {{pdir}} && sh {{dest}} -b -p {{downloadDir}}/miniconda3"}},
	{Name: "gdc-client",
		Site: "github",
		URL: map[string][]string{
			"Linux": []string{"https://github.com/NCI-GDC/gdc-client/releases/download/{{version}}/gdc-client_v{{version}}_Ubuntu_x64.zip"},
			"Mac":   []string{"https://github.com/NCI-GDC/gdc-client/releases/download/{{version}}/gdc-client_v{{version}}_OSX_x64_10.12.6.zip,https://github.com/NCI-GDC/gdc-client/releases/download/v{{version}}/gdc-client_v{{version}}_OSX_x64.zip"},
			"Win":   []string{"https://github.com/NCI-GDC/gdc-client/releases/download/{{version}}/gdc-client_v{{version}}_Windows_x64.zip"},
		}},
}

// BgetFilesPool an object bioinformatics files URL
var BgetFilesPool = []bgetFilesURLType{
	{
		Name: "reffa",
		Site: "ucsc",
		URL:  []string{"https://hgdownload.cse.ucsc.edu/goldenPath/{{version}}/bigZips/{{version}}.fa.gz"},
	},
	{
		Name: "reffa",
		Site: "genecode",
		URL: []string{"http://ftp.ebi.ac.uk/pub/databases/gencode/Gencode_human/release_{{release}}/{{version}}.p12.genome.fa.gz",
			"http://ftp.ebi.ac.uk/pub/databases/gencode/Gencode_human/release_{{release}}/{{version}}_mapping/{{version}}.primary_assembly.genome.fa.gz"},
	},
	{
		Name: "reffa",
		Site: "ensemble",
		URL: []string{"http://ftp.ensembl.org/pub/release-{{release}}/fasta/homo_sapiens/dna/Homo_sapiens.{{version}}.dna.primary_assembly.fa.gz",
			"http://ftp.ensembl.org/pub/release-{{release}}/fasta/homo_sapiens/dna/Homo_sapiens.{{version}}.{{release}}.dna.primary_assembly.fa.gz"},
	},
	{
		Name: "reffa",
		Site: "fusioncatcher",
		URL: []string{"https://svwh.dl.sourceforge.net/project/fusioncatcher/data//human_v{{release}}.tar.gz.aa",
			"https://svwh.dl.sourceforge.net/project/fusioncatcher/data/human_v{{release}}.tar.gz.ab",
			"https://svwh.dl.sourceforge.net/project/fusioncatcher/data/human_v{{release}}.tar.gz.ac",
			"https://svwh.dl.sourceforge.net/project/fusioncatcher/data/human_v{{release}}.tar.gz.ad",
			"https://svwh.dl.sourceforge.net/project/fusioncatcher/data/human_v{{release}}.tar.gz.ae"},
	},
	{
		Name: "reffa",
		Site: "defuse",
		URL: []string{"http://ftp.ensembl.org/pub/release-{{release}}/fasta/homo_sapiens/dna/Homo_sapiens.{{version}}.dna.chromosome.{{chrom}}.fa.gz",
			"http://ftp.ensembl.org/pub/release-{{release}}/gtf/homo_sapiens/Homo_sapiens.{{version}}.{{release}}.gtf.gz",
			"http://hgdownload.cse.ucsc.edu/goldenPath/{{version}}/database/rmsk.txt.gz",
			"http://hgdownload.cse.ucsc.edu/goldenPath/{{version}}/bigZips/est.fa.gz",
			"http://hgdownload.cse.ucsc.edu/goldenPath/{{version}}/database/intronEst.txt.gz",
			"http://ftp.ncbi.nlm.nih.gov/repository/UniGene/Homo_sapiens/Hs.seq.uniq.gz",
			"http://hgdownload.cse.ucsc.edu/goldenPath/{{version}}/database/rmsk.txt.gz"},
	},
}

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
		if BgetFilesPool[f].Name == name && ((*env)["site"] == "" ||
			BgetFilesPool[f].Site == (*env)["site"]) {
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
	BgetFilesPool = append(BgetFilesPool, githubRepos...)
	BgetFilesPool = append(BgetFilesPool, journalsMeta...)
	BgetFilesPool = append(BgetFilesPool, annovarLinks...)
	BgetFilesPool = append(BgetFilesPool, otherFiles...)
}
