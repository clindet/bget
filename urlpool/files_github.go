package urlpool

var githubRepos = []bgetFilesURLType{
	{
		Name:         "test/github_demo",
		URL:          []string{"https://github.com/Miachol/github_demo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "aligner/bwa",
		URL:          []string{"https://github.com/lh3/bwa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "aligner/star",
		URL:          []string{"https://github.com/alexdobin/STAR"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/bcftools",
		URL:          []string{"https://github.com/samtools/bcftools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "aligner/bowtie",
		URL:          []string{"https://github.com/BenLangmead/bowtie"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "aligner/bowtie2",
		URL:          []string{"https://github.com/BenLangmead/bowtie2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "aligner/tophat",
		URL:          []string{"https://github.com/infphilo/tophat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/varscan",
		URL:          []string{"https://github.com/Miachol/varscan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/picard",
		URL:          []string{"https://github.com/broadinstitute/picard"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/vcftools",
		URL:          []string{"https://github.com/vcftools/vcftools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/pindel",
		URL:          []string{"https://github.com/genome/pindel"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/lofreq",
		URL:          []string{"https://github.com/Miachol/lofreq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "aligner/hisat2",
		URL:          []string{"https://github.com/infphilo/hisat2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/freebayes",
		URL:          []string{"https://github.com/ekg/freebayes"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/macs",
		URL:          []string{"https://github.com/taoliu/MACS"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/bedtools2",
		URL:          []string{"https://github.com/arq5x/bedtools2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "other/sparsehash",
		URL:          []string{"https://github.com/sparsehash/sparsehash"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "assemble/abyss",
		URL:          []string{"https://github.com/bcgsc/abyss"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/bamtools",
		URL:          []string{"https://github.com/pezmaster31/bamtools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/breakdancer",
		URL:          []string{"https://github.com/genome/breakdancer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/htseq",
		URL:          []string{"https://github.com/simon-anders/htseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/seqtk",
		URL:          []string{"https://github.com/ndaniel/seqtk"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/delly",
		URL:          []string{"https://github.com/dellytools/delly"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "aligner/tmap",
		URL:          []string{"git://github.com/iontorrent/TMAP"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/somatic-sniper",
		URL:          []string{"https://github.com/genome/somatic-sniper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/bamutil",
		URL:          []string{"https://github.com/statgen/bamUtil"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/vcflib",
		URL:          []string{"https://github.com/vcflib/vcflib"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/jvarkit",
		URL:          []string{"https://github.com/lindenb/jvarkit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "qc/fastx_toolkit",
		URL:          []string{"https://github.com/agordon/fastx_toolkit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "other/libgtextutils",
		URL:          []string{"https://github.com/agordon/libgtextutils"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "assemble/trinityrnaseq",
		URL:          []string{"https://github.com/trinityrnaseq/trinityrnaseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "assemble/oases",
		URL:          []string{"https://github.com/dzerbino/oases"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "aligner/rum",
		URL:          []string{"https://github.com/itmat/rum"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "base/igraph",
		URL:          []string{"https://github.com/igraph/igraph"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "base/pxz",
		URL:          []string{"https://github.com/jnovy/pxz"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/cnvkit",
		URL:          []string{"https://github.com/etal/cnvkit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/speedseq",
		URL:          []string{"https://github.com/hall-lab/speedseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/cnvnator",
		URL:          []string{"https://github.com/abyzovlab/CNVnator"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/asap",
		URL:          []string{"https://github.com/DeplanckeLab/ASAP"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/mdseq",
		URL:          []string{"https://github.com/zjdaye/MDSeq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/sclvm",
		URL:          []string{"https://github.com/PMBio/scLVM"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/f-sclvm",
		URL:          []string{"https://github.com/PMBio/f-scLVM"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/bearscc",
		URL:          []string{"https://github.com/Miachol/bearscc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/singlesplice",
		URL:          []string{"https://github.com/jw156605/SingleSplice"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/isop",
		URL:          []string{"https://github.com/nghiavtr/ISOP"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/brie",
		URL:          []string{"https://github.com/huangyh09/brie"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/outrigger",
		URL:          []string{"https://github.com/YeoLab/outrigger"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/raceid",
		URL:          []string{"https://github.com/dgrun/RaceID"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/backspin",
		URL:          []string{"https://github.com/linnarsson-lab/BackSPIN"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/zifa",
		URL:          []string{"https://github.com/epierson9/ZIFA"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/seurat",
		URL:          []string{"https://github.com/satijalab/seurat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/rca",
		URL:          []string{"https://github.com/GIS-SP-Group/RCA"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/mimosca",
		URL:          []string{"https://github.com/asncd/MIMOSCA"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/tracer",
		URL:          []string{"https://github.com/teichlab/tracer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/scnorm",
		URL:          []string{"https://github.com/rhondabacher/SCnorm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/sleuth",
		URL:          []string{"https://github.com/pachterlab/sleuth"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "anno/oncotator",
		URL:          []string{"https://github.com/broadinstitute/oncotator"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "anno/ensembl-vep",
		URL:          []string{"https://github.com/Ensembl/ensembl-vep"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/fastq-tools",
		URL:          []string{"https://github.com/dcjones/fastq-tools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/arnapipe",
		URL:          []string{"https://github.com/HudsonAlpha/aRNAPipe"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "qc/trimgalore",
		URL:          []string{"https://github.com/FelixKrueger/TrimGalore"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "qc/multiqc",
		URL:          []string{"https://github.com/ewels/MultiQC"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/strawberry",
		URL:          []string{"https://github.com/ruolin/strawberry"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/fastp",
		URL:          []string{"https://github.com/OpenGene/fastp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "aligner/olego",
		URL:          []string{"https://github.com/chaolinzhanglab/olego"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "qc/chronqc",
		URL:          []string{"https://github.com/nilesh-tawari/ChronQC"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "aligner/dart",
		URL:          []string{"https://github.com/hsinnan75/DART"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "aligner/rhat",
		URL:          []string{"https://github.com/HIT-Bioinformatics/rHAT"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "anno/giggle",
		URL:          []string{"https://github.com/ryanlayer/giggle"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/facets",
		URL:          []string{"https://github.com/mskcc/facets"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/resm",
		URL:          []string{"https://github.com/deweylab/RSEM"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/radia",
		URL:          []string{"https://github.com/aradenbaugh/radia"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/manta",
		URL:          []string{"https://github.com/Illumina/manta"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/jaffa",
		URL:          []string{"https://github.com/Oshlack/JAFFA"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "assemble/marvel",
		URL:          []string{"https://github.com/schloi/MARVEL"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/picky",
		URL:          []string{"https://github.com/TheJacksonLaboratory/Picky"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/freec",
		URL:          []string{"https://github.com/BoevaLab/FREEC"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "aligner/taxmaps",
		URL:          []string{"https://github.com/nygenome/taxmaps"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/svaba",
		URL:          []string{"https://github.com/walaj/svaba"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/rop",
		URL:          []string{"https://github.com/smangul1/rop"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/mrnn",
		URL:          []string{"https://github.com/hendrixlab/mRNN"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "base/spack",
		URL:          []string{"https://github.com/spack/spack"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "base/bpipe",
		URL:          []string{"https://github.com/ssadedin/bpipe"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/mscentipede",
		URL:          []string{"https://github.com/rajanil/msCentipede"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/chia-pet2",
		URL:          []string{"https://github.com/GuipengLi/ChIA-PET2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/chromtime",
		URL:          []string{"https://github.com/ernstlab/ChromTime"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/dreg",
		URL:          []string{"https://github.com/Danko-Lab/dREG"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/genomedisco",
		URL:          []string{"https://github.com/kundajelab/genomedisco"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/mistic",
		URL:          []string{"https://github.com/iric-soft/MiSTIC"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/mosdepth",
		URL:          []string{"https://github.com/brentp/mosdepth"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "assemble/mecat",
		URL:          []string{"https://github.com/xiaochuanle/MECAT"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/vg",
		URL:          []string{"https://github.com/vgteam/vg"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "assemble/stringtie",
		URL:          []string{"https://github.com/gpertea/stringtie"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "app/iontorrent-suite",
		URL:          []string{"https://github.com/iontorrent/TS"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/fusioncatcher",
		URL:          []string{"https://github.com/ndaniel/fusioncatcher"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "anno/svscore",
		URL:          []string{"https://github.com/lganel/SVScore"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/svtools",
		URL:          []string{"https://github.com/hall-lab/svtools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/iseq",
		URL:          []string{"https://github.com/JhuangLab/iseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/deepvariant",
		URL:          []string{"https://github.com/google/deepvariant"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/sv2",
		URL:          []string{"https://github.com/dantaki/SV2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/mutscan",
		URL:          []string{"https://github.com/OpenGene/MutScan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/breakmer",
		URL:          []string{"https://github.com/ccgd-profile/BreaKmer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "qc/detin",
		URL:          []string{"https://github.com/broadinstitute/deTiN"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/cn_learn",
		URL:          []string{"https://github.com/girirajanlab/CN_Learn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/wham",
		URL:          []string{"https://github.com/zeeev/wham"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "anno/gemini",
		URL:          []string{"https://github.com/arq5x/gemini"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "anno/pcgr",
		URL:          []string{"https://github.com/sigven/pcgr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "anno/annovarr",
		URL:          []string{"https://github.com/JhuangLab/annovarR"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "anno/bystro",
		URL:          []string{"https://github.com/akotlar/bystro"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "anno/pathopredictor",
		URL:          []string{"https://github.com/samesense/pathopredictor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/agfusion",
		URL:          []string{"https://github.com/murphycj/AGFusion"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "web/genomeuplot",
		URL:          []string{"https://github.com/gaitat/GenomeUPlot"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/breakpointsurveyor",
		URL:          []string{"https://github.com/ding-lab/BreakPointSurveyor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "rlib/chimeraviz",
		URL:          []string{"https://github.com/stianlagstad/chimeraviz"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "web/protein/pv",
		URL:          []string{"https://github.com/biasmv/pv"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "other/g2s",
		URL:          []string{"https://github.com/genome-nexus/g2s"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "web/ngb",
		URL:          []string{"https://github.com/epam/NGB"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/ncboost",
		URL:          []string{"https://github.com/RausellLab/NCBoost"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/leafcutter",
		URL:          []string{"https://github.com/davidaknowles/leafcutter"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/irfinder",
		URL:          []string{"https://github.com/williamritchie/IRFinder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/mmsplice",
		URL:          []string{"https://github.com/gagneurlab/MMSplice"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/ballgown",
		URL:          []string{"https://github.com/alyssafrazee/ballgown"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/mixcr",
		URL:          []string{"https://github.com/milaboratory/mixcr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/topconfects",
		URL:          []string{"https://github.com/pfh/topconfects"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/plier",
		URL:          []string{"https://github.com/wgmao/PLIER"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/viral-ngs",
		URL:          []string{"https://github.com/broadinstitute/viral-ngs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/qap",
		URL:          []string{"https://github.com/mingjiewang/qap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/vifi",
		URL:          []string{"https://github.com/namphuon/ViFi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/micropro",
		URL:          []string{"https://github.com/zifanzhu/MicroPro"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/feast",
		URL:          []string{"https://github.com/cozygene/FEAST"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/mcorr",
		URL:          []string{"https://github.com/kussell-lab/mcorr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/dropclust",
		URL:          []string{"https://github.com/debsin/dropClust"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/trendsceek",
		URL:          []string{"https://github.com/edsgard/trendsceek"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "doc/awesome-single-cell",
		URL:          []string{"https://github.com/seandavi/awesome-single-cell"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/saver",
		URL:          []string{"https://github.com/mohuangx/SAVER"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/cellsius",
		URL:          []string{"https://github.com/Novartis/CellSIUS"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/scrabble",
		URL:          []string{"https://github.com/software-github/SCRABBLE"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/melissa",
		URL:          []string{"https://github.com/andreaskapou/Melissa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "other/paga",
		URL:          []string{"https://github.com/theislab/paga"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/clonealign",
		URL:          []string{"https://github.com/kieranrcampbell/clonealign"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/cellfishing.jl",
		URL:          []string{"https://github.com/bicycle1885/CellFishing.jl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/viper",
		URL:          []string{"https://github.com/ChenMengjie/VIPER"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/scgen",
		URL:          []string{"https://github.com/theislab/scgen"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/conos",
		URL:          []string{"https://github.com/hms-dbmi/conos"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/magic",
		URL:          []string{"https://github.com/KrishnaswamyLab/MAGIC"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/zinbwave",
		URL:          []string{"https://github.com/drisso/zinbwave"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/simlr_py",
		URL:          []string{"https://github.com/bowang87/SIMLR_PY"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/dca",
		URL:          []string{"https://github.com/theislab/dca"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/scvi",
		URL:          []string{"https://github.com/YosefLab/scVI"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/phenograph",
		URL:          []string{"https://github.com/jacoblevine/PhenoGraph"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "doc/splatter-paper",
		URL:          []string{"https://github.com/Oshlack/splatter-paper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/deepnovo-dia",
		URL:          []string{"https://github.com/nh2tran/DeepNovo-DIA"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/caveman",
		URL:          []string{"https://github.com/funpopgen/CaVEMaN"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "base/bcbio-nextgen",
		URL:          []string{"https://github.com/chapmanb/bcbio-nextgen"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "app/orange3",
		URL:          []string{"https://github.com/biolab/orange3"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/bedops",
		URL:          []string{"https://github.com/bedops/bedops"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/maftools",
		URL:          []string{"https://github.com/PoisonAlien/maftools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/samblaster",
		URL:          []string{"https://github.com/GregoryFaust/samblaster"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "base/ngstk",
		URL:          []string{"https://github.com/JhuangLab/ngstk"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "base/bioinstaller",
		URL:          []string{"https://github.com/JhuangLab/BioInstaller"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/chromhmm",
		URL:          []string{"https://github.com/jernst98/ChromHMM"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/biobloom",
		URL:          []string{"https://github.com/bcgsc/biobloom"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/phenopredict",
		URL:          []string{"https://github.com/leekgroup/phenopredict"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "other/recount",
		URL:          []string{"https://github.com/leekgroup/recount"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "stats/lsmm",
		URL:          []string{"https://github.com/mingjingsi/LSMM"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/vcf2maf",
		URL:          []string{"https://github.com/mskcc/vcf2maf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "web/r2d3",
		URL:          []string{"https://github.com/rstudio/r2d3"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "rlib/liteq",
		URL:          []string{"https://github.com/r-lib/liteq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "web/relaxed",
		URL:          []string{"https://github.com/RelaxedJS/ReLaXed"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "rlib/dash",
		URL:          []string{"https://github.com/jonocarroll/dash"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "rlib/threadpool",
		URL:          []string{"https://github.com/rdpeng/threadpool"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/rapid",
		URL:          []string{"https://github.com/ZhiGroup/RaPID"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "other/sellerslab-gemini",
		URL:          []string{"https://github.com/sellerslab/gemini"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "other/confined",
		URL:          []string{"https://github.com/cozygene/CONFINED"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/marginphase",
		URL:          []string{"https://github.com/benedictpaten/marginPhase"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/chicmaxima",
		URL:          []string{"https://github.com/yousra291987/ChiCMaxima"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/circbrain",
		URL:          []string{"https://github.com/yangence/circBrain"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "aligner/bazam",
		URL:          []string{"https://github.com/ssadedin/bazam"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/conbase",
		URL:          []string{"https://github.com/conbase/conbase"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "qc/3dchromatin_replicateqc",
		URL:          []string{"https://github.com/kundajelab/3DChromatin_ReplicateQC"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "stats/i-boost",
		URL:          []string{"https://github.com/alexwky/I-Boost"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/bin3c",
		URL:          []string{"https://github.com/cerebis/bin3C"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/dstruct",
		URL:          []string{"https://github.com/AviranLab/dStruct"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/skmer",
		URL:          []string{"https://github.com/shahab-sarmashghi/Skmer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/iguide",
		URL:          []string{"https://github.com/cnobles/iGUIDE"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/plyranges",
		URL:          []string{"https://github.com/sa-lee/plyranges"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/forge",
		URL:          []string{"https://github.com/langmead-lab/FORGe"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/se-mei",
		URL:          []string{"https://github.com/dpryan79/SE-MEI"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/anchor",
		URL:          []string{"https://github.com/GuanLab/Anchor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "mut/advntr",
		URL:          []string{"https://github.com/mehrdadbakhtiari/adVNTR"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "stats/ldsc",
		URL:          []string{"https://github.com/bulik/ldsc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "other/bigstitcher",
		URL:          []string{"https://github.com/PreibischLab/BigStitcher"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/ivtnmr",
		URL:          []string{"https://github.com/systemsnmr/ivtnmr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/divers",
		URL:          []string{"https://github.com/hym0405/DIVERS"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/prosit",
		URL:          []string{"https://github.com/kusterlab/prosit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "exp/deepcell-tf",
		URL:          []string{"https://github.com/vanvalenlab/deepcell-tf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "image/cdeep3m",
		URL:          []string{"https://github.com/CRBS/cdeep3m"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "image/unet-segmentation",
		URL:          []string{"https://github.com/lmb-freiburg/Unet-Segmentation"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "image/cellprofiler",
		URL:          []string{"https://github.com/CellProfiler/CellProfiler"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "image/mask_rcnn",
		URL:          []string{"https://github.com/matterport/Mask_RCNN"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "image/leap",
		URL:          []string{"https://github.com/talmo/leap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "image/in-silico-labeling",
		URL:          []string{"https://github.com/google/in-silico-labeling"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "doc/trackviewer.documentation",
		URL:          []string{"https://github.com/jianhong/trackViewer.documentation"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/cistopic",
		URL:          []string{"https://github.com/aertslab/cistopic"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/selene",
		URL:          []string{"https://github.com/FunctionLab/selene"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "assemble/sda",
		URL:          []string{"https://github.com/mvollger/SDA"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "image/fmriprep",
		URL:          []string{"https://github.com/poldracklab/fmriprep"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "other/unifrac",
		URL:          []string{"https://github.com/biocore/unifrac"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "rlib/gvmap",
		URL:          []string{"https://github.com/ytdai/gvmap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "rlib/easysvg",
		URL:          []string{"https://github.com/ytdai/easySVG"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "rlib/geogrid",
		URL:          []string{"https://github.com/jbaileyh/geogrid"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "web/clustergrammer",
		URL:          []string{"https://github.com/MaayanLab/clustergrammer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/chromvar",
		URL:          []string{"https://github.com/GreenleafLab/chromVAR"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "web/echarts",
		URL:          []string{"https://github.com/ecomfe/echarts"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "web/plotly.js",
		URL:          []string{"https://github.com/plotly/plotly.js"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "stats/qvalue",
		URL:          []string{"https://github.com/StoreyLab/qvalue"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "rlib/genvisr",
		URL:          []string{"https://github.com/griffithlab/GenVisR"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "rlib/r-color-palettes",
		URL:          []string{"https://github.com/EmilHvitfeldt/r-color-palettes"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "web/opencpu",
		URL:          []string{"https://github.com/opencpu/opencpu"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "rlib/ggthemr",
		URL:          []string{"https://github.com/cttobin/ggthemr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "rlib/paletter",
		URL:          []string{"https://github.com/AndreaCirilloAC/paletter"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "rlib/ggdag",
		URL:          []string{"https://github.com/malcolmbarrett/ggdag"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "rlib/ggseqlogo",
		URL:          []string{"https://github.com/omarwagih/ggseqlogo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "web/three.js",
		URL:          []string{"https://github.com/mrdoob/three.js"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "web/higlass",
		URL:          []string{"https://github.com/hms-dbmi/higlass"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "app/babun",
		URL:          []string{"https://github.com/babun/babun"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "app/cmder",
		URL:          []string{"https://github.com/cmderdev/cmder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "base/tbb",
		URL:          []string{"https://github.com/01org/tbb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "seq/biopython",
		URL:          []string{"https://github.com/biopython/biopython"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
}
