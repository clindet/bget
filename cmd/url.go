package cmd

import (
	"strconv"
	"strings"

	butils "github.com/JhuangLab/butils"
	"github.com/spf13/cobra"
)

type envToolsURLType struct {
	Name    string
	Version string
	URL     map[string]string
}

type envFilesURLType struct {
	Name string
	Site string
	URL  []string
}

var urlCmd = &cobra.Command{
	Use:   "url [url1 url2 url3...]",
	Short: "Can be used to access URLs via Golang http, wget, curl, axel and git, and rsync.",
	Long:  `Can be used to access URLs via Golang http, wget, curl, axel and git, and rsync. More see here https://github.com/JhuangLab/bget.`,
	Run: func(cmd *cobra.Command, args []string) {
		urlCmdRunOptions(cmd)
	},
}

var bgetToolsURLs = []envToolsURLType{
	{Name: "spack", URL: map[string]string{
		"Linux": "https://github.com/spack/spack",
		"Mac":   "https://github.com/spack/spack",
		"Win":   "https://github.com/spack/spack",
	}},
	{Name: "miniconda2", URL: map[string]string{
		"Linux": "https://repo.anaconda.com/miniconda/Miniconda2-{{version}}-Linux-x86_64.sh",
		"Mac":   "https://repo.anaconda.com/miniconda/Miniconda2-{{version}}-MacOSX-x86_64.sh",
		"Win":   "https://repo.anaconda.com/miniconda/Miniconda2-{version}}-Windows-x86_64.exe",
	}},
	{Name: "miniconda3", URL: map[string]string{
		"Linux": "https://repo.anaconda.com/miniconda/Miniconda3-{{version}}-Linux-x86_64.sh",
		"Mac":   "https://repo.anaconda.com/miniconda/Miniconda3-{{version}}-MacOSX-x86_64.sh",
		"Win":   "https://repo.anaconda.com/miniconda/Miniconda3-{version}}-Windows-x86_64.exe",
	}},
}

var bgetFilesURLs = []envFilesURLType{
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
		URL: []string{"https://sourceforge.net/projects/fusioncatcher/files/data/human_v{{release}}.tar.gz.aa",
			"https://sourceforge.net/projects/fusioncatcher/files/data/human_v{{release}}.tar.gz.ab",
			"https://sourceforge.net/projects/fusioncatcher/files/data/human_v{{release}}.tar.gz.ac",
			"https://sourceforge.net/projects/fusioncatcher/files/data/human_v{{release}}.tar.gz.ad",
			"https://sourceforge.net/projects/fusioncatcher/files/data/human_v{{release}}.tar.gz.ae"},
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
	{
		Name: "bwa",
		Site: "github",
		URL:  []string{"https://github.com/lh3/bwa"},
	},
}

func getEnvToolsURL(name string, version string, release string) string {
	tmp := ""
	for i := range bgetToolsURLs {
		if bgetToolsURLs[i].Name == name {
			if osType == "linux" {
				tmp = strings.Replace(bgetToolsURLs[i].URL["Linux"], "{{version}}", version, 100)
				tmp = strings.Replace(tmp, "{{release}}", release, 100)
			} else if osType == "windows" {
				tmp = strings.Replace(bgetToolsURLs[i].URL["Win"], "{{version}}", version, 100)
				tmp = strings.Replace(tmp, "{{release}}", release, 100)
			} else {
				tmp = strings.Replace(bgetToolsURLs[i].URL["Mac"], "{{version}}", version, 100)
				tmp = strings.Replace(tmp, "{{release}}", release, 100)
			}
			return tmp
		}
	}
	return tmp
}

func getEnvFilesURL(name string, site string, version string, release string) []string {
	chrom := []string{}
	for i := 1; i < 23; i++ {
		chrom = append(chrom, strconv.Itoa(i))
	}
	chrom = append(chrom, "X", "Y", "MT")
	for i := range bgetFilesURLs {
		if bgetFilesURLs[i].Name == name && (site == "" || bgetFilesURLs[i].Site == site) {
			rep := []string{}
			for j := range bgetFilesURLs[i].URL {
				tmp := strings.Replace(bgetFilesURLs[i].URL[j], "{{release}}", release, 100)
				version = genomeVersionConvertor(tmp, version)
				tmp = strings.Replace(tmp, "{{version}}", version, 100)
				if butils.StrDetect(tmp, "{{chrom}}") {
					raw := tmp
					for k := range chrom {
						tmp = strings.Replace(raw, "{{chrom}}", chrom[k], 100)
						rep = append(rep, tmp)
					}
				}
				if !butils.StrDetect(tmp, "{{chrom}}") {
					rep = append(rep, tmp)
				}
			}
			return rep
		}
	}
	return nil
}

func genomeVersionConvertor(url string, version string) string {
	if butils.StrDetect(url, "http://hgdownload.cse.ucsc.edu/goldenPath") {
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

func downloadUrls() {
	urls := []string{}
	if bgetClis.urls != "" && strings.Contains(bgetClis.urls, bgetClis.separator) {
		urls = strings.Split(bgetClis.urls, bgetClis.separator)
	} else if bgetClis.urls != "" {
		urls = []string{bgetClis.urls}
	} else if bgetClis.urlsFile != "" {
		urls = butils.ReadLines(bgetClis.urlsFile)
	}
	var destDirArray []string
	for i := range urls {
		urls[i] = strings.TrimSpace(urls[i])
		destDirArray = append(destDirArray, bgetClis.downloadDir)
	}

	HTTPGetURLs(urls, destDirArray, bgetClis.engine, taskID, bgetClis.mirror,
		bgetClis.concurrency, bgetClis.axelThread, overwrite, bgetClis.ignore, quiet, saveLog)
}
func urlCmdRunOptions(cmd *cobra.Command) {
	checkQuiet()
	items := []string{}
	if len(cmd.Flags().Args()) >= 1 {
		items = append(items, cmd.Flags().Args()...)
		bgetClis.urls = strings.Join(items, bgetClis.separator)
	}
	checkDownloadDir(bgetClis.urls != "" || bgetClis.urlsFile != "")

	if bgetClis.urls != "" || bgetClis.urlsFile != "" {
		downloadUrls()
		bgetClis.helpFlags = false
	}
	if bgetClis.helpFlags {
		cmd.Help()
	}
}
func init() {
	urlCmd.Flags().StringVarP(&(bgetClis.urlsFile), "list-file", "l", "", "A file contains URLs for download.")
	urlCmd.Example = `  urls="https://dldir1.qq.com/weixin/Windows/WeChatSetup.exe,http://download.oray.com/pgy/windows/PgyVPN_4.1.0.21693.exe,https://dldir1.qq.com/qqfile/qq/PCQQ9.1.6/25786/QQ9.1.6.25786.exe" && echo $urls | tr "," "\n"> /tmp/urls.list

  bget url ${urls}
  bget url https://dldir1.qq.com/weixin/Windows/WeChatSetup.exe https://dldir1.qq.com/qqfile/qq/PCQQ9.1.6/25786/QQ9.1.6.25786.exe
  bget url ${urls} -t 2 -o /tmp/download
  bget url ${urls} -t 3 -o /tmp/download -f -g wget
  bget url ${urls} -t 3 -o /tmp/download -g wget --ignore
  bget url -l /tmp/urls.list -o /tmp/download -f -t 3`
}
