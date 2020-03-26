package urlpool

var toolsLinks = []bgetToolsURLType{
	{
		Name: "imagej",
		URL: map[string][]string{
			"Linux":   []string{"http://wsr.imagej.net/distros/linux/ij{{version}}-linux64-java8.zip"},
			"Mac":     []string{"http://wsr.imagej.net/distros/osx/ij{{version}}-osx-java8.zip"},
			"Windows": []string{"http://wsr.imagej.net/distros/win/is{{version}}-win-java8.zip"},
		},
		Versions: []string{"150"},
	},
	{
		Name: "zlib",
		URL: map[string][]string{
			"Linux": []string{"https://sourceforge.net/projects/libpng/files/zlib/{{version}}/zlib-{{version}}.tar.gz"},
			"Mac":   []string{"https://sourceforge.net/projects/libpng/files/zlib/{{version}}/zlib-{{version}}.tar.gz"},
		},
		Versions: []string{"1.2.11", "1.2.10", "1.2.9", "1.2.7", "1.2.8", "1.2.6", "1.2.5", "1.2.4", "1.2.3", "1.2.1", "1.1.4", "1.1.3"},
	},
	{
		Name: "vcfanno",
		URL: map[string][]string{
			"Linux": []string{"https://github.com/brentp/vcfanno/releases/download/{{version}}/vcfanno_linux64"},
			"Mac":   []string{"https://github.com/brentp/vcfanno/releases/download/{{version}}/vcfanno_osx"},
		},
	},
	{
		Name: "fatotwobit",
		URL: map[string][]string{
			"Linux": []string{"http://hgdownload.soe.ucsc.edu/admin/exe/linux.x86_64/faToTwoBit"},
			"Mac":   []string{"http://hgdownload.soe.ucsc.edu/admin/exe/macOSX.x86_64/faToTwoBit"},
		},
		Versions: []string{"latest"},
	},
	{
		Name: "aligner/blat",
		URL: map[string][]string{
			"Linux": []string{"http://hgdownload.soe.ucsc.edu/admin/exe/linux.x86_64/blat/blat"},
			"Mac":   []string{"http://hgdownload.soe.ucsc.edu/admin/exe/macOSX.x86_64/blat/blat"},
		},
		Versions: []string{"latest"},
	},
	{
		Name: "aligner/blast",
		URL: map[string][]string{
			"Linux": []string{"http://ftp.ncbi.nlm.nih.gov/blast/executables/blast+/{{version}}/ncbi-blast-{{version}}+-x64-linux.tar.gz"},
			"Mac":   []string{"http://ftp.ncbi.nlm.nih.gov/blast/executables/blast+/{{version}}/ncbi-blast-{{version}}+-x64-macosx.tar.gz", "http://ftp.ncbi.nlm.nih.gov/blast/executables/blast+/{{version}}/ncbi-blast-{{version}}+.dmg"},
			"win":   []string{"http://ftp.ncbi.nlm.nih.gov/blast/executables/blast+/{{version}}/ncbi-blast-{{version}}+-x64-win64.tar.gz", "http://ftp.ncbi.nlm.nih.gov/blast/executables/blast+/{{version}}/ncbi-blast-{{version}}+-win64.exe"},
		},
		Versions: []string{"2.7.1", "2.6.0", "2.5.0", "2.4.0", "2.3.0"},
	},

	{
		Name: "liftover",
		URL: map[string][]string{
			"Linux": []string{"http://hgdownload.soe.ucsc.edu/admin/exe/linux.x86_64/liftOver"},
			"Mac":   []string{"http://hgdownload.soe.ucsc.edu/admin/exe/macOSX.x86_64/liftOver"},
		},
		Versions: []string{"latest"},
	},
	{
		Name: "sratools",
		URL: map[string][]string{
			"Linux":   []string{"https://ftp-trace.ncbi.nlm.nih.gov/sra/sdk/{{version}}/sratoolkit.ubuntu.tar.gz", "https://ftp-trace.ncbi.nlm.nih.gov/sra/sdk/{{version}}/sratoolkit.centos.tar.gz"},
			"Mac":     []string{"https://ftp-trace.ncbi.nlm.nih.gov/sra/sdk/{{version}}/sratoolkit.{{version}}-mac64.tar.gz"},
			"Windows": []string{"https://ftp-trace.ncbi.nlm.nih.gov/sra/sdk/{{version}}/sratoolkit.{{version}}-win64.zip"},
		},
		VersionsAPI: "https://github.com/ncbi/sra-tools",
	},
	{
		Name: "miniconda2",
		URL: map[string][]string{
			"Linux": []string{"https://repo.anaconda.com/miniconda/Miniconda2-{{version}}-Linux-x86_64.sh"},
			"Mac":   []string{"https://repo.anaconda.com/miniconda/Miniconda2-{{version}}-MacOSX-x86_64.sh"},
			"Win":   []string{"https://repo.anaconda.com/miniconda/Miniconda2-{version}}-Windows-x86_64.exe"}},
		PostShellCmd: []string{"cd {{pdir}} && sh {{dest}} -b -p {{downloadDir}}/miniconda2"}},
	{
		Name: "miniconda3",
		URL: map[string][]string{
			"Linux": []string{"https://repo.anaconda.com/miniconda/Miniconda3-{{version}}-Linux-x86_64.sh"},
			"Mac":   []string{"https://repo.anaconda.com/miniconda/Miniconda3-{{version}}-MacOSX-x86_64.sh"},
			"Win":   []string{"https://repo.anaconda.com/miniconda/Miniconda3-{version}}-Windows-x86_64.exe"}},
		PostShellCmd: []string{"cd {{pdir}} && sh {{dest}} -b -p {{downloadDir}}/miniconda3"}},
	{
		Name: "gdc-client",
		URL: map[string][]string{
			"Linux": []string{"https://gdc.cancer.gov/system/files/authenticated%20user/0/gdc-client_v{{version}}_Ubuntu_x64.zip"},
			"Mac":   []string{"https://gdc.cancer.gov/system/files/authenticated%20user/0/gdc-client_v{{version}}_OSX_x64_10.12.6.zip", "https://gdc.cancer.gov/system/files/authenticated%20user/0/gdc-client_v{{version}}_OSX_x64.zip"},
			"Win":   []string{"https://gdc.cancer.gov/system/files/authenticated%20user/0/gdc-client_v{{version}}_Windows_x64.zip"},
		},
		VersionsAPI: "https://github.com/NCI-GDC/gdc-client",
	},
	{
		Name: "golang",
		URL: map[string][]string{
			"Linux": []string{"https://dl.google.com/go/go{{version}}.linux-amd64.tar.gz"},
			"Mac":   []string{"https://dl.google.com/go/go{{version}}.darwin-amd64.pkg"},
			"Win":   []string{"https://dl.google.com/go/go{{version}}.windows-amd64.msi"},
		},
		Versions: []string{"1.13.4", "1.13.3", "1.13.2", "1.13.1", "1.13", "1.12.13", "1.12.12", "1.12.11", "1.12.10", "1.12.9", "1.12.8", "1.12.7", "1.12.6", "1.12.5", "1.12.4", "1.12.3", "1.12.2", "1.12.1", "1.12", "1.11.13", "1.11.12", "1.11.11", "1.11.10", "1.11.9", "1.11.8", "1.11.7", "1.11.6", "1.11.5", "1.11.4", "1.11.3", "1.11.2", "1.11.1", "1.11", "1.10.8", "1.10.7", "1.10.6", "1.10.5", "1.10.4", "1.10.3", "1.10.2", "1.10.1", "1.10", "1.9.7", "1.9.6", "1.9.5", "1.9.4", "1.9.3", "1.9.2", "1.9.1", "1.9", "1.8.7", "1.8.6", "1.8.5", "1.8.4", "1.8.3", "1.8.2", "1.8.1", "1.8", "1.7.6", "1.7.5", "1.7.4", "1.7.3", "1.7.2", "1.7.1", "1.7", "1.6.4", "1.6.3", "1.6.2", "1.6.1", "1.6", "1.5.4", "1.5.3", "1.5.2", "1.5.1", "1.5", "1.4.3", "1.4.2", "1.4.1", "1.4", "1.3.3", "1.3.2", "1.3.1", "1.3", "1.2.2", "1.2.1", "1.2", "1.1.2", "1.1.1", "1.1", "1.0.3", "1.0.2", "1.0.1"},
	},
	{
		Name: "nodejs",
		URL: map[string][]string{
			"Linux": []string{"https://nodejs.org/dist/{{version}}/node-{{version}}-linux-x64.tar.xz"},
			"Mac":   []string{"https://nodejs.org/dist/{{version}}/node-{{version}}.pkg"},
			"Win":   []string{"https://nodejs.org/dist/{{version}}/node-{{version}}-x64.msi"},
		},
		VersionsAPI: "https://github.com/nodejs/node",
	},
	{
		Name: "bget",
		URL: map[string][]string{
			"Linux": []string{"https://github.com/openbiox/bget/releases/download/{{version}}/bget_linux64"},
			"Mac":   []string{"https://github.com/openbiox/bget/releases/download/{{version}}/bget_osx"},
			"Win":   []string{"https://github.com/openbiox/bget/releases/download/{{version}}/bget.exe"},
		},
	},
	{
		Name: "bioctl",
		URL: map[string][]string{
			"Linux": []string{"https://github.com/openbiox/bioctl/releases/download/{{version}}/bioctl_linux64"},
			"Mac":   []string{"https://github.com/openbiox/bioctl/releases/download/{{version}}/bioctl_osx"},
			"Win":   []string{"https://github.com/openbiox/bioctl/releases/download/{{version}}/bioctl.exe"},
		},
	},
	{
		Name: "bioextr",
		URL: map[string][]string{
			"Linux": []string{"https://github.com/openbiox/bioextr/releases/download/{{version}}/bioextr_linux64"},
			"Mac":   []string{"https://github.com/openbiox/bioextr/releases/download/{{version}}/bioextr_osx"},
			"Win":   []string{"https://github.com/openbiox/bioextr/releases/download/{{version}}/bioextr.exe"},
		},
	},
	{
		Name: "xpdf-tools",
		URL: map[string][]string{
			"Linux": []string{"https://xpdfreader-dl.s3.amazonaws.com/xpdf-tools-linux-{{version}}.tar.gz"},
			"Mac":   []string{"https://xpdfreader-dl.s3.amazonaws.com/xpdf-tools-mac-{{version}}.tar.gz"},
			"Win":   []string{"https://xpdfreader-dl.s3.amazonaws.com/xpdf-tools-win-{{version}}.tar.gz"},
		},
		Versions: []string{"4.02"},
	},
	{
		Name: "r",
		URL: map[string][]string{
			"Linux": []string{"https://mirrors.tuna.tsinghua.edu.cn/CRAN/src/base/R-3/R-{{version}}.tar.gz"},
			"Mac":   []string{"https://mirrors.tuna.tsinghua.edu.cn/CRAN/bin/macosx/bin/macosx/R-{{version}}.pkg"},
			"Win":   []string{"https://mirrors.tuna.tsinghua.edu.cn/CRAN/bin/windows/base/R-{{version}}-win.exe"},
		},
		Versions: []string{"3.6.3"},
	},
}
