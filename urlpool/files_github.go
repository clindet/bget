package urlpool

var githubRepos = []bgetFilesURLType{
	{
		Name:         "github_demo",
		URL:          []string{"https://github.com/Miachol/github_demo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bwa",
		URL:          []string{"https://github.com/lh3/bwa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "star",
		URL:          []string{"https://github.com/alexdobin/STAR"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "samtools",
		URL:          []string{"https://github.com/samtools/samtools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bcftools",
		URL:          []string{"https://github.com/samtools/bcftools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bowtie",
		URL:          []string{"https://github.com/BenLangmead/bowtie"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bowtie2",
		URL:          []string{"https://github.com/BenLangmead/bowtie2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "tophat",
		URL:          []string{"https://github.com/infphilo/tophat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "varscan",
		URL:          []string{"https://github.com/Miachol/varscan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "picard",
		URL:          []string{"https://github.com/broadinstitute/picard"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "vcftools",
		URL:          []string{"https://github.com/vcftools/vcftools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "pindel",
		URL:          []string{"https://github.com/genome/pindel"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "lofreq",
		URL:          []string{"https://github.com/Miachol/lofreq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "hisat2",
		URL:          []string{"https://github.com/infphilo/hisat2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "freebayes",
		URL:          []string{"https://github.com/ekg/freebayes"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "macs",
		URL:          []string{"https://github.com/taoliu/MACS/"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bedtools2",
		URL:          []string{"https://github.com/arq5x/bedtools2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "sparsehash",
		URL:          []string{"https://github.com/sparsehash/sparsehash"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "abyss",
		URL:          []string{"https://github.com/bcgsc/abyss"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bamtools",
		URL:          []string{"https://github.com/pezmaster31/bamtools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "breakdancer",
		URL:          []string{"https://github.com/genome/breakdancer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "htseq",
		URL:          []string{"https://github.com/simon-anders/htseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seqtk",
		URL:          []string{"https://github.com/ndaniel/seqtk"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "delly",
		URL:          []string{"https://github.com/dellytools/delly"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "tmap",
		URL:          []string{"git://github.com/iontorrent/TMAP"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "somatic-sniper",
		URL:          []string{"https://github.com/genome/somatic-sniper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bamutil",
		URL:          []string{"https://github.com/statgen/bamUtil"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "vcflib",
		URL:          []string{"https://github.com/vcflib/vcflib"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "jvarkit",
		URL:          []string{"https://github.com/lindenb/jvarkit/"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "fastx_toolkit",
		URL:          []string{"https://github.com/agordon/fastx_toolkit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "libgtextutils",
		URL:          []string{"https://github.com/agordon/libgtextutils"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "trinityrnaseq",
		URL:          []string{"https://github.com/trinityrnaseq/trinityrnaseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "oases",
		URL:          []string{"https://github.com/dzerbino/oases"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "rum",
		URL:          []string{"https://github.com/itmat/rum"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "igraph",
		URL:          []string{"https://github.com/igraph/igraph"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "pxz",
		URL:          []string{"https://github.com/jnovy/pxz"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "cnvkit",
		URL:          []string{"https://github.com/etal/cnvkit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "speedseq",
		URL:          []string{"https://github.com/hall-lab/speedseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "cnvnator",
		URL:          []string{"https://github.com/abyzovlab/CNVnator"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "asap",
		URL:          []string{"https://github.com/DeplanckeLab/ASAP"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mdseq",
		URL:          []string{"https://github.com/zjdaye/MDSeq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "sclvm",
		URL:          []string{"https://github.com/PMBio/scLVM"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "f-sclvm",
		URL:          []string{"https://github.com/PMBio/f-scLVM"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bearscc",
		URL:          []string{"https://github.com/Miachol/bearscc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "singlesplice",
		URL:          []string{"https://github.com/jw156605/SingleSplice"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "isop",
		URL:          []string{"https://github.com/nghiavtr/ISOP"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "brie",
		URL:          []string{"https://github.com/huangyh09/brie"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "outrigger",
		URL:          []string{"https://github.com/YeoLab/outrigger"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "raceid",
		URL:          []string{"https://github.com/dgrun/RaceID"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "backspin",
		URL:          []string{"https://github.com/linnarsson-lab/BackSPIN"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "zifa",
		URL:          []string{"https://github.com/epierson9/ZIFA"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seurat",
		URL:          []string{"https://github.com/satijalab/seurat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "rca",
		URL:          []string{"https://github.com/GIS-SP-Group/RCA"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mimosca",
		URL:          []string{"https://github.com/asncd/MIMOSCA"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "tracer",
		URL:          []string{"https://github.com/teichlab/tracer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "scnorm",
		URL:          []string{"https://github.com/rhondabacher/SCnorm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "sleuth",
		URL:          []string{"https://github.com/pachterlab/sleuth"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "oncotator",
		URL:          []string{"https://github.com/broadinstitute/oncotator"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "ensembl-vep",
		URL:          []string{"https://github.com/Ensembl/ensembl-vep"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "fastq-tools",
		URL:          []string{"https://github.com/dcjones/fastq-tools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "arnapipe",
		URL:          []string{"https://github.com/HudsonAlpha/aRNAPipe"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "trimgalore",
		URL:          []string{"https://github.com/FelixKrueger/TrimGalore"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "multiqc",
		URL:          []string{"https://github.com/ewels/MultiQC"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "strawberry",
		URL:          []string{"https://github.com/ruolin/strawberry"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "fastp",
		URL:          []string{"https://github.com/OpenGene/fastp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "olego",
		URL:          []string{"https://github.com/chaolinzhanglab/olego"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "chronqc",
		URL:          []string{"https://github.com/nilesh-tawari/ChronQC"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "dart",
		URL:          []string{"https://github.com/hsinnan75/DART"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "rhat",
		URL:          []string{"https://github.com/HIT-Bioinformatics/rHAT"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "giggle",
		URL:          []string{"https://github.com/ryanlayer/giggle"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "facets",
		URL:          []string{"https://github.com/mskcc/facets"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "resm",
		URL:          []string{"https://github.com/deweylab/RSEM"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "radia",
		URL:          []string{"https://github.com/aradenbaugh/radia/"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "manta",
		URL:          []string{"https://github.com/Illumina/manta"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "jaffa",
		URL:          []string{"https://github.com/Oshlack/JAFFA"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "marvel",
		URL:          []string{"https://github.com/schloi/MARVEL"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "picky",
		URL:          []string{"https://github.com/TheJacksonLaboratory/Picky"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "freec",
		URL:          []string{"https://github.com/BoevaLab/FREEC"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "taxmaps",
		URL:          []string{"https://github.com/nygenome/taxmaps"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "svaba",
		URL:          []string{"https://github.com/walaj/svaba"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "rop",
		URL:          []string{"https://github.com/smangul1/rop"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mrnn",
		URL:          []string{"https://github.com/hendrixlab/mRNN"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "dna-seq-gatk-variant-calling",
		URL:          []string{"https://github.com/snakemake-workflows/dna-seq-gatk-variant-calling"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "spack",
		URL:          []string{"https://github.com/spack/spack"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bpipe",
		URL:          []string{"https://github.com/ssadedin/bpipe"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mscentipede",
		URL:          []string{"https://github.com/rajanil/msCentipede"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "chia-pet2",
		URL:          []string{"https://github.com/GuipengLi/ChIA-PET2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "chromtime",
		URL:          []string{"https://github.com/ernstlab/ChromTime"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "dreg",
		URL:          []string{"https://github.com/Danko-Lab/dREG"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "genomedisco",
		URL:          []string{"https://github.com/kundajelab/genomedisco"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mistic",
		URL:          []string{"https://github.com/iric-soft/MiSTIC"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mosdepth",
		URL:          []string{"https://github.com/brentp/mosdepth"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mecat",
		URL:          []string{"https://github.com/xiaochuanle/MECAT"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "vg",
		URL:          []string{"https://github.com/vgteam/vg"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "stringtie",
		URL:          []string{"https://github.com/gpertea/stringtie"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "iontorrent/ts",
		URL:          []string{"https://github.com/iontorrent/TS"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "fusioncatcher",
		URL:          []string{"https://github.com/ndaniel/fusioncatcher"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "svscore",
		URL:          []string{"https://github.com/lganel/SVScore"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "svtools",
		URL:          []string{"https://github.com/hall-lab/svtools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "iseq",
		URL:          []string{"https://github.com/JhuangLab/iseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "deepvariant",
		URL:          []string{"https://github.com/google/deepvariant"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "sv2",
		URL:          []string{"https://github.com/dantaki/SV2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mutscan",
		URL:          []string{"https://github.com/OpenGene/MutScan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "breakmer",
		URL:          []string{"https://github.com/ccgd-profile/BreaKmer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "detin",
		URL:          []string{"https://github.com/broadinstitute/deTiN"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "cn_learn",
		URL:          []string{"https://github.com/girirajanlab/CN_Learn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "wham",
		URL:          []string{"https://github.com/zeeev/wham"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "gemini",
		URL:          []string{"https://github.com/arq5x/gemini"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "pcgr",
		URL:          []string{"https://github.com/sigven/pcgr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "annovarr",
		URL:          []string{"https://github.com/JhuangLab/annovarR"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bystro",
		URL:          []string{"https://github.com/akotlar/bystro"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "pathopredictor",
		URL:          []string{"https://github.com/samesense/pathopredictor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "agfusion",
		URL:          []string{"https://github.com/murphycj/AGFusion"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "genomeuplot",
		URL:          []string{"https://github.com/gaitat/GenomeUPlot"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "breakpointsurveyor",
		URL:          []string{"https://github.com/ding-lab/BreakPointSurveyor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "chimeraviz",
		URL:          []string{"https://github.com/stianlagstad/chimeraviz"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "biasmv/pv",
		URL:          []string{"https://github.com/biasmv/pv"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "g2s",
		URL:          []string{"https://github.com/genome-nexus/g2s"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "ngb",
		URL:          []string{"https://github.com/epam/NGB"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "ncboost",
		URL:          []string{"https://github.com/RausellLab/NCBoost"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "leafcutter",
		URL:          []string{"https://github.com/davidaknowles/leafcutter"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "irfinder",
		URL:          []string{"https://github.com/williamritchie/IRFinder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mmsplice",
		URL:          []string{"https://github.com/gagneurlab/MMSplice"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "ballgown",
		URL:          []string{"https://github.com/alyssafrazee/ballgown"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mixcr",
		URL:          []string{"https://github.com/milaboratory/mixcr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "topconfects",
		URL:          []string{"https://github.com/pfh/topconfects"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "plier",
		URL:          []string{"https://github.com/wgmao/PLIER"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "viral-ngs",
		URL:          []string{"https://github.com/broadinstitute/viral-ngs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "qap",
		URL:          []string{"https://github.com/mingjiewang/qap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "vifi",
		URL:          []string{"https://github.com/namphuon/ViFi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "micropro",
		URL:          []string{"https://github.com/zifanzhu/MicroPro"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "feast",
		URL:          []string{"https://github.com/cozygene/FEAST"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mcorr",
		URL:          []string{"https://github.com/kussell-lab/mcorr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "dropclust",
		URL:          []string{"https://github.com/debsin/dropClust"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "trendsceek",
		URL:          []string{"https://github.com/edsgard/trendsceek"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "awesome-single-cell",
		URL:          []string{"https://github.com/seandavi/awesome-single-cell"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "saver",
		URL:          []string{"https://github.com/mohuangx/SAVER"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "cellsius",
		URL:          []string{"https://github.com/Novartis/CellSIUS"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "scrabble",
		URL:          []string{"https://github.com/software-github/SCRABBLE"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "melissa",
		URL:          []string{"https://github.com/andreaskapou/Melissa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "paga",
		URL:          []string{"https://github.com/theislab/paga"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "clonealign",
		URL:          []string{"https://github.com/kieranrcampbell/clonealign"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "cellfishing.jl",
		URL:          []string{"https://github.com/bicycle1885/CellFishing.jl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "viper",
		URL:          []string{"https://github.com/ChenMengjie/VIPER"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "scgen",
		URL:          []string{"https://github.com/theislab/scgen"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "conos",
		URL:          []string{"https://github.com/hms-dbmi/conos"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "magic",
		URL:          []string{"https://github.com/KrishnaswamyLab/MAGIC"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "zinbwave",
		URL:          []string{"https://github.com/drisso/zinbwave"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "simlr_py",
		URL:          []string{"https://github.com/bowang87/SIMLR_PY"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "dca",
		URL:          []string{"https://github.com/theislab/dca"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "scvi",
		URL:          []string{"https://github.com/YosefLab/scVI"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "phenograph",
		URL:          []string{"https://github.com/jacoblevine/PhenoGraph"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "splatter-paper",
		URL:          []string{"https://github.com/Oshlack/splatter-paper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "deepnovo-dia",
		URL:          []string{"https://github.com/nh2tran/DeepNovo-DIA"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "scvi",
		URL:          []string{"https://github.com/YosefLab/scVI"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "caveman",
		URL:          []string{"https://github.com/funpopgen/CaVEMaN"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bcbio-nextgen",
		URL:          []string{"https://github.com/chapmanb/bcbio-nextgen"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "orange3",
		URL:          []string{"https://github.com/biolab/orange3"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "sequana",
		URL:          []string{"https://github.com/sequana/sequana"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bedops",
		URL:          []string{"https://github.com/bedops/bedops"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "maftools",
		URL:          []string{"https://github.com/PoisonAlien/maftools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "samblaster",
		URL:          []string{"https://github.com/GregoryFaust/samblaster"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "ngstk",
		URL:          []string{"https://github.com/JhuangLab/ngstk"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bioinstaller",
		URL:          []string{"https://github.com/JhuangLab/BioInstaller"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "chromhmm",
		URL:          []string{"https://github.com/jernst98/ChromHMM"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "biobloom",
		URL:          []string{"https://github.com/bcgsc/biobloom"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "phenopredict",
		URL:          []string{"https://github.com/leekgroup/phenopredict"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "recount",
		URL:          []string{"https://github.com/leekgroup/recount"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "lsmm",
		URL:          []string{"https://github.com/mingjingsi/LSMM"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "vcf2maf",
		URL:          []string{"https://github.com/mskcc/vcf2maf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "r2d3",
		URL:          []string{"https://github.com/rstudio/r2d3"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "liteq",
		URL:          []string{"https://github.com/r-lib/liteq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "relaxed",
		URL:          []string{"https://github.com/RelaxedJS/ReLaXed"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "dash",
		URL:          []string{"https://github.com/jonocarroll/dash"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "threadpool",
		URL:          []string{"https://github.com/rdpeng/threadpool"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "rapid",
		URL:          []string{"https://github.com/ZhiGroup/RaPID"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "gemini",
		URL:          []string{"https://github.com/sellerslab/gemini"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "confined",
		URL:          []string{"https://github.com/cozygene/CONFINED"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "marginphase",
		URL:          []string{"https://github.com/benedictpaten/marginPhase"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "chicmaxima",
		URL:          []string{"https://github.com/yousra291987/ChiCMaxima"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "circbrain",
		URL:          []string{"https://github.com/yangence/circBrain"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bazam",
		URL:          []string{"https://github.com/ssadedin/bazam"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "conbase",
		URL:          []string{"https://github.com/conbase/conbase"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "3dchromatin_replicateqc",
		URL:          []string{"https://github.com/kundajelab/3DChromatin_ReplicateQC"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "i-boost",
		URL:          []string{"https://github.com/alexwky/I-Boost"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bin3c",
		URL:          []string{"https://github.com/cerebis/bin3C"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "dstruct",
		URL:          []string{"https://github.com/AviranLab/dStruct"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "skmer",
		URL:          []string{"https://github.com/shahab-sarmashghi/Skmer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "iguide",
		URL:          []string{"https://github.com/cnobles/iGUIDE"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "plyranges",
		URL:          []string{"https://github.com/sa-lee/plyranges"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "forge",
		URL:          []string{"https://github.com/langmead-lab/FORGe"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "se-mei",
		URL:          []string{"https://github.com/dpryan79/SE-MEI"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "anchor",
		URL:          []string{"https://github.com/GuanLab/Anchor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "advntr",
		URL:          []string{"https://github.com/mehrdadbakhtiari/adVNTR"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "ldsc",
		URL:          []string{"https://github.com/bulik/ldsc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "bigstitcher",
		URL:          []string{"https://github.com/PreibischLab/BigStitcher"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "ivtnmr",
		URL:          []string{"https://github.com/systemsnmr/ivtnmr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "divers",
		URL:          []string{"https://github.com/hym0405/DIVERS"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "prosit",
		URL:          []string{"https://github.com/kusterlab/prosit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "deepcell-tf",
		URL:          []string{"https://github.com/vanvalenlab/deepcell-tf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "cdeep3m",
		URL:          []string{"https://github.com/CRBS/cdeep3m"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "unet-segmentation",
		URL:          []string{"https://github.com/lmb-freiburg/Unet-Segmentation"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "cellprofiler",
		URL:          []string{"https://github.com/CellProfiler/CellProfiler"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mask_rcnn",
		URL:          []string{"https://github.com/matterport/Mask_RCNN"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "leap",
		URL:          []string{"https://github.com/talmo/leap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "in-silico-labeling",
		URL:          []string{"https://github.com/google/in-silico-labeling"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "trackviewer.documentation",
		URL:          []string{"https://github.com/jianhong/trackViewer.documentation"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "cistopic",
		URL:          []string{"https://github.com/aertslab/cistopic"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "selene",
		URL:          []string{"https://github.com/FunctionLab/selene"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "sda",
		URL:          []string{"https://github.com/mvollger/SDA"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "fmriprep",
		URL:          []string{"https://github.com/poldracklab/fmriprep"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "unifrac",
		URL:          []string{"https://github.com/biocore/unifrac"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "gvmap",
		URL:          []string{"https://github.com/ytdai/gvmap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "easysvg",
		URL:          []string{"https://github.com/ytdai/easySVG"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "hexmapr",
		URL:          []string{"https://github.com/sassalley/hexmapr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "clustergrammer",
		URL:          []string{"https://github.com/MaayanLab/clustergrammer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "chromvar",
		URL:          []string{"https://github.com/GreenleafLab/chromVAR"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "echarts",
		URL:          []string{"https://github.com/ecomfe/echarts"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "plotly.js",
		URL:          []string{"https://github.com/plotly/plotly.js"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "qvalue",
		URL:          []string{"https://github.com/StoreyLab/qvalue"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "genvisr",
		URL:          []string{"https://github.com/griffithlab/GenVisR"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "r-color-palettes",
		URL:          []string{"https://github.com/EmilHvitfeldt/r-color-palettes"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "opencpu",
		URL:          []string{"https://github.com/opencpu/opencpu"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "ggthemr",
		URL:          []string{"https://github.com/cttobin/ggthemr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "paletter",
		URL:          []string{"https://github.com/AndreaCirilloAC/paletter"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "ggdag",
		URL:          []string{"https://github.com/malcolmbarrett/ggdag"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "ggseqlogo",
		URL:          []string{"https://github.com/omarwagih/ggseqlogo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "three.js",
		URL:          []string{"https://github.com/mrdoob/three.js"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "higlass",
		URL:          []string{"https://github.com/hms-dbmi/higlass"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "babun",
		URL:          []string{"https://github.com/babun/babun"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "cmder",
		URL:          []string{"https://github.com/cmderdev/cmder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "tbb",
		URL:          []string{"https://github.com/01org/tbb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "biopython",
		URL:          []string{"https://github.com/biopython/biopython"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
}
