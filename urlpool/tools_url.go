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
		Name: "blat",
		URL: map[string][]string{
			"Linux": []string{"http://hgdownload.soe.ucsc.edu/admin/exe/linux.x86_64/blat/blat"},
			"Mac":   []string{"http://hgdownload.soe.ucsc.edu/admin/exe/macOSX.x86_64/blat/blat"},
		},
		Versions: []string{"latest"},
	},
	{
		Name: "blast",
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
}
