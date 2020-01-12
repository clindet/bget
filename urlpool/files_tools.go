package urlpool

var otherFiles = []bgetFilesURLType{
	{
		Name:     "annovar",
		URL:      []string{"http://www.openbioinformatics.org/annovar/download/0wgxR2rIVP/annovar.{{version}}.tar.gz"},
		Versions: []string{"latest"},
	},
	{
		Name:     "atlas2",
		URL:      []string{"https://sourceforge.net/projects/atlas2/files/Atlas2_v{{version}}.zip"},
		Versions: []string{"1.4.3", "1.4.1"},
	},
	{
		Name:     "autochrom3d",
		URL:      []string{"http://ibi.hzau.edu.cn/3dmodel/AutoChrom3D_{{version}}.tar.gz"},
		Versions: []string{"v1"},
	},
	{
		Name:     "beagle",
		URL:      []string{"http://faculty.washington.edu/browning/beagle/{{version}}"},
		Versions: []string{"beagle.08Jun17.d8b.jar", "beagle_4.1_21Jan17.pdf", "run.beagle.08Jun17.d8b.example", "release_notes", "beagle.170608.zip"},
	},
	{
		Name:     "bzip2",
		URL:      []string{"http://www.bzip.org/{{version}}/bzip2-{{version}}.tar.gz"},
		Versions: []string{"1.0.6"},
	},
	{
		Name:     "cesa",
		URL:      []string{"http://liulab.dfci.harvard.edu/CEAS/src/CEAS-Package-{{version}}.tar.gz"},
		Versions: []string{"1.0.2"},
	},
	{
		Name:     "chromhmm",
		URL:      []string{"http://compbio.mit.edu/ChromHMM/{{version}}.zip"},
		Versions: []string{"ChromHMM"},
	},
	{
		Name: "cromwell",
		URL:  []string{"https://github.com/broadinstitute/cromwell/releases/download/{{version}}/cromwell-{{version}}.jar", "https://github.com/broadinstitute/cromwell/releases/download/{{version}}/womtool-{{version}}.jar"},
	},
	{
		Name:     "curl",
		URL:      []string{"https://curl.haxx.se/download/curl-{{version}}.tar.gz"},
		Versions: []string{"7.54.0", "7.53.1", "7.53.0", "7.52.1", "7.52.0", "7.51.0", "7.50.3", "7.50.2", "7.50.1", "7.50.0", "7.49.1", "7.49.0", "7.48.0", "7.47.1", "7.47.0", "7.46.0", "7.45.0", "7.44.0", "7.43.0", "7.42.1", "7.42.0", "7.41.0", "7.40.0", "7.39.0", "7.38.0", "7.37.1", "7.37.0", "7.36.0", "7.35.0", "7.34.0", "7.33.0", "7.32.0", "7.31.0", "7.30.0", "7.29.0", "7.28.1", "7.28.0", "7.27.0", "7.26.0", "7.25.0", "7.24.0", "7.23.1", "7.23.0", "7.22.0", "7.21.7", "7.21.6", "7.21.5", "7.21.4", "7.21.3", "7.21.2", "7.21.1", "7.21.0", "7.20.1", "7.20.0"},
	},
	{
		Name:     "easeq",
		URL:      []string{"http://easeq.net/download/EaSeq.zip"},
		Versions: []string{"latest"},
	},
	{
		Name:     "assemble/edena",
		URL:      []string{"http://www.genomic.ch/edena/{{version}}.tar.gz", "http://www.genomic.ch/edena/{{version}}.zip"},
		Versions: []string{"EdenaV3.131028", "EdenaV3.130110", "EdenaV3.121122", "Edena3dev120926_linux64", "Edena3Dev120626_linux64", "Edena3Dev120626_linux32", "Edena3Dev120615_linux64", "Edena3Dev120615_linux32", "Edena3Dev110920_linux64", "Edena3Dev110920_linux32", "Edena3Dev110815_linux64", "Edena3Dev110815_linux32", "Edena3Dev110814_linux64", "Edena3Dev110814_linux32", "Edena3Dev110705_linux64", "Edena3Dev110705_linux32", "edena2.1.1_linux32", "edena2.1.1_linux64", "edena2.1.1_windows"},
	},
	{
		Name:     "effusion",
		URL:      []string{"http://www.babbittlab.ucsf.edu/effusion/effusion_{{version}}.tar.gz"},
		Versions: []string{"1.0.0"},
	},
	{
		Name:     "fastqc",
		URL:      []string{"http://www.bioinformatics.babraham.ac.uk/projects/fastqc/fastqc_v{{version}}.zip"},
		Versions: []string{"0.11.8", "0.11.7", "0.11.6", "0.11.5", "0.11.4", "0.11.3", "0.11.2", "0.11.1", "0.10.1", "0.10.0", "0.9.6", "0.9.5", "0.9.4", "0.9.3", "0.9.2", "0.9.1", "0.9.0", "0.8.0", "0.7.2", "0.7.1", "0.7.0", "0.6.1", "0.6.0", "0.5.1", "0.5.0", "0.4.3", "0.4.2", "0.4.1", "0.4", "0.3.1", "0.3", "0.2", "0.1.1", "0.1"},
	},
	{
		Name:     "fusioncatcher",
		URL:      []string{"https://sourceforge.net/projects/fusioncatcher/files/fusioncatcher_{{version}}.zip", "https://sourceforge.net/projects/fusioncatcher/files/bootstrap.py"},
		Versions: []string{"v0.99.7b", "v0.99.7a", "v0.99.6a", "v0.99.5a", "v0.99.4e", "v0.99.4d", "v0.99.4c", "v0.99.4b", "v0.99.4a", "v0.99.3e", "v0.99.3d", "v0.99.3c", "v0.99.3b", "v0.99.3a", "v0.99.2", "v0.99.1", "v0.99.0", "v0.95", "v0.96", "v0.93", "v0.94", "v0.98", "v0.97"},
	},
	{
		Name: "gatk4",
		URL:  []string{"https://github.com/broadinstitute/gatk/releases/download/{{version}}/gatk-{{version}}.zip", "https://github.com/broadinstitute/gatk/archive/{{version}}.zip"},
	},
	{
		Name:     "gatk-bundle",
		URL:      []string{"ftp://gsapubftp-anonymous@ftp.broadinstitute.org/bundle/{{version}}/"},
		Versions: []string{"hg38", "hg19", "hg18", "b37", "b36"},
	},
	{
		Name:     "gmap",
		URL:      []string{"http://research-pub.gene.com/gmap/src/{{version}}.tar.gz"},
		Versions: []string{"gmap-gsnap-2017-05-08", "gmap-gsnap-2017-04-21", "gmap-gsnap-2017-03-17", "gmap-gsnap-2017-02-15", "gmap-gsnap-2016-11-07", "gmap-gsnap-2016-09-14", "gmap-gsnap-2016-06-30.v5", "gmap-gsnap-2015-12-31.v10", "gmap-gsnap-2015-07-23", "gmap-gsnap-2014-12-31.v2", "gmap-gsnap-2014-10-22", "gmap-gsnap-2014-08-20", "gmap-gsnap-2014-06-10", "gmap-gsnap-2014-05-15.v3", "gmap-gsnap-2014-03-28.v2", "gmap-gsnap-2013-10-28", "gmap-gsnap-2013-09-30.v2", "gmap-gsnap-2013-03-31.v5", "gmap-gsnap-2012-07-20.v3", "gmap-gsnap-2012-04-27", "gmap-gsnap-2012-01-11", "gmap-gsnap-2011-12-28", "gmap-gsnap-2011-11-30", "gmap-gsnap-2011-10-16", "gmap-gsnap-2011-09-14", "gmap-gsnap-2011-03-28", "gmap-gsnap-2010-07-27", "gmap-gsnap-2010-03-09", "gmap-2007-09-28", "gmap-2007-06-04", "gmap-2006-12-18", "gmap-2006-04-21", "gmap-2005-10-25", "gmap-2005-07-15", "gmap-2005-06-23", "gmap-2005-05-20", "gmap-2005-05-06", "gmap-2005-03-11"},
	},
	{
		Name: "gridss",
		URL:  []string{"https://github.com/PapenfussLab/gridss"},
	},
	{
		Name:     "igv",
		URL:      []string{"http://data.broadinstitute.org/igv/projects/downloads/2.4/IGV_{{version}}.zip"},
		Versions: []string{"2.4.6"},
	},

	{
		Name:     "interproscan",
		URL:      []string{"ftp://ftp.ebi.ac.uk/pub/software/unix/iprscan/5/{{version}}/interproscan-{{version}}-64-bit.tar.gz"},
		Versions: []string{"5.27-66.0"},
	},
	{
		Name:     "lzo",
		URL:      []string{"http://www.oberhumer.com/opensource/lzo/download/lzo-{{version}}.tar.gz"},
		Versions: []string{"2.00", "2.01", "2.02", "2.03", "2.04", "2.05", "2.06", "2.07", "2.08", "2.09", "2.10"},
	},
	{
		Name:     "lzop",
		URL:      []string{"http://www.lzop.org/download/lzop-{{version}}.tar.gz"},
		Versions: []string{"1.03", "1.02-rc1", "1.01"},
	},
	{
		Name:     "maorm",
		URL:      []string{"http://bcb.dfci.harvard.edu/~gcyuan/MAnorm/MAnorm_data/MAnorm_Linux_R_Package.zip"},
		Versions: []string{"latest"},
	},
	{
		Name:     "mapsplice2",
		URL:      []string{"http://protocols.netlab.uky.edu/~zeng/MapSplice-v{{version}}.zip"},
		Versions: []string{"2.2.1", "2.2.0", "2.1.9", "2.1.8", "2.1.7", "2.1.6", "2.1.5", "2.1.4", "2.1.3", "2.1.2", "2.1.0"},
	},
	{
		Name:     "mutect",
		URL:      []string{"https://github.com/Miachol/gatk_releases/raw/master/mutect{{version}}.zip"},
		Versions: []string{"1.1.7", "1.1.4", "1.1.1"},
	},
	{
		Name: "nextflow",
		URL:  []string{"https://github.com/nextflow-io/nextflow"},
	},
	{
		Name:     "ngs-qc-toolkit",
		URL:      []string{"http://14.139.61.3:8080/ngsqctoolkit/NGSQCToolkit_{{version}}.zip"},
		Versions: []string{"v2.3.3", "v2.3.2", "v2.3.1"},
	},
	{
		Name:     "novoalign",
		URL:      []string{"https://github.com/Miachol/novocraft_releases/raw/master/novocraft{{version}}.tar.gz"},
		Versions: []string{"V3.07.01.Linux3.0", "V3.07.01.Linux2.6", "V3.07.01.MacOSX", "V3.07.00.Linux3.0", "V3.07.00.Linux2.6", "V3.07.00.MacOSX", "V3.06.05.Linux3.0", "V3.06.05.Linux2.6", "V3.06.05.MacOSX", "V3.06.04.Linux3.0", "V3.06.04.Linux2.6", "V3.06.04.MacOSX", "V3.06.03.Linux3.0", "V3.06.03.Linux2.6", "V3.06.03.MacOSX", "V3.06.02.Linux3.0", "V3.06.02.Linux2.6", "V3.06.02.MacOSX", "V3.06.01.Linux3.0", "V3.06.01.Linux2.6", "V3.06.01.MacOSX", "V3.06.00.Linux3.0", "V3.06.00.Linux2.6", "V3.06.00.MacOSX", "V3.05.01.Linux3.0", "V3.05.01.Linux2.6", "V3.05.01.MacOSX", "V3.04.06.Linux3.0", "V3.04.06.Linux2.6", "V3.04.06.MacOSX", "V3.04.04.Linux3.0", "V3.04.04.Linux2.6", "V3.04.04.MacOSX", "V3.04.02.Linux3.0", "V3.04.02.Linux2.6", "V3.04.02.MacOSX", "V3.04.01.Linux3.0", "V3.04.01.Linux2.6", "V3.04.01.MacOSX", "V3.03.02.Linux3.0", "V3.03.02.Linux2.6", "V3.03.02.MacOSX", "V3.03.01.Linux3.0", "V3.03.01.Linux2.6", "V3.03.01.MacOSX", "V3.03.00.Linux3.0", "V3.03.00.Linux2.6", "V3.03.00.MacOSX", "V3.02.13.Linux3.0", "V3.02.13.Linux2.6", "V3.02.13.MacOSX", "V3.02.13.Linux3.0", "V3.02.13.Linux2.6", "V3.02.13.MacOSX", "V3.02.12.Linux3.0", "V3.02.12.Linux2.6", "V3.02.12.MacOSX", "V3.02.12.Linux3.0", "V3.02.12.Linux2.6", "V3.02.12.MacOSX", "V3.02.12.Linux3.0", "V3.02.12.Linux2.6", "V3.02.12.MacOSX", "V3.02.11.Linux3.0", "V3.02.11.Linux2.6", "V3.02.11.MacOSX", "V3.02.10.Linux3.0", "V3.02.10.Linux2.6", "V3.02.10.MacOSX", "V3.02.09.Linux3.0", "V3.02.09.Linux2.6", "V3.02.09.MacOSX", "V3.02.08.Linux3.0", "V3.02.08.Linux2.6", "V3.02.08.MacOSX", "V3.02.07.Linux3.0", "V3.02.07.Linux2.6", "V3.02.07.MacOSX", "V3.02.06.Linux3.0", "V3.02.06.Linux2.6", "V3.02.06.MacOSX"},
	},
	{
		Name:     "other/paradigm",
		URL:      []string{"http://paradigm.five3genomics.com/five3_paradigm_weapi.py"},
		Versions: []string{"latest"},
	},
	{
		Name:     "pcre",
		URL:      []string{"https://ftp.pcre.org/pub/pcre/pcre-{{version}}.tar.gz"},
		Versions: []string{"8.40", "8.39", "8.38", "8.37", "8.36", "8.35", "8.34", "8.33", "8.32", "8.31", "8.30", "8.21", "8.20", "8.13", "8.12", "8.11", "8.10", "8.02", "8.01", "8.00"},
	},
	{
		Name:     "pigz",
		URL:      []string{"http://cdn-fastly.deb.debian.org/debian/pool/main/p/pigz/pigz_{{version}}.orig.tar.gz"},
		Versions: []string{"2.2.4", "2.3.1", "2.3.4"},
	},
	{
		Name:     "prada",
		URL:      []string{"https://sourceforge.net/projects/prada/files/pyPRADA/pyPRADA_{{version}}.tar.gz"},
		Versions: []string{"1.2"},
	},
	{
		Name:     "prinseq",
		URL:      []string{"https://sourceforge.net/projects/prinseq/files/standalone/prinseq-lite-{{version}}.tar.gz"},
		Versions: []string{"0.20.4", "0.20.3", "0.20.2", "0.20.1", "0.20", "0.19.5", "0.19.4", "0.19.3", "0.19.2", "0.19.1", "0.19", "0.18.3", "0.18.2", "0.18.1", "0.18", "0.17.4", "0.17.3", "0.17.2", "0.17.1", "0.17", "0.16", "0.15.1", "0.15", "0.14.4", "0.14.2", "0.14.1", "0.14", "0.13.2", "0.13.1", "0.13", "0.12", "0.11", "0.10", "0.9", "0.8"},
	},
	{
		Name:     "quest",
		URL:      []string{"http://mendel.stanford.edu/sidowlab/downloads/quest/QuEST_{{version}}.tar.gz"},
		Versions: []string{"2.4"},
	},
	{
		Name: "reditools",
		URL:  []string{"https://nchc.dl.sourceforge.net/project/reditools/REDItools-{{version}}.tar.gz"},
	},
	{
		Name:     "rmats",
		URL:      []string{"https://sourceforge.net/projects/rnaseq-mats/files/MATS/rMATS.{{version}}.tgz"},
		Versions: []string{"4.0.1", "3.2.5"},
	},
	{
		Name:     "root",
		URL:      []string{"https://root.cern.ch/download/root_v{{version}}.source.tar.gz"},
		Versions: []string{"6.09.02", "6.08.06", "6.08.04", "6.08.02", "6.08.00", "6.06.08", "6.06.06", "6.06.04", "6.06.02", "6.06.00", "6.05.02", "6.04.18", "6.04.16", "6.04.14", "6.04.12", "6.04.10", "6.04.08", "6.04.06", "6.04.04", "6.04.02", "6.04.00", "6.03.04", "6.03.02", "6.02.12", "6.02.10", "6.02.08", "6.02.05", "6.02.04", "6.02.03", "6.02.02", "6.02.01", "6.02.00", "6.00.02", "6.00.01", "6.00.00", "5.99.06", "5.99.05", "5.99.04", "5.34.36", "5.34.34", "5.34.32", "5.34.30", "5.34.28", "5.34.26", "5.34.25", "5.34.24", "5.34.23", "5.34.22", "5.34.21", "5.34.20", "5.34.19", "5.34.18", "5.34.17", "5.34.14", "5.34.13", "5.34.12", "5.34.11", "5.34.10", "5.34.09", "5.34.08", "5.34.07", "5.34.06", "5.34.05", "5.34.04", "5.34.03", "5.34.02", "5.34.01", "5.34.00", "5.33.02", "5.32.04", "5.32.03", "5.32.02", "5.32.01", "5.32.00", "5.30.06", "5.30.05", "5.30.04", "5.30.03", "5.30.02", "5.30.01", "5.30.00", "5.29.02", "5.28.00", "5.28.00", "5.28.00", "5.28.00", "5.28.00", "5.28.00", "5.28.00", "5.28.00", "5.28.00", "5.27.06", "5.27.04", "5.27.02", "5.26.00", "5.26.00", "5.26.00", "5.25.02", "5.24.00", "5.24.00", "5.23.04", "5.23.02", "5.22.00", "5.22.00", "5.22.00", "5.22.00", "5.22.00", "5.22.00", "5.22.00", "5.22.00", "5.22.00"},
	},
	{
		Name:     "samstat",
		URL:      []string{"https://sourceforge.net/projects/samstat/files/samstat-{{version}}.tar.gz"},
		Versions: []string{"1.5.1", "1.5"},
	},
	{
		Name:     "snpeff",
		URL:      []string{"https://sourceforge.net/projects/snpeff/files/snpEff_{{version}}_core.zip"},
		Versions: []string{"v4.3o", "v4.3m", "v4_3l", "v4_3k", "v4_3i", "v4_3g", "v4_3f", "v4_3e", "v4_3c", "v4_3b", "v4_3a", "v4_3", "v4_2", "v4_1l", "v4_1k", "v4_1j", "v4_1i", "v4_1h", "v4_1g", "v4_1f", "v4_1e", "v4_1d", "v4_1c", "v4_1b", "v4_1", "v4_0", "v3_6", "v3_5", "v3_4", "v3_3", "v3_2", "v3_1", "v3_0", "v3_0a"},
	},
	{
		Name:     "solexaqa",
		URL:      []string{"https://sourceforge.net/projects/solexaqa/files/src/SolexaQA++_v{{version}}.zip", "https://sourceforge.net/projects/solexaqa/files/src/SolexaQA_v.{{version}}.zip"},
		Versions: []string{"3.1.7.1", "3.1.7", "3.1.6", "3.1.5", "3.1.4", "3.1.3", "3.1.2", "3.1.1", "3.1", "3.0", "2.5", "2.4", "2.3", "2.2", "2.1", "2.0", "1.13"},
	},
	{
		Name:     "sqlite",
		URL:      []string{"http://www.sqlite.org/2017/sqlite-autoconf-{{version}}.tar.gz", "http://www.sqlite.org/2016/sqlite-autoconf-{{version}}.tar.gz", "http://www.sqlite.org/2015/sqlite-autoconf-{{version}}.tar.gz", "http://www.sqlite.org/2014/sqlite-autoconf-{{version}}.tar.gz", "http://www.sqlite.org/2013/sqlite-autoconf-{{version}}.tar.gz"},
		Versions: []string{"3180000", "3170000", "3130000"},
	},

	{
		Name:     "srnanalyzer",
		URL:      []string{"http://srnanalyzer.systemsbiology.net/downloads/sRNAnalyzer.tar.gz"},
		Versions: []string{"latest"},
	},
	{
		Name:     "ssaha2",
		URL:      []string{"ftp://ftp.sanger.ac.uk/pub/resources/software/ssaha2/ssaha2_{{version}}.tgz"},
		Versions: []string{"v2.5.1_mac", "v2.5.2_i686", "v2.5.2_ia64", "v2.5.2_x86_64", "v2.5.3_i686", "v2.5.3_ia64", "v2.5.3_x86_64", "v2.5.4_i686", "v2.5.4_MacOSX_i386", "v2.5.4_x86_64", "v2.5.5_i686", "v2.5.5_x86_64"},
	},
	{
		Name:     "strelka",
		URL:      []string{"ftp://strelka@ftp.illumina.com/v1-branch/v{{version}}/strelka_workflow-{{version}}.tar.gz"},
		Versions: []string{"1.0.15", "1.0.4", "1.0.5", "1.0.6", "1.0.7", "1.0.10", "1.0.11", "1.0.12", "1.0.13", "1.0.14"},
	},
	{
		Name:     "subread",
		URL:      []string{"https://sourceforge.net/projects/subread/files/subread-{{version}}/subread-{{version}}-source.tar.gz"},
		Versions: []string{"1.6.0"},
	},
	{
		Name:     "tvc",
		URL:      []string{"http://updates.iontorrent.com/tvc_standalone/tvc-{{version}}.tar.gz"},
		Versions: []string{"5.2.2", "5.2.1", "5.0.3"},
	},
	{
		Name:     "ucsc",
		URL:      []string{"http://hgdownload.cse.ucsc.edu/admin/exe/userApps.{{version}}.src.tgz"},
		Versions: []string{"v305", "v306", "v307", "v308", "v309", "v310", "v311", "v312", "v313", "v314", "v315", "v316", "v317", "v318", "v319", "v320", "v321", "v322", "v323", "v324", "v325", "v326", "v327", "v328", "v329", "v330", "v331", "v332", "v333", "v334", "v335", "v336", "v337", "v338", "v339", "v340", "v341", "v342", "v343", "v344"},
	},
	{
		Name: "useq",
		URL:  []string{"https://github.com/HuntsmanCancerInstitute/USeq/releases/download/USeq_{{version}}/USeq_{{version}}.zip"},
	},
	{
		Name:     "vadir",
		URL:      []string{"ftp://penguin.genomics.cn/pub/10.5524/100001_101000/100360/VaDiR.tar.gz"},
		Versions: []string{"latest"},
	},
	{
		Name:     "velvet",
		URL:      []string{"http://www.ebi.ac.uk/~zerbino/velvet/velvet_{{version}}.tgz"},
		Versions: []string{"1.2.10", "0.7.01", "0.6.01", "0.5.01", "0.4", "0.3"},
	},
	{
		Name:     "xz",
		URL:      []string{"http://tukaani.org/xz/xz-{{version}}.tar.gz"},
		Versions: []string{"5.2.3", "5.2.2"},
	},
}
