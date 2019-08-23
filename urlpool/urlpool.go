package urlpool

import (
	"strconv"
	"strings"

	butils "github.com/JhuangLab/butils"
)

type bgetToolsURLType struct {
	Name         string
	Site         string
	Version      string
	URL          map[string]string
	PostShellCmd []string
}

type bgetFilesURLType struct {
	Name         string
	Site         string
	URL          []string
	PostShellCmd []string
}

// BgetToolsPool an object bioinformatics tools URL
var BgetToolsPool = []bgetToolsURLType{
	{Name: "miniconda2", URL: map[string]string{
		"Linux": "https://repo.anaconda.com/miniconda/Miniconda2-{{version}}-Linux-x86_64.sh",
		"Mac":   "https://repo.anaconda.com/miniconda/Miniconda2-{{version}}-MacOSX-x86_64.sh",
		"Win":   "https://repo.anaconda.com/miniconda/Miniconda2-{version}}-Windows-x86_64.exe"},
		PostShellCmd: []string{"cd {{pdir}} && sh {{dest}} -b -p {{downloadDir}}/miniconda2"}},
	{Name: "miniconda3", URL: map[string]string{
		"Linux": "https://repo.anaconda.com/miniconda/Miniconda3-{{version}}-Linux-x86_64.sh",
		"Mac":   "https://repo.anaconda.com/miniconda/Miniconda3-{{version}}-MacOSX-x86_64.sh",
		"Win":   "https://repo.anaconda.com/miniconda/Miniconda3-{version}}-Windows-x86_64.exe"},
		PostShellCmd: []string{"cd {{pdir}} && sh {{dest}} -b -p {{downloadDir}}/miniconda3"}},
	{Name: "gdc-client",
		Site: "github",
		URL: map[string]string{
			"Linux": "https://github.com/NCI-GDC/gdc-client/releases/download/{{version}}/gdc-client_v{{version}}_Ubuntu_x64.zip",
			"Mac":   "https://github.com/NCI-GDC/gdc-client/releases/download/{{version}}/gdc-client_v{{version}}_OSX_x64_10.12.6.zip",
			"Win":   "https://github.com/NCI-GDC/gdc-client/releases/download/{{version}}/gdc-client_v{{version}}_Windows_x64.zip",
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

func QueryBgetTools(name string, version string, release string, osType string) (urls []string, postShellCmd []string) {
	tmp := ""
	for i := range BgetToolsPool {
		if BgetToolsPool[i].Name == name {
			if osType == "linux" {
				tmp = strings.Replace(BgetToolsPool[i].URL["Linux"], "{{version}}", version, 100)
				tmp = strings.Replace(tmp, "{{release}}", release, 100)
			} else if osType == "windows" {
				tmp = strings.Replace(BgetToolsPool[i].URL["Win"], "{{version}}", version, 100)
				tmp = strings.Replace(tmp, "{{release}}", release, 100)
			} else {
				tmp = strings.Replace(BgetToolsPool[i].URL["Mac"], "{{version}}", version, 100)
				tmp = strings.Replace(tmp, "{{release}}", release, 100)
			}
			urls = append(urls, tmp)
			for j := range BgetToolsPool[i].PostShellCmd {
				tmp = strings.Replace(BgetToolsPool[i].PostShellCmd[j], "{{version}}", version, 100)
				tmp = strings.Replace(tmp, "{{release}}", release, 100)
				postShellCmd = append(postShellCmd, tmp)
			}
		}
	}
	return urls, postShellCmd
}

func QueryBgetFiles(name string, version string, release string, site string) (urls []string, postShellCmd []string) {
	chrom := []string{}
	for i := 1; i < 23; i++ {
		chrom = append(chrom, strconv.Itoa(i))
	}
	chrom = append(chrom, "X", "Y", "MT")
	for i := range BgetFilesPool {
		if BgetFilesPool[i].Name == name && (site == "" || BgetFilesPool[i].Site == site) {
			for j := range BgetFilesPool[i].URL {
				tmp := strings.Replace(BgetFilesPool[i].URL[j], "{{release}}", release, 100)
				version = genomeVersionConvertor(tmp, version)
				tmp = strings.Replace(tmp, "{{version}}", version, 100)
				tmp = strings.Replace(tmp, "{{site}}", site, 100)
				if butils.StrDetect(tmp, "{{chrom}}") {
					raw := tmp
					for k := range chrom {
						tmp = strings.Replace(raw, "{{chrom}}", chrom[k], 100)
						urls = append(urls, tmp)
					}
				}
				if !butils.StrDetect(tmp, "{{chrom}}") {
					urls = append(urls, tmp)
				}
			}
			for j := range BgetFilesPool[i].PostShellCmd {
				tmp := strings.Replace(BgetFilesPool[i].PostShellCmd[j], "{{version}}", version, 100)
				tmp = strings.Replace(tmp, "{{release}}", release, 100)
				tmp = strings.Replace(tmp, "{{site}}", site, 100)
				postShellCmd = append(postShellCmd, tmp)
			}
		}
	}
	return urls, postShellCmd
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

func init() {
	BgetFilesPool = append(BgetFilesPool, githubReposPool...)
}
