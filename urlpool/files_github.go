package urlpool

var githubRepos = []bgetFilesURLType{
	{
		Name:         "test/github-demo",
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
		Name:         "qc/fastx-toolkit",
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
		Name:         "mut/cn-learn",
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
		Name:         "exp/simlr-py",
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
		Name:         "qc/3dchromatin-replicateqc",
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
		Name:         "image/mask-rcnn",
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
		Name:         "doc/trackviewer",
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
	{
		Name:         "doc/awosome-bioinformatics",
		URL:          []string{"https://github.com/openbiox/awosome-bioinformatics"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/immclass2019",
		URL:          []string{"https://github.com/klinkelab/immclass2019"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/edgescaping",
		URL:          []string{"https://github.com/bhusain/edgescaping"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nacho",
		URL:          []string{"https://github.com/mcanouil/nacho"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/imap",
		URL:          []string{"https://github.com/jkimlab/imap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/treeexp-tutorial",
		URL:          []string{"https://github.com/jingwyang/treeexp-tutorial"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rnaindel",
		URL:          []string{"https://github.com/stjude/rnaindel"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pygtftk",
		URL:          []string{"https://github.com/dputhier/pygtftk"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hypocotyl-unet",
		URL:          []string{"https://github.com/biomag-lab/hypocotyl-unet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dligand2",
		URL:          []string{"https://github.com/sysu-yanglab/dligand2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/protein-recon-mcriemman",
		URL:          []string{"https://github.com/xubiaopeng/protein_recon_mcriemman"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ephagen",
		URL:          []string{"https://github.com/m4merg/ephagen"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/spm",
		URL:          []string{"https://github.com/dalton386/spm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/flexor",
		URL:          []string{"https://github.com/flexorsoftware/flexor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pnc",
		URL:          []string{"https://github.com/nwpu-903pr/pnc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/crtl",
		URL:          []string{"https://github.com/wangshanshancqu/crtl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/x-cnn",
		URL:          []string{"https://github.com/ernstlab/x-cnn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/blant",
		URL:          []string{"https://github.com/waynebhayes/blant"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/simurg",
		URL:          []string{"https://github.com/iferres/simurg"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/grailbio-bio",
		URL:          []string{"https://github.com/grailbio/bio"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/omicpred",
		URL:          []string{"https://github.com/yaluwen/omicpred"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phydelity",
		URL:          []string{"https://github.com/alvinxhan/phydelity"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mr-demuxy",
		URL:          []string{"https://github.com/lefeverde/mr_demuxy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ramf",
		URL:          []string{"https://github.com/mchiapello/ramf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/modelexplorer",
		URL:          []string{"https://github.com/theangryfox/modelexplorer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vari-merge",
		URL:          []string{"https://github.com/cosmo-team/cosmo/tree/vari-merge"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hops",
		URL:          []string{"https://github.com/rondolab/hops"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ssrsc",
		URL:          []string{"https://github.com/csjunxu/ssrsc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/howdesbt",
		URL:          []string{"https://github.com/medvedevgroup/howdesbt"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mtbgt",
		URL:          []string{"https://github.com/kamelang/mtbgt"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ga-bagging-svm",
		URL:          []string{"https://github.com/qust-aibbdrc/ga-bagging-svm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rbbt-workflows",
		URL:          []string{"https://github.com/rbbt-workflows"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gppf",
		URL:          []string{"https://github.com/algolab/gppf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mixmpln",
		URL:          []string{"https://github.com/sahatava/mixmpln"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hintra",
		URL:          []string{"https://github.com/sahandk/hintra"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dbgap2x",
		URL:          []string{"https://github.com/gversmee/dbgap2x"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sigprofilermatrixgenerator",
		URL:          []string{"https://github.com/alexandrovlab/sigprofilermatrixgenerator"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/genome-topology-network",
		URL:          []string{"https://github.com/0232/genome_topology_network"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/malini",
		URL:          []string{"https://github.com/pradlanka/malini"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/schinter",
		URL:          []string{"https://github.com/bmilab/schinter"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dnam-based-age-predictor",
		URL:          []string{"https://github.com/qzhang314/dnam-based-age-predictor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/merlot",
		URL:          []string{"https://github.com/soedinglab/merlot"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mitoimp",
		URL:          []string{"https://github.com/omics-tools/mitoimp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/micropro",
		URL:          []string{"https://github.com/zifanzhu/micropro"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gwasrapidd",
		URL:          []string{"https://github.com/ramiromagno/gwasrapidd"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepathology",
		URL:          []string{"https://github.com/sharifbioinf/deepathology"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/submito-xgboost",
		URL:          []string{"https://github.com/qust-aibbdrc/submito-xgboost"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepseenet",
		URL:          []string{"https://github.com/ncbi-nlp/deepseenet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biobert-pretrained",
		URL:          []string{"https://github.com/naver/biobert-pretrained"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biobert",
		URL:          []string{"https://github.com/dmis-lab/biobert"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/crossmapper",
		URL:          []string{"https://github.com/gabaldonlab/crossmapper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/alleleanalyzer",
		URL:          []string{"https://github.com/keoughkath/alleleanalyzer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tfea.chip-downloads",
		URL:          []string{"https://github.com/lauraps1/tfea.chip_downloads"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/crispritz",
		URL:          []string{"https://github.com/pinellolab/crispritzhttps://github.com/infomics/crispritz"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pll-modules",
		URL:          []string{"https://github.com/ddarriba/pll-modules"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/raxml-ng",
		URL:          []string{"https://github.com/amkozlov/raxml-ng"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tbe",
		URL:          []string{"https://github.com/lutteropp/raxml-ng/tree/tbe"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/diffuse",
		URL:          []string{"https://github.com/haochenucr/diffuse"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/brapes",
		URL:          []string{"https://github.com/yoseflab/brapes"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mind",
		URL:          []string{"https://github.com/randel/mind"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/align3d",
		URL:          []string{"https://github.com/heltilda/align3d"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pattrec",
		URL:          []string{"https://github.com/irotero/pattrec"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/data-of-genome-rearrangement",
		URL:          []string{"https://github.com/wangjuanimu/data-of-genome-rearrangement"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/leaf",
		URL:          []string{"https://github.com/uwrit/leaf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rcsm",
		URL:          []string{"https://github.com/jasonlinchina/rcsm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/stripes",
		URL:          []string{"https://github.com/carlborggenomics/stripes"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ndd",
		URL:          []string{"https://github.com/nrohani/ndd"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tn-test",
		URL:          []string{"https://github.com/jessemzhang/tn_test"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rosids",
		URL:          []string{"https://github.com/deisejpg/rosids"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/theta",
		URL:          []string{"https://github.com/viro-tls/theta"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sigma",
		URL:          []string{"https://github.com/lrgr/sigma"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ld",
		URL:          []string{"https://github.com/tanviralambd/ld"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/chipseqspikeinfree",
		URL:          []string{"https://github.com/stjude/chipseqspikeinfree"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/spectrum",
		URL:          []string{"https://github.com/birl/spectrum"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pourrna",
		URL:          []string{"https://github.com/viennarna/pourrna"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mdine",
		URL:          []string{"https://github.com/kevinmcgregor/mdine"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/synquant",
		URL:          []string{"https://github.com/yu-lab-vt/synquant"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/spedeimportance",
		URL:          []string{"https://github.com/lm-ugent/spedeimportance"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cwsdtwnano",
		URL:          []string{"https://github.com/icthrm/cwsdtwnano"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/super-resolution-adversarial-defense",
		URL:          []string{"https://github.com/aamir-mustafa/super-resolution-adversarial-defense"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/evidenceaggregateddriverranking",
		URL:          []string{"https://github.com/sage-bionetworks/evidenceaggregateddriverranking"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vapper",
		URL:          []string{"https://github.com/pgb-liv/vapper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vertical-tracts",
		URL:          []string{"https://github.com/brainlife/vertical_tracts"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mipepid",
		URL:          []string{"https://github.com/mindai/mipepid"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mmvec",
		URL:          []string{"https://github.com/biocore/mmvec"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ritan",
		URL:          []string{"https://github.com/mtzimmer/ritan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dispot",
		URL:          []string{"https://github.com/korkinlab/dispot"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dde-bd",
		URL:          []string{"https://github.com/cbskust/dde_bd"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rasp",
		URL:          []string{"https://github.com/sculab/rasp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/simka",
		URL:          []string{"https://github.com/gatb/simka"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hh-suite",
		URL:          []string{"https://github.com/soedinglab/hh-suite"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/maizesnpdb",
		URL:          []string{"https://github.com/venyao/maizesnpdb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/croco",
		URL:          []string{"https://github.com/cschmidtlab/croco"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/visor",
		URL:          []string{"https://github.com/davidebolo1993/visor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scdesign",
		URL:          []string{"https://github.com/vivianstats/scdesign"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/infiniumpurify",
		URL:          []string{"https://github.com/xiaoqizheng/infiniumpurify"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/targetmine-gradle",
		URL:          []string{"https://github.com/chenyian-nibio/targetmine-gradle"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/waternn",
		URL:          []string{"https://github.com/wizard1203/waternn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hasn",
		URL:          []string{"https://github.com/shenjianbing/hasn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/contextual-regression-for-circrna",
		URL:          []string{"https://github.com/chl556/contextual_regression_for_circrna"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/romop",
		URL:          []string{"https://github.com/benglicksberg/romop"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rsmlm",
		URL:          []string{"https://github.com/jeremypike/rsmlm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rsmlm-tutorials",
		URL:          []string{"https://github.com/jeremypike/rsmlm-tutorials"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/telescope",
		URL:          []string{"https://github.com/mlbendall/telescope"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metumpx-bin",
		URL:          []string{"https://github.com/hasaniqbal777/metumpx-bin"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/high-quality-ellipse-detection",
		URL:          []string{"https://github.com/alanlusun/high-quality-ellipse-detection"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gbwt",
		URL:          []string{"https://github.com/jltsiren/gbwt"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gcsa2",
		URL:          []string{"https://github.com/jltsiren/gcsa2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/allelic-inclusion",
		URL:          []string{"https://github.com/jasonacarter/allelic_inclusion"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/slr",
		URL:          []string{"https://github.com/luojunwei/slr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tmtscca",
		URL:          []string{"https://github.com/dulei323/tmtscca"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ragp",
		URL:          []string{"https://github.com/missuse/ragp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/spdrank",
		URL:          []string{"https://github.com/akiyamalab/spdrank"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hupan",
		URL:          []string{"https://github.com/sjtu-cgm/hupan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/anatomy-modality-decomposition",
		URL:          []string{"https://github.com/agis85/anatomy_modality_decomposition"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/squigglekit",
		URL:          []string{"https://github.com/psy-fer/squigglekit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phylomagnet",
		URL:          []string{"https://github.com/maxemil/phylomagnet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dpre",
		URL:          []string{"https://github.com/loaloaf/dpre"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pclean-release",
		URL:          []string{"https://github.com/aimeed90/pclean_release"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pathwaymatcher",
		URL:          []string{"https://github.com/pathwayanalysisplatform/pathwaymatcher"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nnlda",
		URL:          []string{"https://github.com/gao793583308/nnlda"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bonita",
		URL:          []string{"https://github.com/thakar-lab/bonita"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/srhtests",
		URL:          []string{"https://github.com/roblanf/srhtests"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mct",
		URL:          []string{"https://github.com/elkebir-group/mct"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cobia",
		URL:          []string{"https://github.com/bertrand-lab/cobia"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ptuneos",
		URL:          []string{"https://github.com/bm2-lab/ptuneos"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/uniconsig",
		URL:          []string{"https://github.com/wangxlab/uniconsig"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sequence-prediction-using-cnn-and-lstms",
		URL:          []string{"https://github.com/rajkumar1501/sequence-prediction-using-cnn-and-lstms"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/njmerge",
		URL:          []string{"https://github.com/ekmolloy/njmerge"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scanfold",
		URL:          []string{"https://github.com/moss-lab/scanfold"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepdna",
		URL:          []string{"https://github.com/rongjiewang/deepdna"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepram",
		URL:          []string{"https://github.com/medchaabane/deepram"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/chaperism",
		URL:          []string{"https://github.com/bioinflab/chaperism"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ularcirc",
		URL:          []string{"https://github.com/vccri/ularcirc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/differential-privacy-genomic-inference-attack",
		URL:          []string{"https://github.com/nourmadhoun/differential-privacy-genomic-inference-attack"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/regbase",
		URL:          []string{"https://github.com/mulinlab/regbase"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/discorhythm",
		URL:          []string{"https://github.com/matthewcarlucci/discorhythm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hatspil",
		URL:          []string{"https://github.com/dodomorandi/hatspil"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/integratedphasing",
		URL:          []string{"https://github.com/vibansal/integratedphasing"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nexus",
		URL:          []string{"https://github.com/priyamdas2/nexus"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/thermonet",
		URL:          []string{"https://github.com/suyufeng/thermonet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/treecluster",
		URL:          []string{"https://github.com/niemasd/treecluster"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/microaggregation-based-anonymization-tool",
		URL:          []string{"https://github.com/crisesurv/microaggregation-based_anonymization_tool"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/multitable-review",
		URL:          []string{"https://github.com/krisrs1128/multitable_review"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/beam-propagation-method",
		URL:          []string{"https://github.com/bunpc/beam-propagation-method"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/episafari",
		URL:          []string{"https://github.com/harmancilab/episafari"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rnn-for-membrane-protein-types-prediction",
		URL:          []string{"https://github.com/yellowcardd/rnn-for-membrane-protein-types-prediction"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/zing",
		URL:          []string{"https://github.com/weng-lab/zing"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/organsegrstn",
		URL:          []string{"https://github.com/198808xc/organsegrstn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ajia",
		URL:          []string{"https://github.com/lyotvincent/ajia"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/samdude",
		URL:          []string{"https://github.com/ihwang/samdude"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ge-lda-survival",
		URL:          []string{"https://github.com/nitsanluke/ge-lda-survival"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/online-chem",
		URL:          []string{"https://github.com/bigchem/online-chem"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cepics",
		URL:          []string{"https://github.com/gaolabxdu/cepics"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fisheye-orb-slam",
		URL:          []string{"https://github.com/lsyads/fisheye-orb-slam"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/chordomics",
		URL:          []string{"https://github.com/kevinmcdonnell6/chordomics"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/swpepnovo",
		URL:          []string{"https://github.com/chuangli99/swpepnovo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/syntenybrowser",
		URL:          []string{"https://github.com/thejacksonlaboratory/syntenybrowser"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ribotricer",
		URL:          []string{"https://github.com/smithlabcode/ribotricer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ribotricer-results",
		URL:          []string{"https://github.com/smithlabcode/ribotricer-results"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/modeltest",
		URL:          []string{"https://github.com/ddarriba/modeltest"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/amplisolve",
		URL:          []string{"https://github.com/dkleftogi/amplisolve"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/circcode",
		URL:          []string{"https://github.com/pssun/circcode"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lhqxinghun/bioinformatics",
		URL:          []string{"https://github.com/lhqxinghun/bioinformatics"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bamfa",
		URL:          []string{"https://github.com/markusheinonen/bamfa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/omhismb2019",
		URL:          []string{"https://github.com/kingsford-group/omhismb2019"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dipartite",
		URL:          []string{"https://github.com/mohammad-vahed/dipartite"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/circmeta",
		URL:          []string{"https://github.com/lichen-lab/circmeta"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biwalklda",
		URL:          []string{"https://github.com/screamer/biwalklda"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sdg",
		URL:          []string{"https://github.com/bioinfologics/sdg"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/efp-seq-browser",
		URL:          []string{"https://github.com/bioanalyticresource/efp-seq-browser"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/baynorm",
		URL:          []string{"https://github.com/wt215/baynorm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/baynorm-papercode",
		URL:          []string{"https://github.com/wt215/baynorm_papercode"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bnnr",
		URL:          []string{"https://github.com/bioinformaticscsu/bnnr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/chromdragonn",
		URL:          []string{"https://github.com/kundajelab/chromdragonn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/clinical-sentences",
		URL:          []string{"https://github.com/nlpie/clinical-sentences"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mutationtcn",
		URL:          []string{"https://github.com/ha01994/mutationtcn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nicebot",
		URL:          []string{"https://github.com/tntlfreiburg/nicebot"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ionspattern",
		URL:          []string{"https://github.com/vitek-lab/ionspattern"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/qubic2",
		URL:          []string{"https://github.com/osu-bmbl/qubic2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cometsc",
		URL:          []string{"https://github.com/msingerlab/cometsc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/meltos",
		URL:          []string{"https://github.com/ih-lab/meltos"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/taba",
		URL:          []string{"https://github.com/azevedolab/taba"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ltmgsca",
		URL:          []string{"https://github.com/zy26/ltmgsca"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phyloctcflooping",
		URL:          []string{"https://github.com/morphos30/phyloctcflooping"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mitre",
		URL:          []string{"https://github.com/gerberlab/mitre"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepcon",
		URL:          []string{"https://github.com/badriadhikari/deepcon"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/breastcancerclassifier",
		URL:          []string{"https://github.com/nyukat/breastcancerclassifier"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dcm",
		URL:          []string{"https://github.com/liuji93/dcm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mcs",
		URL:          []string{"https://github.com/rezamash/mcs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sourmash",
		URL:          []string{"https://github.com/dib-lab/sourmash"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/funvar",
		URL:          []string{"https://github.com/lasigebiotm/funvar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cner",
		URL:          []string{"https://github.com/ge11232002/cner"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/stra-net",
		URL:          []string{"https://github.com/ashleylqx/stra-net"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pluto",
		URL:          []string{"https://github.com/weipeneghu/pluto"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/shaker",
		URL:          []string{"https://github.com/backofenlab/shaker"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scipion",
		URL:          []string{"https://github.com/i2pc/scipion"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hurdlenormal",
		URL:          []string{"https://github.com/amcdavid/hurdlenormal"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mines",
		URL:          []string{"https://github.com/yeolab/mines"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/maizedig",
		URL:          []string{"https://github.com/maize-genetics-and-genomics-database/maizedig"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bionitio",
		URL:          []string{"https://github.com/bionitio-team/bionitio"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/krab-znf",
		URL:          []string{"https://github.com/mi2datalab/krab_znf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/visfeature",
		URL:          []string{"https://github.com/wangjun1996/visfeature"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/end2end-all-conv",
		URL:          []string{"https://github.com/lishen/end2end-all-conv"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/newton-method-mds",
		URL:          []string{"https://github.com/skypes/newton-method-mds"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pepvis",
		URL:          []string{"https://github.com/inpacdb/pepvis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/blmrm",
		URL:          []string{"https://github.com/jingxiemizzou/blmrm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sora",
		URL:          []string{"https://github.com/biohpc/sora"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cardiacpbpk",
		URL:          []string{"https://github.com/jszlek/cardiacpbpk"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ngly",
		URL:          []string{"https://github.com/bioinformaticsml/ngly"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/set-cover-tools",
		URL:          []string{"https://github.com/yichaoou/set_cover_tools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/driversub",
		URL:          []string{"https://github.com/jianingxi/driversub"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/aten",
		URL:          []string{"https://github.com/ningshi/aten"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cline",
		URL:          []string{"https://github.com/rodolfoallendes/cline"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/joint-marker-free-alignment",
		URL:          []string{"https://github.com/icthrm/joint-marker-free-alignment"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/awfisher",
		URL:          []string{"https://github.com/caleb-huo/awfisher"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/strainhub",
		URL:          []string{"https://github.com/abschneider/strainhub"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mimi",
		URL:          []string{"https://github.com/graemefox/mimi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/apples",
		URL:          []string{"https://github.com/balabanmetin/apples"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mark2curedata",
		URL:          []string{"https://github.com/sulab/mark2curedata"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/m2c-rel-nb",
		URL:          []string{"https://github.com/gtsueng/m2c_rel_nb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pymethylprocess",
		URL:          []string{"https://github.com/christensen-lab-dartmouth/pymethylprocess"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ebi-metabolights",
		URL:          []string{"https://github.com/ebi-metabolights"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hyper",
		URL:          []string{"https://github.com/montilab/hyper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phosphoproteome-prediction",
		URL:          []string{"https://github.com/guanlab/phosphoproteome_prediction"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/immucc",
		URL:          []string{"https://github.com/wuaipinglab/immucc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lstd",
		URL:          []string{"https://github.com/cassie94/lstd/tree/lstd"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/yalff",
		URL:          []string{"https://github.com/yhhshb/yalff"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gkmexplain",
		URL:          []string{"https://github.com/kundajelab/gkmexplain"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ccbgpipe",
		URL:          []string{"https://github.com/jade-nhri/ccbgpipe"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bjass",
		URL:          []string{"https://github.com/yiucla/bjass"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepligand",
		URL:          []string{"https://github.com/gifford-lab/deepligand"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mlrmbo",
		URL:          []string{"https://github.com/mlr-org/mlrmbo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/abus-code",
		URL:          []string{"https://github.com/nawang0226/abus_code"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lgcm",
		URL:          []string{"https://github.com/aiying0512/lgcm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/epivan",
		URL:          []string{"https://github.com/hzy95/epivan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gsl",
		URL:          []string{"https://github.com/fjr9516/gsl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tcsm",
		URL:          []string{"https://github.com/lrgr/tcsm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dncon2",
		URL:          []string{"https://github.com/multicom-toolbox/dncon2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/proxdist",
		URL:          []string{"https://github.com/klkeys/proxdist"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/crm",
		URL:          []string{"https://github.com/icthrm/crm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/zombi",
		URL:          []string{"https://github.com/aadavin/zombi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ewca",
		URL:          []string{"https://github.com/rongquanwang/ewca"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/flint",
		URL:          []string{"https://github.com/camilo-v/flint"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tcgenerators",
		URL:          []string{"https://github.com/bielcardona/tcgenerators"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/brm",
		URL:          []string{"https://github.com/huanglikun/brm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vipr-mat-peptide",
		URL:          []string{"https://github.com/virusbrc/vipr_mat_peptide"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/assexon",
		URL:          []string{"https://github.com/yhadevol/assexon"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bgsc",
		URL:          []string{"https://github.com/grosselab/bgsc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/neep",
		URL:          []string{"https://github.com/thecodingdoc/neep"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sinbad",
		URL:          []string{"https://github.com/scwest/sinbad"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cnape",
		URL:          []string{"https://github.com/wanglabhkust/cnape"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/asja",
		URL:          []string{"https://github.com/huanglab-fudan/asja"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ncphlda",
		URL:          []string{"https://github.com/bryanze/ncphlda"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pyro-nn",
		URL:          []string{"https://github.com/csyben/pyro-nn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rbv",
		URL:          []string{"https://github.com/whitneywhitford/rbv"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/eagle",
		URL:          []string{"https://github.com/evansgao/eagle"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sap01-tgs-lite-supplem",
		URL:          []string{"https://github.com/sap01/tgs-lite-supplem"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ananas",
		URL:          []string{"https://github.com/lesniak43/ananas"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/circacompare",
		URL:          []string{"https://github.com/rwparsons/circacompare"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/harmony",
		URL:          []string{"https://github.com/immunogenomics/harmony"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deephlapan",
		URL:          []string{"https://github.com/jiujiezz/deephlapan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/duett",
		URL:          []string{"https://github.com/bagherilab/duett"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/reglinker",
		URL:          []string{"https://github.com/murali-group/reglinker"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rnaeditingindexer",
		URL:          []string{"https://github.com/a2iediting/rnaeditingindexer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tada",
		URL:          []string{"https://github.com/tada-alg/tada"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/denoptim",
		URL:          []string{"https://github.com/denoptim-project/denoptim"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lightdock",
		URL:          []string{"https://github.com/brianjimenez/lightdock"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lightdock-bm5",
		URL:          []string{"https://github.com/brianjimenez/lightdock_bm5"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/amplimap",
		URL:          []string{"https://github.com/koelling/amplimap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cerebro",
		URL:          []string{"https://github.com/romanhaa/cerebro"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cerebroapp",
		URL:          []string{"https://github.com/romanhaa/cerebroapp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/qpa",
		URL:          []string{"https://github.com/github-gs/qpa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scrnaseq-cell-cluster-labeling",
		URL:          []string{"https://github.com/jdime/scrnaseq_cell_cluster_labeling"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/distanced",
		URL:          []string{"https://github.com/thackmann/distanced"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/combat-tb",
		URL:          []string{"https://github.com/combat-tb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/doepipeline",
		URL:          []string{"https://github.com/clicumu/doepipeline"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/warp",
		URL:          []string{"https://github.com/cramerlab/warp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bionev",
		URL:          []string{"https://github.com/xiangyue9607/bionev"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/atc",
		URL:          []string{"https://github.com/dqwei-lab/atc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/uorf-tools",
		URL:          []string{"https://github.com/biochemistry1-ffm/uorf-tools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ann-solo",
		URL:          []string{"https://github.com/bittremieux/ann-solo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mtmr-net",
		URL:          []string{"https://github.com/captainwilliam/mtmr-net"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/g3lollipop.js",
		URL:          []string{"https://github.com/g3viz/g3lollipop.js"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hpccs",
		URL:          []string{"https://github.com/cepid-cces/hpccs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/outlier-in-blast-hits",
		URL:          []string{"https://github.com/shahnidhi/outlier_in_blast_hits"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mirureader",
		URL:          []string{"https://github.com/phglab/mirureader"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/prism",
		URL:          []string{"https://github.com/dohlee/prism"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tidehunter",
		URL:          []string{"https://github.com/yangao07/tidehunter"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/treehouse",
		URL:          []string{"https://github.com/jlsteenwyk/treehouse"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mastermsm",
		URL:          []string{"https://github.com/daviddesancho/mastermsm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/matscholar",
		URL:          []string{"https://github.com/materialsintelligence/matscholar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/canet",
		URL:          []string{"https://github.com/xmengli999/canet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/miams",
		URL:          []string{"https://github.com/bialimed/miams"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/response-logic",
		URL:          []string{"https://github.com/grosstor/response-logic"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/response-logic-projects",
		URL:          []string{"https://github.com/grosstor/response-logic-projects"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/prodigy",
		URL:          []string{"https://github.com/shamir-lab/prodigy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mtgosc",
		URL:          []string{"https://github.com/ne1s0n/mtgosc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/csmd",
		URL:          []string{"https://github.com/liuyu8721/csmd"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/f3das",
		URL:          []string{"https://github.com/mabessa/f3das"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/esp-dnn",
		URL:          []string{"https://github.com/astexuk/esp_dnn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/qs-net",
		URL:          []string{"https://github.com/tmyiri/qs-net"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kpm",
		URL:          []string{"https://github.com/yantaoshen/kpm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hypercubeme",
		URL:          []string{"https://github.com/ivankovlab/hypercubeme"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scmarker",
		URL:          []string{"https://github.com/kchen-lab/scmarker"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/matched-forest",
		URL:          []string{"https://github.com/nooshinsh/matched_forest"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/padimi",
		URL:          []string{"https://github.com/bio-ontology-research-group/padimi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/model-discovery-tool",
		URL:          []string{"https://github.com/dewancse/model-discovery-tool"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/medication-qa-medinfo2019",
		URL:          []string{"https://github.com/abachaa/medication_qa_medinfo2019"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/genomealignmenttools",
		URL:          []string{"https://github.com/hillerlab/genomealignmenttools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/openfmrianalysis",
		URL:          []string{"https://github.com/timvanmourik/openfmrianalysis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lca-cnn",
		URL:          []string{"https://github.com/billzyx/lca-cnn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/seqsero2",
		URL:          []string{"https://github.com/denglab/seqsero2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepmspeptide",
		URL:          []string{"https://github.com/vsegurar/deepmspeptide"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metabmf",
		URL:          []string{"https://github.com/didi10384/metabmf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/xitosbml",
		URL:          []string{"https://github.com/spatialsimulator/xitosbml"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dove",
		URL:          []string{"https://github.com/kiharalab/dovehttp://kiharalab.org/dove"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ags-and-acn-tools",
		URL:          []string{"https://github.com/pereiramemo/ags-and-acn-tools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/workflowr",
		URL:          []string{"https://github.com/jdblischak/workflowr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/patchdca",
		URL:          []string{"https://github.com/biomlboston/patchdca"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scrnaseq-benchmark",
		URL:          []string{"https://github.com/tabdelaal/scrnaseq_benchmark"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/comm",
		URL:          []string{"https://github.com/gordonliu810822/comm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/treeshapestats",
		URL:          []string{"https://github.com/wgs-tb/treeshapestats"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/linearfold",
		URL:          []string{"https://github.com/linearfold/linearfold"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sandres",
		URL:          []string{"https://github.com/azevedolab/sandres"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rapidaci",
		URL:          []string{"https://github.com/manuellamothe/rapidaci"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lem",
		URL:          []string{"https://github.com/ewaszczurek/lem"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hlmethy",
		URL:          []string{"https://github.com/liuze-nwafu/hlmethy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/reqtl",
		URL:          []string{"https://github.com/horvathlab/reqtl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/evntextrc",
		URL:          []string{"https://github.com/cskyan/evntextrc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scgen",
		URL:          []string{"https://github.com/theislab/scgen"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pyar",
		URL:          []string{"https://github.com/anooplab/pyar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/precise",
		URL:          []string{"https://github.com/nki-ccb/precise"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rpowerlib",
		URL:          []string{"https://github.com/samplesizeshop/rpowerlib"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/wipp",
		URL:          []string{"https://github.com/bihealth/wipp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/atlas-rat",
		URL:          []string{"https://github.com/neuropoly/atlas-rat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mass",
		URL:          []string{"https://github.com/gillianlynnryan/mass"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/evolclust",
		URL:          []string{"https://github.com/gabaldonlab/evolclust"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hifresp",
		URL:          []string{"https://github.com/chunquanlipathway/hifresp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vpot",
		URL:          []string{"https://github.com/vccri/vpot"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/orthofinder",
		URL:          []string{"https://github.com/davidemms/orthofinder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/maizete-variation",
		URL:          []string{"https://github.com/snanderson/maizete_variation"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/artifusion",
		URL:          []string{"https://github.com/tron-bioinformatics/artifusion"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hiceekr",
		URL:          []string{"https://github.com/lucidif/hiceekr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/studyprotocolsandbox",
		URL:          []string{"https://github.com/ohdsi/studyprotocolsandbox"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/data-in-brief-influence",
		URL:          []string{"https://github.com/mmrojasgracia/data-in-brief_influence"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tfanalysis",
		URL:          []string{"https://github.com/schulzlab/tfanalysis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/chroma-clade",
		URL:          []string{"https://github.com/chrismonit/chroma_clade"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/massif",
		URL:          []string{"https://github.com/schulzlab/massif"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/swiftortho",
		URL:          []string{"https://github.com/rinoahu/swiftortho"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ppnid",
		URL:          []string{"https://github.com/xueqing4083/ppnid"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/moli",
		URL:          []string{"https://github.com/hosseinshn/moli"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/shamsaraj",
		URL:          []string{"https://github.com/shamsaraj"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pavian",
		URL:          []string{"https://github.com/fbreitwieser/pavian"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scistree",
		URL:          []string{"https://github.com/yufengwudcs/scistree"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/npa",
		URL:          []string{"https://github.com/philipmorrisintl/npa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bwmr",
		URL:          []string{"https://github.com/jiazhao97/bwmr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/psi",
		URL:          []string{"https://github.com/cartoonist/psi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/spades",
		URL:          []string{"https://github.com/ablab/spades"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cnn-brain-strains",
		URL:          []string{"https://github.com/jilab-biomechanics/cnn-brain-strains"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metaerg",
		URL:          []string{"https://github.com/xiaoli-dong/metaerg"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nuclei-segmentation",
		URL:          []string{"https://github.com/easycui/nuclei_segmentation"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lstrap-lite",
		URL:          []string{"https://github.com/mutwil/lstrap-lite"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/knot-pull",
		URL:          []string{"https://github.com/dzarmola/knot_pull"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/neural-circuits-vr",
		URL:          []string{"https://github.com/ivan-vishniakou/neural-circuits-vr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gsso",
		URL:          []string{"https://github.com/superraptor/gsso"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gavin-yinld",
		URL:          []string{"https://github.com/gavin-yinld"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/blastgui",
		URL:          []string{"https://github.com/byemaxx/blastgui"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/epooling",
		URL:          []string{"https://github.com/gao-lab/epooling"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nitumid",
		URL:          []string{"https://github.com/tdw1221/nitumid"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cytokit",
		URL:          []string{"https://github.com/hammerlab/cytokit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sem-cpp",
		URL:          []string{"https://github.com/boyle-lab/sem_cpp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mprascore",
		URL:          []string{"https://github.com/abhisheknrl/mprascore"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nonparametricsummarypsm",
		URL:          []string{"https://github.com/magstra/nonparametricsummarypsm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gpseudoclust",
		URL:          []string{"https://github.com/magstra/gpseudoclust"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tnm",
		URL:          []string{"https://github.com/ugobas/tnm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/epigraph",
		URL:          []string{"https://github.com/complexorganizationoflivingmatter/epigraph"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/seq-ppi",
		URL:          []string{"https://github.com/muhaochen/seq_ppi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/immune-deconvolution-benchmark",
		URL:          []string{"https://github.com/grst/immune_deconvolution_benchmark"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pastas",
		URL:          []string{"https://github.com/pastas/pastas"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tapes",
		URL:          []string{"https://github.com/a-xavier/tapes"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/surfelmeshing",
		URL:          []string{"https://github.com/puzzlepaint/surfelmeshing"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/als-deeplearning",
		URL:          []string{"https://github.com/byin-cwi/als-deeplearning"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/decometdia",
		URL:          []string{"https://github.com/zhumslab/decometdia"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scssim",
		URL:          []string{"https://github.com/qasimyu/scssim"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cloops",
		URL:          []string{"https://github.com/yaqiangcao/cloops"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dynalogo",
		URL:          []string{"https://github.com/lafontaine-uchc/dynalogo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sle",
		URL:          []string{"https://github.com/rabbitpei/sle"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/imspectr",
		URL:          []string{"https://github.com/martijn-cordes/imspectr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/iatc-nrakel",
		URL:          []string{"https://github.com/zhou256/iatc-nrakel"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mountainclimber",
		URL:          []string{"https://github.com/gxiaolab/mountainclimber"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/iso2flux",
		URL:          []string{"https://github.com/cfoguet/iso2flux"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/momi-g",
		URL:          []string{"https://github.com/momi-g/momi-g"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dcnlp",
		URL:          []string{"https://github.com/pjward5656/dcnlp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/elmeri",
		URL:          []string{"https://github.com/leenasalmela/elmeri"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/struo",
		URL:          []string{"https://github.com/leylabmpi/struo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/t-cell-classification",
		URL:          []string{"https://github.com/gitter-lab/t-cell-classification"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gtdbtk",
		URL:          []string{"https://github.com/ecogenomics/gtdbtk"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dingo",
		URL:          []string{"https://github.com/radoslav180/dingo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dmdeepm6a1.0",
		URL:          []string{"https://github.com/nwpu-903pr/dmdeepm6a1.0"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepppisp",
		URL:          []string{"https://github.com/csubiogroup/deepppisp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/swapcounter",
		URL:          []string{"https://github.com/mengjintao/swapcounter"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/la-isla-de-tomato",
		URL:          []string{"https://github.com/mmjulkowska/la_isla_de_tomato"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/linkhd",
		URL:          []string{"https://github.com/lauzingaretti/linkhd"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mpies",
		URL:          []string{"https://github.com/johanneswerner/mpies"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/addo",
		URL:          []string{"https://github.com/leileicui/addo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/drdiff",
		URL:          []string{"https://github.com/ronaldolab/drdiff"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/drugz",
		URL:          []string{"https://github.com/hart-lab/drugz"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/435271",
		URL:          []string{"https://github.com/lotteanna/defence_adaptation,https://doi.org/10.1101/435271"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mtbnn",
		URL:          []string{"https://github.com/zoesgithub/mtbnn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/xmanv2",
		URL:          []string{"https://github.com/lazarlab/xmanv2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cobamp",
		URL:          []string{"https://github.com/biosystemsum/cobamp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/moomin",
		URL:          []string{"https://github.com/htpusa/moomin"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ia-lab",
		URL:          []string{"https://github.com/amcorrigan/ia-lab"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metgem",
		URL:          []string{"https://github.com/metgem/metgem"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/u-pass",
		URL:          []string{"https://github.com/pill-gz/u-pass"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scgeatoolbox",
		URL:          []string{"https://github.com/jamesjcai/scgeatoolbox"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/magi-s",
		URL:          []string{"https://github.com/jchow32/magi-s"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/suicidalaltruismdissertation",
		URL:          []string{"https://github.com/anyaevostinar/suicidalaltruismdissertation"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/netml",
		URL:          []string{"https://github.com/compbiolabucf/netml"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/timeseriesnem",
		URL:          []string{"https://github.com/cbg-ethz/timeseriesnem"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/iseeu",
		URL:          []string{"https://github.com/williamcaicedo/iseeu"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dr-thermo",
		URL:          []string{"https://github.com/samansalike/dr-thermo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/contaminationx",
		URL:          []string{"https://github.com/sapfo/contaminationx"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scifil",
		URL:          []string{"https://github.com/compbel/scifil"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/quartetscores",
		URL:          []string{"https://github.com/lutteropp/quartetscores"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/extrarg",
		URL:          []string{"https://github.com/gaarangoa/extrarg"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/semanticprogramming",
		URL:          []string{"https://github.com/semanticprogramming"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/disseqt.jl",
		URL:          []string{"https://github.com/rasmushenningsson/disseqt.jl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/l1em",
		URL:          []string{"https://github.com/fenyolab/l1em"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pydca",
		URL:          []string{"https://github.com/kit-mbs/pydca"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bioinformatics-sourcecode",
		URL:          []string{"https://github.com/lauramoraes/bioinformatics-sourcecode"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pyhsiclasso",
		URL:          []string{"https://github.com/riken-aip/pyhsiclasso"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dp-representation-transfer",
		URL:          []string{"https://github.com/dpbayes/dp-representation-transfer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/strassembly",
		URL:          []string{"https://github.com/gulfemd/strassembly"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/salsa",
		URL:          []string{"https://github.com/machinegun/salsa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/echo",
		URL:          []string{"https://github.com/delosh653/echo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scrnabatchqc",
		URL:          []string{"https://github.com/liuqivandy/scrnabatchqc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/few-shot-segmentation",
		URL:          []string{"https://github.com/abhi4ssj/few-shot-segmentation"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hicgan",
		URL:          []string{"https://github.com/kimmo1019/hicgan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fertility-gru",
		URL:          []string{"https://github.com/khanhlee/fertility-gru"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ssasu",
		URL:          []string{"https://github.com/brossetti/ssasu"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/baerhunter",
		URL:          []string{"https://github.com/irilenia/baerhunter"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/robokop",
		URL:          []string{"https://github.com/ncats-gamma/robokop"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/szczepankiewicz-dib-2019",
		URL:          []string{"https://github.com/filip-szczepankiewicz/szczepankiewicz_dib_2019"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bagse",
		URL:          []string{"https://github.com/xqwen/bagse"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kalign",
		URL:          []string{"https://github.com/timolassmann/kalign"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/jmztab-m",
		URL:          []string{"https://github.com/lifs-tools/jmztab-m"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lifs-tools",
		URL:          []string{"https://github.com/lifs-tools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/genome-sketching",
		URL:          []string{"https://github.com/will-rowe/genome-sketching"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ipa",
		URL:          []string{"https://github.com/francescodc87/ipa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sprites2",
		URL:          []string{"https://github.com/zhangzhen/sprites2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fumoso",
		URL:          []string{"https://github.com/aresio/fumoso"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bcrystal",
		URL:          []string{"https://github.com/raghvendra5688/bcrystal"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/multiguidescan",
		URL:          []string{"https://github.com/bioinfomaticscsu/multiguidescan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/switchablenorms",
		URL:          []string{"https://github.com/switchablenorms"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mirgff3",
		URL:          []string{"https://github.com/mirtop/mirgff3"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mirtop",
		URL:          []string{"https://github.com/mirtop/mirtop"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/growthestimate",
		URL:          []string{"https://github.com/rungruangsaktorrissenmanoonpong/growthestimate"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepimpute",
		URL:          []string{"https://github.com/lanagarmire/deepimpute"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gcmdr",
		URL:          []string{"https://github.com/yahuang1991polyu/gcmdr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/npyc-toolbox",
		URL:          []string{"https://github.com/phenomecentre/npyc-toolbox"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/npyc-toolbox-tutorials",
		URL:          []string{"https://github.com/phenomecentre/npyc-toolbox-tutorials"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dropclust",
		URL:          []string{"https://github.com/debsin/dropclust"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/identificationwitharme",
		URL:          []string{"https://github.com/almarguet/identificationwitharme"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/antibody-2019",
		URL:          []string{"https://github.com/gifford-lab/antibody-2019"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/virbin",
		URL:          []string{"https://github.com/chjiao/virbin"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mspac",
		URL:          []string{"https://github.com/oscarlr/mspac"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tangle",
		URL:          []string{"https://github.com/samuelefiorini/tangle"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bbknn",
		URL:          []string{"https://github.com/teichlab/bbknn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tersect",
		URL:          []string{"https://github.com/tomkurowski/tersect"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/treemerge",
		URL:          []string{"https://github.com/ekmolloy/treemerge"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/multimodalprognosis",
		URL:          []string{"https://github.com/gevaertlab/multimodalprognosis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/selfish",
		URL:          []string{"https://github.com/ucrbioinfo/selfish"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/npd-micro-saccade-detection",
		URL:          []string{"https://github.com/hz-zhu/npd-micro-saccade-detection"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/wenda",
		URL:          []string{"https://github.com/pfeiferlabtue/wenda"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tfregulomer",
		URL:          []string{"https://github.com/benoukraflab/tfregulomer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vic",
		URL:          []string{"https://github.com/hglab/vic"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nextpolish",
		URL:          []string{"https://github.com/nextomics/nextpolish"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/snappy",
		URL:          []string{"https://github.com/pmmaraujo/snappy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/psiz",
		URL:          []string{"https://github.com/roads/psiz"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rsnapsim",
		URL:          []string{"https://github.com/munskygroup/rsnapsim"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/flashweave.jl",
		URL:          []string{"https://github.com/meringlab/flashweave.jl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cfnet",
		URL:          []string{"https://github.com/bchidest/cfnet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lionessr",
		URL:          []string{"https://github.com/kuijjerlab/lionessr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/t-lex3",
		URL:          []string{"https://github.com/gonzalezlab/t-lex3"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dna-rchitect",
		URL:          []string{"https://github.com/alosdiallo/dna_rchitect"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/actinn",
		URL:          []string{"https://github.com/mafeiyang/actinn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cancer-subtyping",
		URL:          []string{"https://github.com/tjgu/cancer_subtyping"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gfakluge",
		URL:          []string{"https://github.com/edawson/gfakluge"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/splicesites-ann-cgr",
		URL:          []string{"https://github.com/thoang3/portfolio/tree/splicesites_ann_cgr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/splicelauncher",
		URL:          []string{"https://github.com/raphaelleman/splicelauncher"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fldgen",
		URL:          []string{"https://github.com/jgcri/fldgen"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bvp-pred-unb",
		URL:          []string{"https://github.com/muhammad-arif-nust/bvp_pred_unb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/praline",
		URL:          []string{"https://github.com/ibivu/praline"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/uslcount",
		URL:          []string{"https://github.com/mikpom/uslcount"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ragoo",
		URL:          []string{"https://github.com/malonge/ragoo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/designmetabolicrnalabeling",
		URL:          []string{"https://github.com/dieterich-lab/designmetabolicrnalabeling"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mexcowalk",
		URL:          []string{"https://github.com/abu-compbio/mexcowalk"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bioner-cross-sharing",
		URL:          []string{"https://github.com/joglelew/bioner-cross-sharing"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/overlap",
		URL:          []string{"https://github.com/rachelnethery/overlap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lemon-tree",
		URL:          []string{"https://github.com/eb00/lemon-tree"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/predictability-of-cancer-evolution",
		URL:          []string{"https://github.com/cbg-ethz/predictability_of_cancer_evolution"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sctree-test",
		URL:          []string{"https://github.com/xqbai/sctree-test"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/promodule",
		URL:          []string{"https://github.com/chupan1218/promodule"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gokit",
		URL:          []string{"https://github.com/gokit1/gokit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cellsim",
		URL:          []string{"https://github.com/lileijie1992/cellsim"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/clickx",
		URL:          []string{"https://github.com/liulab-csrc/clickx"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/exmachina",
		URL:          []string{"https://github.com/shuichiro-makigaki/exmachina"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/maps",
		URL:          []string{"https://github.com/ijuric/maps"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dspred",
		URL:          []string{"https://github.com/yusufzaferaydin/dspred"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lsm-worker",
		URL:          []string{"https://github.com/jmanj/lsm_worker"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rescue",
		URL:          []string{"https://github.com/seasamgo/rescue"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sscclust",
		URL:          []string{"https://github.com/japrin/sscclust"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/opal",
		URL:          []string{"https://github.com/cami-challenge/opal"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepmetapsicov",
		URL:          []string{"https://github.com/psipred/deepmetapsicov"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/genhap",
		URL:          []string{"https://github.com/andrea-tango/genhap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ann-glycolysis-flux-prediction",
		URL:          []string{"https://github.com/dsimb/ann-glycolysis-flux-prediction"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/groupk",
		URL:          []string{"https://github.com/strideradu/groupk"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/r4kappa",
		URL:          []string{"https://github.com/lptolik/r4kappa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biomedical-qa",
		URL:          []string{"https://github.com/jinzanxia/biomedical-qa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/aneuvis",
		URL:          []string{"https://github.com/dpique/aneuvis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/doc2hpo",
		URL:          []string{"https://github.com/stormliucong/doc2hpo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cam",
		URL:          []string{"https://github.com/ridgelab/cam"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/salmon",
		URL:          []string{"https://github.com/huangzhii/salmon"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mctcodes",
		URL:          []string{"https://github.com/iullah1980/mctcodes"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/samovar",
		URL:          []string{"https://github.com/cdarby/samovar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepmito",
		URL:          []string{"https://github.com/bolognabiocomp/deepmito"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/feast",
		URL:          []string{"https://github.com/cozygene/feast"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/puffin",
		URL:          []string{"https://github.com/gifford-lab/puffin"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/arcashla",
		URL:          []string{"https://github.com/rabadanlab/arcashla"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/x2h",
		URL:          []string{"https://github.com/lyotvincent/x2h"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mobiusklein",
		URL:          []string{"https://github.com/mobiusklein"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pdeep",
		URL:          []string{"https://github.com/pfindstudio/pdeep"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepdom",
		URL:          []string{"https://github.com/yuexujiang/deepdom"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vankrevelenlocal",
		URL:          []string{"https://github.com/hegemanlab/vankrevelenlocal"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/summaryauc",
		URL:          []string{"https://github.com/lsncibb/summaryauc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tad-fusion-score",
		URL:          []string{"https://github.com/hormozdiarilab/tad-fusion-score"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lemon",
		URL:          []string{"https://github.com/chopralab/lemon"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cd-trace",
		URL:          []string{"https://github.com/coamo2/cd-trace"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/senet",
		URL:          []string{"https://github.com/hujie-frank/senet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pacl",
		URL:          []string{"https://github.com/tmallava/pacl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/path2surv",
		URL:          []string{"https://github.com/mehmetgonen/path2surv"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sst",
		URL:          []string{"https://github.com/shijies/sst"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/manga",
		URL:          []string{"https://github.com/greco-lab/manga"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/home",
		URL:          []string{"https://github.com/listerlab/home"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cuddi",
		URL:          []string{"https://github.com/chengusf/cuddi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepexpression",
		URL:          []string{"https://github.com/wanwenzeng/deepexpression"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nanopype",
		URL:          []string{"https://github.com/giesselmann/nanopype"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bactsnp",
		URL:          []string{"https://github.com/iekadn/bactsnp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/multiphate",
		URL:          []string{"https://github.com/carolzhou/multiphate"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kmcex",
		URL:          []string{"https://github.com/lzhlab/kmcex"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ancestry-viz",
		URL:          []string{"https://github.com/hagax8/ancestry_viz"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/emwer",
		URL:          []string{"https://github.com/kojikoji/emwer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/funmap2",
		URL:          []string{"https://github.com/wzhy2000/funmap2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ipsa",
		URL:          []string{"https://github.com/coongroup/ipsa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/stillbirth",
		URL:          []string{"https://github.com/migariane/stillbirth"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/myelinj",
		URL:          []string{"https://github.com/barnettlab/myelinj"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pixel",
		URL:          []string{"https://github.com/candihub/pixel"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/wdstar",
		URL:          []string{"https://github.com/alekseyenko/wdstar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tsd",
		URL:          []string{"https://github.com/menggf/tsd"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/qparse",
		URL:          []string{"https://github.com/b3rse/qparse"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/alarm",
		URL:          []string{"https://github.com/masilab/alarm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gappadder",
		URL:          []string{"https://github.com/reedwarbler/gappadder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sema",
		URL:          []string{"https://github.com/msolmazm/sema"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ohvarfinder",
		URL:          []string{"https://github.com/takumorizo/ohvarfinder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/toast",
		URL:          []string{"https://github.com/ziyili20/toast"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/distruct",
		URL:          []string{"https://github.com/kit-mbs/distruct"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/patientexplorer",
		URL:          []string{"https://github.com/benglicksberg/patientexplorer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deep-learning-examples",
		URL:          []string{"https://github.com/lykaust15/deep_learning_examples"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/teachopencadd",
		URL:          []string{"https://github.com/volkamerlab/teachopencadd"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rna-secondary-structure-database",
		URL:          []string{"https://github.com/linyuwangphd/rna-secondary-structure-database"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tcc-gui",
		URL:          []string{"https://github.com/swsoyee/tcc-gui"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fluka-ct",
		URL:          []string{"https://github.com/chezhia/fluka_ct"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ancis-pytorch",
		URL:          []string{"https://github.com/yijingru/ancis-pytorch"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sasm-vgwas",
		URL:          []string{"https://github.com/meiyan88/sasm-vgwas"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rapid",
		URL:          []string{"https://github.com/schulzlab/rapid"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cooler",
		URL:          []string{"https://github.com/mirnylab/cooler"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/wcmc-research-informatics",
		URL:          []string{"https://github.com/wcmc-research-informatics"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/epee",
		URL:          []string{"https://github.com/cobanoglu-lab/epee"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sleepeegnet",
		URL:          []string{"https://github.com/sajadmo/sleepeegnet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mctl",
		URL:          []string{"https://github.com/wangshanshancqu/mctl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kneeanalysis",
		URL:          []string{"https://github.com/pingjunchen/kneeanalysis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tawny-karyotype",
		URL:          []string{"https://github.com/jaydchan/tawny-karyotype"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sss-test",
		URL:          []string{"https://github.com/waltercostamb/sss-test"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ntedit",
		URL:          []string{"https://github.com/bcgsc/ntedit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ebolavirussdpsbioinformatics2019",
		URL:          []string{"https://github.com/wasslab/ebolavirussdpsbioinformatics2019"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sbl",
		URL:          []string{"https://github.com/meiyuecomputbio/sbl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deephistone",
		URL:          []string{"https://github.com/qijinyin/deephistone"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dmsc",
		URL:          []string{"https://github.com/nwpu-903pr/dmsc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pocketpipe",
		URL:          []string{"https://github.com/inpacdb/pocketpipe"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/divvier",
		URL:          []string{"https://github.com/simonwhelan/divvier"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deconvseq",
		URL:          []string{"https://github.com/rosedu1/deconvseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/datasets",
		URL:          []string{"https://github.com/shengli0201/datasets"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fssemr",
		URL:          []string{"https://github.com/ivis4ml/fssemr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/2stepqa",
		URL:          []string{"https://github.com/xiangxuyu/2stepqa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/easycodeml",
		URL:          []string{"https://github.com/bioeasy/easycodeml"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepsignal",
		URL:          []string{"https://github.com/bioinfomaticscsu/deepsignal"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/qanalysis-project",
		URL:          []string{"https://github.com/nlp-bigdatalab/qanalysis-project"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/csi-utr",
		URL:          []string{"https://github.com/uoflbioinformatics/csi-utr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gift",
		URL:          []string{"https://github.com/leesael/gift"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hypergraphdynamiccorrelation",
		URL:          []string{"https://github.com/yunchuankong/hypergraphdynamiccorrelation"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rescore",
		URL:          []string{"https://github.com/compomics/rescore"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pats",
		URL:          []string{"https://github.com/blaboaklandu/pats"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cocoscore",
		URL:          []string{"https://github.com/jungealexander/cocoscore"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vomm",
		URL:          []string{"https://github.com/jnalanko/vomm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/flas",
		URL:          []string{"https://github.com/baoe/flas"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pcmf",
		URL:          []string{"https://github.com/gdurif/pcmf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mctandem",
		URL:          []string{"https://github.com/logiczy/mctandem"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/d2r-codes",
		URL:          []string{"https://github.com/xuegonglab/d2r_codes"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rnaredprint",
		URL:          []string{"https://github.com/yannponty/rnaredprint"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/snakepipes",
		URL:          []string{"https://github.com/maxplanck-ie/snakepipes"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nanor",
		URL:          []string{"https://github.com/davidebolo1993/nanor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/resimnet",
		URL:          []string{"https://github.com/dmis-lab/resimnet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/artmap",
		URL:          []string{"https://github.com/rihalab/artmap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/duphold",
		URL:          []string{"https://github.com/brentp/duphold"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/multiclass-segmentation",
		URL:          []string{"https://github.com/neuropoly/multiclass-segmentation"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/yeast-segmentation",
		URL:          []string{"https://github.com/alexxijielu/yeast_segmentation"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/image-engine",
		URL:          []string{"https://github.com/vusionmed/image-engine"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/grnte",
		URL:          []string{"https://github.com/jccastrog/grnte"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lipase-reclassification",
		URL:          []string{"https://github.com/thh32/lipase_reclassification"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metaqubic",
		URL:          []string{"https://github.com/osu-bmbl/metaqubic"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pwrewas",
		URL:          []string{"https://github.com/stefangraw/pwrewas"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/locnuclei",
		URL:          []string{"https://github.com/rostlab/locnuclei"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phyml",
		URL:          []string{"https://github.com/stephaneguindon/phyml"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ivis",
		URL:          []string{"https://github.com/beringresearch/ivis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pi2",
		URL:          []string{"https://github.com/arttumiettinen/pi2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/grein",
		URL:          []string{"https://github.com/uc-bd2k/grein"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scrublet",
		URL:          []string{"https://github.com/allonkleinlab/scrublet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/crossbar-net",
		URL:          []string{"https://github.com/qianyu1226/crossbar-net"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/keanu",
		URL:          []string{"https://github.com/igbb/keanu"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/g2p",
		URL:          []string{"https://github.com/xiaoleiliubio/g2p"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepslide",
		URL:          []string{"https://github.com/bmirds/deepslide"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/caslocusanno",
		URL:          []string{"https://github.com/riversdong/caslocusanno"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/varscot",
		URL:          []string{"https://github.com/bauerlab/varscot"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/aivar",
		URL:          []string{"https://github.com/topgene/aivar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ivus-ultrasonic",
		URL:          []string{"https://github.com/kulbear/ivus-ultrasonic"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rea",
		URL:          []string{"https://github.com/4dsoftware/rea"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/heritability",
		URL:          []string{"https://github.com/paulschmidtgit/heritability"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/acp-dl",
		URL:          []string{"https://github.com/haichengyi/acp-dl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepcircode",
		URL:          []string{"https://github.com/biodatalearning/deepcircode"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/solidbin",
		URL:          []string{"https://github.com/sufforest/solidbin"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/stytra",
		URL:          []string{"https://github.com/portugueslab/stytra"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cregulome",
		URL:          []string{"https://github.com/ropensci/cregulome"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/colocr",
		URL:          []string{"https://github.com/ropensci/colocr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/faot",
		URL:          []string{"https://github.com/caolman/faot"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/f-anogan",
		URL:          []string{"https://github.com/tschlegl/f-anogan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tractaviewer",
		URL:          []string{"https://github.com/neilpearson-lilly/tractaviewer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ngsrelate",
		URL:          []string{"https://github.com/angsd/ngsrelate"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/galaxy",
		URL:          []string{"https://github.com/galaxyproject/galaxy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cloudauthz",
		URL:          []string{"https://github.com/galaxyproject/cloudauthz"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ppa-assembler",
		URL:          []string{"https://github.com/yaobaiwei/ppa-assembler"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scmut",
		URL:          []string{"https://github.com/nghiavtr/scmut"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cytofmerge",
		URL:          []string{"https://github.com/tabdelaal/cytofmerge"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bioshake",
		URL:          []string{"https://github.com/papenfusslab/bioshake"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mrlr",
		URL:          []string{"https://github.com/chonglab/mrlr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/chipulate",
		URL:          []string{"https://github.com/vishakad/chipulate"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/absnf",
		URL:          []string{"https://github.com/pfruan/absnf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/joa",
		URL:          []string{"https://github.com/burcakotlu/joa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/plier",
		URL:          []string{"https://github.com/wgmao/plier"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gtshark",
		URL:          []string{"https://github.com/refresh-bio/gtshark"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tegs",
		URL:          []string{"https://github.com/wzhangwhu/tegs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/do-ms",
		URL:          []string{"https://github.com/slavovlab/do-ms"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/prosampler",
		URL:          []string{"https://github.com/zhengchangsulab/prosampler"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/indra-pathway-map",
		URL:          []string{"https://github.com/sorgerlab/indra_pathway_map"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/browniecorrector",
		URL:          []string{"https://github.com/biointec/browniecorrector"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/umib",
		URL:          []string{"https://github.com/gutmicrobes/umib"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/orna",
		URL:          []string{"https://github.com/schulzlab/orna"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepfe-ppi",
		URL:          []string{"https://github.com/xal2019/deepfe-ppi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepconv-dti",
		URL:          []string{"https://github.com/gist-csbl/deepconv-dti"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/chicdiff",
		URL:          []string{"https://github.com/regulatorygenomicsgroup/chicdiff"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/expansionhunter",
		URL:          []string{"https://github.com/illumina/expansionhunter"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tibanna",
		URL:          []string{"https://github.com/4dn-dcic/tibanna"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hdx-viewer",
		URL:          []string{"https://github.com/david-bouyssie/hdx-viewer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mycorrhiza",
		URL:          []string{"https://github.com/jgeofil/mycorrhiza"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/splicev",
		URL:          []string{"https://github.com/flemingtonlab/splicev"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/enigma",
		URL:          []string{"https://github.com/abikoushi/enigma"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hierarchical3dgenome",
		URL:          []string{"https://github.com/bdm-lab/hierarchical3dgenome"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metaassemblyeval",
		URL:          []string{"https://github.com/ziyewang/metaassemblyeval"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dairydb",
		URL:          []string{"https://github.com/marcomeola/dairydb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vcf-plotein",
		URL:          []string{"https://github.com/raulossio/vcf-plotein"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tpot-fss",
		URL:          []string{"https://github.com/lelaboratoire/tpot-fss"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tpot",
		URL:          []string{"https://github.com/epistasislab/tpot"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pygenprop",
		URL:          []string{"https://github.com/micromeda/pygenprop"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/saber",
		URL:          []string{"https://github.com/baderlab/saber"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/towards-reliable-bioner",
		URL:          []string{"https://github.com/baderlab/towards-reliable-bioner"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/two-layer-qrcode",
		URL:          []string{"https://github.com/yuantailing/two-layer-qrcode"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/terapca",
		URL:          []string{"https://github.com/aritra90/terapca"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tormes",
		URL:          []string{"https://github.com/nmquijada/tormes"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/etcomp",
		URL:          []string{"https://github.com/behinger/etcomp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pdx-analysis-workflows",
		URL:          []string{"https://github.com/thejacksonlaboratory/pdx-analysis-workflows"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/darts",
		URL:          []string{"https://github.com/xinglab/darts"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pysradb",
		URL:          []string{"https://github.com/saketkc/pysradb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/amda",
		URL:          []string{"https://github.com/xyz5074/amda"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/r.sambada",
		URL:          []string{"https://github.com/solanged/r.sambada"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rmut",
		URL:          []string{"https://github.com/csclab/rmut"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/csvd",
		URL:          []string{"https://github.com/vguillemot/csvd"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phylopi",
		URL:          []string{"https://github.com/armandbester/phylopi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/linker",
		URL:          []string{"https://github.com/mikelhernaez/linker"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/jama16-retina-replication",
		URL:          []string{"https://github.com/mikevoets/jama16-retina-replication"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gene-surrounder",
		URL:          []string{"https://github.com/sahildshah1/gene-surrounder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metalab",
		URL:          []string{"https://github.com/nmikolajewicz/metalab"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/subnetwork-inference",
		URL:          []string{"https://github.com/craven-biostat-lab/subnetwork_inference"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/subdyquency",
		URL:          []string{"https://github.com/weiba/subdyquency"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bts-dsn",
		URL:          []string{"https://github.com/guomugong/bts-dsn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/uspikehunter",
		URL:          []string{"https://github.com/uspikehunter"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tardis",
		URL:          []string{"https://github.com/bilkentcompgen/tardis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mplnclust",
		URL:          []string{"https://github.com/anjalisilva/mplnclust"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pipred",
		URL:          []string{"https://github.com/labstructbioinf/pipred"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/defor",
		URL:          []string{"https://github.com/drzh/defor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/juchmme",
		URL:          []string{"https://github.com/pbagos/juchmme"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/smash",
		URL:          []string{"https://github.com/sun-lab/smash"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pedigreenet",
		URL:          []string{"https://github.com/maize-genetics-and-genomics-database/pedigreenet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ppr-meta",
		URL:          []string{"https://github.com/zhenchengfang/ppr-meta"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hicap",
		URL:          []string{"https://github.com/scwatts/hicap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/flowgrid",
		URL:          []string{"https://github.com/vccri/flowgrid"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/aphid",
		URL:          []string{"https://github.com/shaunpwilkinson/aphid"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rbapy",
		URL:          []string{"https://github.com/sysbioinra/rbapy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gaynor-sun-gbj-breast-cancer",
		URL:          []string{"https://github.com/ryanrsun/gaynor_sun_gbj_breast_cancer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phyclip",
		URL:          []string{"https://github.com/alvinxhan/phyclip"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hppi-tensorflow",
		URL:          []string{"https://github.com/smalltalkman/hppi-tensorflow"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/multimodal-medical-image-synthesis",
		URL:          []string{"https://github.com/hustlinyi/multimodal-medical-image-synthesis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rerconverge",
		URL:          []string{"https://github.com/nclark-lab/rerconverge"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phanotate",
		URL:          []string{"https://github.com/deprekate/phanotate"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mhn",
		URL:          []string{"https://github.com/rudischill/mhn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rodin",
		URL:          []string{"https://github.com/pnnl-comp-mass-spec/rodin"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lipidminion",
		URL:          []string{"https://github.com/pnnl-comp-mass-spec/lipidminion"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nitpicker",
		URL:          []string{"https://github.com/ezer/nitpicker"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/icn3d",
		URL:          []string{"https://github.com/ncbi/icn3d"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scmatch",
		URL:          []string{"https://github.com/forrest-lab/scmatch"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/augmentor",
		URL:          []string{"https://github.com/mdbloice/augmentor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/augmentor.jl",
		URL:          []string{"https://github.com/evizero/augmentor.jl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/raamlab",
		URL:          []string{"https://github.com/bioinfo0706/raamlab"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/micop",
		URL:          []string{"https://github.com/smangul1/micop"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ggm-shrinkage",
		URL:          []string{"https://github.com/v-bernal/ggm-shrinkage"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/heteroplasmyworkflow",
		URL:          []string{"https://github.com/vtphan/heteroplasmyworkflow"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/functree-ng",
		URL:          []string{"https://github.com/yamada-lab/functree-ng"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scbs-map",
		URL:          []string{"https://github.com/wupengomics/scbs-map"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pyfeat",
		URL:          []string{"https://github.com/mrzresearcharena/pyfeat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pyfeat",
		URL:          []string{"https://github.com/mrzresearcharena/pyfeat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/confined",
		URL:          []string{"https://github.com/cozygene/confined"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/prankwebapp",
		URL:          []string{"https://github.com/jendelel/prankwebapp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/p2rank",
		URL:          []string{"https://github.com/rdk/p2rank"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/3dchromatin-replicateqc",
		URL:          []string{"https://github.com/kundajelab/3dchromatin_replicateqc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metapheno",
		URL:          []string{"https://github.com/nlapier2/metapheno"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kekulescope",
		URL:          []string{"https://github.com/isidroc/kekulescope"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/predpsi-svr",
		URL:          []string{"https://github.com/chenkenbio/predpsi-svr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pinetree",
		URL:          []string{"https://github.com/benjaminjack/pinetree"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/adjutant",
		URL:          []string{"https://github.com/amcrisan/adjutant"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fromage",
		URL:          []string{"https://github.com/crespo-otero-group/fromage"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepred",
		URL:          []string{"https://github.com/cansyl/deepred"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/snp2sim",
		URL:          []string{"https://github.com/mccoymd/snp2sim"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/yalla",
		URL:          []string{"https://github.com/germannp/yalla"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phnmnl",
		URL:          []string{"https://github.com/phnmnl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cgrtools",
		URL:          []string{"https://github.com/cimm-kzn/cgrtools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/networkinference",
		URL:          []string{"https://github.com/bagherilab/networkinference"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lncrna-disease-link",
		URL:          []string{"https://github.com/pengeace/lncrna-disease-link"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ddsae",
		URL:          []string{"https://github.com/zhang-informatics/ddsae"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mitoz",
		URL:          []string{"https://github.com/linzhi2013/mitoz"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gpa-package",
		URL:          []string{"https://github.com/iid-dth/gpa-package"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/relz-fs",
		URL:          []string{"https://github.com/vsepulve/relz_fs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/polya-prediction-lrm-dnn",
		URL:          []string{"https://github.com/emang-kaust/polya_prediction_lrm_dnn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/simphony",
		URL:          []string{"https://github.com/hugheylab/simphony"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/decent",
		URL:          []string{"https://github.com/cz-ye/decent"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gma",
		URL:          []string{"https://github.com/chaoning/gma"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/circdeep",
		URL:          []string{"https://github.com/uoflbioinformatics/circdeep"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tumorcnv",
		URL:          []string{"https://github.com/yongzhuang/tumorcnv"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rupee",
		URL:          []string{"https://github.com/rayoub/rupee"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pyrod",
		URL:          []string{"https://github.com/schallerdavid/pyrod"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bel-enrichment",
		URL:          []string{"https://github.com/bel-enrichment/bel-enrichment"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nurse",
		URL:          []string{"https://github.com/jianxigao/nurse"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ngo2019-authortopic",
		URL:          []string{"https://github.com/thomasyeolab/"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ar-pred-source",
		URL:          []string{"https://github.com/sambitmishra0628/ar-pred_source"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mlsf",
		URL:          []string{"https://github.com/hongjianli/mlsf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/simkern",
		URL:          []string{"https://github.com/davidcraft/simkern"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cliquems",
		URL:          []string{"https://github.com/osenan/cliquems"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sieve-ub",
		URL:          []string{"https://github.com/biodataganache/sieve-ub"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ssdfa",
		URL:          []string{"https://github.com/bcrafton/ssdfa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/proteoformer",
		URL:          []string{"https://github.com/biobix/proteoformer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ploidyinfer",
		URL:          []string{"https://github.com/huangkang1987/ploidyinfer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dnea",
		URL:          []string{"https://github.com/wiggie/dnea"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/slantbrainseg",
		URL:          []string{"https://github.com/masilab/slantbrainseg"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/haploconduct",
		URL:          []string{"https://github.com/haploconduct/haploconduct"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tankche1",
		URL:          []string{"https://github.com/tankche1"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sniper",
		URL:          []string{"https://github.com/suzhixi/sniper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mhc-typer",
		URL:          []string{"https://github.com/huangkang1987/mhc-typer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/downloads",
		URL:          []string{"https://github.com/tbl-uiuc/downloads"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/eecog-comp",
		URL:          []string{"https://github.com/vincent-wq/eecog-comp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/efindsite",
		URL:          []string{"https://github.com/michal-brylinski/efindsite"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/betaturn18",
		URL:          []string{"https://github.com/sh-maxim/betaturn18"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/spliceogen",
		URL:          []string{"https://github.com/vccri/spliceogen"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/electron-micrograph-denoiser",
		URL:          []string{"https://github.com/jeffrey-ede/electron-micrograph-denoiser"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/crisprdisco",
		URL:          []string{"https://github.com/crisprlab/crisprdisco"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nlpar",
		URL:          []string{"https://github.com/usnavalresearchlaboratory/nlpar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/texomer",
		URL:          []string{"https://github.com/kchen-lab/texomer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gladiatox",
		URL:          []string{"https://github.com/philipmorrisintl/gladiatox"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mgrfe-garfe",
		URL:          []string{"https://github.com/pengeace/mgrfe-garfe"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/apero",
		URL:          []string{"https://github.com/simon-leonard/apero"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bacterial-colonization-model",
		URL:          []string{"https://github.com/mjarvenpaa/bacterial-colonization-model"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ecogems",
		URL:          []string{"https://github.com/venyao/ecogems"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gtotree",
		URL:          []string{"https://github.com/astrobiomike/gtotree"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/autoimmune-research",
		URL:          []string{"https://github.com/jieunjung511/autoimmune-research"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/av-segmentation",
		URL:          []string{"https://github.com/rubenhx/av-segmentation"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mirtime",
		URL:          []string{"https://github.com/mirtime/mirtime"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mepurity",
		URL:          []string{"https://github.com/xjtu-omics/mepurity"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/psi-sigma",
		URL:          []string{"https://github.com/wososa/psi-sigma"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/st-viewer",
		URL:          []string{"https://github.com/spatialtranscriptomicsresearch/st_viewer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/chfs",
		URL:          []string{"https://github.com/kazukisakura/chfs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/neopredpipe",
		URL:          []string{"https://github.com/mathonco/neopredpipe"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/squat",
		URL:          []string{"https://github.com/luke831215/squat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hulk",
		URL:          []string{"https://github.com/will-rowe/hulk"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metachip",
		URL:          []string{"https://github.com/songweizhi/metachip"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vpac",
		URL:          []string{"https://github.com/shengquanchen/vpac"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/goeckslab",
		URL:          []string{"https://github.com/goeckslab"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deep-collaborative-filtering",
		URL:          []string{"https://github.com/xzenglab/deep-collaborative-filtering"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hmmpe",
		URL:          []string{"https://github.com/emdadi/hmmpe"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vimco",
		URL:          []string{"https://github.com/xingjieshi/vimco"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fastbaps",
		URL:          []string{"https://github.com/gtonkinhill/fastbaps"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cnspector",
		URL:          []string{"https://github.com/papenfusslab/cnspector"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hla-la",
		URL:          []string{"https://github.com/diltheylab/hla-la"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bamixchecker",
		URL:          []string{"https://github.com/heinc1010/bamixchecker"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/georacle",
		URL:          []string{"https://github.com/vccri/georacle"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/geneticscreen.jl",
		URL:          []string{"https://github.com/janewliang/geneticscreen.jl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cram-js",
		URL:          []string{"https://github.com/gmod/cram-js"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cistopic",
		URL:          []string{"https://github.com/aertslab/cistopic"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lincs-rnaseq-cpp",
		URL:          []string{"https://github.com/biodepot/lincs_rnaseq_cpp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pga",
		URL:          []string{"https://github.com/quxiaojian/pga"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/c-intersecture",
		URL:          []string{"https://github.com/nuriddinovma/c-intersecture"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/studyprotocolsandbox",
		URL:          []string{"https://github.com/ohdsi/studyprotocolsandbox"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/blisar",
		URL:          []string{"https://github.com/soufianeajana/blisar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/self-organization",
		URL:          []string{"https://github.com/hfooladi/self_organization"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/look4trs",
		URL:          []string{"https://github.com/tulsabioinformaticstoolsmith/look4trs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/suretypesc",
		URL:          []string{"https://github.com/puko818/suretypesc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/msf",
		URL:          []string{"https://github.com/modulated-subgraph-finder/msf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bfmem",
		URL:          []string{"https://github.com/yuansliu/bfmem"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mp",
		URL:          []string{"https://github.com/smirarab/astral/tree/mp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/shiny-seq",
		URL:          []string{"https://github.com/schultzelab/shiny-seq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/galaxy-rna-workbench",
		URL:          []string{"https://github.com/bgruening/galaxy-rna-workbench"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/spar-k",
		URL:          []string{"https://github.com/romaingroux/spar-k"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/protein-solubility",
		URL:          []string{"https://github.com/xiaomizhou616/protein_solubility"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/funpred-3.0",
		URL:          []string{"https://github.com/sovansaha/funpred-3.0"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/evolstruct-phogly",
		URL:          []string{"https://github.com/abelavit/evolstruct-phogly"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/apinet",
		URL:          []string{"https://github.com/han-siyu/apinet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metaseek",
		URL:          []string{"https://github.com/metaseek-sequencing-data-discovery/metaseek"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/glsnn",
		URL:          []string{"https://github.com/htzhaoecust/glsnn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/capc-map",
		URL:          []string{"https://github.com/cbrackley/capc-map"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/circsplice",
		URL:          []string{"https://github.com/genefeng/circsplice"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/repviz",
		URL:          []string{"https://github.com/elolab/repviz"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nsclc-subtype-classification",
		URL:          []string{"https://github.com/ransulab/nsclc-subtype-classification"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hergepred",
		URL:          []string{"https://github.com/yangkuoone/hergepred"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scanvis",
		URL:          []string{"https://github.com/nygenome/scanvis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/impact",
		URL:          []string{"https://github.com/lmse/impact"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/diffexpy",
		URL:          []string{"https://github.com/bagherilab/diffexpy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/featureselect",
		URL:          []string{"https://github.com/lbbsoft/featureselect"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ngsld",
		URL:          []string{"https://github.com/fgvieira/ngsld"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/idssr",
		URL:          []string{"https://github.com/allsummerking/idssr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/noisecancellingrepeatfinder",
		URL:          []string{"https://github.com/makovalab-psu/noisecancellingrepeatfinder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/imperfect-gold-standard",
		URL:          []string{"https://github.com/ksny/imperfect-gold-standard"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/og-consistency-pipeline",
		URL:          []string{"https://github.com/meringlab/og_consistency_pipeline"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dnascan",
		URL:          []string{"https://github.com/khp-informatics/dnascan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/orca",
		URL:          []string{"https://github.com/bcgsc/orca"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/single-cell-review",
		URL:          []string{"https://github.com/kuanglab/single-cell-review"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/adcp",
		URL:          []string{"https://github.com/ccsb-scripps/adcp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sodapop",
		URL:          []string{"https://github.com/louisgt/sodapop"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hiv-gag-immunogens",
		URL:          []string{"https://github.com/faraz107/hiv-gag-immunogens"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mgicnn",
		URL:          []string{"https://github.com/ku-milab/mgicnn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/shouji",
		URL:          []string{"https://github.com/cmu-safari/shouji"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cytobackbone",
		URL:          []string{"https://github.com/tchitchek-lab/cytobackbone"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ddot",
		URL:          []string{"https://github.com/idekerlab/ddot"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/uncover",
		URL:          []string{"https://github.com/vandinlab/uncover"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pass",
		URL:          []string{"https://github.com/wupengomics/pass"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/longitudinal-microbiome-analysis-public",
		URL:          []string{"https://github.com/jlugomar/longitudinal_microbiome_analysis_public"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/trader",
		URL:          []string{"https://github.com/lbbsoft/trader"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nx4",
		URL:          []string{"https://github.com/nx4/nx4"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sparsemcmm",
		URL:          []string{"https://github.com/chanw0/sparsemcmm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mixed-effect-gams",
		URL:          []string{"https://github.com/eric-pedersen/mixed-effect-gams"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biodiscml",
		URL:          []string{"https://github.com/mickaelleclercq/biodiscml"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dognet",
		URL:          []string{"https://github.com/kulikovv/dognet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/degsm",
		URL:          []string{"https://github.com/hitbc/degsm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hocnnlb",
		URL:          []string{"https://github.com/nwpu-903pr/hocnnlb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ms-empire",
		URL:          []string{"https://github.com/zimmerlab/ms-empire"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mwas-biomarkers",
		URL:          []string{"https://github.com/microava/mwas-biomarkers"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fp2vec",
		URL:          []string{"https://github.com/wsjeon92/fp2vec"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pharmaconer",
		URL:          []string{"https://github.com/plantl-sanidad/pharmaconer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/swift-t-variant-calling",
		URL:          []string{"https://github.com/ncsa/swift-t-variant-calling"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/meta-gdbp",
		URL:          []string{"https://github.com/ransulab/meta-gdbp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ebwt2snp",
		URL:          []string{"https://github.com/nicolaprezza/ebwt2snp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sbgntikz",
		URL:          []string{"https://github.com/adrienrougny/sbgntikz"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scout",
		URL:          []string{"https://github.com/statway/scout"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/enimpute",
		URL:          []string{"https://github.com/zhangxf-ccnu/enimpute"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/simug",
		URL:          []string{"https://github.com/yjx1217/simug"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/respre",
		URL:          []string{"https://github.com/leeyang/respre"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/abmda",
		URL:          []string{"https://github.com/githubcode007/abmda"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/colocalizr",
		URL:          []string{"https://github.com/kroemerlab/colocalizr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tar-vir",
		URL:          []string{"https://github.com/chjiao/tar-vir"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/clustermq",
		URL:          []string{"https://github.com/mschubert/clustermq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/clustermq-performance",
		URL:          []string{"https://github.com/mschubert/clustermq-performance"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nanodj",
		URL:          []string{"https://github.com/genomicsiter/nanodj"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/boolean-t2dm",
		URL:          []string{"https://github.com/jiezheng-shanghaitech/boolean-t2dm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sumer",
		URL:          []string{"https://github.com/bzhanglab/sumer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepdr",
		URL:          []string{"https://github.com/chengf-lab/deepdr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/quantilebootstrap",
		URL:          []string{"https://github.com/owatson/quantilebootstrap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dnascent",
		URL:          []string{"https://github.com/mboemo/dnascent"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/survindel",
		URL:          []string{"https://github.com/mesh89/survindel"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/depecher",
		URL:          []string{"https://github.com/theorell/depecher"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/masscomp",
		URL:          []string{"https://github.com/iochoa/masscomp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/iscore",
		URL:          []string{"https://github.com/deeprank/iscore"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cusnn",
		URL:          []string{"https://github.com/tudelft/cusnn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/3d-printed-radiographic-test-tools",
		URL:          []string{"https://github.com/upstate3dlab/3d-printed-radiographic-test-tools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/e-prospect",
		URL:          []string{"https://github.com/max-uhlich/e-prospect"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cissticp",
		URL:          []string{"https://github.com/ayushisinha/cissticp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/paga",
		URL:          []string{"https://github.com/theislab/paga"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scanneo",
		URL:          []string{"https://github.com/ylab-hi/scanneo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/neuromeasure",
		URL:          []string{"https://github.com/edwardslabneurosci/neuromeasure"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lambda",
		URL:          []string{"https://github.com/tsteelejohnson91/lambda"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sunbeam",
		URL:          []string{"https://github.com/sunbeam-labs/sunbeam"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/disdrugpred",
		URL:          []string{"https://github.com/pingxuan-hlju/disdrugpred"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/janus",
		URL:          []string{"https://github.com/ccqc/janus"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fastqpuri",
		URL:          []string{"https://github.com/jengelmann/fastqpuri"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/genomewarp",
		URL:          []string{"https://github.com/verilylifesciences/genomewarp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/magi",
		URL:          []string{"https://github.com/biorack/magi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/founder-sequences",
		URL:          []string{"https://github.com/tsnorri/founder-sequences"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hyasp",
		URL:          []string{"https://github.com/cchauve/hyasp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/and",
		URL:          []string{"https://github.com/amorgani/and"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/advanced-multiloops",
		URL:          []string{"https://github.com/maxhwardg/advanced_multiloops"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/riginv",
		URL:          []string{"https://github.com/utbioinf/riginv"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bamse",
		URL:          []string{"https://github.com/hoseint/bamse"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/random-forest-protein-ligand-decoy-detection",
		URL:          []string{"https://github.com/junpei000/random_forest_protein_ligand_decoy_detection"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fardeep",
		URL:          []string{"https://github.com/yuninghao/fardeep"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dna-nn",
		URL:          []string{"https://github.com/lh3/dna-nn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hapmc",
		URL:          []string{"https://github.com/smajidian/hapmc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ped",
		URL:          []string{"https://github.com/akiomiyao/ped"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mio",
		URL:          []string{"https://github.com/wrede/mio"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bio3d",
		URL:          []string{"https://github.com/grantlab/bio3d"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/confindr",
		URL:          []string{"https://github.com/olc-bioinformatics/confindr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/solrplantapi",
		URL:          []string{"https://github.com/bcbi/solrplantapi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nojah",
		URL:          []string{"https://github.com/bbisr-shinyapps/nojah"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pldist",
		URL:          []string{"https://github.com/aplantin/pldist"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/viruses-classifier",
		URL:          []string{"https://github.com/wojciech-galan/viruses_classifier"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gemini",
		URL:          []string{"https://github.com/sellerslab/gemini"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bio-gradient-descent",
		URL:          []string{"https://github.com/yukinoi/bio_gradient_descent"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/capssa",
		URL:          []string{"https://github.com/yjjang/capssa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deviate",
		URL:          []string{"https://github.com/w-l/deviate"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bofdat",
		URL:          []string{"https://github.com/jclachance/bofdat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/epic2",
		URL:          []string{"https://github.com/biocore-ntnu/epic2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/digitalcellsorter",
		URL:          []string{"https://github.com/sdomanskyi/digitalcellsorter"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cnet",
		URL:          []string{"https://github.com/bsml320/cnet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pypore",
		URL:          []string{"https://github.com/rsemeraro/pypore"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cbig",
		URL:          []string{"https://github.com/thomasyeolab/cbig"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/reactidr",
		URL:          []string{"https://github.com/carushi/reactidr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/convlstmforgr",
		URL:          []string{"https://github.com/guangmingzhu/convlstmforgr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/transposons",
		URL:          []string{"https://github.com/alexshein/transposons"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/drvae",
		URL:          []string{"https://github.com/rampasek/drvae"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/effect-of-the-operational-parameters",
		URL:          []string{"https://github.com/1kovalevskiy/effect-of-the-operational-parameters"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phylostratr",
		URL:          []string{"https://github.com/arendsee/phylostratr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/splicefamalign",
		URL:          []string{"https://github.com/udes-cobius/splicefamalign"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cuedemotion",
		URL:          []string{"https://github.com/naiqixiao/cuedemotion"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hayai-annotation-plants",
		URL:          []string{"https://github.com/kdri-genomics/hayai-annotation-plants"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mda-cnn",
		URL:          []string{"https://github.com/issingjessica/mda-cnn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gne",
		URL:          []string{"https://github.com/kckishan/gne"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gemtools",
		URL:          []string{"https://github.com/sgreer77/gemtools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lncdiff",
		URL:          []string{"https://github.com/qianli10000/lncdiff"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/optical",
		URL:          []string{"https://github.com/shiukumar/optical"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/correlation-between-rna-seq-and-rrbs",
		URL:          []string{"https://github.com/wjlim/correlation-between-rna-seq-and-rrbs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hsesumo",
		URL:          []string{"https://github.com/yosvanylopez/hsesumo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mirfix",
		URL:          []string{"https://github.com/bierinformatik/mirfix"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pbtest-r-package",
		URL:          []string{"https://github.com/yunzhang813/pbtest-r-package"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/grep",
		URL:          []string{"https://github.com/saorisakaue/grep"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hollytibble",
		URL:          []string{"https://github.com/hollytibble"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/svtools",
		URL:          []string{"https://github.com/hall-lab/svtools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gpqtlmapping",
		URL:          []string{"https://github.com/jpvanhat/gpqtlmapping"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/remap",
		URL:          []string{"https://github.com/hansaimlim/remap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sarks",
		URL:          []string{"https://github.com/denniscwylie/sarks"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metabolomics-filtering",
		URL:          []string{"https://github.com/courtneyschiffman/metabolomics-filtering"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/igloo",
		URL:          []string{"https://github.com/zerotonin/igloo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dearwxpub",
		URL:          []string{"https://github.com/deargen/dearwxpub"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/taska",
		URL:          []string{"https://github.com/bioinformatics-ua/taska"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metamsd",
		URL:          []string{"https://github.com/soyoungryu/metamsd"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/christopherblum",
		URL:          []string{"https://github.com/christopherblum"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/isocor",
		URL:          []string{"https://github.com/metasys-lisbp/isocor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/graphaligner",
		URL:          []string{"https://github.com/maickrau/graphaligner"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deeppasta",
		URL:          []string{"https://github.com/arefeen/deeppasta"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cesar2.0",
		URL:          []string{"https://github.com/hillerlab/cesar2.0"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cymira",
		URL:          []string{"https://github.com/dbsloan/cymira"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mcee-2.0",
		URL:          []string{"https://github.com/chentianlu/mcee-2.0"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/epanet-dev",
		URL:          []string{"https://github.com/openwateranalytics/epanet-dev"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ms-lda",
		URL:          []string{"https://github.com/qkirikigaku/ms_lda"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/epilog",
		URL:          []string{"https://github.com/epilog-tool/epilog"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pepfoot",
		URL:          []string{"https://github.com/jbellamycarter/pepfoot"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/molecular-clock",
		URL:          []string{"https://github.com/ugobas/molecular_clock"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tft",
		URL:          []string{"https://github.com/lanec-unifesspa/tft"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/raincloudplots",
		URL:          []string{"https://github.com/raincloudplots/raincloudplots"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/antigenpredictor",
		URL:          []string{"https://github.com/srautonu/antigenpredictor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/jaxbd2k-shortcourse",
		URL:          []string{"https://github.com/thejacksonlaboratory/jaxbd2k-shortcourse"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deplogo",
		URL:          []string{"https://github.com/jstacs/deplogo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scae-ir-spectral-imaging",
		URL:          []string{"https://github.com/arnrau/scae_ir_spectral_imaging"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bdmma",
		URL:          []string{"https://github.com/daizhenwei/bdmma/bdmma"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bdmma-macos",
		URL:          []string{"https://github.com/daizhenwei/bdmma/bdmma_macos"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rnmrfit",
		URL:          []string{"https://github.com/ssokolen/rnmrfit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lep",
		URL:          []string{"https://github.com/daviddaigithub/lep"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/togogenome",
		URL:          []string{"https://github.com/togogenome"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/togostanza",
		URL:          []string{"https://github.com/togostanza"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bigred",
		URL:          []string{"https://github.com/ac2278/bigred"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lncrna-homologs",
		URL:          []string{"https://github.com/bioinformatics-sannio/lncrna-homologs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/abdominal-mr-phantom",
		URL:          []string{"https://github.com/seiberlichlab/abdominal_mr_phantom"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/msc-tree-resampling",
		URL:          []string{"https://github.com/dbsloan/msc_tree_resampling"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/spin-scenario",
		URL:          []string{"https://github.com/spin-scenario"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pnnl",
		URL:          []string{"https://github.com/pnnl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/flycop",
		URL:          []string{"https://github.com/beatrizgj/flycop"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/querymed",
		URL:          []string{"https://github.com/yannrivault/querymed"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pandelos",
		URL:          []string{"https://github.com/giugnolab/pandelos"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bioinfo",
		URL:          []string{"https://github.com/bioinfproject/bioinfo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/coac",
		URL:          []string{"https://github.com/chengf-lab/coac"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/edge-in-tcga",
		URL:          []string{"https://github.com/andreamrau/edge-in-tcga"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/copal",
		URL:          []string{"https://github.com/cmbi/copal"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pals",
		URL:          []string{"https://github.com/npnl/pals"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/simulatingcochlearimplants",
		URL:          []string{"https://github.com/clips/simulatingcochlearimplants"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metacluster",
		URL:          []string{"https://github.com/mbanf/metacluster"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/simgwas",
		URL:          []string{"https://github.com/chr1swallace/simgwas"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/wisecondorx",
		URL:          []string{"https://github.com/centerformedicalgeneticsghent/wisecondorx"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pcr-error-correction",
		URL:          []string{"https://github.com/crukmi-computationalbiology/pcr_error_correction"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/whisper",
		URL:          []string{"https://github.com/refresh-bio/whisper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gotrapper",
		URL:          []string{"https://github.com/biogenetools/gotrapper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/popcluster",
		URL:          []string{"https://github.com/gurinovich/popcluster"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bispark",
		URL:          []string{"https://github.com/bhi-kimlab/bispark"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ncgocr",
		URL:          []string{"https://github.com/jeroyang/ncgocr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sinnlrr",
		URL:          []string{"https://github.com/zrq0123/sinnlrr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scent",
		URL:          []string{"https://github.com/aet21/scent"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mirfinder",
		URL:          []string{"https://github.com/wangying0128/mirfinder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/camisim",
		URL:          []string{"https://github.com/cami-challenge/camisim"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/decontam",
		URL:          []string{"https://github.com/benjjneb/decontam"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/densitypath",
		URL:          []string{"https://github.com/ucasdp/densitypath"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/caspo-ts",
		URL:          []string{"https://github.com/misbahch6/caspo-ts"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gbs-snp-crop",
		URL:          []string{"https://github.com/halelab/gbs-snp-crop"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gbs-pacecar",
		URL:          []string{"https://github.com/halelab/gbs-pacecar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/plaidoh",
		URL:          []string{"https://github.com/sarahpyfrom/plaidoh"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ddseeker",
		URL:          []string{"https://github.com/cgplab/ddseeker"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/peakplotter",
		URL:          []string{"https://github.com/wtsi-team144/peakplotter"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/transformphenotype",
		URL:          []string{"https://github.com/wtsi-team144/transformphenotype"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/spagi",
		URL:          []string{"https://github.com/vccri/spagi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/map3k8-thyroid-spheres-v-3.0",
		URL:          []string{"https://github.com/francescopappalardo/map3k8-thyroid-spheres-v-3.0"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rna3dcnn",
		URL:          []string{"https://github.com/lijunrna/rna3dcnn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/eklipse",
		URL:          []string{"https://github.com/dooguypapua/eklipse"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/b-lore",
		URL:          []string{"https://github.com/soedinglab/b-lore"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/usefulplants-indicator",
		URL:          []string{"https://github.com/ciat-dapa/usefulplants-indicator"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/multivarnetwork",
		URL:          []string{"https://github.com/jchiquet/multivarnetwork"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/npnn",
		URL:          []string{"https://github.com/htzhaoecust/npnn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gtm-generativetopographicmapping",
		URL:          []string{"https://github.com/hkaneko1985/gtm-generativetopographicmapping"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/irecspot",
		URL:          []string{"https://github.com/mrzresearcharena/irecspot"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kf-arabidopsis-grna",
		URL:          []string{"https://github.com/bartongroup/kf_arabidopsis-grna"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/protspam",
		URL:          []string{"https://github.com/jschellh/protspam"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dppi",
		URL:          []string{"https://github.com/hashemifar/dppi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/guv-ap",
		URL:          []string{"https://github.com/ag-roemer/guv-ap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tetyper",
		URL:          []string{"https://github.com/aesheppard/tetyper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/antibodyinterfaceprediction",
		URL:          []string{"https://github.com/sebastiandaberdaku/antibodyinterfaceprediction"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/c3d",
		URL:          []string{"https://github.com/lupienlaborganization/c3d"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/combine-cnn-enhancer-and-promoters",
		URL:          []string{"https://github.com/zzumn/combine-cnn-enhancer-and-promoters"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/el-smurf",
		URL:          []string{"https://github.com/qust-aibbdrc/el-smurf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mic-meme",
		URL:          []string{"https://github.com/hkwkevin28/mic-meme"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/emclarity",
		URL:          []string{"https://github.com/bhimes/emclarity"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/iguide",
		URL:          []string{"https://github.com/cnobles/iguide"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/neubi",
		URL:          []string{"https://github.com/nafizh/neubi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bixgboost",
		URL:          []string{"https://github.com/zrq0123/bixgboost"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sedef",
		URL:          []string{"https://github.com/vpc-ccg/sedef"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/curatio",
		URL:          []string{"https://github.com/qiwenkang/curatio"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pose",
		URL:          []string{"https://github.com/cdcgov/pose"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/diffnetfdr",
		URL:          []string{"https://github.com/zhangxf-ccnu/diffnetfdr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ancestral-blocks-reconstruction",
		URL:          []string{"https://github.com/nguyenngochuy91/ancestral-blocks-reconstruction"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pavooc",
		URL:          []string{"https://github.com/moritzschaefer/pavooc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/targetclone",
		URL:          []string{"https://github.com/umcugenetics/targetclone"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mqc",
		URL:          []string{"https://github.com/biobix/mqc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mixsubhap",
		URL:          []string{"https://github.com/yixuanwang1120/mixsubhap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepseqpan",
		URL:          []string{"https://github.com/pcpliu/deepseqpan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/polydfe",
		URL:          []string{"https://github.com/paula-tataru/polydfe"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biolitmap",
		URL:          []string{"https://github.com/inab/biolitmap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deept3",
		URL:          []string{"https://github.com/lje00006/deept3"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/optimir",
		URL:          []string{"https://github.com/florianthibord/optimir"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/chiapop",
		URL:          []string{"https://github.com/wh90999/chiapop"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/spind",
		URL:          []string{"https://github.com/liulab-csrc/spind"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/helmsman",
		URL:          []string{"https://github.com/carjed/helmsman"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sphyr",
		URL:          []string{"https://github.com/elkebir-group/sphyr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/recap",
		URL:          []string{"https://github.com/theodorejperkins/recap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/omgraph",
		URL:          []string{"https://github.com/kingufl/omgraph"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/neuroinformatics",
		URL:          []string{"https://github.com/neuroinformatics"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fastlzerospikeinference",
		URL:          []string{"https://github.com/jewellsean/fastlzerospikeinference"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cysmotifsearcher",
		URL:          []string{"https://github.com/fallandar/cysmotifsearcher"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/adaptive-geometric-search-for-protein-design",
		URL:          []string{"https://github.com/jiangtian/adaptive-geometric-search-for-protein-design"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/elemcor",
		URL:          []string{"https://github.com/4dsoftware/elemcor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mrgbp",
		URL:          []string{"https://github.com/skouchaki/mrgbp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepaffinity",
		URL:          []string{"https://github.com/shen-lab/deepaffinity"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sw-tandem",
		URL:          []string{"https://github.com/logic09/sw-tandem"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bacpacs",
		URL:          []string{"https://github.com/barashe/bacpacs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/identification-of-brain-based-disorders",
		URL:          []string{"https://github.com/xinyuzhao/identification-of-brain-based-disorders"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rappas",
		URL:          []string{"https://github.com/blinard-bioinfo/rappas"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pwc-net",
		URL:          []string{"https://github.com/nvlabs/pwc-net"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tadoss",
		URL:          []string{"https://github.com/lafita/tadoss"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepdta",
		URL:          []string{"https://github.com/hkmztrk/deepdta"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/squadbookchapter",
		URL:          []string{"https://github.com/caramirezal/squadbookchapter"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/beyondbinaryparcellationdata",
		URL:          []string{"https://github.com/rainerboegle/beyondbinaryparcellationdata"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rmetl",
		URL:          []string{"https://github.com/hitbc/rmetl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ebgsea",
		URL:          []string{"https://github.com/aet21/ebgsea"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dlprb",
		URL:          []string{"https://github.com/ilanbb/dlprb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kusterlab",
		URL:          []string{"https://github.com/kusterlab"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pact",
		URL:          []string{"https://github.com/jklesmith/pact"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lcslcis",
		URL:          []string{"https://github.com/holtzy/lcslcis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bermuda",
		URL:          []string{"https://github.com/abikoushi/bermuda"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ccc",
		URL:          []string{"https://github.com/lucanard/ccc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/opal-plus",
		URL:          []string{"https://github.com/roneshsharma/opal-plus"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sccg",
		URL:          []string{"https://github.com/jhchen5/sccg"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sifrproject",
		URL:          []string{"https://github.com/sifrproject"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mcimpute-scrnaseq",
		URL:          []string{"https://github.com/aanchalmongia/mcimpute_scrnaseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/imaging-of-drugs-distribution-and-quantifications",
		URL:          []string{"https://github.com/francescafalcetta/imaging_of_drugs_distribution_and_quantifications"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ppis-wdsvm",
		URL:          []string{"https://github.com/qust-aibbdrc/ppis-wdsvm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/branchinggps",
		URL:          []string{"https://github.com/cap76/branchinggps"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/med4way",
		URL:          []string{"https://github.com/anddis/med4way"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ipct",
		URL:          []string{"https://github.com/muhammadshoaib/ipct"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/prosstt",
		URL:          []string{"https://github.com/soedinglab/prosstt"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/snp-select",
		URL:          []string{"https://github.com/transgenomicsosu/snp-select"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/antivpp",
		URL:          []string{"https://github.com/bio-coding/antivpp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hierarchical-hotnet",
		URL:          []string{"https://github.com/raphael-group/hierarchical-hotnet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ptolemy",
		URL:          []string{"https://github.com/abeellab/ptolemy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/opa2vec",
		URL:          []string{"https://github.com/bio-ontology-research-group/opa2vec"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/extramp",
		URL:          []string{"https://github.com/ridgelab/extramp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vargenius",
		URL:          []string{"https://github.com/frankmusacchia/vargenius"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/krakenuniq",
		URL:          []string{"https://github.com/fbreitwieser/krakenuniq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tner",
		URL:          []string{"https://github.com/ctdna/tner"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/boost-hic",
		URL:          []string{"https://github.com/leopoldc/boost-hic"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/garleek",
		URL:          []string{"https://github.com/insilichem/garleek"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bulkvis",
		URL:          []string{"https://github.com/looselab/bulkvis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/abis",
		URL:          []string{"https://github.com/giannimonaco/abis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/benchmark-models",
		URL:          []string{"https://github.com/benchmarking-initiative/benchmark-models"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ogfsc",
		URL:          []string{"https://github.com/xzouprojects/ogfsc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ptm-x",
		URL:          []string{"https://github.com/huangyh09/ptm-x"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ontologyeval",
		URL:          []string{"https://github.com/schulzlab/ontologyeval"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/raven",
		URL:          []string{"https://github.com/sysbiochalmers/raven"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/streptomyces-coelicolor-gem",
		URL:          []string{"https://github.com/sysbiochalmers/streptomyces_coelicolor-gem"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/substra",
		URL:          []string{"https://github.com/sahandk/substra"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sia",
		URL:          []string{"https://github.com/erechtheus/sia"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rinbix",
		URL:          []string{"https://github.com/insilico/rinbix"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/priorknowledgeepistasisrank",
		URL:          []string{"https://github.com/insilico/priorknowledgeepistasisrank"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biograph",
		URL:          []string{"https://github.com/icarpa-tblab/biograph"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mitefinder",
		URL:          []string{"https://github.com/screamer/mitefinder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nrps-linker",
		URL:          []string{"https://github.com/swfarag/nrps-linker"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gridlmm",
		URL:          []string{"https://github.com/deruncie/gridlmm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hnematzadeh",
		URL:          []string{"https://github.com/hnematzadeh"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/q2-longitudinal",
		URL:          []string{"https://github.com/qiime2/q2-longitudinal"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/abb",
		URL:          []string{"https://github.com/francesc-muyas/abb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/synergy-screen",
		URL:          []string{"https://github.com/arnaudmgh/synergy-screen"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepphos",
		URL:          []string{"https://github.com/ustchilab/deepphos"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cddforfmri",
		URL:          []string{"https://github.com/namgillee/cddforfmri"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pnerf",
		URL:          []string{"https://github.com/aqlaboratory/pnerf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dbtora",
		URL:          []string{"https://github.com/ime-tmp-ffm/dbtora"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fuseq",
		URL:          []string{"https://github.com/nghiavtr/fuseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deicode",
		URL:          []string{"https://github.com/biocore/deicode"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/largegopred",
		URL:          []string{"https://github.com/gauravpandeylab/largegopred"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/csbfinder",
		URL:          []string{"https://github.com/dinasv/csbfinder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kbet",
		URL:          []string{"https://github.com/theislab/kbet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mtbseq-source",
		URL:          []string{"https://github.com/ngs-fzb/mtbseq_source"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/equivariantnetworks",
		URL:          []string{"https://github.com/luntergroup/equivariantnetworks"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lacroixlaurent",
		URL:          []string{"https://github.com/lacroixlaurent"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/wgd",
		URL:          []string{"https://github.com/arzwa/wgd"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/clust",
		URL:          []string{"https://github.com/baselabujamous/clust"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dgmdl",
		URL:          []string{"https://github.com/luoping1004/dgmdl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/destin",
		URL:          []string{"https://github.com/urrutiag/destin"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biojava-tutorial",
		URL:          []string{"https://github.com/biojava/biojava-tutorial"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scvi",
		URL:          []string{"https://github.com/yoseflab/scvi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vep-plugins",
		URL:          []string{"https://github.com/ensembl/vep_plugins"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ctc-screening-cnn",
		URL:          []string{"https://github.com/chenyzstju/ctc_screening_cnn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/uvp",
		URL:          []string{"https://github.com/cptr-reseqtb/uvp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vhost-classifier",
		URL:          []string{"https://github.com/kzra/vhost-classifier"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ads-hcspark",
		URL:          []string{"https://github.com/scut-ccnl/ads-hcspark"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nipter",
		URL:          []string{"https://github.com/molgenis/nipter"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phenotypeseeker",
		URL:          []string{"https://github.com/bioinfo-ut/phenotypeseeker"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biomethyl",
		URL:          []string{"https://github.com/yuewangpanda/biomethyl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gfplain",
		URL:          []string{"https://github.com/fnardi/gfplain"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/slimenrich",
		URL:          []string{"https://github.com/slimsuite/slimenrich"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/olga",
		URL:          []string{"https://github.com/zsethna/olga"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/g4catchall",
		URL:          []string{"https://github.com/odoluca/g4catchall"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hpc-workflow-manager",
		URL:          []string{"https://github.com/fiji-hpc/hpc-workflow-manager"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pose-regularization",
		URL:          []string{"https://github.com/ifl-camp/pose_regularization"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/procell",
		URL:          []string{"https://github.com/aresio/procell"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bart-bma",
		URL:          []string{"https://github.com/belindahernandez/bart-bma"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ordino",
		URL:          []string{"https://github.com/caleydo/ordino"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biocompute-objects",
		URL:          []string{"https://github.com/biocompute-objects"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gmrad",
		URL:          []string{"https://github.com/tongchf/gmrad"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/invertinggan",
		URL:          []string{"https://github.com/tonicreswell/invertinggan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/genecnv",
		URL:          []string{"https://github.com/vkozareva/genecnv"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scfba",
		URL:          []string{"https://github.com/bimib-disco/scfba"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dhaka",
		URL:          []string{"https://github.com/microsoftgenomics/dhaka"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/seqsleepnet",
		URL:          []string{"https://github.com/pquochuy/seqsleepnet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/agrp",
		URL:          []string{"https://github.com/hqwang126/agrp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mmcnn",
		URL:          []string{"https://github.com/psychopa4/mmcnn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gpds",
		URL:          []string{"https://github.com/cap76/gpds"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/asd-genes-prediction",
		URL:          []string{"https://github.com/muh-asif/asd-genes-prediction"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/argdit",
		URL:          []string{"https://github.com/phglab/argdit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cnr",
		URL:          []string{"https://github.com/nki-ccb/cnr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cnr-analyses",
		URL:          []string{"https://github.com/nki-ccb/cnr-analyses"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/webmolcs",
		URL:          []string{"https://github.com/reymond-group/webmolcs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/caffe-3d-faster-rcnn",
		URL:          []string{"https://github.com/superxuang/caffe_3d_faster_rcnn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rgv",
		URL:          []string{"https://github.com/fchalmel/rgv"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/shinycnv",
		URL:          []string{"https://github.com/gzhmat/shinycnv"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/glycanformatconverter",
		URL:          []string{"https://github.com/glycoinfo/glycanformatconverter"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/grassmanncluster",
		URL:          []string{"https://github.com/michaelsharpnack/grassmanncluster"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cafemocha",
		URL:          []string{"https://github.com/binaypanda/cafemocha"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/abra2",
		URL:          []string{"https://github.com/mozack/abra2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scdnet",
		URL:          []string{"https://github.com/chenlabgccri/scdnet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pymcpsc",
		URL:          []string{"https://github.com/xulesc/pymcpsc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pdb-tools",
		URL:          []string{"https://github.com/haddocking/pdb-tools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sian",
		URL:          []string{"https://github.com/pogudingleb/sian"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dnadataaugmentation",
		URL:          []string{"https://github.com/zhanglabtools/dnadataaugmentation"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/miso-example",
		URL:          []string{"https://github.com/yonghuidong/miso_example"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/silentmutations",
		URL:          []string{"https://github.com/desiro/silentmutations"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ldjump",
		URL:          []string{"https://github.com/phhermann/ldjump"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/modsara2",
		URL:          []string{"https://github.com/feifeixiaousc/modsara2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mer",
		URL:          []string{"https://github.com/lasigebiotm/mer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/repdenovo",
		URL:          []string{"https://github.com/reedwarbler/repdenovo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nemo",
		URL:          []string{"https://github.com/shamir-lab/nemo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phat",
		URL:          []string{"https://github.com/chgibb/phat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/prosave",
		URL:          []string{"https://github.com/mahajanlab/prosave"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/offtargetpredict",
		URL:          []string{"https://github.com/penn-hui/offtargetpredict"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/comorbidity4j",
		URL:          []string{"https://github.com/fra82/comorbidity4j"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/topas",
		URL:          []string{"https://github.com/bjyoontamu/topas"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/colliderapp",
		URL:          []string{"https://github.com/migariane/colliderapp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metasmc",
		URL:          []string{"https://github.com/tarjxvf/metasmc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lwlrtest",
		URL:          []string{"https://github.com/ramonoller/lwlrtest"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepubi",
		URL:          []string{"https://github.com/sunmile/deepubi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/stablede",
		URL:          []string{"https://github.com/linbingqing/stablede"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/distributedlearningpredictor",
		URL:          []string{"https://github.com/ziyili20/distributedlearningpredictor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/imputation",
		URL:          []string{"https://github.com/joewuca/imputation"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dpm-lgcp",
		URL:          []string{"https://github.com/bsharmi/dpm-lgcp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/depop",
		URL:          []string{"https://github.com/lab9arriam/depop"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fastderain",
		URL:          []string{"https://github.com/taixiangjiang/fastderain"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/multi-csar",
		URL:          []string{"https://github.com/ablab-nthu/multi-csar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/flexgeo",
		URL:          []string{"https://github.com/amarinhosn/flexgeo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/beam",
		URL:          []string{"https://github.com/sayakamiura/beam"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/script-hmm-replication-timing",
		URL:          []string{"https://github.com/pouletaxel/script_hmm_replication_timing"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/signature-sj",
		URL:          []string{"https://github.com/vyacheslav-tsivina/signature-sj"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cafu",
		URL:          []string{"https://github.com/cma2015/cafu"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sda",
		URL:          []string{"https://github.com/mvollger/sda"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/superbubble",
		URL:          []string{"https://github.com/fabianexe/superbubble"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/capsnet-ptm",
		URL:          []string{"https://github.com/duolinwang/capsnet_ptm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tlcr-rl",
		URL:          []string{"https://github.com/junjun-jiang/tlcr-rl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/clinical-citation-sentiment",
		URL:          []string{"https://github.com/kilicogluh/clinical-citation-sentiment"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/neodti",
		URL:          []string{"https://github.com/fangpingwan/neodti"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mapoptics",
		URL:          []string{"https://github.com/fadymohareb/mapoptics"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/delta-statistic",
		URL:          []string{"https://github.com/mrborges23/delta_statistic"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/edetect",
		URL:          []string{"https://github.com/zi-lab/edetect"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vcmicrobiome",
		URL:          []string{"https://github.com/jingzhai63/vcmicrobiome"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/psbvb",
		URL:          []string{"https://github.com/lauzingaretti/psbvb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bdbg",
		URL:          []string{"https://github.com/rongjiewang/bdbg"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/icr142-benchmarker",
		URL:          []string{"https://github.com/rahmanteamdevelopment/icr142_benchmarker"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/favites",
		URL:          []string{"https://github.com/niemasd/favites"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gospel",
		URL:          []string{"https://github.com/infinite-point/gospel"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/quantinemo",
		URL:          []string{"https://github.com/jgx65/quantinemo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/squeezemeta",
		URL:          []string{"https://github.com/jtamames/squeezemeta"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/varanto",
		URL:          []string{"https://github.com/oqe/varanto"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/foresee",
		URL:          []string{"https://github.com/jrc-combine/foresee"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/algorithm-performance-analysis",
		URL:          []string{"https://github.com/atyryshkina/algorithm-performance-analysis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sid",
		URL:          []string{"https://github.com/cbskust/sid"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepdrug3d",
		URL:          []string{"https://github.com/pulimeng/deepdrug3d"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mxfold",
		URL:          []string{"https://github.com/keio-bioinformatics/mxfold"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ardiss",
		URL:          []string{"https://github.com/borgwardtlab/ardiss"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepcseqsite",
		URL:          []string{"https://github.com/yfcuifaith/deepcseqsite"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepacet",
		URL:          []string{"https://github.com/sunmile/deepacet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hfufs",
		URL:          []string{"https://github.com/swainechen/hfufs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/codeml-modl",
		URL:          []string{"https://github.com/jehops/codeml_modl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/magpy",
		URL:          []string{"https://github.com/watsonlab/magpy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/myelinq",
		URL:          []string{"https://github.com/parham-ap/myelinq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/poplddecay",
		URL:          []string{"https://github.com/bgi-shenzhen/poplddecay"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lorisnanni",
		URL:          []string{"https://github.com/lorisnanni"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vapid",
		URL:          []string{"https://github.com/rcs333/vapid"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/laggedorderedlassonetwork",
		URL:          []string{"https://github.com/pn51/laggedorderedlassonetwork"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kurtosis-conservation",
		URL:          []string{"https://github.com/alexander-nash/kurtosis_conservation"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/naf",
		URL:          []string{"https://github.com/kirillkryukov/naf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/molti-dream",
		URL:          []string{"https://github.com/gilles-didier/molti-dream"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/airnet-pytorch",
		URL:          []string{"https://github.com/soeaver/airnet-pytorch"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deerect-polya",
		URL:          []string{"https://github.com/likesum/deerect-polya"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/seevis",
		URL:          []string{"https://github.com/ghattab/seevis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bridge",
		URL:          []string{"https://github.com/scientificomputing/bridge"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mitodix",
		URL:          []string{"https://github.com/topazgl/mitodix"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/stochastic-diffusion",
		URL:          []string{"https://github.com/ronaldolab/stochastic_diffusion"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tf2exp",
		URL:          []string{"https://github.com/wqshi/tf2exp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gff3toolkit",
		URL:          []string{"https://github.com/nal-i5k/gff3toolkit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/magic",
		URL:          []string{"https://github.com/gimelbrantlab/magic"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/diffpop",
		URL:          []string{"https://github.com/ferlicjl/diffpop"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mirlocator",
		URL:          []string{"https://github.com/cma2015/mirlocator"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cnefinder",
		URL:          []string{"https://github.com/lorrainea/cnefinder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fastp",
		URL:          []string{"https://github.com/opengene/fastp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cnvalidator",
		URL:          []string{"https://github.com/kuhnerlab/cnvalidator"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/interface-classifier",
		URL:          []string{"https://github.com/haddocking/interface-classifier"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bmf-qsar",
		URL:          []string{"https://github.com/grisonifr/bmf_qsar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cwdtw",
		URL:          []string{"https://github.com/realbigws/cwdtw"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hivmmer",
		URL:          []string{"https://github.com/kantorlab/hivmmer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/swne",
		URL:          []string{"https://github.com/yanwu2014/swne"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bgsa",
		URL:          []string{"https://github.com/sdu-hpcl/bgsa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dgcox",
		URL:          []string{"https://github.com/luigiaugugliaro/dgcox"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/spurf",
		URL:          []string{"https://github.com/krdav/spurf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/methylimp",
		URL:          []string{"https://github.com/pdilena/methylimp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/drugcombinationprediction",
		URL:          []string{"https://github.com/roosevelt-pku/drugcombinationprediction"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deeplightfieldssr",
		URL:          []string{"https://github.com/spatialsr/deeplightfieldssr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/taiji",
		URL:          []string{"https://github.com/guanlab/taiji"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/integrate",
		URL:          []string{"https://github.com/haleyeidem/integrate"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cluesv1",
		URL:          []string{"https://github.com/wuchao1984/cluesv1"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/skpcr",
		URL:          []string{"https://github.com/weikanggong/skpcr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/xolotl",
		URL:          []string{"https://github.com/sg-s/xolotl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pconsc4",
		URL:          []string{"https://github.com/elofssonlab/pconsc4"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepcas9",
		URL:          []string{"https://github.com/lje00006/deepcas9"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ibpp",
		URL:          []string{"https://github.com/hahatcdg/ibpp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dialignr",
		URL:          []string{"https://github.com/roestlab/dialignr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ntp-ersn",
		URL:          []string{"https://github.com/mohamedlahdour/ntp-ersn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phoglystruct",
		URL:          []string{"https://github.com/abelavit/phoglystruct"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/logloss-beraf",
		URL:          []string{"https://github.com/bioinformatics-ibch/logloss-beraf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/funmappone",
		URL:          []string{"https://github.com/grecolab/funmappone"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lzerospikeinference",
		URL:          []string{"https://github.com/jewellsean/lzerospikeinference"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/circtools",
		URL:          []string{"https://github.com/dieterich-lab/circtools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mashmap",
		URL:          []string{"https://github.com/marbl/mashmap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/psims",
		URL:          []string{"https://github.com/mobiusklein/psims"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ssgsea2.0",
		URL:          []string{"https://github.com/broadinstitute/ssgsea2.0"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tide-for-ptm-search",
		URL:          []string{"https://github.com/tide-for-ptm-search/tide-for-ptm-search"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/minicom",
		URL:          []string{"https://github.com/yuansliu/minicom"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sjaracne",
		URL:          []string{"https://github.com/jyyulab/sjaracne"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/polster",
		URL:          []string{"https://github.com/yoshida-lab/polster"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gsmodutils",
		URL:          []string{"https://github.com/sbrcnottingham/gsmodutils"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pmd",
		URL:          []string{"https://github.com/yufree/pmd"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fta",
		URL:          []string{"https://github.com/mpgerstl/fta"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/casmap",
		URL:          []string{"https://github.com/borgwardtlab/casmap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/off-target-prediction",
		URL:          []string{"https://github.com/michaellinn/off_target_prediction"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/off_target_prediction",
		URL:          []string{"https://github.com/michaellinn/off_target_prediction"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mionsite",
		URL:          []string{"https://github.com/liangqiaogu/mionsite"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/libra-x",
		URL:          []string{"https://github.com/quantum-dynamics-hub/libra-x"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phytoassembly",
		URL:          []string{"https://github.com/cpolano/phytoassembly"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/brwhnha",
		URL:          []string{"https://github.com/myl446/brwhnha"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/protein-folding-decoy-set",
		URL:          []string{"https://github.com/junpei000/protein_folding-decoy-set"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/crisflash",
		URL:          []string{"https://github.com/crisflash"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/smurf",
		URL:          []string{"https://github.com/lupienlab/smurf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/jbcb2018",
		URL:          []string{"https://github.com/guoyanb/jbcb2018"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/evidenceontology",
		URL:          []string{"https://github.com/evidenceontology"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pargenes",
		URL:          []string{"https://github.com/benoitmorel/pargenes"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/interminer",
		URL:          []string{"https://github.com/intermine/interminer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/grandprix",
		URL:          []string{"https://github.com/manchesterbioinference/grandprix"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ontologies",
		URL:          []string{"https://github.com/biosemantics/ontologies"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lstmvoter",
		URL:          []string{"https://github.com/texttechnologylab/lstmvoter"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nbfpeg",
		URL:          []string{"https://github.com/ramanlab/nbfpeg"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sieve-score",
		URL:          []string{"https://github.com/sekijima-lab/sieve-score"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bcool",
		URL:          []string{"https://github.com/malfoy/bcool"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pctpred",
		URL:          []string{"https://github.com/liulab-hzau/pctpred"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/demixtallmaterials",
		URL:          []string{"https://github.com/wwylab/demixtallmaterials"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepubiquitylation",
		URL:          []string{"https://github.com/jiagenlee/deepubiquitylation"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bioinstaller",
		URL:          []string{"https://github.com/jhuanglab/bioinstaller"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/multiwayregression",
		URL:          []string{"https://github.com/lockef/multiwayregression"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fit-sne",
		URL:          []string{"https://github.com/klugerlab/fit-sne"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/t-sne-heatmaps",
		URL:          []string{"https://github.com/klugerlab/t-sne-heatmaps"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/compath",
		URL:          []string{"https://github.com/compath"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lilikoi",
		URL:          []string{"https://github.com/lanagarmire/lilikoi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepisofun",
		URL:          []string{"https://github.com/dls03/deepisofun"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/clairvoyante",
		URL:          []string{"https://github.com/aquaskyline/clairvoyante"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pytanfinder",
		URL:          []string{"https://github.com/kirovez/pytanfinder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/m3drop",
		URL:          []string{"https://github.com/tallulandrews/m3drop"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/predgly",
		URL:          []string{"https://github.com/yujialinncu/predgly"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/regmex",
		URL:          []string{"https://github.com/muhligs/regmex"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/brca-analyzer",
		URL:          []string{"https://github.com/aakechin/brca-analyzer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/computational",
		URL:          []string{"https://github.com/computational"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/iham",
		URL:          []string{"https://github.com/dessimozlab/iham"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pyham",
		URL:          []string{"https://github.com/dessimozlab/pyham"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/factorial-hmm",
		URL:          []string{"https://github.com/regevs/factorial_hmm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ngtas-pipeline",
		URL:          []string{"https://github.com/cclab-brca/ngtas_pipeline"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/medestrand",
		URL:          []string{"https://github.com/jxu1234/medestrand"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mnem",
		URL:          []string{"https://github.com/cbg-ethz/mnem"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/microrna-nhpred",
		URL:          []string{"https://github.com/myl446/microrna-nhpred"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rl-gensvm",
		URL:          []string{"https://github.com/qust-aibbdrc/rl-gensvm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pascal",
		URL:          []string{"https://github.com/adspit/pascal"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/segmitos-mitosis-detection",
		URL:          []string{"https://github.com/chaoli977/segmitos_mitosis_detection"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cgan",
		URL:          []string{"https://github.com/divelab/cgan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ukbrest",
		URL:          []string{"https://github.com/hakyimlab/ukbrest"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nclcomparator",
		URL:          []string{"https://github.com/treeslab/nclcomparator"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lrf-dtis",
		URL:          []string{"https://github.com/qust-aibbdrc/lrf-dtis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tron",
		URL:          []string{"https://github.com/davidssmith/tron"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pia",
		URL:          []string{"https://github.com/mpc-bioinformatics/pia"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/epidish",
		URL:          []string{"https://github.com/sjczheng/epidish"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/outcome-prediction-nse",
		URL:          []string{"https://github.com/emilymuller1991/outcome_prediction_nse"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/intlim",
		URL:          []string{"https://github.com/mathelab/intlim"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/flea-pipeline",
		URL:          []string{"https://github.com/veg/flea-pipeline"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/flea-web-app",
		URL:          []string{"https://github.com/veg/flea-web-app"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/erd2",
		URL:          []string{"https://github.com/tyongkim/erd2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/additive-fnnrw",
		URL:          []string{"https://github.com/littleq1991/additive_fnnrw"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/calib",
		URL:          []string{"https://github.com/vpc-ccg/calib"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/svict",
		URL:          []string{"https://github.com/vpc-ccg/svict"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phenomenet-vp",
		URL:          []string{"https://github.com/bio-ontology-research-group/phenomenet-vp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/squire",
		URL:          []string{"https://github.com/wyang17/squire"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/miscoto",
		URL:          []string{"https://github.com/cfrioux/miscoto"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/drugthatgene",
		URL:          []string{"https://github.com/pinellolab/drugthatgene"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mec",
		URL:          []string{"https://github.com/bioinfomaticscsu/mec"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/driverml",
		URL:          []string{"https://github.com/helloyihan/driverml"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/proteomicsbrowser",
		URL:          []string{"https://github.com/peng-gang/proteomicsbrowser"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rse",
		URL:          []string{"https://github.com/ecomol/rse"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/prottrace",
		URL:          []string{"https://github.com/bionf/prottrace"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepfold",
		URL:          []string{"https://github.com/largelymfs/deepfold"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/akt-selective",
		URL:          []string{"https://github.com/undwivedi/akt-selective"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nrlmfb",
		URL:          []string{"https://github.com/akiyamalab/nrlmfb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/emep",
		URL:          []string{"https://github.com/lixt314/emep"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rlscore",
		URL:          []string{"https://github.com/aatapa/rlscore"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/comoto-pasteur-fr",
		URL:          []string{"https://github.com/comoto-pasteur-fr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/betaboost",
		URL:          []string{"https://github.com/boost-r/betaboost"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mcsmrt",
		URL:          []string{"https://github.com/jpearl01/mcsmrt"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/microbiota-resistome",
		URL:          []string{"https://github.com/lpenguin/microbiota-resistome"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/clustermap",
		URL:          []string{"https://github.com/xgaoo/clustermap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pam",
		URL:          []string{"https://github.com/eelhaik/pam"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/superencoder",
		URL:          []string{"https://github.com/hestiri/superencoder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/drugtargetor",
		URL:          []string{"https://github.com/hagax8/drugtargetor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sgtk",
		URL:          []string{"https://github.com/olga24912/sgtk"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pyseer",
		URL:          []string{"https://github.com/mgalardini/pyseer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/psite",
		URL:          []string{"https://github.com/hchyang/psite"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/etoxpred",
		URL:          []string{"https://github.com/pulimeng/etoxpred"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/georgettetanner",
		URL:          []string{"https://github.com/georgettetanner"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dets",
		URL:          []string{"https://github.com/bsml320/dets"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mirtardann",
		URL:          []string{"https://github.com/xuelab/mirtardann"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mcorr",
		URL:          []string{"https://github.com/kussell-lab/mcorr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/intarna",
		URL:          []string{"https://github.com/backofenlab/intarna"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepcrystal",
		URL:          []string{"https://github.com/elbasir/deepcrystal"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bin3c",
		URL:          []string{"https://github.com/cerebis/bin3c"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ogt-prediction",
		URL:          []string{"https://github.com/davidbsauer/ogt_prediction"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lcv",
		URL:          []string{"https://github.com/eyalsim/lcv"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pasnet",
		URL:          []string{"https://github.com/datax-jiehao/pasnet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tpmcalculator",
		URL:          []string{"https://github.com/ncbi/tpmcalculator"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/idhi-mirw",
		URL:          []string{"https://github.com/nwpu-903pr/idhi-mirw"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/slimer",
		URL:          []string{"https://github.com/sysbiochalmers/slimer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gfaviz",
		URL:          []string{"https://github.com/ggonnella/gfaviz"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/breakid",
		URL:          []string{"https://github.com/sinoncology/breakid"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nvr",
		URL:          []string{"https://github.com/kenlaulab/nvr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/svim",
		URL:          []string{"https://github.com/eldariont/svim"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/clumsid",
		URL:          []string{"https://github.com/tdepke/clumsid"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rgaat-v2",
		URL:          []string{"https://github.com/wushyer/rgaat_v2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lordfast",
		URL:          []string{"https://github.com/vpc-ccg/lordfast"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/astrap",
		URL:          []string{"https://github.com/bmilab/astrap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/align-linguistic-alignment",
		URL:          []string{"https://github.com/nickduran/align-linguistic-alignment"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/grimm",
		URL:          []string{"https://github.com/nmdp-bioinformatics/grimm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tumor-clones",
		URL:          []string{"https://github.com/moyanre/tumor_clones"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/orthoscope",
		URL:          []string{"https://github.com/jun-inoue/orthoscope"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gnn",
		URL:          []string{"https://github.com/ibpa/gnn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deep-resp-forest",
		URL:          []string{"https://github.com/ransulab/deep-resp-forest"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deuteros",
		URL:          []string{"https://github.com/andymlau/deuteros"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/aces",
		URL:          []string{"https://github.com/grabherrgroup/aces"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/moff",
		URL:          []string{"https://github.com/compomics/moff"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/barnaba",
		URL:          []string{"https://github.com/srnas/barnaba"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rcppocc",
		URL:          []string{"https://github.com/allanclark/rcppocc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/imagecn",
		URL:          []string{"https://github.com/zhangchenlab/imagecn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mcmc-ce-codes",
		URL:          []string{"https://github.com/shilab2017/mcmc-ce-codes"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gorampage",
		URL:          []string{"https://github.com/brendelgroup/gorampage"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/parentage",
		URL:          []string{"https://github.com/huangkang1987/parentage"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepbinner",
		URL:          []string{"https://github.com/rrwick/deepbinner"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/batmeth2",
		URL:          []string{"https://github.com/guoliangli-hzau/batmeth2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gemprot",
		URL:          []string{"https://github.com/taniacuppens/gemprot"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gencof",
		URL:          []string{"https://github.com/mattczajkowski/gencof"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bioposdep",
		URL:          []string{"https://github.com/datquocnguyen/bioposdep"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/clusterdv",
		URL:          []string{"https://github.com/jcbmarques/clusterdv"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deblender",
		URL:          []string{"https://github.com/kondim1983/deblender"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/smudge",
		URL:          []string{"https://github.com/bio-ontology-research-group/smudge"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dfim",
		URL:          []string{"https://github.com/kundajelab/dfim"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepcoil",
		URL:          []string{"https://github.com/labstructbioinf/deepcoil"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mhfp",
		URL:          []string{"https://github.com/reymond-group/mhfp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/condo",
		URL:          []string{"https://github.com/gicsaw/condo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/genepheno",
		URL:          []string{"https://github.com/bio-ontology-research-group/genepheno"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gopher",
		URL:          []string{"https://github.com/thejacksonlaboratory/gopher"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mirtrace",
		URL:          []string{"https://github.com/friedlanderlab/mirtrace"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cirgo",
		URL:          []string{"https://github.com/irinavkuznetsova/cirgo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/algorithmcomparison",
		URL:          []string{"https://github.com/felipejcolon/algorithmcomparison"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nucleosee",
		URL:          []string{"https://github.com/rodrigosantamaria/nucleosee"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/racat",
		URL:          []string{"https://github.com/ellipfaehlerumcg/racat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/github-bioinformatics",
		URL:          []string{"https://github.com/pamelarussell/github-bioinformatics"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/surrogateminimaldepth",
		URL:          []string{"https://github.com/stephanseifert/surrogateminimaldepth"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/voxelmorph",
		URL:          []string{"https://github.com/voxelmorph/voxelmorph"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/promoterpredict",
		URL:          []string{"https://github.com/promoterpredict"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biokeen",
		URL:          []string{"https://github.com/smartdataanalytics/biokeen"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pykeen",
		URL:          []string{"https://github.com/smartdataanalytics/pykeen"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sumsec",
		URL:          []string{"https://github.com/yosvanylopez/sumsec"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phenogocon",
		URL:          []string{"https://github.com/bio-ontology-research-group/phenogocon"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/text-image-deblur",
		URL:          []string{"https://github.com/shenjianbing/text-image-deblur"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/benchmarking-tsdiscretizations",
		URL:          []string{"https://github.com/veraliconaresearchgroup/benchmarking_tsdiscretizations"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ngmerge",
		URL:          []string{"https://github.com/harvardinformatics/ngmerge"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/psearch",
		URL:          []string{"https://github.com/meddwl/psearch"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tunnex",
		URL:          []string{"https://github.com/prs-group/tunnex"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/libgtftk",
		URL:          []string{"https://github.com/dputhier/libgtftk"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepm6aseq",
		URL:          []string{"https://github.com/rreybeyb/deepm6aseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mdpbiome",
		URL:          []string{"https://github.com/beatrizgj/mdpbiome"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/proteinpocketdetection",
		URL:          []string{"https://github.com/rdzhao/proteinpocketdetection"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bmdexpress-2",
		URL:          []string{"https://github.com/auerbachs/bmdexpress-2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/assemblosis",
		URL:          []string{"https://github.com/vetscience/assemblosis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/typeloader",
		URL:          []string{"https://github.com/dkms-lsl/typeloader"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/calour",
		URL:          []string{"https://github.com/biocore/calour"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/caars",
		URL:          []string{"https://github.com/carinerey/caars"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rsfp",
		URL:          []string{"https://github.com/ek1203/rsfp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pirnn",
		URL:          []string{"https://github.com/bioinfolabmu/pirnn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/papago",
		URL:          []string{"https://github.com/adalca/papago"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bionet-mining",
		URL:          []string{"https://github.com/fabiennel/bionet-mining"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/himc",
		URL:          []string{"https://github.com/vserch/himc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/uncurl-python",
		URL:          []string{"https://github.com/yjzhang/uncurl_python"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/strelka",
		URL:          []string{"https://github.com/illumina/strelka"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/comsa",
		URL:          []string{"https://github.com/refresh-bio/comsa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pathemb",
		URL:          []string{"https://github.com/zhangjiaobxy/pathemb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/atac-pipe",
		URL:          []string{"https://github.com/qukunlab/atac-pipe"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/docker-galaxy-hicexplorer",
		URL:          []string{"https://github.com/deeptools/docker-galaxy-hicexplorer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/care-rcortex",
		URL:          []string{"https://github.com/fannygrosselin/care-rcortex"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/minimap2",
		URL:          []string{"https://github.com/lh3/minimap2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/opal",
		URL:          []string{"https://github.com/yunwilliamyu/opal"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pupil-size",
		URL:          []string{"https://github.com/elios-s/pupil-size"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mcindoor20000",
		URL:          []string{"https://github.com/bircatmcri/mcindoor20000"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rnapuzzler",
		URL:          []string{"https://github.com/dwiegreffe/rnapuzzler"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/imapsplice",
		URL:          []string{"https://github.com/liubioinfo/imapsplice"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phewas",
		URL:          []string{"https://github.com/phewas/phewas"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ehr",
		URL:          []string{"https://github.com/choileena/ehr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cmfsm",
		URL:          []string{"https://github.com/ysycloud/cmfsm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/virulign",
		URL:          []string{"https://github.com/rega-cev/virulign"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/straingems",
		URL:          []string{"https://github.com/hust-ningkang-lab/straingems"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/netsmooth",
		URL:          []string{"https://github.com/bimsbbioinfo/netsmooth"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/psucce",
		URL:          []string{"https://github.com/ningq669/psucce"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/micmic",
		URL:          []string{"https://github.com/zhangjlab/micmic"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pdb2entropy",
		URL:          []string{"https://github.com/federico-fogolari/pdb2entropy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pdb2trent",
		URL:          []string{"https://github.com/federico-fogolari/pdb2trent"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nifpred",
		URL:          []string{"https://github.com/prabinameher/nifpred"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biocppi-extraction",
		URL:          []string{"https://github.com/bionlproc/biocppi_extraction"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/disease-similarity-fusion",
		URL:          []string{"https://github.com/e-oerton/disease-similarity-fusion"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/charger",
		URL:          []string{"https://github.com/ding-lab/charger"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bilouvain",
		URL:          []string{"https://github.com/paolapesantez/bilouvain"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/onto2vec",
		URL:          []string{"https://github.com/bio-ontology-research-group/onto2vec"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lfc",
		URL:          []string{"https://github.com/erhard-lab/lfc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/matryoshka",
		URL:          []string{"https://github.com/combine-lab/matryoshka"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/geck",
		URL:          []string{"https://github.com/sbg/geck"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lrdkt",
		URL:          []string{"https://github.com/shaohuilin/lrdkt"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/stretch",
		URL:          []string{"https://github.com/oshlack/stretch"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/q2mm",
		URL:          []string{"https://github.com/q2mm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/undwive",
		URL:          []string{"https://github.com/undwive"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/s-predixcan",
		URL:          []string{"https://github.com/hakyimlab/s-predixcan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metaxcan",
		URL:          []string{"https://github.com/hakyimlab/metaxcan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fastspar",
		URL:          []string{"https://github.com/scwatts/fastspar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gapreduce",
		URL:          []string{"https://github.com/bioinfomaticscsu/gapreduce"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/genefuse",
		URL:          []string{"https://github.com/opengene/genefuse"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/aalto-ics-kepaco",
		URL:          []string{"https://github.com/aalto-ics-kepaco"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/drimpute",
		URL:          []string{"https://github.com/gongx030/drimpute"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/z-transform-method",
		URL:          []string{"https://github.com/spawaskar-cora/z-transform-method"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cartoon-network",
		URL:          []string{"https://github.com/rcalinjageman/cartoon_network"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/stk",
		URL:          []string{"https://github.com/supramolecular-toolkit/stk"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cobratoolbox",
		URL:          []string{"https://github.com/opencobra/cobratoolbox"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cadrres",
		URL:          []string{"https://github.com/csb5/cadrres"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/omic-features-successful-targets",
		URL:          []string{"https://github.com/arouillard/omic-features-successful-targets"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pywindow",
		URL:          []string{"https://github.com/jelfsmaterialsgroup/pywindow"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mipup",
		URL:          []string{"https://github.com/zhero9/mipup"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/simd",
		URL:          []string{"https://github.com/focuspaka/simd"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dnbfa",
		URL:          []string{"https://github.com/siamakz/dnbfa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biostructmap",
		URL:          []string{"https://github.com/andrewguy/biostructmap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ideeps",
		URL:          []string{"https://github.com/xypan1232/ideeps"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gooogle",
		URL:          []string{"https://github.com/himelmallick/gooogle"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dnabinding",
		URL:          []string{"https://github.com/srautonu/dnabinding"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hnmf",
		URL:          []string{"https://github.com/yuanluo/hnmf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/joker",
		URL:          []string{"https://github.com/williamfp7/joker"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/chimericognizer",
		URL:          []string{"https://github.com/ucrbioinfo/chimericognizer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pvalue-weighting-matlab",
		URL:          []string{"https://github.com/dobriban/pvalue_weighting_matlab"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/qtc",
		URL:          []string{"https://github.com/jssong-lab/qtc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bruno",
		URL:          []string{"https://github.com/parbliss/bruno"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pytag",
		URL:          []string{"https://github.com/kociorges/pytag"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/panisa",
		URL:          []string{"https://github.com/bvalot/panisa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bacteriamslf",
		URL:          []string{"https://github.com/dipcarbon/bacteriamslf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cryfa",
		URL:          []string{"https://github.com/pratas/cryfa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/agotron-detector",
		URL:          []string{"https://github.com/ncrnalab/agotron_detector"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/whatshap",
		URL:          []string{"https://github.com/whatshap/whatshap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/negbio",
		URL:          []string{"https://github.com/ncbi-nlp/negbio"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/smilesvecproteinrepresentation",
		URL:          []string{"https://github.com/hkmztrk/smilesvecproteinrepresentation"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metqy",
		URL:          []string{"https://github.com/oss-lab/metqy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/framework",
		URL:          []string{"https://github.com/asaim/framework"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/litemol",
		URL:          []string{"https://github.com/dsehnal/litemol"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dogfinder",
		URL:          []string{"https://github.com/shalgilab/dogfinder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nb-distribution",
		URL:          []string{"https://github.com/czc/nb_distribution"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/limbr",
		URL:          []string{"https://github.com/aleccrowell/limbr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pdv",
		URL:          []string{"https://github.com/wenbostar/pdv"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mirge",
		URL:          []string{"https://github.com/mhalushka/mirge"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/famplex",
		URL:          []string{"https://github.com/sorgerlab/famplex"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/btw",
		URL:          []string{"https://github.com/vpylro/btw"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biotoolscompose",
		URL:          []string{"https://github.com/bio-tools/biotoolscompose"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/unity",
		URL:          []string{"https://github.com/bogdanlab/unity"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hpo2go",
		URL:          []string{"https://github.com/cansyl/hpo2go"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nextstrain",
		URL:          []string{"https://github.com/nextstrain"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/xispec",
		URL:          []string{"https://github.com/rappsilber-laboratory/xispec"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mass-simulator",
		URL:          []string{"https://github.com/pcdslab/mass-simulator"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/image-quality-evaluation",
		URL:          []string{"https://github.com/ezimic/image-quality-evaluation"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/glutton",
		URL:          []string{"https://github.com/yeeho/glutton"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tapas",
		URL:          []string{"https://github.com/arefeen/tapas"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mob-suite",
		URL:          []string{"https://github.com/phac-nml/mob-suite"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/springsalad",
		URL:          []string{"https://github.com/jmasison/springsalad"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rnadetect",
		URL:          []string{"https://github.com/bjyoontamu/rnadetect"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/divpop",
		URL:          []string{"https://github.com/wheelerlab/divpop"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/stereoseg",
		URL:          []string{"https://github.com/shenjianbing/stereoseg"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dynamicdbg",
		URL:          []string{"https://github.com/csirac/dynamicdbg"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/minion-qc",
		URL:          []string{"https://github.com/roblanf/minion_qc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gsea-incontext",
		URL:          []string{"https://github.com/costellolab/gsea-incontext"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mode-task",
		URL:          []string{"https://github.com/rubi-za/mode-task"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pisces",
		URL:          []string{"https://github.com/illumina/pisces"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/snp-sites",
		URL:          []string{"https://github.com/sanger-pathogens/snp-sites"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/genomemining",
		URL:          []string{"https://github.com/batlabucd/genomemining"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/integreat",
		URL:          []string{"https://github.com/faryabib/integreat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/presm",
		URL:          []string{"https://github.com/precisionomics/presm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/submodulartrack",
		URL:          []string{"https://github.com/shenjianbing/submodulartrack"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/super-i2b2",
		URL:          []string{"https://github.com/wcmc-research-informatics/super-i2b2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ravensoftware",
		URL:          []string{"https://github.com/jsgerke/ravensoftware"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/submatrixselectionsvd.jl",
		URL:          []string{"https://github.com/rasmushenningsson/submatrixselectionsvd.jl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nbzimm",
		URL:          []string{"https://github.com//nyiuab//nbzimm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/goatools",
		URL:          []string{"https://github.com/tanghaibao/goatools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ewas-twin-simulation",
		URL:          []string{"https://github.com/zickyls/ewas-twin-simulation"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lpg",
		URL:          []string{"https://github.com/shufeyangyi2015310117/lpg"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pcnem",
		URL:          []string{"https://github.com/cbg-ethz/pcnem"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/13check-rna",
		URL:          []string{"https://github.com/bios-imasl/13check_rna"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/perfect",
		URL:          []string{"https://github.com/katiasmirn/perfect"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/musica",
		URL:          []string{"https://github.com/marcos-diazg/musica"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sopang",
		URL:          []string{"https://github.com/mralexsee/sopang"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/physiboss",
		URL:          []string{"https://github.com/sysbio-curie/physiboss"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/verse",
		URL:          []string{"https://github.com/spandanagella/verse"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/regent",
		URL:          []string{"https://github.com/wmmthu/regent"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/candi",
		URL:          []string{"https://github.com/mbadge/candi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/novocaller",
		URL:          []string{"https://github.com/bgm-cwg/novocaller"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/viromic-diversity",
		URL:          []string{"https://github.com/jorgevazcast/viromic-diversity"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/treesgs",
		URL:          []string{"https://github.com/kuanglab/treesgs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/knotty",
		URL:          []string{"https://github.com/hosnajabbari/knotty"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/index-app-public",
		URL:          []string{"https://github.com/drdanl/index-app-public"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/terraphast",
		URL:          []string{"https://github.com/terraphast"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rm-seq",
		URL:          []string{"https://github.com/rguerillot/rm-seq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hogmmnc",
		URL:          []string{"https://github.com/scutbioinformatics/hogmmnc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/epsdc",
		URL:          []string{"https://github.com/kdding/epsdc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/icnv",
		URL:          []string{"https://github.com/zhouzilu/icnv"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/classifier-selection-code",
		URL:          []string{"https://github.com/timodeist/classifier_selection_code"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/modentify",
		URL:          []string{"https://github.com/krumsieklab/modentify"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/camstyle",
		URL:          []string{"https://github.com/zhunzhong07/camstyle"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cgvr",
		URL:          []string{"https://github.com/xbjin/cgvr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gevitanalysisrelease",
		URL:          []string{"https://github.com/amcrisan/gevitanalysisrelease"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dfm",
		URL:          []string{"https://github.com/lingxiao-he/dfm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/qiaseq-dna",
		URL:          []string{"https://github.com/qiaseq/qiaseq-dna"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/effusion",
		URL:          []string{"https://github.com/babbittlab/effusion"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/immunedb",
		URL:          []string{"https://github.com/arosenfeld/immunedb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deephint",
		URL:          []string{"https://github.com/nonnerdling/deephint"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/flnadd",
		URL:          []string{"https://github.com/scalderazzo/flnadd"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bth",
		URL:          []string{"https://github.com/b2du/bth"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepmirtar-sda",
		URL:          []string{"https://github.com/bjoux2/deepmirtar_sda"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/semgen",
		URL:          []string{"https://github.com/sembioprocess/semgen"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/browniealigner",
		URL:          []string{"https://github.com/biointec/browniealigner"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/promotepredictor",
		URL:          []string{"https://github.com/maqin2001/promotepredictor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sonics",
		URL:          []string{"https://github.com/kzkedzierska/sonics"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/conspecifix",
		URL:          []string{"https://github.com/bobay-ochman/conspecifix"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/timing2",
		URL:          []string{"https://github.com/roysamlab/timing2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mdnnmd",
		URL:          []string{"https://github.com/ustc-hilab/mdnnmd"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dlbi",
		URL:          []string{"https://github.com/lykaust15/dlbi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hemodonacion",
		URL:          []string{"https://github.com/fanavarro/hemodonacion"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/emerald",
		URL:          []string{"https://github.com/statgen/emerald"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/equivalent-junctions",
		URL:          []string{"https://github.com/salzmanlab/equivalent-junctions"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/psepssm-dcca-lfda",
		URL:          []string{"https://github.com/qust-bsbrc/psepssm-dcca-lfda"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/co-fuse",
		URL:          []string{"https://github.com/sakrapee/co-fuse"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/model-fitting-cbs",
		URL:          []string{"https://github.com/ruwant/model-fitting-cbs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/enzynet",
		URL:          []string{"https://github.com/shervinea/enzynet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tepic",
		URL:          []string{"https://github.com/schulzlab/tepic"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/discoversl",
		URL:          []string{"https://github.com/shaoli86/discoversl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/triomdr",
		URL:          []string{"https://github.com/triomdr/triomdr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/treeshrink",
		URL:          []string{"https://github.com/uym2/treeshrink"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/odin",
		URL:          []string{"https://github.com/hormozdiarilab/odin"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/couplenet",
		URL:          []string{"https://github.com/tshizys/couplenet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/matataki",
		URL:          []string{"https://github.com/informationsea/matataki"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rhe-reg",
		URL:          []string{"https://github.com/sriramlab/rhe-reg"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tivan",
		URL:          []string{"https://github.com/lichen-lab/tivan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/skesa",
		URL:          []string{"https://github.com/ncbi/skesa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tcrclusteringpaper",
		URL:          []string{"https://github.com/pmeysman/tcrclusteringpaper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/esprit2",
		URL:          []string{"https://github.com/dessimozlab/esprit2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/conics",
		URL:          []string{"https://github.com/diazlab/conics"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/locustreeinference",
		URL:          []string{"https://github.com/mciach/locustreeinference"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gmxapi",
		URL:          []string{"https://github.com/kassonlab/gmxapi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mypmfs",
		URL:          []string{"https://github.com/bibip-impmc/mypmfs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepsnr",
		URL:          []string{"https://github.com/sirajulsalekin/deepsnr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lvq-knn",
		URL:          []string{"https://github.com/ab1989/lvq-knn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/anonimme",
		URL:          []string{"https://github.com/bristena-op/anonimme"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sv-plauditdemonstration",
		URL:          []string{"https://github.com/jbelyeu/sv-plauditdemonstration"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pyabc",
		URL:          []string{"https://github.com/icb-dcm/pyabc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gmalign",
		URL:          []string{"https://github.com/yzlwhu/gmalign"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/anticancer-peptides-review",
		URL:          []string{"https://github.com/shoombuatong2527/anticancer-peptides-review"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/virtual-reality-chemical-space",
		URL:          []string{"https://github.com/reymond-group/virtual-reality-chemical-space"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mrbait",
		URL:          []string{"https://github.com/tkchafin/mrbait"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gsbc",
		URL:          []string{"https://github.com/mehmetgonen/gsbc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/3dpatch",
		URL:          []string{"https://github.com/davidjakubec/3dpatch"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/flexidot",
		URL:          []string{"https://github.com/molbio-dresden/flexidot"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/insilicoseq",
		URL:          []string{"https://github.com/hadrieng/insilicoseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rmtl",
		URL:          []string{"https://github.com/transbiozi/rmtl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metagrn",
		URL:          []string{"https://github.com/zheng-lab/metagrn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/3dface",
		URL:          []string{"https://github.com/juyong/3dface"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/detectron",
		URL:          []string{"https://github.com/facebookresearch/detectron"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nested-sampling",
		URL:          []string{"https://github.com/beast2-dev/nested-sampling"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pea",
		URL:          []string{"https://github.com/cma2015/pea"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/isofishr",
		URL:          []string{"https://github.com/maltewillmes/isofishr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/umons-taichi",
		URL:          []string{"https://github.com/numediart/umons-taichi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/safeclustering",
		URL:          []string{"https://github.com/yycunc/safeclustering"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/consensus",
		URL:          []string{"https://github.com/timpeters82/consensus"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/iswathx",
		URL:          []string{"https://github.com/znoor/iswathx"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/imtrbm",
		URL:          []string{"https://github.com/liuying201705/imtrbm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/crispr",
		URL:          []string{"https://github.com/alexzsx/crispr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biostructurem",
		URL:          []string{"https://github.com/yuan-yu/biostructurem"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/graph-annotation",
		URL:          []string{"https://github.com/ratschlab/graph_annotation"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/parkour",
		URL:          []string{"https://github.com/maxplanck-ie/parkour"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metagenomic-benchmark",
		URL:          []string{"https://github.com/ales-ibt/metagenomic-benchmark"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bipspi",
		URL:          []string{"https://github.com/bioinsilico/bipspi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/epiprofile2.0-family",
		URL:          []string{"https://github.com/zfyuan/epiprofile2.0_family"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/microbiome-helper",
		URL:          []string{"https://github.com/mlangill/microbiome_helper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pantools",
		URL:          []string{"https://github.com/sheikhizadeh/pantools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biocreativevi-bioid-assignment",
		URL:          []string{"https://github.com/turkunlp/biocreativevi_bioid_assignment"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tobmi",
		URL:          []string{"https://github.com/xuesidong/tobmi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/visioprot-ms",
		URL:          []string{"https://github.com/mlocardpaulet/visioprot-ms"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gpca",
		URL:          []string{"https://github.com/yangzi4/gpca"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/neuromorphovis",
		URL:          []string{"https://github.com/bluebrain/neuromorphovis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mea",
		URL:          []string{"https://github.com/julienrichardalbert/mea"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/coptea",
		URL:          []string{"https://github.com/wulingyun/coptea"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/trca-ssvep",
		URL:          []string{"https://github.com/mnakanishi/trca-ssvep"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cellmissy",
		URL:          []string{"https://github.com/compomics/cellmissy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/haoran2014",
		URL:          []string{"https://github.com/haoran2014"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/crisprmatch",
		URL:          []string{"https://github.com/zhangtaolab/crisprmatch"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dislncrf",
		URL:          []string{"https://github.com/xypan1232/dislncrf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/theprecisionlasso",
		URL:          []string{"https://github.com/haohanwang/theprecisionlasso"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/radaa",
		URL:          []string{"https://github.com/sciseim/radaa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gm-cml",
		URL:          []string{"https://github.com/wangyida/gm-cml"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lncmirsrn",
		URL:          []string{"https://github.com/zhangjunpeng411/lncmirsrn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/shiver",
		URL:          []string{"https://github.com/chrishiv/shiver"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sm-engine",
		URL:          []string{"https://github.com/anasilviacs/sm-engine"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rescore-metabolites",
		URL:          []string{"https://github.com/anasilviacs/rescore-metabolites"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bactdating",
		URL:          []string{"https://github.com/xavierdidelot/bactdating"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pangenomepipeline",
		URL:          []string{"https://github.com/jcventerinstitute/pangenomepipeline"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/novo-stitch",
		URL:          []string{"https://github.com/ucrbioinfo/novo_stitch"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/simextargid",
		URL:          []string{"https://github.com/josielhayes/simextargid"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scrnaseq-clustering-comparison",
		URL:          []string{"https://github.com/markrobinsonuzh/scrnaseq_clustering_comparison"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fluidigm2purc",
		URL:          []string{"https://github.com/pblischak/fluidigm2purc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mendelprob",
		URL:          []string{"https://github.com/statgenetics/mendelprob"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tsvdb",
		URL:          []string{"https://github.com/wenjie1991/tsvdb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dinamo",
		URL:          []string{"https://github.com/bonsai-team/dinamo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/seroba",
		URL:          []string{"https://github.com/sanger-pathogens/seroba"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/inpactor",
		URL:          []string{"https://github.com/simonorozcoarias/inpactor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/groot",
		URL:          []string{"https://github.com/will-rowe/groot"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ecpred",
		URL:          []string{"https://github.com/cansyl/ecpred"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/spike-compression-autoencoder",
		URL:          []string{"https://github.com/tong-wu-umn/spike-compression-autoencoder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dinar",
		URL:          []string{"https://github.com/nib-si/dinar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/poptrt",
		URL:          []string{"https://github.com/eden-kramer-lab/poptrt"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bindash",
		URL:          []string{"https://github.com/zhaoxiaofei/bindash"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/grpslopemc",
		URL:          []string{"https://github.com/agisga/grpslopemc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vbt-trioanalysis",
		URL:          []string{"https://github.com/sbg/vbt-trioanalysis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ac-diamond",
		URL:          []string{"https://github.com/maihj/ac-diamond"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deeppath",
		URL:          []string{"https://github.com/ncoudray/deeppath"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/creamino",
		URL:          []string{"https://github.com/arcesunibo/creamino"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/posdatools",
		URL:          []string{"https://github.com/uams-dbmi/posdatools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/macau",
		URL:          []string{"https://github.com/jaak-s/macau"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bayesiandatafusion.jl",
		URL:          []string{"https://github.com/jaak-s/bayesiandatafusion.jl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/eqtloverlapper",
		URL:          []string{"https://github.com/imb-computational-genomics-lab/eqtloverlapper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/neuroner",
		URL:          []string{"https://github.com/franck-dernoncourt/neuroner"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/transfer-learning-bner-bioinformatics-2018",
		URL:          []string{"https://github.com/baderlab/transfer-learning-bner-bioinformatics-2018"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/peakc",
		URL:          []string{"https://github.com/dewitlab/peakc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/trainbenchmark",
		URL:          []string{"https://github.com/ftsrg/trainbenchmark"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bootejtk",
		URL:          []string{"https://github.com/alanlhutchison/bootejtk"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rehunt",
		URL:          []string{"https://github.com/yuhuei/rehunt"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bowtie-scaling",
		URL:          []string{"https://github.com/benlangmead/bowtie-scaling"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/brainimager",
		URL:          []string{"https://github.com/saralinker/brainimager"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/catfish",
		URL:          []string{"https://github.com/kingsford-group/catfish"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pysfd",
		URL:          []string{"https://github.com/markovmodel/pysfd"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kmer-db",
		URL:          []string{"https://github.com/refresh-bio/kmer-db"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/slide",
		URL:          []string{"https://github.com/soumitag/slide"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ngl",
		URL:          []string{"https://github.com/arose/ngl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mmtf-javascript",
		URL:          []string{"https://github.com/rcsb/mmtf-javascript"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/polyclustr",
		URL:          []string{"https://github.com/syspremed/polyclustr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/para-dpmm",
		URL:          []string{"https://github.com/tiehangd/para_dpmm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gpseudorank",
		URL:          []string{"https://github.com/magstra/gpseudorank"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gatk",
		URL:          []string{"https://github.com/broadinstitute/gatk"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lancet",
		URL:          []string{"https://github.com/nygenome/lancet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/glyxtoolms",
		URL:          []string{"https://github.com/glyxera/glyxtoolms"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sixparam",
		URL:          []string{"https://github.com/ebraun68/sixparam"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gwc",
		URL:          []string{"https://github.com/yuzhounh/gwc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/regcombims",
		URL:          []string{"https://github.com/nhpatterson/regcombims"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bicolor",
		URL:          []string{"https://github.com/yuansliu/bicolor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nbss",
		URL:          []string{"https://github.com/wzhang1984/nbss"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/csdm",
		URL:          []string{"https://github.com/fengli28/csdm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/artgan",
		URL:          []string{"https://github.com/cs-chan/artgan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/eeg-spindles",
		URL:          []string{"https://github.com/vislab/eeg-spindles"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/justorthologs",
		URL:          []string{"https://github.com/ridgelab/justorthologs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/comparisonfastafiles",
		URL:          []string{"https://github.com/ridgelab/justorthologs/comparisonfastafiles"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/transcriptclean",
		URL:          []string{"https://github.com/dewyman/transcriptclean"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/multimodal-network-diffusion",
		URL:          []string{"https://github.com/lichtargelab/multimodal-network-diffusion"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bpp",
		URL:          []string{"https://github.com/bpp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metaomics",
		URL:          []string{"https://github.com/metaomics/metaomics"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pyphi",
		URL:          []string{"https://github.com/wmayner/pyphi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/q2-feature-classifier",
		URL:          []string{"https://github.com/qiime2/q2-feature-classifier"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tax-credit-data",
		URL:          []string{"https://github.com/caporaso-lab/tax-credit-data"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lzw-kernel",
		URL:          []string{"https://github.com/kfattila/lzw-kernel"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/att-chemprot",
		URL:          []string{"https://github.com/ohnlp/att-chemprot"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/emblmygff3",
		URL:          []string{"https://github.com/nbisweden/emblmygff3"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/treegrafter",
		URL:          []string{"https://github.com/pantherdb/treegrafter"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ctcf-mp",
		URL:          []string{"https://github.com/ma-compbio/ctcf-mp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phylorad",
		URL:          []string{"https://github.com/fanhuan/phylorad"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/personalized-regression",
		URL:          []string{"https://github.com/blengerich/personalized_regression"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/arcas",
		URL:          []string{"https://github.com/uva-peirce-cottler-lab/arcas"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ampumi",
		URL:          []string{"https://github.com/pinellolab/ampumi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metapoap",
		URL:          []string{"https://github.com/lmward/metapoap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nrl2drp",
		URL:          []string{"https://github.com/ustc-hilab/nrl2drp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/motif-scraper",
		URL:          []string{"https://github.com/robersonlab/motif_scraper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gatc",
		URL:          []string{"https://github.com/udem-lbit/gatc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nmflibrary",
		URL:          []string{"https://github.com/ramkikannan/nmflibrary"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/chiimp",
		URL:          []string{"https://github.com/shawhahnlab/chiimp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dot2dot",
		URL:          []string{"https://github.com/gege7177/dot2dot"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/is-cellr",
		URL:          []string{"https://github.com/immcore/is-cellr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hyppo",
		URL:          []string{"https://github.com/manuellafond/hyppo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/confit",
		URL:          []string{"https://github.com/lgai/confit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/findgco",
		URL:          []string{"https://github.com/tongchf/findgco"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/genotype-corrector",
		URL:          []string{"https://github.com/freemao/genotype-corrector"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/presogenesis",
		URL:          []string{"https://github.com/mrb20045/presogenesis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/grouper",
		URL:          []string{"https://github.com/combine-lab/grouper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kanizsa-prime",
		URL:          []string{"https://github.com/zenben/kanizsa_prime"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sppider",
		URL:          []string{"https://github.com/glbrc/sppider"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/microbial-conditions-ontology",
		URL:          []string{"https://github.com/microbial-conditions-ontology"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mgpfusion",
		URL:          []string{"https://github.com/emmijokinen/mgpfusion"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/recursive-chemprot",
		URL:          []string{"https://github.com/arwhirang/recursive_chemprot"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sgspls",
		URL:          []string{"https://github.com/matt-sutton/sgspls"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vigen",
		URL:          []string{"https://github.com/icbi/vigen"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/localtadsim",
		URL:          []string{"https://github.com/kingsford-group/localtadsim"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nvenn",
		URL:          []string{"https://github.com/vqf/nvenn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sarw-lnkf",
		URL:          []string{"https://github.com/luciano-censoni/sarw-lnkf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fgcs17",
		URL:          []string{"https://github.com/sguera/fgcs17"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/repeatcraftp",
		URL:          []string{"https://github.com/niccw/repeatcraftp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sereega",
		URL:          []string{"https://github.com/lrkrol/sereega"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/aptablocks",
		URL:          []string{"https://github.com/wyjhxq/aptablocks"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/evcouplings",
		URL:          []string{"https://github.com/debbiemarkslab/evcouplings"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/marvel",
		URL:          []string{"https://github.com/laboratoriobioinformatica/marvel"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/prequal",
		URL:          []string{"https://github.com/simonwhelan/prequal"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vargeno",
		URL:          []string{"https://github.com/medvedevgroup/vargeno"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bindpredict",
		URL:          []string{"https://github.com/rostlab/bindpredict"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/epicom",
		URL:          []string{"https://github.com/cthoyt/epicom"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lll-fva",
		URL:          []string{"https://github.com/maranasgroup/lll-fva"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/copmem",
		URL:          []string{"https://github.com/wbieniec/copmem"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metaboanalystr",
		URL:          []string{"https://github.com/xia-lab/metaboanalystr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/knowledge-elicitation-for-precision-medicine",
		URL:          []string{"https://github.com/aaltopml/knowledge-elicitation-for-precision-medicine"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepsequence",
		URL:          []string{"https://github.com/debbiemarkslab/deepsequence"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gimpute",
		URL:          []string{"https://github.com/transbiozi/gimpute"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/saturn",
		URL:          []string{"https://github.com/ddamerell53/saturn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/counting-consistent-sub-dag",
		URL:          []string{"https://github.com/shawn-peng/counting-consistent-sub-dag"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ncddetect2",
		URL:          []string{"https://github.com/tobiasmadsen/ncddetect2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/piblup",
		URL:          []string{"https://github.com/huiminkang/piblup"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tusv",
		URL:          []string{"https://github.com/jaebird123/tusv"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/maize-tissue-specific-grn",
		URL:          []string{"https://github.com/timedreamer/maize_tissue-specific_grn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rwr-mh",
		URL:          []string{"https://github.com/alberto-valdeolivas/rwr-mh"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/meffil",
		URL:          []string{"https://github.com/perishky/meffil"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/snpenrichr",
		URL:          []string{"https://github.com/kartiek/snpenrichr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deep-plant",
		URL:          []string{"https://github.com/cs-chan/deep-plant"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/s3m",
		URL:          []string{"https://github.com/borgwardtlab/s3m"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tksamc",
		URL:          []string{"https://github.com/contessoto/tksamc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/irspot-sf",
		URL:          []string{"https://github.com/abdlmaruf/irspot-sf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/icb-docker",
		URL:          []string{"https://github.com/bordin89/icb_docker"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/coverview",
		URL:          []string{"https://github.com/rahmanteamdevelopment/coverview"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/chexmix",
		URL:          []string{"https://github.com/seqcode/chexmix"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mitophast",
		URL:          []string{"https://github.com/mht85/mitophast"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/trac",
		URL:          []string{"https://github.com/nauenlab/trac"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cosr",
		URL:          []string{"https://github.com/xiehq/cosr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/slin",
		URL:          []string{"https://github.com/ghoresh11/slin"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ltr-retriever",
		URL:          []string{"https://github.com/oushujun/ltr_retriever"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/designmpra",
		URL:          []string{"https://github.com/andrewghazi/designmpra"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mpradesigntools",
		URL:          []string{"https://github.com/andrewghazi/mpradesigntools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/jiyuanhu",
		URL:          []string{"https://github.com/jiyuanhu"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/uqsa",
		URL:          []string{"https://github.com/alexjau/uqsa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ranger",
		URL:          []string{"https://github.com/imbs-hl/ranger"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cmap",
		URL:          []string{"https://github.com/cmap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/canalizingpower",
		URL:          []string{"https://github.com/eunjikim-angie/canalizingpower"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/isotree",
		URL:          []string{"https://github.com/jane110111107/isotree"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/yamda",
		URL:          []string{"https://github.com/daquang/yamda"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cmb-labeler",
		URL:          []string{"https://github.com/lupolab-ucsf/cmb_labeler"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepcropping",
		URL:          []string{"https://github.com/shenjianbing/deepcropping"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/crumble",
		URL:          []string{"https://github.com/jkbonfield/crumble"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pycotools",
		URL:          []string{"https://github.com/ciaranwelsh/pycotools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mmpdb",
		URL:          []string{"https://github.com/rdkit/mmpdb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mad-hype",
		URL:          []string{"https://github.com/birnbaumlab/mad-hype"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/likelihoodfreephylogenetics",
		URL:          []string{"https://github.com/simonll/likelihoodfreephylogenetics"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/alifreefold",
		URL:          []string{"https://github.com/udes-cobius/alifreefold"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rifraf.jl",
		URL:          []string{"https://github.com/murrellgroup/rifraf.jl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cath-tools",
		URL:          []string{"https://github.com/uclorengogroup/cath-tools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scop",
		URL:          []string{"https://github.com/bioinfomaticscsu/scop"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ribowaltz",
		URL:          []string{"https://github.com/labtranslationalarchitectomics/ribowaltz"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dcars",
		URL:          []string{"https://github.com/shazanfar/dcars"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dgenies",
		URL:          []string{"https://github.com/genotoul-bioinfo/dgenies"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tecandidates",
		URL:          []string{"https://github.com/tecandidates/tecandidates"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/propip",
		URL:          []string{"https://github.com/acg-team/propip"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/checkmyindex",
		URL:          []string{"https://github.com/pf2-pasteur-fr/checkmyindex"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ngspanpipe",
		URL:          []string{"https://github.com/biomedinformatics/ngspanpipe"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dmcm",
		URL:          []string{"https://github.com/xinguolu/dmcm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hg-color",
		URL:          []string{"https://github.com/morispi/hg-color"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/networkannotation",
		URL:          []string{"https://github.com/spatkar94/networkannotation"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cumulative-fourier-power-spectrum",
		URL:          []string{"https://github.com/yaulabtsinghua/cumulative-fourier-power-spectrum"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepnf",
		URL:          []string{"https://github.com/vgligorijevic/deepnf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/racipe-1.0",
		URL:          []string{"https://github.com/simonhb1990/racipe-1.0"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tisan",
		URL:          []string{"https://github.com/kevinvervier/tisan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/xenofilter",
		URL:          []string{"https://github.com/peeperlab/xenofilter"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/polyqtl",
		URL:          []string{"https://github.com/jxzb1988/polyqtl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/reverse-engineering-bc-grns",
		URL:          []string{"https://github.com/eleonoralusito/reverse_engineering_bc_grns"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pufferfish",
		URL:          []string{"https://github.com/combine-lab/pufferfish"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hypergate",
		URL:          []string{"https://github.com/ebecht/hypergate"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cobralab-atlases",
		URL:          []string{"https://github.com/cobralab/atlases"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/genomeflow",
		URL:          []string{"https://github.com/jianlin-cheng/genomeflow"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bc6pm-hrnn",
		URL:          []string{"https://github.com/afergadis/bc6pm-hrnn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/s2b",
		URL:          []string{"https://github.com/frpinto/s2b"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/spot-a-neonatalrabbit",
		URL:          []string{"https://github.com/gift-surg/spot-a-neonatalrabbit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/islandpath",
		URL:          []string{"https://github.com/brinkmanlab/islandpath"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sskat",
		URL:          []string{"https://github.com/jchen1981/sskat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/happytools",
		URL:          []string{"https://github.com/tarskin/happytools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gedfn",
		URL:          []string{"https://github.com/yunchuankong/gedfn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/glycresoft",
		URL:          []string{"https://github.com/bostonuniversitycbms/glycresoft"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/simpel",
		URL:          []string{"https://github.com/precisionomics/simpel"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fbm",
		URL:          []string{"https://github.com/chpgenetics/fbm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/epem-gmm",
		URL:          []string{"https://github.com/darlenelu72/epem-gmm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/longo",
		URL:          []string{"https://github.com/biohpc/longo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/macarthur-lab/clinvar",
		URL:          []string{"https://github.com/macarthur-lab/clinvar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/quagmir",
		URL:          []string{"https://github.com/gu-lab-rbl-nci/quagmir"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hmc-grt",
		URL:          []string{"https://github.com/ubioptronics/hmc-grt"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ontologyevaluate",
		URL:          []string{"https://github.com/anonymousresearcher1/ontologyevaluate"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/arl-eegmodels",
		URL:          []string{"https://github.com/vlawhern/arl-eegmodels"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pingpongpro",
		URL:          []string{"https://github.com/suhrig/pingpongpro"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/miphy",
		URL:          []string{"https://github.com/dave-the-scientist/miphy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metawrap",
		URL:          []string{"https://github.com/bxlab/metawrap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deconvolution-of-essential-gene-signitures",
		URL:          []string{"https://github.com/jmjung83/deconvolution_of_essential_gene_signitures"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sputnik",
		URL:          []string{"https://github.com/paoloinglese/sputnik"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/saarclust",
		URL:          []string{"https://github.com/daewoooo/saarclust"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ssp-lsdr",
		URL:          []string{"https://github.com/stamakro/ssp-lsdr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/marsi",
		URL:          []string{"https://github.com/biosustain/marsi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lncadeep",
		URL:          []string{"https://github.com/cyang235/lncadeep"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/asgart",
		URL:          []string{"https://github.com/delehef/asgart"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/checkmyblob",
		URL:          []string{"https://github.com/dabrze/checkmyblob"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tciapathfindersignificance",
		URL:          []string{"https://github.com/pamelarussell/tciapathfindersignificance"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deeprtplus",
		URL:          []string{"https://github.com/horsepurve/deeprtplus"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gsia",
		URL:          []string{"https://github.com/ucsd-ccbb/gsia"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bel2abm",
		URL:          []string{"https://github.com/pybel/bel2abm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/masashitsubaki",
		URL:          []string{"https://github.com/masashitsubaki"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/grimon",
		URL:          []string{"https://github.com/mkanai/grimon"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/clonefinderapi",
		URL:          []string{"https://github.com/gstecher/clonefinderapi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dream-gene-essentiality-challenge",
		URL:          []string{"https://github.com/guanlab/dream-gene-essentiality-challenge"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mascot",
		URL:          []string{"https://github.com/nicfel/mascot"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ncyc",
		URL:          []string{"https://github.com/qichao1984/ncyc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/jarvis",
		URL:          []string{"https://github.com/usnistgov/jarvis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/steadystateconc",
		URL:          []string{"https://github.com/bioprocessdesignlab/steadystateconc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bmsim",
		URL:          []string{"https://github.com/pingchen09990102/bmsim"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lm-lstm-crf",
		URL:          []string{"https://github.com/yuzhimanhua/lm-lstm-crf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cglasso",
		URL:          []string{"https://github.com/luigiaugugliaro/cglasso"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/isdb",
		URL:          []string{"https://github.com/mullinslab/isdb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/libstructural",
		URL:          []string{"https://github.com/sys-bio/libstructural"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/officechromatography",
		URL:          []string{"https://github.com/officechromatography"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dante",
		URL:          []string{"https://github.com/jbudis/dante"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/moabb",
		URL:          []string{"https://github.com/neurotechx/moabb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metago",
		URL:          []string{"https://github.com/vvsmileyx/metago"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/upstrap",
		URL:          []string{"https://github.com/ccrainic/upstrap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/oyster-river-protocol",
		URL:          []string{"https://github.com/macmanes-lab/oyster_river_protocol"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cpi-em",
		URL:          []string{"https://github.com/vishakad/cpi-em"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/axe",
		URL:          []string{"https://github.com/kdmurray91/axe"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/catana",
		URL:          []string{"https://github.com/shiauck/catana"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metquest",
		URL:          []string{"https://github.com/ramanlab/metquest"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/imcmda",
		URL:          []string{"https://github.com/imcmdasourcecode/imcmda"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/robustclusteringpatientsubtyping",
		URL:          []string{"https://github.com/angy89/robustclusteringpatientsubtyping"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gappa",
		URL:          []string{"https://github.com/lczech/gappa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kubenow-plugin",
		URL:          []string{"https://github.com/phnmnl/kubenow-plugin"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lc-ms-pachyderm",
		URL:          []string{"https://github.com/pharmbio/lc-ms-pachyderm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fnbtools",
		URL:          []string{"https://github.com/noble-research-institute/fnbtools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tsnet",
		URL:          []string{"https://github.com/petraf01/tsnet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tripoly",
		URL:          []string{"https://github.com/ehsanmotazedi/tripoly"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fpkit",
		URL:          []string{"https://github.com/davidbajusz/fpkit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/squiggle",
		URL:          []string{"https://github.com/lab41/squiggle"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/oddt",
		URL:          []string{"https://github.com/oddt/oddt"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gpu-daemon",
		URL:          []string{"https://github.com/pcdslab/gpu-daemon"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pcdslab",
		URL:          []string{"https://github.com/pcdslab"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ggsashimi",
		URL:          []string{"https://github.com/guigolab/ggsashimi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biotite",
		URL:          []string{"https://github.com/biotite-dev/biotite"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ptempest",
		URL:          []string{"https://github.com/ruleworld/ptempest"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/epa-ng",
		URL:          []string{"https://github.com/pbdas/epa-ng"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mlstar",
		URL:          []string{"https://github.com/iferres/mlstar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/get-phylomarkers",
		URL:          []string{"https://github.com/vinuesa/get_phylomarkers"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/able",
		URL:          []string{"https://github.com/champost/able"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/varan-gie",
		URL:          []string{"https://github.com/popitsch/varan-gie"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mocaprecovery",
		URL:          []string{"https://github.com/numediart/mocaprecovery"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sigmat",
		URL:          []string{"https://github.com/jinfengxiao/sigmat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tensqr",
		URL:          []string{"https://github.com/soyeona/tensqr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/stasnet",
		URL:          []string{"https://github.com/molsysbio/stasnet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metdiff",
		URL:          []string{"https://github.com/compgenomics/metdiff"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/popdemog",
		URL:          []string{"https://github.com/yingzhou001/popdemog"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/subrecon",
		URL:          []string{"https://github.com/chrismonit/subrecon"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/remma",
		URL:          []string{"https://github.com/chaoning/remma"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dbc",
		URL:          []string{"https://github.com/bchidest/dbc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/eagle",
		URL:          []string{"https://github.com/tony-kuo/eagle"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/jami",
		URL:          []string{"https://github.com/schulzlab/jami"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hicomet",
		URL:          []string{"https://github.com/taehoonlee/hicomet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tsapa",
		URL:          []string{"https://github.com/bmilab/tsapa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pbxplore",
		URL:          []string{"https://github.com/pierrepo/pbxplore"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/eplmi",
		URL:          []string{"https://github.com/yahuang1991polyu/eplmi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pseaac-psepssm-wd",
		URL:          []string{"https://github.com/qust-bsbrc/pseaac-psepssm-wd"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kaptive-web",
		URL:          []string{"https://github.com/kelwyres/kaptive-web"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/genomeuplot",
		URL:          []string{"https://github.com/gaitat/genomeuplot"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/eeg-annotate",
		URL:          []string{"https://github.com/vislab/eeg-annotate"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ramp-backend",
		URL:          []string{"https://github.com/mathelab/ramp-backend"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/harc",
		URL:          []string{"https://github.com/shubhamchandak94/harc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/betaserpentine",
		URL:          []string{"https://github.com/stanislavspbgu/betaserpentine"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phenos",
		URL:          []string{"https://github.com/gact/phenos"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hippodeep",
		URL:          []string{"https://github.com/bthyreau/hippodeep"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/peath",
		URL:          []string{"https://github.com/jcna99/peath"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/jrgui",
		URL:          []string{"https://github.com/curieshicy/jrgui"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dmrfinder",
		URL:          []string{"https://github.com/jsh58/dmrfinder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/r2dgc",
		URL:          []string{"https://github.com/rramaker/r2dgc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/facets",
		URL:          []string{"https://github.com/mcantor2/facets"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vecscreen-plus-taxonomy",
		URL:          []string{"https://github.com/aaschaffer/vecscreen_plus_taxonomy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/iamcomparison",
		URL:          []string{"https://github.com/thallingerlab/iamcomparison"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rrbssim",
		URL:          []string{"https://github.com/xwbio/rrbssim"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/docker-rrbssim",
		URL:          []string{"https://github.com/xwbio/docker-rrbssim"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/quimp",
		URL:          []string{"https://github.com/celldynamics/quimp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rfathm6a",
		URL:          []string{"https://github.com/nongdaxiaofeng/rfathm6a"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cnn-smoothie",
		URL:          []string{"https://github.com/ih-_lab/cnn_smoothie"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/robustsparsecorrelation",
		URL:          []string{"https://github.com/angy89/robustsparsecorrelation"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/thetamater",
		URL:          []string{"https://github.com/radamsrha/thetamater"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metabridge-shiny",
		URL:          []string{"https://github.com/samhinshaw/metabridge_shiny"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/irscope",
		URL:          []string{"https://github.com/limpfrog/irscope"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/chewbbaca",
		URL:          []string{"https://github.com/b-ummi/chewbbaca"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/emase",
		URL:          []string{"https://github.com/churchill-lab/emase"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cellnomenclaturestudy",
		URL:          []string{"https://github.com/shenay/cellnomenclaturestudy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/xomeblender",
		URL:          []string{"https://github.com/rsemeraro/xomeblender"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dcaxl",
		URL:          []string{"https://github.com/mcubeg/dcaxl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dexsi",
		URL:          []string{"https://github.com/dexsi/dexsi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dp-gp-cluster",
		URL:          []string{"https://github.com/princetonuniversity/dp_gp_cluster"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fastore",
		URL:          []string{"https://github.com/refresh-bio/fastore"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gated-xnor",
		URL:          []string{"https://github.com/acrossv/gated-xnor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bdchemo",
		URL:          []string{"https://github.com/yiyiliu1/bdchemo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/consexpression",
		URL:          []string{"https://github.com/costasilvati/consexpression"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/microbiomedda",
		URL:          []string{"https://github.com/jchen1981/microbiomedda"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/morph-bulk",
		URL:          []string{"https://github.com/arzwa/morph-bulk"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/b-mis-normalization",
		URL:          []string{"https://github.com/ingallslabuw/b-mis-normalization"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/montra",
		URL:          []string{"https://github.com/bioinformatics-ua/montra"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nquire",
		URL:          []string{"https://github.com/clwgg/nquire"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pyefp",
		URL:          []string{"https://github.com/ale-odinokov/pyefp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phyloprofile",
		URL:          []string{"https://github.com/bionf/phyloprofile"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nimbus",
		URL:          []string{"https://github.com/erasmus-center-for-biomics/nimbus"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/maxent-ppi",
		URL:          []string{"https://github.com/ima23/maxent-ppi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/predcrp",
		URL:          []string{"https://github.com/nctuiclab/predcrp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dmbc",
		URL:          []string{"https://github.com/qunfengdong/dmbc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gwa-tutorial",
		URL:          []string{"https://github.com/mareesat/gwa_tutorial"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fsqn",
		URL:          []string{"https://github.com/jenniferfranks/fsqn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/frcnn-cad",
		URL:          []string{"https://github.com/riblidezso/frcnn_cad"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cyversewarwick",
		URL:          []string{"https://github.com/cyversewarwick"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sechidis",
		URL:          []string{"https://github.com/sechidis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/htdp",
		URL:          []string{"https://github.com/pmadanecki/htdp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/epslasso",
		URL:          []string{"https://github.com/xu1912/epslasso"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gsimp",
		URL:          []string{"https://github.com/wanderum/gsimp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/quentin",
		URL:          []string{"https://github.com/skumsp/quentin"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kpax3.jl",
		URL:          []string{"https://github.com/albertopessia/kpax3.jl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/structurefold2",
		URL:          []string{"https://github.com/structurefold2/structurefold2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/snyder-birth-death-code",
		URL:          []string{"https://github.com/kpzoo/snyder-birth-death-code"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pysster",
		URL:          []string{"https://github.com/budach/pysster"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/side-effects",
		URL:          []string{"https://github.com/poleksic/side-effects"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nanopack",
		URL:          []string{"https://github.com/wdecoster/nanopack"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/multimodal-brain-synthesis",
		URL:          []string{"https://github.com/agis85/multimodal_brain_synthesis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scvdmc",
		URL:          []string{"https://github.com/kuanglab/scvdmc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mccortex",
		URL:          []string{"https://github.com/mcveanlab/mccortex"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/glep",
		URL:          []string{"https://github.com/lzhlab/glep"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/replong",
		URL:          []string{"https://github.com/ruiguo-bio/replong"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/protein2genetree",
		URL:          []string{"https://github.com/udes-cobius/protein2genetree"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/superproteintree",
		URL:          []string{"https://github.com/udes-cobius/superproteintree"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/microbiome-fixed-reference",
		URL:          []string{"https://github.com/nci-biostats/microbiome-fixed-reference"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bamgineer",
		URL:          []string{"https://github.com/pughlab/bamgineer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mayonlppapcdss",
		URL:          []string{"https://github.com/ohnlp/mayonlppapcdss"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/regbenchmark",
		URL:          []string{"https://github.com/bioimageinformaticstampere/regbenchmark"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gi-cluster",
		URL:          []string{"https://github.com/icelu/gi_cluster"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/coge",
		URL:          []string{"https://github.com/lyonslab/coge"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gram-cnn",
		URL:          []string{"https://github.com/valdersoul/gram-cnn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/poap",
		URL:          []string{"https://github.com/inpacdb/poap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deopen",
		URL:          []string{"https://github.com/kimmo1019/deopen"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fitild",
		URL:          []string{"https://github.com/ogotoh/fitild"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/althap",
		URL:          []string{"https://github.com/realabolfazl/althap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/comer",
		URL:          []string{"https://github.com/minmarg/comer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rdolphin",
		URL:          []string{"https://github.com/danielcanueto/rdolphin"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/datamonkey-js",
		URL:          []string{"https://github.com/veg/datamonkey-js"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/parapred",
		URL:          []string{"https://github.com/eliberis/parapred"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/arioc",
		URL:          []string{"https://github.com/rwilton/arioc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/primer3-masker",
		URL:          []string{"https://github.com/bioinfo-ut/primer3_masker"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/primer3",
		URL:          []string{"https://github.com/primer3-org/primer3"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tpglda",
		URL:          []string{"https://github.com/ustc-hilab/tpglda"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/genssi",
		URL:          []string{"https://github.com/genssi-developer/genssi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/datasets",
		URL:          []string{"https://github.com/wgs-standards-and-analysis/datasets"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/chopstitch",
		URL:          []string{"https://github.com/bcgsc/chopstitch"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tripal-elasticsearch",
		URL:          []string{"https://github.com/tripal/tripal_elasticsearch"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tripal-analysis-expression",
		URL:          []string{"https://github.com/tripal/tripal_analysis_expression"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/elm",
		URL:          []string{"https://github.com/quantitativeimaging/elm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/radscripts",
		URL:          []string{"https://github.com/lyl8086/radscripts"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/wgs-analysis-demo",
		URL:          []string{"https://github.com/genetalks/wgs_analysis_demo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mp-lamp",
		URL:          []string{"https://github.com/tsudalab/mp-lamp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hvisaclassifier",
		URL:          []string{"https://github.com/bioprojects/hvisaclassifier"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/inform",
		URL:          []string{"https://github.com/greco-lab/inform"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phyloscanner",
		URL:          []string{"https://github.com/bdi-pathogens/phyloscanner"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rnaseqeval",
		URL:          []string{"https://github.com/kkrizanovic/rnaseqeval"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lsmm",
		URL:          []string{"https://github.com/mingjingsi/lsmm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mapexr",
		URL:          []string{"https://github.com/bmannakee/mapexr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metacherchant",
		URL:          []string{"https://github.com/ctlab/metacherchant"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nglview",
		URL:          []string{"https://github.com/arose/nglview"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/aselux",
		URL:          []string{"https://github.com/abl0719/aselux"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pegasus",
		URL:          []string{"https://github.com/mhayes20/pegasus"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/marathon",
		URL:          []string{"https://github.com/yuchaojiang/marathon"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tsam",
		URL:          []string{"https://github.com/penn-hui/tsam"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/plasmidseeker",
		URL:          []string{"https://github.com/bioinfo-ut/plasmidseeker"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nethypgeom",
		URL:          []string{"https://github.com/galanisl/nethypgeom"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hierarchical-rnns-model-for-ddi-extraction",
		URL:          []string{"https://github.com/zhangyijia1979/hierarchical-rnns-model-for-ddi-extraction"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vifi",
		URL:          []string{"https://github.com/namphuon/vifi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/svdquest",
		URL:          []string{"https://github.com/pranjalv123/svdquest"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mutscan",
		URL:          []string{"https://github.com/opengene/mutscan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ggrasp",
		URL:          []string{"https://github.com/jcventerinstitute/ggrasp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fem",
		URL:          []string{"https://github.com/haowenz/fem"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bret-analyzer",
		URL:          []string{"https://github.com/ychastagnier/bret-analyzer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/drwh",
		URL:          []string{"https://github.com/imagine-bdd/drwh"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gep2pep",
		URL:          []string{"https://github.com/franapoli/gep2pep"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gemse",
		URL:          []string{"https://github.com/genometric/gemse"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pqhe",
		URL:          []string{"https://github.com/biotangweiqi/pqhe"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sequana",
		URL:          []string{"https://github.com/sequana/sequana"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rbiomirgs",
		URL:          []string{"https://github.com/jzhangc/rbiomirgs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/polyainv",
		URL:          []string{"https://github.com/mghamilton/polyainv"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biomedical-corpora",
		URL:          []string{"https://github.com/dterg/biomedical_corpora"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mfmd",
		URL:          []string{"https://github.com/jadermcg/mfmd"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/canary",
		URL:          []string{"https://github.com/papenfusslab/canary"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/picky",
		URL:          []string{"https://github.com/thejacksonlaboratory/picky"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pynbs",
		URL:          []string{"https://github.com/idekerlab/pynbs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cscoretool",
		URL:          []string{"https://github.com/scoutzxb/cscoretool"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/indecut",
		URL:          []string{"https://github.com/megrawlab/indecut"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/clustertad",
		URL:          []string{"https://github.com/bdm-lab/clustertad"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bcigepred",
		URL:          []string{"https://github.com/brsaran/bcigepred"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/findgse",
		URL:          []string{"https://github.com/schneebergerlab/findgse"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/onto2graph",
		URL:          []string{"https://github.com/bio-ontology-research-group/onto2graph"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/chronqc",
		URL:          []string{"https://github.com/nilesh-tawari/chronqc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lme4qtl",
		URL:          []string{"https://github.com/variani/lme4qtl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/universal-tts",
		URL:          []string{"https://github.com/vanya-antonov/universal_tts"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/htssip",
		URL:          []string{"https://github.com/buckleylab/htssip"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/simclda",
		URL:          []string{"https://github.com//bioinfomaticscsu/simclda"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/garfield-ngs",
		URL:          []string{"https://github.com/gedoardo83/garfield-ngs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hyde",
		URL:          []string{"https://github.com/pblischak/hyde"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sve",
		URL:          []string{"https://github.com/thejacksonlaboratory/sve"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ifeature",
		URL:          []string{"https://github.com/superzchen/ifeature"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/exranges",
		URL:          []string{"https://github.com/dohertylab/exranges"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/viewbs",
		URL:          []string{"https://github.com/xie186/viewbs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/calq",
		URL:          []string{"https://github.com/voges/calq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/neptune",
		URL:          []string{"https://github.com/phac-nml/neptune"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bemkl-rbps",
		URL:          []string{"https://github.com/mehr-een/bemkl-rbps"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/xmwas",
		URL:          []string{"https://github.com/kuppal2/xmwas"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rwen",
		URL:          []string{"https://github.com/kiwtir/rwen"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biopyramid",
		URL:          []string{"https://github.com/jarny/biopyramid"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/adversarial-relation-classification",
		URL:          []string{"https://github.com/bionlproc/adversarial-relation-classification"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/l0adridge",
		URL:          []string{"https://github.com/liuzqx/l0adridge"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/xmabio",
		URL:          []string{"https://github.com/xmabio"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/parsnip",
		URL:          []string{"https://github.com/redarawi/parsnip"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sports1.0",
		URL:          []string{"https://github.com/junchaoshi/sports1.0"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/casian",
		URL:          []string{"https://github.com/mmahsa/casian"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/graftm-gpkgs",
		URL:          []string{"https://github.com/geronimp/graftm_gpkgs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/drugdiseasenet",
		URL:          []string{"https://github.com/azampvd/drugdiseasenet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tagfinder",
		URL:          []string{"https://github.com/jamigo/tagfinder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pytc",
		URL:          []string{"https://github.com/harmslab/pytc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mtb-report",
		URL:          []string{"https://github.com/jperera-bel/mtb-report"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ngsphy",
		URL:          []string{"https://github.com/merlyescalona/ngsphy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fmoc",
		URL:          []string{"https://github.com/ythuang0522/fmoc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hts-nim",
		URL:          []string{"https://github.com/brentp/hts-nim"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hts-nim-tools",
		URL:          []string{"https://github.com/brentp/hts-nim-tools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/benchmarkncvtools",
		URL:          []string{"https://github.com/oncostat/benchmarkncvtools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hasappy",
		URL:          []string{"https://github.com/gdiminin/hasappy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/medline-code-v2",
		URL:          []string{"https://github.com/vishrawas/medline-code_v2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/netics",
		URL:          []string{"https://github.com/cbg-ethz/netics"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pepkalc",
		URL:          []string{"https://github.com/peptoneinc/pepkalc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/spectre",
		URL:          []string{"https://github.com/maplesond/spectre"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ndexr",
		URL:          []string{"https://github.com/frankkramer-lab/ndexr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ssages-public",
		URL:          []string{"https://github.com/miccom/ssages-public"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scikit-ribo",
		URL:          []string{"https://github.com/schatzlab/scikit-ribo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/curatr",
		URL:          []string{"https://github.com/alexandrovteam/curatr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/trna-read-mapping",
		URL:          []string{"https://github.com/annehoffmann/trna-read-mapping"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/floc",
		URL:          []string{"https://github.com/fij/floc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fast-gep",
		URL:          []string{"https://github.com/jizhang-nz/fast-gep"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ssbio",
		URL:          []string{"https://github.com/sbrg/ssbio"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/negationdistribution",
		URL:          []string{"https://github.com/kevinbretonnelcohen/negationdistribution"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/acnviewer",
		URL:          []string{"https://github.com/fjd-ceph/acnviewer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/roicnn",
		URL:          []string{"https://github.com/zijingmao/roicnn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phesant",
		URL:          []string{"https://github.com/mrcieu/phesant"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bner",
		URL:          []string{"https://github.com/lvchen1989/bner"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mirvial",
		URL:          []string{"https://github.com/systemsbiologyofjianghanuniversity/mirvial"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lrsim",
		URL:          []string{"https://github.com/aquaskyline/lrsim"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/acnc-dame",
		URL:          []string{"https://github.com/bdpiccolo/acnc-dame"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ideepe",
		URL:          []string{"https://github.com/xypan1232/ideepe"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ctap",
		URL:          []string{"https://github.com/bwrc/ctap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vapr",
		URL:          []string{"https://github.com/ucsd-ccbb/vapr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/happe",
		URL:          []string{"https://github.com/lcnhappe/happe"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phylomad",
		URL:          []string{"https://github.com/duchene/phylomad"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tmcrys",
		URL:          []string{"https://github.com/brgenzim/tmcrys"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gtz",
		URL:          []string{"https://github.com/genetalks/gtz"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/frogs",
		URL:          []string{"https://github.com/geraldinepascal/frogs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pacom",
		URL:          []string{"https://github.com/smdb21/pacom"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/simpleitk-notebooks",
		URL:          []string{"https://github.com/insightsoftwareconsortium/simpleitk-notebooks"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/andy-s-algorithm",
		URL:          []string{"https://github.com/andlaw1841/andy-s-algorithm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/partitionuce",
		URL:          []string{"https://github.com/tagliacollo/partitionuce"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pfinderuce-swsc-en",
		URL:          []string{"https://github.com/tagliacollo/pfinderuce-swsc-en"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/quantnorm",
		URL:          []string{"https://github.com/tengfei-emory/quantnorm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gofeat",
		URL:          []string{"https://github.com/fabriciopa/gofeat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sammi",
		URL:          []string{"https://github.com/nelsongil92/sammi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/smcgenedeconv",
		URL:          []string{"https://github.com/moyanre/smcgenedeconv"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/glass",
		URL:          []string{"https://github.com/infspiredbat/glass"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tisk1.0",
		URL:          []string{"https://github.com/maglab-uconn/tisk1.0"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sparse-lowrank-regression",
		URL:          []string{"https://github.com/littleq1991/sparse_lowrank_regression"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/icopydav",
		URL:          []string{"https://github.com/vogetihrsh/icopydav"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/isgpt",
		URL:          []string{"https://github.com/srautonu/isgpt"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/prince",
		URL:          []string{"https://github.com/fosterlab/prince"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biocaddie",
		URL:          []string{"https://github.com/emory-irlab/biocaddie"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/althapalignr",
		URL:          []string{"https://github.com/jknightlab/althapalignr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/optpipe",
		URL:          []string{"https://github.com/andrashartmann/optpipe"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pycgm",
		URL:          []string{"https://github.com/cadop/pycgm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kirp-topological-features",
		URL:          []string{"https://github.com/chengjun583/kirp-topological-features"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/smrt",
		URL:          []string{"https://github.com/guofeieileen/smrt"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dto",
		URL:          []string{"https://github.com/drugtargetontology/dto"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/under",
		URL:          []string{"https://github.com/nigyta/dfast_core/under"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sc2p",
		URL:          []string{"https://github.com/haowulab/sc2p"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepsimulator",
		URL:          []string{"https://github.com/lykaust15/deepsimulator"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/uv-posit",
		URL:          []string{"https://github.com/jarosenb/uv_posit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hmtgo",
		URL:          []string{"https://github.com/k22liang/hmtgo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/embeddings-reproduction",
		URL:          []string{"https://github.com/fhalab/embeddings_reproduction"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gasol",
		URL:          []string{"https://github.com/accsc/gasol"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/diffgraph",
		URL:          []string{"https://github.com/zhangxf-ccnu/diffgraph"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dendrosplit",
		URL:          []string{"https://github.com/jessemzhang/dendrosplit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/starchip",
		URL:          []string{"https://github.com/losiclab/starchip"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fineradstructure",
		URL:          []string{"https://github.com/millanek/fineradstructure"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cornas",
		URL:          []string{"https://github.com/joel-lzb/cornas"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metagenome-pfam-score",
		URL:          []string{"https://github.com/eead-csic-compbio/metagenome_pfam_score"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/strawberry",
		URL:          []string{"https://github.com/ruolin/strawberry"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/synima",
		URL:          []string{"https://github.com/rhysf/synima"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gmm-marker-tracking",
		URL:          []string{"https://github.com/icthrm/gmm-marker-tracking"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/variantsdwh",
		URL:          []string{"https://github.com/zsi-bio/variantsdwh"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/warfarin-cohort",
		URL:          []string{"https://github.com/assafgo/warfarin-cohort"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mtb-sig-miner",
		URL:          []string{"https://github.com/zoozeal/mtb-sig-miner"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pea-m5c",
		URL:          []string{"https://github.com/cma2015/pea-m5c"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nitrilase",
		URL:          []string{"https://github.com/rover2380/nitrilase"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mytai",
		URL:          []string{"https://github.com/hajkd/mytai"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bsvf",
		URL:          []string{"https://github.com/bgi-sz/bsvf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/jvarkit",
		URL:          []string{"https://github.com/lindenb/jvarkit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fda-ars",
		URL:          []string{"https://github.com/lhncbc/fda-ars"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gtc",
		URL:          []string{"https://github.com/refresh-bio/gtc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/monster",
		URL:          []string{"https://github.com/quackenbushlab/monster"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/att-chemdner",
		URL:          []string{"https://github.com/lingluodlut/att-chemdner"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rvolcano",
		URL:          []string{"https://github.com/nishithkumarpaul/rvolcano"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ivc",
		URL:          []string{"https://github.com/namsyvo/ivc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/libgaba",
		URL:          []string{"https://github.com/ocxtal/libgaba"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/traveler",
		URL:          []string{"https://github.com/davidhoksza/traveler"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/chemts",
		URL:          []string{"https://github.com/tsudalab/chemts"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/alo-algorithm-for-kidney-exchanges",
		URL:          []string{"https://github.com/sarael-metwally/alo_algorithm_for_kidney_exchanges"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/isocket",
		URL:          []string{"https://github.com/woolfson-group/isocket"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepcov",
		URL:          []string{"https://github.com/psipred/deepcov"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/protein-sequence-analysis",
		URL:          []string{"https://github.com/manimkn89/protein-sequence-analysis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tc-recon",
		URL:          []string{"https://github.com/nojgaard/tc-recon"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nitrogenuptake2016",
		URL:          []string{"https://github.com/troyhill/nitrogenuptake2016"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/meta-proteome-analyzer",
		URL:          []string{"https://github.com/compomics/meta-proteome-analyzer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scepath",
		URL:          []string{"https://github.com/sqjin/scepath"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rcupcake",
		URL:          []string{"https://github.com/hms-dbmi/rcupcake"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rcupcake-case-studies",
		URL:          []string{"https://github.com/hms-dbmi/rcupcake-case-studies"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/emdunifrac",
		URL:          []string{"https://github.com/dkoslicki/emdunifrac"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/reactome-pengine",
		URL:          []string{"https://github.com/samwalrus/reactome-pengine"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/survival",
		URL:          []string{"https://github.com/batuff/survival"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/recovery",
		URL:          []string{"https://github.com/makovalab-psu/recovery"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/wscunmix",
		URL:          []string{"https://github.com/tedroman/wscunmix"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/matrixepistasis",
		URL:          []string{"https://github.com/fanglab/matrixepistasis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dr2di",
		URL:          []string{"https://github.com/huayu1111/dr2di"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/alpha",
		URL:          []string{"https://github.com/chilleo/alpha"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nessie",
		URL:          []string{"https://github.com/b3rse/nessie"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fab-phmm",
		URL:          []string{"https://github.com/bigsea-t/fab-phmm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/giggle",
		URL:          []string{"https://github.com/ryanlayer/giggle"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lis-context-viewer",
		URL:          []string{"https://github.com/legumeinfo/lis_context_viewer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metabodiff",
		URL:          []string{"https://github.com/andreasmock/metabodiff"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/knn",
		URL:          []string{"https://github.com/syang11/knn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scbgfengchao",
		URL:          []string{"https://github.com/scbgfengchao"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fundus-fractal-analysis",
		URL:          []string{"https://github.com/ignaciorlando/fundus-fractal-analysis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hvsm",
		URL:          []string{"https://github.com/kejia1215/hvsm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/spv-signaling-pathway-visualizer-v1.0",
		URL:          []string{"https://github.com/sinnefa/spv_signaling_pathway_visualizer_v1.0"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cancerdiscover",
		URL:          []string{"https://github.com/helikarlab/cancerdiscover"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bpbi",
		URL:          []string{"https://github.com/tsudalab/bpbi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rop",
		URL:          []string{"https://github.com/smangul1/rop"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/giremi",
		URL:          []string{"https://github.com/zhqingit/giremi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phenorank",
		URL:          []string{"https://github.com/alexjcornish/phenorank"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/delta",
		URL:          []string{"https://github.com/zhangzhwlab/delta"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/goleft",
		URL:          []string{"https://github.com/brentp/goleft"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/morfpred-plus",
		URL:          []string{"https://github.com/roneshsharma/morfpred-plus"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tcga-assembler-2",
		URL:          []string{"https://github.com/compgenome365/tcga-assembler-2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fmlc",
		URL:          []string{"https://github.com/fastmlc/fmlc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/readtools",
		URL:          []string{"https://github.com/magicdgs/readtools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dcmqi",
		URL:          []string{"https://github.com/qiicr/dcmqi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/axondeepseg",
		URL:          []string{"https://github.com/neuropoly/axondeepseg"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pureclip",
		URL:          []string{"https://github.com/skrakau/pureclip"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/exttada",
		URL:          []string{"https://github.com/hoangtn/exttada"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/supcp",
		URL:          []string{"https://github.com/lockef/supcp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/methcomp",
		URL:          []string{"https://github.com/jianhao2016/methcomp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/squeakr",
		URL:          []string{"https://github.com/splatlab/squeakr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metabocraft",
		URL:          []string{"https://github.com/argymeg/metabocraft"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/chemogenomicalg4dtipred",
		URL:          []string{"https://github.com/minghao2016/chemogenomicalg4dtipred"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hpviewer",
		URL:          []string{"https://github.com/yuhanh/hpviewer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/chipwig-v2",
		URL:          []string{"https://github.com/vidarmehr/chipwig-v2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scuba",
		URL:          []string{"https://github.com/gzampieri/scuba"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dataset-retrieval-pipeline",
		URL:          []string{"https://github.com/w2wei/dataset_retrieval_pipeline"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/diagram",
		URL:          []string{"https://github.com/reactome-pwp/diagram"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mosdepth",
		URL:          []string{"https://github.com/brentp/mosdepth"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biodataome",
		URL:          []string{"https://github.com/mensxmachina/biodataome"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kernelized-rank-learning",
		URL:          []string{"https://github.com/borgwardtlab/kernelized-rank-learning"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/spm-code",
		URL:          []string{"https://github.com/lecea/spm-code"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gemmer",
		URL:          []string{"https://github.com/barberislab/gemmer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cpdock",
		URL:          []string{"https://github.com/nemo8130/cpdock"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/haystack-bio",
		URL:          []string{"https://github.com/pinellolab/haystack_bio"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/limitfluxtocore",
		URL:          []string{"https://github.com/jbei/limitfluxtocore"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/clifinder",
		URL:          []string{"https://github.com/gred-clermont/clifinder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/morpholex-en",
		URL:          []string{"https://github.com/hugomailhot/morpholex-en"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hicnvtranslocation",
		URL:          []string{"https://github.com/ay-lab/hicnvtranslocation"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hictranshi-c",
		URL:          []string{"https://github.com/ay-lab/hictranshi-c"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/avesim",
		URL:          []string{"https://github.com/ay-lab/avesim"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cox-nnet",
		URL:          []string{"https://github.com/lanagarmire/cox-nnet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/speaq",
		URL:          []string{"https://github.com/beirnaert/speaq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mds2",
		URL:          []string{"https://github.com/sbbi/mds2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/flps",
		URL:          []string{"https://github.com/pmharrison/flps"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mcts-rna",
		URL:          []string{"https://github.com/tsudalab/mcts-rna"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ngscloud",
		URL:          []string{"https://github.com/ggfhf/ngscloud"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bnbr",
		URL:          []string{"https://github.com/siamakz/bnbr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cogstack",
		URL:          []string{"https://github.com/cogstack"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/semehr",
		URL:          []string{"https://github.com/cogstack/semehr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/jesse",
		URL:          []string{"https://github.com/biointec/jesse"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/confold2",
		URL:          []string{"https://github.com/multicom-toolbox/confold2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hybphylomaker",
		URL:          []string{"https://github.com/tomas-fer/hybphylomaker"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/allo",
		URL:          []string{"https://github.com/fibonaccirabbits/allo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/knoto-id",
		URL:          []string{"https://github.com/sib-swiss/knoto-id"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/diversity",
		URL:          []string{"https://github.com/narlikarlab/diversity"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/specter",
		URL:          []string{"https://github.com/rpeckner-broad/specter"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ramwas",
		URL:          []string{"https://github.com/andreyshabalin/ramwas"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ishspsy-project",
		URL:          []string{"https://github.com/ishspsy/project"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kinact",
		URL:          []string{"https://github.com/saezlab/kinact"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lenup",
		URL:          []string{"https://github.com/biomedbit/lenup"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sprint",
		URL:          []string{"https://github.com/lucian-ilie/sprint"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rdb2c",
		URL:          []string{"https://github.com/wzmao/rdb2c"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pehaplo",
		URL:          []string{"https://github.com/chjiao/pehaplo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ezqsar",
		URL:          []string{"https://github.com/shamsaraj/ezqsar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/concise",
		URL:          []string{"https://github.com/gagneurlab/concise"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/manuscript-avsec-bioinformatics-2017",
		URL:          []string{"https://github.com/gagneurlab/manuscript_avsec_bioinformatics_2017"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/detect-v2",
		URL:          []string{"https://github.com/parkinsonlab/detect-v2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/opencapsule",
		URL:          []string{"https://github.com/jhegemann/opencapsule"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/slicemap",
		URL:          []string{"https://github.com/mbarbie1/slicemap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/microbiome-v2.1",
		URL:          []string{"https://github.com/chaplenf/microbiome-v2.1"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/isop",
		URL:          []string{"https://github.com/nghiavtr/isop"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sv2",
		URL:          []string{"https://github.com/dantaki/sv2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/regulatorynetworkgmpmodel",
		URL:          []string{"https://github.com/caramirezal/regulatorynetworkgmpmodel"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mrrad",
		URL:          []string{"https://github.com/lasigebiotm/mrrad"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bartender-1.1",
		URL:          []string{"https://github.com/laozzzzz/bartender-1.1"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/seqan",
		URL:          []string{"https://github.com/seqan/seqan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cnn",
		URL:          []string{"https://github.com/alanfchina/cnn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/orchid",
		URL:          []string{"https://github.com/wittelab/orchid"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/megagta",
		URL:          []string{"https://github.com/hku-bal/megagta"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sorting-cancer-karyotypes",
		URL:          []string{"https://github.com/shamir-lab/sorting-cancer-karyotypes"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pasdqc",
		URL:          []string{"https://github.com/parklab/pasdqc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/larvalign",
		URL:          []string{"https://github.com/larvalign/larvalign"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/prefmd",
		URL:          []string{"https://github.com/feiglab/prefmd"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pames",
		URL:          []string{"https://github.com/cgplab/pames"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/leaf-gp",
		URL:          []string{"https://github.com/crop-phenomics-group/leaf-gp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/peitho",
		URL:          []string{"https://github.com/michaelphstumpf/peitho"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/imagepy",
		URL:          []string{"https://github.com/image-py/imagepy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/moloc",
		URL:          []string{"https://github.com/clagiamba/moloc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/multimed",
		URL:          []string{"https://github.com/siminab/multimed"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/data2dynamics",
		URL:          []string{"https://github.com/data2dynamics"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/eztree",
		URL:          []string{"https://github.com/yuwwu/eztree"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sim-tmle-tutorial",
		URL:          []string{"https://github.com/migariane/sim-tmle-tutorial"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/netminer",
		URL:          []string{"https://github.com/czllab/netminer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/genegraphics",
		URL:          []string{"https://github.com/katlabs/genegraphics"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/jsms",
		URL:          []string{"https://github.com/optimusmoose/jsms"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/genomedisco",
		URL:          []string{"https://github.com/kundajelab/genomedisco"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/price",
		URL:          []string{"https://github.com/erhard-lab/price"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/crenet",
		URL:          []string{"https://github.com/kouroshz/crenethttps://github.com/kouroshz/crenet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/alidetection",
		URL:          []string{"https://github.com/qizhangstat/alidetection"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/arcs",
		URL:          []string{"https://github.com/bcgsc/arcs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/idingo",
		URL:          []string{"https://github.com/minjinha/idingo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/plasmidtron",
		URL:          []string{"https://github.com/sanger-pathogens/plasmidtron"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/skygrowth",
		URL:          []string{"https://github.com/mrc-ide/skygrowth"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pychimera",
		URL:          []string{"https://github.com/insilichem/pychimera"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mentalist",
		URL:          []string{"https://github.com/wgs-tb/mentalist"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepdiffraction",
		URL:          []string{"https://github.com/subangstrom/deepdiffraction"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gdcrnatools",
		URL:          []string{"https://github.com/jialab-ucr/gdcrnatools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/logical-modelling-pipeline",
		URL:          []string{"https://github.com/sysbio-curie/logical_modelling_pipeline"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/neuromatic",
		URL:          []string{"https://github.com/silverlabucl/neuromatic"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/baldr",
		URL:          []string{"https://github.com/bosingerlab/baldr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/idlp",
		URL:          []string{"https://github.com/nkiip/idlp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hgtsim",
		URL:          []string{"https://github.com/songweizhi/hgtsim"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/perf",
		URL:          []string{"https://github.com/rkmlab/perf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/3dmax",
		URL:          []string{"https://github.com/bdm-lab/3dmax"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/planteome",
		URL:          []string{"https://github.com/planteome"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fetalfraction-snpimpute",
		URL:          []string{"https://github.com/kmj403/fetalfraction-snpimpute"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/miraimplications",
		URL:          []string{"https://github.com/comprna/miraimplications"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/badlands-model",
		URL:          []string{"https://github.com/badlands-model"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/penndiff",
		URL:          []string{"https://github.com/tigerhu15/penndiff"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/aether",
		URL:          []string{"https://github.com/kosticlab/aether"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/saint2",
		URL:          []string{"https://github.com/sauloho/saint2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/abscan",
		URL:          []string{"https://github.com/csw407/abscan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/args-oap-v2.0",
		URL:          []string{"https://github.com/xiaole99/args-oap-v2.0"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cjbitseq",
		URL:          []string{"https://github.com/mqbssppe/cjbitseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/asar",
		URL:          []string{"https://github.com/askarbek-orakov/asar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/capsim",
		URL:          []string{"https://github.com/devika1/capsim"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/virapipe",
		URL:          []string{"https://github.com/ngseq/virapipe"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gaze3dfix",
		URL:          []string{"https://github.com/applied-cognition-research/gaze3dfix"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/riboviz",
		URL:          []string{"https://github.com/shahpr/riboviz"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/q2-fragment-insertion",
		URL:          []string{"https://github.com/biocore/q2-fragment-insertion"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ngmlr",
		URL:          []string{"https://github.com/philres/ngmlr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sniffles",
		URL:          []string{"https://github.com/fritzsedlazeck/sniffles"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/probannopy",
		URL:          []string{"https://github.com/pricelab/probannopy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/repository",
		URL:          []string{"https://github.com/pocppsus/repository"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tongchf",
		URL:          []string{"https://github.com/tongchf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/serpent",
		URL:          []string{"https://github.com/comprna/serpent"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/matam",
		URL:          []string{"https://github.com/bonsai-team/matam"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hmmerctter",
		URL:          []string{"https://github.com/bbcmdp/hmmerctter"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/spark-cpvs",
		URL:          []string{"https://github.com/laeeq80/spark-cpvs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/spherical",
		URL:          []string{"https://github.com/thh32/spherical"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pcgr",
		URL:          []string{"https://github.com/sigven/pcgr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/integrate-vis",
		URL:          []string{"https://github.com/chrismaherlab/integrate-vis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/red-lesion-detection",
		URL:          []string{"https://github.com/ignaciorlando/red-lesion-detection"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/uea-srna-workbench",
		URL:          []string{"https://github.com/srnaworkbenchuea/uea_srna_workbench"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ssp",
		URL:          []string{"https://github.com/rnakato/ssp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/flib-coevo",
		URL:          []string{"https://github.com/sauloho/flib-coevo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/flowermorphology",
		URL:          []string{"https://github.com/deyneko/flowermorphology"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/visnormsc",
		URL:          []string{"https://github.com/solo7773/visnormsc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/quadmutex",
		URL:          []string{"https://github.com/bokhariy/quadmutex"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ngrhmda",
		URL:          []string{"https://github.com/yahuang1991/ngrhmda"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/orbslamm",
		URL:          []string{"https://github.com/hdaoud/orbslamm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ctdquerier",
		URL:          []string{"https://github.com/isglobal-brge/ctdquerier"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dasc",
		URL:          []string{"https://github.com/zhanglabnku/dasc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gac",
		URL:          []string{"https://github.com/manalirupji/gac"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mea-tools",
		URL:          []string{"https://github.com/dbridges/mea-tools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pmfa",
		URL:          []string{"https://github.com/aalto-ics-kepaco/pmfa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/treewas",
		URL:          []string{"https://github.com/caitiecollins/treewas"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/shinycircos",
		URL:          []string{"https://github.com/venyao/shinycircos"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/spcovr",
		URL:          []string{"https://github.com/katrijnvandeun/spcovr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hicaptools",
		URL:          []string{"https://github.com/sahlenlab/hicaptools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vcselection",
		URL:          []string{"https://github.com/jingzhai63/vcselection"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sage2",
		URL:          []string{"https://github.com/lucian-ilie/sage2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rss",
		URL:          []string{"https://github.com/stephenslab/rss"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/medtest",
		URL:          []string{"https://github.com/jchen1981/medtest"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/simlrgithub",
		URL:          []string{"https://github.com/batzogloulabsu/simlrgithub"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cobasi",
		URL:          []string{"https://github.com/laura-gomez/cobasi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/drl4cellmovement",
		URL:          []string{"https://github.com/zwang84/drl4cellmovement"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/narx-qoe-release",
		URL:          []string{"https://github.com/christosbampis/narx_qoe_release"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/flowlearn",
		URL:          []string{"https://github.com/mlux86/flowlearn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/secretsanta",
		URL:          []string{"https://github.com/gogleva/secretsanta"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cpinsim",
		URL:          []string{"https://github.com/biancastoecker/cpinsim"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bwas",
		URL:          []string{"https://github.com/weikanggong/bwas"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bioont-search-benchmark",
		URL:          []string{"https://github.com/danielapoliveira/bioont-search-benchmark"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/crossplan",
		URL:          []string{"https://github.com/murali-group/crossplan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bayesianpgmm",
		URL:          []string{"https://github.com/lockef/bayesianpgmm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ght",
		URL:          []string{"https://github.com/zhaoni153/ght"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/wenpingd",
		URL:          []string{"https://github.com/wenpingd"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rnaplonc",
		URL:          []string{"https://github.com/tatiannenegri/rnaplonc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hnmdrp",
		URL:          []string{"https://github.com/ustc-hilab/hnmdrp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/isescan",
		URL:          []string{"https://github.com/xiezhq/isescan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/data-prism",
		URL:          []string{"https://github.com/agresearch/data_prism"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pymzml",
		URL:          []string{"https://github.com/pymzml/pymzml"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/iterassemble",
		URL:          []string{"https://github.com/looselab/iterassemble"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/galaxydocker",
		URL:          []string{"https://github.com/carpem/galaxydocker"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/esigen",
		URL:          []string{"https://github.com/insilichem/esigen"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/snapperdb",
		URL:          []string{"https://github.com/phe-bioinformatics/snapperdb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/snapperdb-references",
		URL:          []string{"https://github.com/phe-bioinformatics/snapperdb_references"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bives-statsgenerator",
		URL:          []string{"https://github.com/binfalse/bives-statsgenerator"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scanpav",
		URL:          []string{"https://github.com/wtsi-hpag/scanpav"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pdhs-elm",
		URL:          []string{"https://github.com/shanxinzhang/pdhs-elm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ddpcrclust",
		URL:          []string{"https://github.com/bgbrink/ddpcrclust"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ddpcrvis",
		URL:          []string{"https://github.com/bgbrink/ddpcrvis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gfcom",
		URL:          []string{"https://github.com/jiyuanhu/gfcom"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phenotypesimulator",
		URL:          []string{"https://github.com/hannahvmeyer/phenotypesimulator"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kestrel",
		URL:          []string{"https://github.com/paudano/kestrel"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kescases",
		URL:          []string{"https://github.com/paudano/kescases"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/p5-bpwrapper",
		URL:          []string{"https://github.com/bioperl/p5-bpwrapper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/guoweifeng",
		URL:          []string{"https://github.com/wilfongguo/guoweifeng"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/postplsr",
		URL:          []string{"https://github.com/pn51/postplsr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/krait",
		URL:          []string{"https://github.com/lmdu/krait"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dynadup",
		URL:          []string{"https://github.com/smirarab/dynadup"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/viper",
		URL:          []string{"https://github.com/marwoes/viper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/d3ner",
		URL:          []string{"https://github.com/aidantee/d3ner"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/accumulate",
		URL:          []string{"https://github.com/dwinter/accumulate"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gcevobase",
		URL:          []string{"https://github.com/nextgenbioinformatics/gcevobase"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scanpy",
		URL:          []string{"https://github.com/theislab/scanpy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/anndata",
		URL:          []string{"https://github.com/theislab/anndata"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gfusion",
		URL:          []string{"https://github.com/xiaofengsong/gfusion"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/walking-rdf-and-owl",
		URL:          []string{"https://github.com/bio-ontology-research-group/walking-rdf-and-owl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rfscorevs",
		URL:          []string{"https://github.com/oddt/rfscorevs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rfscorevs-binary",
		URL:          []string{"https://github.com/oddt/rfscorevs_binary"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kmerind",
		URL:          []string{"https://github.com/parbliss/kmerind"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/blca",
		URL:          []string{"https://github.com/qunfengdong/blca"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gpmicrobiome",
		URL:          []string{"https://github.com/tare/gpmicrobiome"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phylogeotool",
		URL:          []string{"https://github.com/rega-cev/phylogeotool"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/neo-pep-tool",
		URL:          []string{"https://github.com/pgb-liv/neo-pep-tool"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/snovault",
		URL:          []string{"https://github.com/encode-dcc/snovault"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/encoded",
		URL:          []string{"https://github.com/encode-dcc/encoded"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pathway-mapper",
		URL:          []string{"https://github.com/ivis-at-bilkent/pathway-mapper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/japsa",
		URL:          []string{"https://github.com/mdcao/japsa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lhqxinghun-bioinformatics",
		URL:          []string{"https://github.com/lhqxinghun/bioinformatics"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/differential-mutation-analysis",
		URL:          []string{"https://github.com/singh-lab/differential-mutation-analysis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gaudi",
		URL:          []string{"https://github.com/insilichem/gaudi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dbh",
		URL:          []string{"https://github.com/nwpu134/dbh"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ohana",
		URL:          []string{"https://github.com/jade-cheng/ohana"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/minimizers",
		URL:          []string{"https://github.com/gmarcais/minimizers"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nanatex",
		URL:          []string{"https://github.com/ryusukemomota/nanatex"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rlowpc",
		URL:          []string{"https://github.com/wyguo/rlowpc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pepx",
		URL:          []string{"https://github.com/calipho-sib/pepx"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gppfst",
		URL:          []string{"https://github.com/radamsrha/gppfst"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/3dec",
		URL:          []string{"https://github.com/flishwnag/3dec"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/quasar",
		URL:          []string{"https://github.com/piquelab/quasar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/truc",
		URL:          []string{"https://github.com/oarnaiz/truc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dugongbioinformatics",
		URL:          []string{"https://github.com/dugongbioinformatics"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/casper",
		URL:          []string{"https://github.com/trinhlab/casper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/splicify",
		URL:          []string{"https://github.com/nki-tgo/splicify"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nephele",
		URL:          []string{"https://github.com/niaid/nephele"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/derep-np",
		URL:          []string{"https://github.com/clzani/derep-np"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/protein-ml",
		URL:          []string{"https://github.com/sbrg/protein_ml"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vdjdb-db",
		URL:          []string{"https://github.com/antigenomics/vdjdb-db"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/superangle",
		URL:          []string{"https://github.com/subangstrom/superangle"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dqbioinformatics",
		URL:          []string{"https://github.com/rbouadjenek/dqbioinformatics"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/segmentation-software",
		URL:          []string{"https://github.com/pswain/segmentation-software"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/paragsea",
		URL:          []string{"https://github.com/ysycloud/paragsea"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ismb2017-lstm",
		URL:          []string{"https://github.com/minxueric/ismb2017_lstm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/crisprpred",
		URL:          []string{"https://github.com/khaled-buet/crisprpred"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phyloligo",
		URL:          []string{"https://github.com/itsmeludo/phyloligo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/m2align",
		URL:          []string{"https://github.com/khaosresearch/m2align"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/wft4galaxy",
		URL:          []string{"https://github.com/phnmnl/wft4galaxy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cufusion",
		URL:          []string{"https://github.com/zhangxaochen/cufusion"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/flexbar",
		URL:          []string{"https://github.com/seqan/flexbar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/haplomerger2",
		URL:          []string{"https://github.com/mapleforest/haplomerger2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/faucet",
		URL:          []string{"https://github.com/shamir-lab/faucet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pulser",
		URL:          []string{"https://github.com/dieterich-lab/pulser"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mvlr",
		URL:          []string{"https://github.com/suleimank/mvlr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/glanet",
		URL:          []string{"https://github.com/burcakotlu/glanet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepchem",
		URL:          []string{"https://github.com/deepchem/deepchem"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/workflow-of-chipseq",
		URL:          []string{"https://github.com/ddhb/workflow_of_chipseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hiclipr",
		URL:          []string{"https://github.com/luslab/hiclipr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mintmap",
		URL:          []string{"https://github.com/tju-cmc-org/mintmap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/aqua",
		URL:          []string{"https://github.com/tparidae/aqua"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kpic2",
		URL:          []string{"https://github.com/hcji/kpic2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biomake",
		URL:          []string{"https://github.com/evoldoers/biomake"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/msqrob",
		URL:          []string{"https://github.com/statomics/msqrob"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/oxsa",
		URL:          []string{"https://github.com/oxsatoolbox/oxsa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/offsetbasedgraph",
		URL:          []string{"https://github.com/uio-cels/offsetbasedgraph"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/genomicgraphcoords",
		URL:          []string{"https://github.com/uio-cels/genomicgraphcoords"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mrgsea",
		URL:          []string{"https://github.com/zaedpolsl/mrgsea"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/casas",
		URL:          []string{"https://github.com/manalirupji/casas"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gfapy",
		URL:          []string{"https://github.com/ggonnella/gfapy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/keratocyte",
		URL:          []string{"https://github.com/danielprieto/keratocyte"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mackinac",
		URL:          []string{"https://github.com/mmundy42/mackinac"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ngscheckmate",
		URL:          []string{"https://github.com/parklab/ngscheckmate"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bronchomix",
		URL:          []string{"https://github.com/drdanielgartner/bronchomix"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/opennft-demo",
		URL:          []string{"https://github.com/opennft/opennft_demo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cpr",
		URL:          []string{"https://github.com/mathcom/cpr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tigeri",
		URL:          []string{"https://github.com/namshik/tigeri"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/addit",
		URL:          []string{"https://github.com/ndbl/addit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/epinem",
		URL:          []string{"https://github.com/cbg-ethz/epinem"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/regresshaplo",
		URL:          []string{"https://github.com/sleviyang/regresshaplo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/umediation",
		URL:          []string{"https://github.com/sharonlutz/umediation"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sbml-mod-ws",
		URL:          []string{"https://github.com/molecularbioinformatics/sbml-mod-ws"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bioqueue",
		URL:          []string{"https://github.com/liyao001/bioqueue"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gpica",
		URL:          []string{"https://github.com/dittehald/gpica"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/titer",
		URL:          []string{"https://github.com/zhangsaithu/titer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/positive-selection",
		URL:          []string{"https://github.com/robinvanderlee/positive-selection"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fimtrack",
		URL:          []string{"https://github.com/i-git/fimtrack"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/musitedeep",
		URL:          []string{"https://github.com/duolinwang/musitedeep"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/recov",
		URL:          []string{"https://github.com/datduong/recov"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cloudneo",
		URL:          []string{"https://github.com/thejacksonlaboratory/cloudneo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/wfes",
		URL:          []string{"https://github.com/dekoning-lab/wfes"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/irap",
		URL:          []string{"https://github.com/nunofonseca/irap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phylip-enhance",
		URL:          []string{"https://github.com/shimadamk/phylip_enhance"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/conffuse",
		URL:          []string{"https://github.com/zhiqin-huang/conffuse"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/countsimqc",
		URL:          []string{"https://github.com/csoneson/countsimqc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pgrnafinder",
		URL:          []string{"https://github.com/xiexiaowei/pgrnafinder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/netprot",
		URL:          []string{"https://github.com/gohwils/netprot"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kmerpyramid",
		URL:          []string{"https://github.com/jkruppa/kmerpyramid"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ad-zcc",
		URL:          []string{"https://github.com/afteich/ad_zcc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metabolomicstools",
		URL:          []string{"https://github.com/raspicer/metabolomicstools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/eviltools",
		URL:          []string{"https://github.com/fickludd/eviltools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mptm",
		URL:          []string{"https://github.com/ustc-hilab/mptm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mdplot",
		URL:          []string{"https://github.com/mdplot/mdplot"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/picalculax",
		URL:          []string{"https://github.com/ebjerrum/picalculax"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/heap",
		URL:          []string{"https://github.com/meiji-bioinf/heap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/siapopr",
		URL:          []string{"https://github.com/olliemcdonald/siapopr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gsu",
		URL:          []string{"https://github.com/changshuaiwei/gsu"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metakallisto",
		URL:          []string{"https://github.com/pachterlab/metakallisto"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pmanalyzer",
		URL:          []string{"https://github.com/dacuevas/pmanalyzer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pmanalyzerweb",
		URL:          []string{"https://github.com/dacuevas/pmanalyzerweb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/coda",
		URL:          []string{"https://github.com/kundajelab/coda"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/debgr",
		URL:          []string{"https://github.com/splatlab/debgr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/regnet",
		URL:          []string{"https://github.com/seifemi/regnet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/draco-stem",
		URL:          []string{"https://github.com/virtualplants/draco-stem"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/batcheffectremoval",
		URL:          []string{"https://github.com/ushaham/batcheffectremoval"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biren",
		URL:          []string{"https://github.com/wenjiegroup/biren"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mina",
		URL:          []string{"https://github.com/sciencetoolkit/mina"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/t2ga",
		URL:          []string{"https://github.com/roqe/t2ga"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/adaptivehm",
		URL:          []string{"https://github.com/benliemory/adaptivehm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nucdiff",
		URL:          []string{"https://github.com/uio-cels/nucdiff"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepem3d",
		URL:          []string{"https://github.com/divelab/deepem3d"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/stego",
		URL:          []string{"https://github.com/dschlauch/stego"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/garlic",
		URL:          []string{"https://github.com/szpiech/garlic"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/if",
		URL:          []string{"https://github.com/biocryst/if"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cerena",
		URL:          []string{"https://github.com/cerenadevelopers/cerena"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/regate",
		URL:          []string{"https://github.com/c3bi-pasteur-fr/regate"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/amylogramanalysis",
		URL:          []string{"https://github.com/michbur/amylogramanalysis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepcnf-auc",
		URL:          []string{"https://github.com/realbigws/deepcnf_auc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hybriddetective",
		URL:          []string{"https://github.com/bwringe/hybriddetective"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gennet",
		URL:          []string{"https://github.com/raquele/gennet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deisom",
		URL:          []string{"https://github.com/hao-peng/deisom"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/stopgap",
		URL:          []string{"https://github.com/statgenprd/stopgap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/novoplasty",
		URL:          []string{"https://github.com/ndierckx/novoplasty"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/complexviewer",
		URL:          []string{"https://github.com/micommunity/complexviewer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ksea",
		URL:          []string{"https://github.com/casecpb/ksea"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hlaprofiler",
		URL:          []string{"https://github.com/expressionanalysis/hlaprofiler"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mvp-aaa-codelabs",
		URL:          []string{"https://github.com/stanfordbioinformatics/mvp_aaa_codelabs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hirgc",
		URL:          []string{"https://github.com/yuansliu/hirgc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fqc",
		URL:          []string{"https://github.com/pnnl/fqc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/strainseeker",
		URL:          []string{"https://github.com/bioinfo-ut/strainseeker"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hlama",
		URL:          []string{"https://github.com/bihealth/hlama"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/smlocalizer",
		URL:          []string{"https://github.com/kristofferbernhem/smlocalizer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/craftgbd",
		URL:          []string{"https://github.com/craftgbd/craftgbd"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/findr",
		URL:          []string{"https://github.com/lingfeiwang/findr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mappi-dat",
		URL:          []string{"https://github.com/compomics/mappi-dat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/genmap-comparator",
		URL:          []string{"https://github.com/holtzy/genmap-comparator"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hetero-rp",
		URL:          []string{"https://github.com/younglululu/hetero-rp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ngs-pipe",
		URL:          []string{"https://github.com/cbg-ethz/ngs-pipe"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/grnvbem",
		URL:          []string{"https://github.com/mscastillo/grnvbem"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/chromdet",
		URL:          []string{"https://github.com/david-juan/chromdet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/modmaps3d",
		URL:          []string{"https://github.com/moleculardistancemaps/modmaps3d"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lolopicker",
		URL:          []string{"https://github.com/jcarrotzhang/lolopicker"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/svmine",
		URL:          []string{"https://github.com/xyc0813/svmine"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rnaseqview",
		URL:          []string{"https://github.com/ncbi-hackathons/rnaseqview"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fastq-brew",
		URL:          []string{"https://github.com/dohalloran/fastq_brew"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/condition-dependent-correlation-subgroups-ccs",
		URL:          []string{"https://github.com/abhatta3/condition-dependent-correlation-subgroups-ccs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/badock",
		URL:          []string{"https://github.com/badocksbi/badock"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nss-of-lwir-and-vissible-images",
		URL:          []string{"https://github.com/ujemd/nss-of-lwir-and-vissible-images"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/breakpointsurveyor",
		URL:          []string{"https://github.com/ding-lab/breakpointsurveyor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tsis",
		URL:          []string{"https://github.com/wyguo/tsis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tadbit",
		URL:          []string{"https://github.com/3dgenomes/tadbit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rqt",
		URL:          []string{"https://github.com/izhbannikov/rqt"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lipidfinder",
		URL:          []string{"https://github.com/cjbrasher/lipidfinder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/emolfrag",
		URL:          []string{"https://github.com/liutairan/emolfrag"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/orthopara-constraintchecker",
		URL:          []string{"https://github.com/udem-lbit/orthopara-constraintchecker"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/teacheng",
		URL:          []string{"https://github.com/knoweng/teacheng"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metacache",
		URL:          []string{"https://github.com/muellan/metacache"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gromacs-mic",
		URL:          []string{"https://github.com/tianhe2/gromacs-mic"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/xifdr",
		URL:          []string{"https://github.com/rappsilberlab/xifdr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/boltzmannmachines.jl",
		URL:          []string{"https://github.com/binderh/boltzmannmachines.jl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/drug-target-prediction",
		URL:          []string{"https://github.com/zongnansu1982/drug-target-prediction"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pdna",
		URL:          []string{"https://github.com/zhangxf-ccnu/pdna"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/varcap",
		URL:          []string{"https://github.com/ma2o/varcap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/chimeraviz",
		URL:          []string{"https://github.com/stianlagstad/chimeraviz"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lmmo",
		URL:          []string{"https://github.com/ekffar/lmmo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/arachne",
		URL:          []string{"https://github.com/leonidsavtchenko/arachne"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/toolbox-romano-et-al",
		URL:          []string{"https://github.com/zebrain-lab/toolbox-romano-et-al"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ms2ldaviz",
		URL:          []string{"https://github.com/sdrogers/ms2ldaviz"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/wdnfinder",
		URL:          []string{"https://github.com/dustincys/wdnfinder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/librfn",
		URL:          []string{"https://github.com/bioinf-jku/librfn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cmsa",
		URL:          []string{"https://github.com/wangvsa/cmsa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fungap",
		URL:          []string{"https://github.com/compsynbiolab-koreauniv/fungap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biopax.viz",
		URL:          []string{"https://github.com/cgu-certh/biopax.viz"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ppi",
		URL:          []string{"https://github.com/cyinbox/ppi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/qsmooth",
		URL:          []string{"https://github.com/stephaniehicks/qsmooth"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/genomescope",
		URL:          []string{"https://github.com/schatzlab/genomescope"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hic-spector",
		URL:          []string{"https://github.com/gersteinlab/hic-spector"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/haploforge",
		URL:          []string{"https://github.com/mtekman/haploforge"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cnt-ilp",
		URL:          []string{"https://github.com/raphael-group/cnt-ilp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/xenophile",
		URL:          []string{"https://github.com/mgleeming/xenophile"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phylo-node",
		URL:          []string{"https://github.com/dohalloran/phylo-node"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tetools",
		URL:          []string{"https://github.com/l-modolo/tetools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pimp",
		URL:          []string{"https://github.com/ronandaly/pimp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/probabilitymapviewer",
		URL:          []string{"https://github.com/crbs/probabilitymapviewer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metacrast",
		URL:          []string{"https://github.com/molleraj/metacrast"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/chimerscope",
		URL:          []string{"https://github.com/chimerscope/chimerscope"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/openbehavior",
		URL:          []string{"https://github.com/chen42/openbehavior"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/compms2miner",
		URL:          []string{"https://github.com/wmbedmands/compms2miner"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sitar",
		URL:          []string{"https://github.com/zhenwendai/sitar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sccaller",
		URL:          []string{"https://github.com/biosinodx/sccaller"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gatekeeper",
		URL:          []string{"https://github.com/bilkentcompgen/gatekeeper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kart",
		URL:          []string{"https://github.com/hsinnan75/kart"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pcrduplicates",
		URL:          []string{"https://github.com/vibansal/pcrduplicates"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/atropos",
		URL:          []string{"https://github.com/jdidion/atropos"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cluster-wgs-data",
		URL:          []string{"https://github.com/heidefier/cluster_wgs_data"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/e2fm",
		URL:          []string{"https://github.com/montecuollo/e2fm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pastaspark",
		URL:          []string{"https://github.com/citiususc/pastaspark"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gemsplice-code",
		URL:          []string{"https://github.com/gemsplice/gemsplice_code"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/isa-tools",
		URL:          []string{"https://github.com/isa-tools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/collector",
		URL:          []string{"https://github.com/phi-grib/collector"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/miic",
		URL:          []string{"https://github.com/miicteam/miic"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hgvs",
		URL:          []string{"https://github.com/biocommons/hgvs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/multiplierz",
		URL:          []string{"https://github.com/blaisproteomics/multiplierz"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepboost",
		URL:          []string{"https://github.com/dongfanghong/deepboost"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ngmaster",
		URL:          []string{"https://github.com/mdu-phl/ngmaster"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/netprophet-2.0",
		URL:          []string{"https://github.com/yiming-kang/netprophet_2.0"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pheno4j",
		URL:          []string{"https://github.com/phenopolis/pheno4j"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fuento",
		URL:          []string{"https://github.com/davidweichselbaum/fuento"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/funnel-gsea-r-package",
		URL:          []string{"https://github.com/yunzhang813/funnel-gsea-r-package"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phelim",
		URL:          []string{"https://github.com/andreariba/phelim"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mcclintock",
		URL:          []string{"https://github.com/bergmanlab/mcclintock"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/reactionminer",
		URL:          []string{"https://github.com/ramanlab/reactionminer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sequence2vec",
		URL:          []string{"https://github.com/ramzan1990/sequence2vec"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/structure-threader",
		URL:          []string{"https://github.com/stuntspt/structure_threader"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/shinyheatmap",
		URL:          []string{"https://github.com/bohdan-khomtchouk/shinyheatmap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fastheatmap",
		URL:          []string{"https://github.com/bohdan-khomtchouk/fastheatmap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ezmap",
		URL:          []string{"https://github.com/dekoning-lab/ezmap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fsepsa",
		URL:          []string{"https://github.com/udes-cobius/fsepsa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/axonpacking",
		URL:          []string{"https://github.com/neuropoly/axonpacking"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/qbioimages",
		URL:          []string{"https://github.com/blazi13/qbioimages"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/neoepitope",
		URL:          []string{"https://github.com/zhanglabstjude/neoepitope"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ntcard",
		URL:          []string{"https://github.com/bcgsc/ntcard"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/strandscript",
		URL:          []string{"https://github.com/seasky002002/strandscript"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/aozan",
		URL:          []string{"https://github.com/genomicpariscentre/aozan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/snaptron",
		URL:          []string{"https://github.com/christopherwilks/snaptron"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/snaptron-experiments",
		URL:          []string{"https://github.com/christopherwilks/snaptron-experiments"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/redcap-ddp",
		URL:          []string{"https://github.com/wcmc-research-informatics/redcap-ddp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/eqtlseq",
		URL:          []string{"https://github.com/dvav/eqtlseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phyc",
		URL:          []string{"https://github.com/ymatts/phyc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/combination-index",
		URL:          []string{"https://github.com/rbbt-workflows/combination_index"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/synthesizer",
		URL:          []string{"https://github.com/lisagandy/synthesizer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/km2gcn",
		URL:          []string{"https://github.com/juanbot/km2gcn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/stp-gui",
		URL:          []string{"https://github.com/elettrascicomp/stp-gui"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rl-skat",
		URL:          []string{"https://github.com/cozygene/rl-skat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metatopics",
		URL:          []string{"https://github.com/bm2-lab/metatopics"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/reptile",
		URL:          []string{"https://github.com/yupenghe/reptile"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/napeasy",
		URL:          []string{"https://github.com/khp-informatics/napeasy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fmtree",
		URL:          []string{"https://github.com/chhylp123/fmtree"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sioor",
		URL:          []string{"https://github.com/hongyu-miao/sioor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/knowledgediscovery",
		URL:          []string{"https://github.com/jakelever/knowledgediscovery"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sexcmd",
		URL:          []string{"https://github.com/lovemun/sexcmd"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metcirc",
		URL:          []string{"https://github.com/plantdefensemetabolism/metcirc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/annotatr",
		URL:          []string{"https://github.com/rcavalcante/annotatr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/partie",
		URL:          []string{"https://github.com/linsalrob/partie"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/svpv",
		URL:          []string{"https://github.com/vccri/svpv"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mirmmr",
		URL:          []string{"https://github.com/ding-lab/mirmmr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tea",
		URL:          []string{"https://github.com/solgenomics/tea"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/spectrassembler",
		URL:          []string{"https://github.com/antrec/spectrassembler"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hopland",
		URL:          []string{"https://github.com/netland-ntu/hopland"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tagger",
		URL:          []string{"https://github.com/glample/tagger"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/corpora",
		URL:          []string{"https://github.com/corposaurus/corpora"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nrc",
		URL:          []string{"https://github.com/icarpa-tblab/nrc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/upsetr",
		URL:          []string{"https://github.com/hms-dbmi/upsetr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pathos",
		URL:          []string{"https://github.com/papenfusslab/pathos"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ritornello",
		URL:          []string{"https://github.com/klugerlab/ritornello"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/reoa",
		URL:          []string{"https://github.com/pathint/reoa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/btyper",
		URL:          []string{"https://github.com/lmc297/btyper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bminerimportancebacillus",
		URL:          []string{"https://github.com/lmc297/bminerimportancebacillus"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/psi-blastexb",
		URL:          []string{"https://github.com/kyungtaeklim/psi-blastexb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/omaat",
		URL:          []string{"https://github.com/biorack/omaat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/isomir2function",
		URL:          []string{"https://github.com/347033139/isomir2function"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mtipatchsearch",
		URL:          []string{"https://github.com/mtipatchsearch"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/jepegmix2",
		URL:          []string{"https://github.com/chatzinakos/jepegmix2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/section-sort",
		URL:          []string{"https://github.com/saalfeldlab/section-sort"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/z-spacing",
		URL:          []string{"https://github.com/saalfeldlab/z-spacing"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/z-spacing-spark",
		URL:          []string{"https://github.com/saalfeldlab/z-spacing-spark"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scode",
		URL:          []string{"https://github.com/hmatsu1226/scode"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bsevaluationremotesceneir",
		URL:          []string{"https://github.com/jerryyaogl/bsevaluationremotesceneir"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sieve",
		URL:          []string{"https://github.com/nkullman/sieve"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sailfish-cir",
		URL:          []string{"https://github.com/zerodel/sailfish-cir"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/drseq2",
		URL:          []string{"https://github.com/chengchenzhao/drseq2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/genesurv",
		URL:          []string{"https://github.com/selcukorkmaz/genesurv"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rtm-gwas",
		URL:          []string{"https://github.com/njau-sri/rtm-gwas"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dangertrack",
		URL:          []string{"https://github.com/dcgenomics/dangertrack"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rmfilter",
		URL:          []string{"https://github.com/hitbc/rmfilter"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tfce-mediation",
		URL:          []string{"https://github.com/trislett/tfce_mediation"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/imagebox",
		URL:          []string{"https://github.com/sbu-bmi/imagebox"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nandb",
		URL:          []string{"https://github.com/rorynolan/nandb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/md-task",
		URL:          []string{"https://github.com/rubi-za/md-task"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biocontainers",
		URL:          []string{"https://github.com/biocontainers"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/remilo",
		URL:          []string{"https://github.com/songc001/remilo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/friends-and-family",
		URL:          []string{"https://github.com/deondejager/friends-and-family"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tissue-enrichment-test",
		URL:          []string{"https://github.com/gregor-mendel-institute/tissue-enrichment-test"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/docks",
		URL:          []string{"https://github.com/shamir-lab/docks"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/piil",
		URL:          []string{"https://github.com/behroozt/piil"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/james-e-barrett",
		URL:          []string{"https://github.com/james-e-barrett"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/isoem2",
		URL:          []string{"https://github.com/mandricigor/isoem2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pdhs-svm",
		URL:          []string{"https://github.com/shanxinzhang/pdhs-svm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lasagne4bio",
		URL:          []string{"https://github.com/vanessajurtz/lasagne4bio"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kwip",
		URL:          []string{"https://github.com/kdmurray91/kwip"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tsnad",
		URL:          []string{"https://github.com/jiujiezz/tsnad"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/edlib",
		URL:          []string{"https://github.com/martinsos/edlib"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/crispr-dav",
		URL:          []string{"https://github.com/pinetree1/crispr-dav"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/doritool",
		URL:          []string{"https://github.com/doritool"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ols-client",
		URL:          []string{"https://github.com/pride-utilities/ols-client"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ols-dialog",
		URL:          []string{"https://github.com/pride-toolsuite/ols-dialog"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/posigene",
		URL:          []string{"https://github.com/gengit/posigene"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kmer-ssr",
		URL:          []string{"https://github.com/ridgelab/kmer-ssr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/16gt",
		URL:          []string{"https://github.com/aquaskyline/16gt"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/unicycler",
		URL:          []string{"https://github.com/rrwick/unicycler"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/secureqc",
		URL:          []string{"https://github.com/acs6610987/secureqc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/inspiired",
		URL:          []string{"https://github.com/bushmanlab/inspiired"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gene-is",
		URL:          []string{"https://github.com/g100dkfz/gene-is"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/self-blm",
		URL:          []string{"https://github.com/gist-csbl/self-blm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mdts",
		URL:          []string{"https://github.com/tsudalab/mdts"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/spacescanner",
		URL:          []string{"https://github.com/atiselsts/spacescanner"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pad",
		URL:          []string{"https://github.com/vccri/pad"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/flashpca",
		URL:          []string{"https://github.com/gabraham/flashpca"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/factorizations",
		URL:          []string{"https://github.com/combine-lab/salmon/tree/factorizations"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/genometester4",
		URL:          []string{"https://github.com/bioinfo-ut/genometester4"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/node.dating",
		URL:          []string{"https://github.com/brj1/node.dating"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dectdec",
		URL:          []string{"https://github.com/microdect/dectdec"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/graphspace",
		URL:          []string{"https://github.com/murali-group/graphspace"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vcftotree-3.0.0",
		URL:          []string{"https://github.com/duoduoo/vcftotree_3.0.0"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/figs",
		URL:          []string{"https://github.com/thakar-lab/figs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hla-bind",
		URL:          []string{"https://github.com/uci-cbcl/hla-bind"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/aascatterplot",
		URL:          []string{"https://github.com/whittakerlab/aascatterplot"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cnanalysis450k",
		URL:          []string{"https://github.com/mknoll/cnanalysis450k"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cidr",
		URL:          []string{"https://github.com/vccri/cidr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepgo",
		URL:          []string{"https://github.com/bio-ontology-research-group/deepgo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bamclipper",
		URL:          []string{"https://github.com/tommyau/bamclipper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/spm",
		URL:          []string{"https://github.com/izhbannikov/spm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rnaseq-protocol",
		URL:          []string{"https://github.com/mwschmid/rnaseq_protocol"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bib",
		URL:          []string{"https://github.com/probic/bib"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/index",
		URL:          []string{"https://github.com/a-mir/index"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fastbma",
		URL:          []string{"https://github.com/lhhunghimself/fastbma"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/data-example",
		URL:          []string{"https://github.com/jhroth/data-example"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/celltrans",
		URL:          []string{"https://github.com/tbuder/celltrans"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rmsi",
		URL:          []string{"https://github.com/prafols/rmsi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/jointrn",
		URL:          []string{"https://github.com/louyinxia/jointrn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ovarcall",
		URL:          []string{"https://github.com/takumorizo/ovarcall"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/panoptes",
		URL:          []string{"https://github.com/cggh/panoptes"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/chmannot",
		URL:          []string{"https://github.com/cskyan/chmannot"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sincerities",
		URL:          []string{"https://github.com/cabsel/sincerities"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mapseq",
		URL:          []string{"https://github.com/jfmrod/mapseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/d2sbin",
		URL:          []string{"https://github.com/kunwangkun/d2sbin"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fasp",
		URL:          []string{"https://github.com/vtcbil/fasp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/visid",
		URL:          []string{"https://github.com/gabora/visid"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cafe",
		URL:          []string{"https://github.com/younglululu/cafe"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kinship-privacy",
		URL:          []string{"https://github.com/tastanlab/kinship-privacy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/appscangeo",
		URL:          []string{"https://github.com/stantonlabdartmouth/appscangeo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/powsimr",
		URL:          []string{"https://github.com/bvieth/powsimr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hmeta-d",
		URL:          []string{"https://github.com/smfleming/hmeta-d"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cgheliparm",
		URL:          []string{"https://github.com/ifaust83/cgheliparm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mistic",
		URL:          []string{"https://github.com/iric-soft/mistic"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/compomics-utilities",
		URL:          []string{"https://github.com/compomics/compomics-utilities"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/armsd",
		URL:          []string{"https://github.com/armsd/armsd"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/zerosum",
		URL:          []string{"https://github.com/rehbergt/zerosum"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phenopolis",
		URL:          []string{"https://github.com/phenopolis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/miergolf",
		URL:          []string{"https://github.com/gibbsdavidl/miergolf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/genocore",
		URL:          []string{"https://github.com/lovemun/genocore"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/red-ml",
		URL:          []string{"https://github.com/bgired/red-ml"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lw-fqzip2",
		URL:          []string{"https://github.com/zhuzxlab/lw-fqzip2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/isown",
		URL:          []string{"https://github.com/ikalatskaya/isown"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rarefaction",
		URL:          []string{"https://github.com/hildebra/rarefaction"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/uncertainty-components",
		URL:          []string{"https://github.com/c-zrv/uncertainty_components"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/4d-nucleome-analysis-toolbox",
		URL:          []string{"https://github.com/laseaman/4d_nucleome_analysis_toolbox"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/popfba",
		URL:          []string{"https://github.com/bimib-disco/popfba"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/st-pipeline",
		URL:          []string{"https://github.com/spatialtranscriptomicsresearch/st_pipeline"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rmats-dvr",
		URL:          []string{"https://github.com/xinglab/rmats-dvr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/csar",
		URL:          []string{"https://github.com/ablab-nthu/csar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cognate",
		URL:          []string{"https://github.com/zfmk/cognate"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cfdnapattern",
		URL:          []string{"https://github.com/opengene/cfdnapattern"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dwcox",
		URL:          []string{"https://github.com/jinfengxiao/dwcox"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/oncosimul",
		URL:          []string{"https://github.com/rdiaz02/oncosimul"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/snvphyl-galaxy",
		URL:          []string{"https://github.com/phac-nml/snvphyl-galaxy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/clonevol",
		URL:          []string{"https://github.com/chrismaherlab/clonevol"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/coretracker",
		URL:          []string{"https://github.com/udem-lbit/coretracker"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/esmo",
		URL:          []string{"https://github.com/xiongli2016/esmo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/riblast",
		URL:          []string{"https://github.com/fukunagatsu/riblast"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/smartr",
		URL:          []string{"https://github.com/transmart/smartr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/aspc",
		URL:          []string{"https://github.com/jasonzyx/aspc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mousetrap-os",
		URL:          []string{"https://github.com/pascalkieslich/mousetrap-os"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kbranches",
		URL:          []string{"https://github.com/theislab/kbranches"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ploidyngs",
		URL:          []string{"https://github.com/diriano/ploidyngs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/canvas",
		URL:          []string{"https://github.com/illumina/canvas"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/umis",
		URL:          []string{"https://github.com/vals/umis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kollector",
		URL:          []string{"https://github.com/bcgsc/kollector"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deploid",
		URL:          []string{"https://github.com/mcveanlab/deploid"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deploid-r",
		URL:          []string{"https://github.com/mcveanlab/deploid-r"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pydream",
		URL:          []string{"https://github.com/lolab-vu/pydream"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/iht.jl",
		URL:          []string{"https://github.com/klkeys/iht.jl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kgs",
		URL:          []string{"https://github.com/excitedstates/kgs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/proteomodlr",
		URL:          []string{"https://github.com/kentsisresearchgroup/proteomodlr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/insilico-subtyping",
		URL:          []string{"https://github.com/superphy/insilico-subtyping"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/jdinac",
		URL:          []string{"https://github.com/jijiadong/jdinac"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ggtk",
		URL:          []string{"https://github.com/paulbible/ggtk"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pyscestoolbox",
		URL:          []string{"https://github.com/pysces/pyscestoolbox"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepfr",
		URL:          []string{"https://github.com/zhujianwei31415/deepfr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/wish",
		URL:          []string{"https://github.com/soedinglab/wish"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pixie",
		URL:          []string{"https://github.com/pnnl-comp-mass-spec/pixie"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ims-informed-library",
		URL:          []string{"https://github.com/pnnl-comp-mass-spec/ims-informed-library"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/chem-preview",
		URL:          []string{"https://github.com/wallerlab/chem-preview"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/binning-refiner",
		URL:          []string{"https://github.com/songweizhi/binning_refiner"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dptdt",
		URL:          []string{"https://github.com/mwgrassgreen/dptdt"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/episurg",
		URL:          []string{"https://github.com/episurg/episurg"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/explobatch",
		URL:          []string{"https://github.com/syspremed/explobatch"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pamir",
		URL:          []string{"https://github.com/vpc-ccg/pamir"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metacomp",
		URL:          []string{"https://github.com/pzhaipku/metacomp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hhconpred",
		URL:          []string{"https://github.com/dpxiong/hhconpred"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cupsoda",
		URL:          []string{"https://github.com/aresio/cupsoda"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/toner",
		URL:          []string{"https://github.com/pavitakae/toner"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/osprey-refactor",
		URL:          []string{"https://github.com/donaldlab/osprey_refactor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cytocompare",
		URL:          []string{"https://github.com/tchitchek-lab/cytocompare"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/geno-diver",
		URL:          []string{"https://github.com/jeremyhoward/geno-diver"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vcnet",
		URL:          []string{"https://github.com/wangzengmiao/vcnet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metagen",
		URL:          []string{"https://github.com/bioalgs/metagen"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ethseq",
		URL:          []string{"https://github.com/aromanel/ethseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/grnmf",
		URL:          []string{"https://github.com/xiao-hn/grnmf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/trove",
		URL:          []string{"https://github.com/trove2017/trove"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/suget",
		URL:          []string{"https://github.com/udem-lbit/suget"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scrat",
		URL:          []string{"https://github.com/zji90/scrat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pycgtool",
		URL:          []string{"https://github.com/jag1g13/pycgtool"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepcytof",
		URL:          []string{"https://github.com/klugerlab/deepcytof"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/salivaprint",
		URL:          []string{"https://github.com/salivatec/salivaprint"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kfits",
		URL:          []string{"https://github.com/odedrim/kfits"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/runbng",
		URL:          []string{"https://github.com/appliedbioinformatics/runbng"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/wuetal2017",
		URL:          []string{"https://github.com/cartwrightlab/wuetal2017"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mrcnn-scene-recognition",
		URL:          []string{"https://github.com/wanglimin/mrcnn-scene-recognition"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vari",
		URL:          []string{"https://github.com/cosmo-team/cosmo/tree/vari"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/consensx.itk.ppke.hu",
		URL:          []string{"https://github.com/ppke-bioinf/consensx.itk.ppke.hu"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metasra-pipeline",
		URL:          []string{"https://github.com/deweylab/metasra-pipeline"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/webgivi",
		URL:          []string{"https://github.com/sunliang3361/webgivi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/viennarna",
		URL:          []string{"https://github.com/viennarna"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/stells2",
		URL:          []string{"https://github.com/yufengwudcs/stells2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/minimds",
		URL:          []string{"https://github.com/seqcode/minimds"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ibrel",
		URL:          []string{"https://github.com/andrelamurias/ibrel"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phylonetworks.jl",
		URL:          []string{"https://github.com/crsl4/phylonetworks.jl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nuclitrack",
		URL:          []string{"https://github.com/samocooper/nuclitrack"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/agennt",
		URL:          []string{"https://github.com/kandlinf/agennt"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lrcstats",
		URL:          []string{"https://github.com/cchauve/lrcstats"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hcsa",
		URL:          []string{"https://github.com/kekegg/hcsa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mathbiocu",
		URL:          []string{"https://github.com/mathbiocu"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/twisst",
		URL:          []string{"https://github.com/simonhmartin/twisst"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/refbool",
		URL:          []string{"https://github.com/saschajung/refbool"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/spark-vs",
		URL:          []string{"https://github.com/mcapuccini/spark-vs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/polyploid-genotyping",
		URL:          []string{"https://github.com/pblischak/polyploid-genotyping"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bpp",
		URL:          []string{"https://github.com/zhqingit/bpp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/igess",
		URL:          []string{"https://github.com/daviddaigithub/igess"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ringo",
		URL:          []string{"https://github.com/pedrofeijao/ringo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kronos",
		URL:          []string{"https://github.com/jtaghiyar/kronos"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ananke",
		URL:          []string{"https://github.com/beiko-lab/ananke"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/malax",
		URL:          []string{"https://github.com/omerwe/malax"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/transcriptome-assembly-pipeline",
		URL:          []string{"https://github.com/biomendi/transcriptome-assembly-pipeline"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/earomatic",
		URL:          []string{"https://github.com/michal-brylinski/earomatic"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/eigen-scripts",
		URL:          []string{"https://github.com/danbuchan/eigen_scripts"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/reparation",
		URL:          []string{"https://github.com/biobix/reparation"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/icepop",
		URL:          []string{"https://github.com/ewijaya/icepop"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/survivalcausaltree",
		URL:          []string{"https://github.com/weijiazhang24/survivalcausaltree"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/g-msr",
		URL:          []string{"https://github.com/pcdslab/g-msr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sparseiso",
		URL:          []string{"https://github.com/henryxushi/sparseiso"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pep",
		URL:          []string{"https://github.com/ma-compbio/pep"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/halc",
		URL:          []string{"https://github.com/lanl001/halc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lyve-set",
		URL:          []string{"https://github.com/lskatz/lyve-set"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/virgena",
		URL:          []string{"https://github.com/gfedonin/virgena"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepbound",
		URL:          []string{"https://github.com/realbigws/deepbound"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phandango",
		URL:          []string{"https://github.com/jameshadfield/phandango"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lc50",
		URL:          []string{"https://github.com/ahproctor/lc50"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dtwscore",
		URL:          []string{"https://github.com/xiaoxiaoxier/dtwscore"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/proteosign",
		URL:          []string{"https://github.com/yorgodillo/proteosign"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/methrafo",
		URL:          []string{"https://github.com/phoenixding/methrafo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phyd3",
		URL:          []string{"https://github.com/vibbits/phyd3"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pipeline-for-isoseq",
		URL:          []string{"https://github.com/nextomics/pipeline-for-isoseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/visjs2jupyter",
		URL:          []string{"https://github.com/ucsd-ccbb/visjs2jupyter"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/iphloc-es",
		URL:          []string{"https://github.com/swakkhar/iphloc-es"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/homblocks",
		URL:          []string{"https://github.com/fenghen360/homblocks"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/homblocks",
		URL:          []string{"https://github.com/fenghen360/homblocks"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/htsvis",
		URL:          []string{"https://github.com/boutroslab/htsvis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pcasa",
		URL:          []string{"https://github.com/atfrank/pcasa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/subcellular-localization",
		URL:          []string{"https://github.com/jjalmagro/subcellular_localization"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mecat",
		URL:          []string{"https://github.com/xiaochuanle/mecat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/integrative",
		URL:          []string{"https://github.com/xqwen/integrative"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sbc",
		URL:          []string{"https://github.com/ashar799/sbc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dpuc2",
		URL:          []string{"https://github.com/alexviiia/dpuc2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kaptive",
		URL:          []string{"https://github.com/katholt/kaptive"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dart",
		URL:          []string{"https://github.com/hsinnan75/dart"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/trs",
		URL:          []string{"https://github.com/giorgosminas/trs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ichseg",
		URL:          []string{"https://github.com/muschellij2/ichseg"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bigld",
		URL:          []string{"https://github.com/sunnyeesl/bigld"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deblur",
		URL:          []string{"https://github.com/biocore/deblur"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/h-blast",
		URL:          []string{"https://github.com/yeyke/h-blast"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cgdm",
		URL:          []string{"https://github.com/evanswang/cgdm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phylo-io",
		URL:          []string{"https://github.com/dessimozlab/phylo-io"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/panviz",
		URL:          []string{"https://github.com/thomasp85/panviz"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/iform",
		URL:          []string{"https://github.com/wenjiegroup/iform"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/snpfilt",
		URL:          []string{"https://github.com/lanlab/snpfilt"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fishplot",
		URL:          []string{"https://github.com/chrisamiller/fishplot"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sequence-search-tool-for-antimicrobial-resistance-sstar-",
		URL:          []string{"https://github.com/tomdeman-bio/sequence-search-tool-for-antimicrobial-resistance-sstar-"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cymer",
		URL:          []string{"https://github.com/bmuchmore/cymer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ractip",
		URL:          []string{"https://github.com/satoken/ractip"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vashr",
		URL:          []string{"https://github.com/mengyin/vashr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/combiningdependentpvalues",
		URL:          []string{"https://github.com/ilyalab/combiningdependentpvalues"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sampl5-dc-surface-empirical",
		URL:          []string{"https://github.com/diogomart/sampl5-dc-surface-empirical"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/romulus",
		URL:          []string{"https://github.com/ajank/romulus"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/predixcan",
		URL:          []string{"https://github.com/hakyimlab/predixcan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ggv",
		URL:          []string{"https://github.com/novembrelab/ggv"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ggv-api",
		URL:          []string{"https://github.com/novembrelab/ggv-api"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tappm-cli",
		URL:          []string{"https://github.com/davecao/tappm_cli"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ribodiff",
		URL:          []string{"https://github.com/ratschlab/ribodiff"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/peakxus",
		URL:          []string{"https://github.com/hartonen/peakxus"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mosbat",
		URL:          []string{"https://github.com/csglab/mosbat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bnrr",
		URL:          []string{"https://github.com/wangronglu/bnrr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/retrotransposons-spread",
		URL:          []string{"https://github.com/sergemoulin/retrotransposons-spread"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/obcs",
		URL:          []string{"https://github.com/obcs/obcs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phage",
		URL:          []string{"https://github.com/jglamine/phage"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/transphylo",
		URL:          []string{"https://github.com/xavierdidelot/transphylo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/topological-measures-wide-analysis",
		URL:          []string{"https://github.com/biomedical-cybernetics/topological_measures_wide_analysis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/besst",
		URL:          []string{"https://github.com/ksahlin/besst"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nest",
		URL:          []string{"https://github.com/thomastaus/nest"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/snpclust",
		URL:          []string{"https://github.com/thomaschln/snpclust"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bioattribution",
		URL:          []string{"https://github.com/wchangmitre/bioattribution"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mycofier",
		URL:          []string{"https://github.com/ldelgado-serrano/mycofier"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/medici",
		URL:          []string{"https://github.com/cooperlab/medici"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/blasst",
		URL:          []string{"https://github.com/vislab/blasst"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sonar",
		URL:          []string{"https://github.com/scharch/sonar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rd-analyzer",
		URL:          []string{"https://github.com/xiaeryu/rd-analyzer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/parasor",
		URL:          []string{"https://github.com/carushi/parasor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/parfit",
		URL:          []string{"https://github.com/fzahari/parfit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sigmod",
		URL:          []string{"https://github.com/yuanlongliu/sigmod"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phyinformr",
		URL:          []string{"https://github.com/carolinafishes/phyinformr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/altre",
		URL:          []string{"https://github.com/mathelab/altre"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/peakerror",
		URL:          []string{"https://github.com/tdhock/peakerror"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/seqlm",
		URL:          []string{"https://github.com/raivokolde/seqlm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fsva",
		URL:          []string{"https://github.com/topwood91/fsva"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/para-suite",
		URL:          []string{"https://github.com/akloetgen/para-suite"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/para-suite-aligner",
		URL:          []string{"https://github.com/akloetgen/para-suite_aligner"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mefit",
		URL:          []string{"https://github.com/nisheth/mefit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rtfbs-db",
		URL:          []string{"https://github.com/danko-lab/rtfbs_db"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hmm-based-method",
		URL:          []string{"https://github.com/rabbitpei/hmm_based-method"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/atlases",
		URL:          []string{"https://github.com/cobralab/atlases"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/msparsm",
		URL:          []string{"https://github.com/cmontemuino/msparsm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deeppicker-python",
		URL:          []string{"https://github.com/nejyeah/deeppicker-python"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tabsat",
		URL:          []string{"https://github.com/tadkeys/tabsat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rapmap",
		URL:          []string{"https://github.com/combine-lab/rapmap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/asdpex",
		URL:          []string{"https://github.com/charite/asdpex"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metapred2cs",
		URL:          []string{"https://github.com/martinjvickers/metapred2cs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/immboost",
		URL:          []string{"https://github.com/weiyangedward/immboost"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mispu",
		URL:          []string{"https://github.com/chongwu-biostat/mispu"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/riborex",
		URL:          []string{"https://github.com/smithlabcode/riborex"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metapalette",
		URL:          []string{"https://github.com/dkoslicki/metapalette"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/axonseg",
		URL:          []string{"https://github.com/neuropoly/axonseg"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lsue",
		URL:          []string{"https://github.com/ekffar/lsue"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/liquid",
		URL:          []string{"https://github.com/pnnl-comp-mass-spec/liquid"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bayescat",
		URL:          []string{"https://github.com/heejungshim/bayescat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/imgui-electron-packages",
		URL:          []string{"https://github.com/jaredgk/imgui-electron-packages"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cocacola",
		URL:          []string{"https://github.com/younglululu/cocacola"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/recycler",
		URL:          []string{"https://github.com/shamir-lab/recycler"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/browsevcf",
		URL:          []string{"https://github.com/bsgoxford/browsevcf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biomsef",
		URL:          []string{"https://github.com/agjacome/biomsef"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pathscore",
		URL:          []string{"https://github.com/sggaffney/pathscore"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/4dassign",
		URL:          []string{"https://github.com/thomasexner/4dassign"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/orfm",
		URL:          []string{"https://github.com/wwood/orfm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cudagsea",
		URL:          []string{"https://github.com/gravitino/cudagsea"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/descs-standalone",
		URL:          []string{"https://github.com/mantczak/descs-standalone"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/read-split-run",
		URL:          []string{"https://github.com/xuric/read-split-run"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/chemps2",
		URL:          []string{"https://github.com/sebwouters/chemps2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gbgfa",
		URL:          []string{"https://github.com/olganikolova/gbgfa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gap",
		URL:          []string{"https://github.com/anand-bhaskar/gap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gwalpha",
		URL:          []string{"https://github.com/aflevel/gwalpha"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metlab",
		URL:          []string{"https://github.com/norling/metlab"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biopartsdb",
		URL:          []string{"https://github.com/baderzone/biopartsdb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/eval",
		URL:          []string{"https://github.com/genecodeq/eval"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cell-maps",
		URL:          []string{"https://github.com/opencb/cell-maps"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/chipsad",
		URL:          []string{"https://github.com/silviamicroarray/chipsad"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kat",
		URL:          []string{"https://github.com/tgac/kat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cuclark",
		URL:          []string{"https://github.com/funatiq/cuclark"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/come",
		URL:          []string{"https://github.com/lulab/come"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/powerfit",
		URL:          []string{"https://github.com/haddocking/powerfit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tscan",
		URL:          []string{"https://github.com/zji90/tscan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/marvin",
		URL:          []string{"https://github.com/illumina/marvin"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nssrfbinary",
		URL:          []string{"https://github.com/zhangjiaobxy/nssrfbinary"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nssrfpackage",
		URL:          []string{"https://github.com/zhangjiaobxy/nssrfpackage"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/doraemon",
		URL:          []string{"https://github.com/abcgrp/doraemon"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mvbc",
		URL:          []string{"https://github.com/javonsun/mvbc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/reactiondecoder",
		URL:          []string{"https://github.com/asad/reactiondecoder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/chia-pet2",
		URL:          []string{"https://github.com/guipengli/chia-pet2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/prediction-prostate-surveillance",
		URL:          []string{"https://github.com/rycoley/prediction-prostate-surveillance"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fast-lmm",
		URL:          []string{"https://github.com/microsoftgenomics/fast-lmm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hipred",
		URL:          []string{"https://github.com/hashihab/hipred"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kent",
		URL:          []string{"https://github.com/ucscgenomebrowser/kent"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kreationsupplementary",
		URL:          []string{"https://github.com/schulzlab/kreationsupplementary"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lncscore",
		URL:          []string{"https://github.com/wglab/lncscore"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/region-templates",
		URL:          []string{"https://github.com/sbu-bmi/region-templates"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/methyliftover",
		URL:          []string{"https://github.com/christensen-lab-dartmouth/methyliftover"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sbd-evaluation",
		URL:          []string{"https://github.com/drgriffis/sbd-evaluation"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/h-popg",
		URL:          []string{"https://github.com/minzhuxie/h-popg"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/diseaseextract",
		URL:          []string{"https://github.com/tcrnbioinformatics/diseaseextract"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nanocall",
		URL:          []string{"https://github.com/mateidavid/nanocall"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pirnapredictor",
		URL:          []string{"https://github.com/zw9977129/pirnapredictor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lims",
		URL:          []string{"https://github.com/jianwei-zhang/lims"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/annopeak",
		URL:          []string{"https://github.com/xingtang2014/annopeak"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ac-pca",
		URL:          []string{"https://github.com/linzx06/ac-pca"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/microbe-id",
		URL:          []string{"https://github.com/grunwaldlab/microbe-id"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/castin",
		URL:          []string{"https://github.com/tmd-gpat/castin"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phenotype-prediction-pipeline",
		URL:          []string{"https://github.com/clabuzze/phenotype-prediction-pipeline"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dpetstep",
		URL:          []string{"https://github.com/crossschmidtlein/dpetstep"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepsnvminer",
		URL:          []string{"https://github.com/mattmattmattmatt/deepsnvminer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/variantbam",
		URL:          []string{"https://github.com/jwalabroad/variantbam"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mid-correct",
		URL:          []string{"https://github.com/seliv55/mid_correct"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dmdtoolkit",
		URL:          []string{"https://github.com/zhoujp111/dmdtoolkit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/targetspecificgatksequencingpipeline",
		URL:          []string{"https://github.com/christopher-gillies/targetspecificgatksequencingpipeline"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/feelnc",
		URL:          []string{"https://github.com/tderrien/feelnc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sbml-diff",
		URL:          []string{"https://github.com/jamesscottbrown/sbml-diff"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fastrfs",
		URL:          []string{"https://github.com/pranjalv123/fastrfs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/boss",
		URL:          []string{"https://github.com/bioinfomaticscsu/boss"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bcgtree",
		URL:          []string{"https://github.com/iimog/bcgtree"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/architect",
		URL:          []string{"https://github.com/kuleshov/architect"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/seqenv",
		URL:          []string{"https://github.com/xapple/seqenv"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/seqlib",
		URL:          []string{"https://github.com/walaj/seqlib"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cnara",
		URL:          []string{"https://github.com/baudisgroup/cnara"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/valor",
		URL:          []string{"https://github.com/bilkentcompgen/valor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/prefersim",
		URL:          []string{"https://github.com/lohmuellerlab/prefersim"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phraider",
		URL:          []string{"https://github.com/karroje/phraider"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/disambiguate",
		URL:          []string{"https://github.com/astrazeneca-ngs/disambiguate"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dada2",
		URL:          []string{"https://github.com/benjjneb/dada2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/smrt-pipeline",
		URL:          []string{"https://github.com/ndliberial/smrt_pipeline"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cellprofiler-analyst",
		URL:          []string{"https://github.com/cellprofiler/cellprofiler-analyst"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/diffsplicing",
		URL:          []string{"https://github.com/probic/diffsplicing"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/anaquin",
		URL:          []string{"https://github.com/student-t/anaquin"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kdsnp",
		URL:          []string{"https://github.com/hirotosaigo/kdsnp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bdss",
		URL:          []string{"https://github.com/feltus/bdss"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/samfire",
		URL:          []string{"https://github.com/cjri/samfire"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/paladin",
		URL:          []string{"https://github.com/twestbrookunh/paladin"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pysm",
		URL:          []string{"https://github.com/alexandrovteam/pysm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/methylflow",
		URL:          []string{"https://github.com/hcorrada/methylflow"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bprmeth",
		URL:          []string{"https://github.com/andreaskapou/bprmeth"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/asciigenome",
		URL:          []string{"https://github.com/dariober/asciigenome"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/zorro",
		URL:          []string{"https://github.com/cina/zorro"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rna-interactions-benchmark",
		URL:          []string{"https://github.com/ucancompbio/rna_interactions_benchmark"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/m2c",
		URL:          []string{"https://github.com/ilyalab/m2c"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vcf-kit",
		URL:          []string{"https://github.com/andersenlab/vcf-kit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/eqp-cluster",
		URL:          []string{"https://github.com/novartis/eqp-cluster"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/assemblytics",
		URL:          []string{"https://github.com/marianattestad/assemblytics"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dinosaur",
		URL:          []string{"https://github.com/fickludd/dinosaur"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/foldatlas",
		URL:          []string{"https://github.com/mnori/foldatlas"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/snowball",
		URL:          []string{"https://github.com/hzi-bifo/snowball"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nvt",
		URL:          []string{"https://github.com/edert/nvt"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gsa4mirna",
		URL:          []string{"https://github.com/dmontaner-papers/gsa4mirna"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/iwgs",
		URL:          []string{"https://github.com/zhouxiaofan1983/iwgs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kinmut2",
		URL:          []string{"https://github.com/rbbt-workflows/kinmut2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/variant-validation",
		URL:          []string{"https://github.com/suyashss/variant_validation"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scoary",
		URL:          []string{"https://github.com/admiralenola/scoary"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phenomescape",
		URL:          []string{"https://github.com/soulj/phenomescape"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/solarius",
		URL:          []string{"https://github.com/ugcd/solarius"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/arnapipe",
		URL:          []string{"https://github.com/hudsonalpha/arnapipe"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dagitty",
		URL:          []string{"https://github.com/jtextor/dagitty"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/caspo",
		URL:          []string{"https://github.com/bioasp/caspo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sativa",
		URL:          []string{"https://github.com/amkozlov/sativa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gneiss",
		URL:          []string{"https://github.com/biocore/gneiss"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lamsa",
		URL:          []string{"https://github.com/hitbc/lamsa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metaphinder",
		URL:          []string{"https://github.com/vanessajurtz/metaphinder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cocos",
		URL:          []string{"https://github.com/butkiem/cocos"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/awol-mrf",
		URL:          []string{"https://github.com/cobralab/awol-mrf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sa-ssr",
		URL:          []string{"https://github.com/ridgelab/sa-ssr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mash",
		URL:          []string{"https://github.com/marbl/mash"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/saspy",
		URL:          []string{"https://github.com/emblsaxs/saspy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/shiny-revecor",
		URL:          []string{"https://github.com/yiluheihei/shiny-revecor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/epiabc",
		URL:          []string{"https://github.com/kypraios/epiabc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hotspot",
		URL:          []string{"https://github.com/evolbioinf/hotspot"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pathdeseq",
		URL:          []string{"https://github.com/malathisidona/pathdeseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/microbiomegwas",
		URL:          []string{"https://github.com/lsncibb/microbiomegwas"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pgen-genomicvariations-workflow",
		URL:          []string{"https://github.com/pegasus-isi/pgen-genomicvariations-workflow"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biojs-io-biom",
		URL:          []string{"https://github.com/molbiodiv/biojs-io-biom"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mz.unity",
		URL:          []string{"https://github.com/nathaniel-mahieu/mz.unity"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/megan-ce",
		URL:          []string{"https://github.com/danielhuson/megan-ce"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/egcpi",
		URL:          []string{"https://github.com/hetiantian1985/egcpi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/alignerboost",
		URL:          []string{"https://github.com/grice-lab/alignerboost"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bpsc",
		URL:          []string{"https://github.com/nghiavtr/bpsc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/drug-drug-interaction",
		URL:          []string{"https://github.com/zw9977129/drug-drug-interaction"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/remap",
		URL:          []string{"https://github.com/hansaimlim/remap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vdjviz",
		URL:          []string{"https://github.com/antigenomics/vdjviz"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/medusa",
		URL:          []string{"https://github.com/marinkaz/medusa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/assocplots",
		URL:          []string{"https://github.com/khramts/assocplots"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nsg",
		URL:          []string{"https://github.com/samhykim/nsg"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/discoal",
		URL:          []string{"https://github.com/kern-lab/discoal"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/historian",
		URL:          []string{"https://github.com/evoldoers/historian"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/trig",
		URL:          []string{"https://github.com/tllab/trig"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metaplotr",
		URL:          []string{"https://github.com/olarerin/metaplotr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phyfoo",
		URL:          []string{"https://github.com/mgledi/phyfoo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/multiwayclassification",
		URL:          []string{"https://github.com/lockef/multiwayclassification"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/means",
		URL:          []string{"https://github.com/theosysbio/means"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cafe-plugin",
		URL:          []string{"https://github.com/huiliucode/cafe_plugin"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/memo",
		URL:          []string{"https://github.com/memo-toolbox/memo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dnapi",
		URL:          []string{"https://github.com/jnktsj/dnapi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bereta",
		URL:          []string{"https://github.com/kms1041/bereta"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hybpiper",
		URL:          []string{"https://github.com/mossmatters/hybpiper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sexy",
		URL:          []string{"https://github.com/christopher-amos-lab/sexy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/network-classifier",
		URL:          []string{"https://github.com/adrinjalali/network-classifier"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/molyso",
		URL:          []string{"https://github.com/modsim/molyso"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/genepopedit",
		URL:          []string{"https://github.com/rystanley/genepopedit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scellsupplementary",
		URL:          []string{"https://github.com/diazlab/scellsupplementary"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/repurpose",
		URL:          []string{"https://github.com/emreg00/repurpose"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ashr",
		URL:          []string{"https://github.com/stephens999/ashr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pyhla",
		URL:          []string{"https://github.com/felixfan/pyhla"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/partitionfinder",
		URL:          []string{"https://github.com/brettc/partitionfinder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ismags",
		URL:          []string{"https://github.com/biointec/ismags"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kover",
		URL:          []string{"https://github.com/aldro61/kover"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mswaldhmp",
		URL:          []string{"https://github.com/mafed/mswaldhmp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biomartr",
		URL:          []string{"https://github.com/hajkd/biomartr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vsearch",
		URL:          []string{"https://github.com/torognes/vsearch"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/almostsignificant",
		URL:          []string{"https://github.com/bartongroup/almostsignificant"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/markyt",
		URL:          []string{"https://github.com/sing-group/markyt"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/zerone",
		URL:          []string{"https://github.com/nanakiksc/zerone"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/organelle-pba",
		URL:          []string{"https://github.com/aubombarely/organelle_pba"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ogcleaner",
		URL:          []string{"https://github.com/byucsl/ogcleaner"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cobratoolbox",
		URL:          []string{"https://github.com/opencobra/cobratoolbox"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pychemia",
		URL:          []string{"https://github.com/materialsdiscovery/pychemia"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/boiler",
		URL:          []string{"https://github.com/jpritt/boiler"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rambl",
		URL:          []string{"https://github.com/homopolymer/rambl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mixedperfectphylogeny",
		URL:          []string{"https://github.com/alexandrutomescu/mixedperfectphylogeny"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/neat-genreads",
		URL:          []string{"https://github.com/zstephens/neat-genreads"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/simbac",
		URL:          []string{"https://github.com/tbrown91/simbac"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gtrac",
		URL:          []string{"https://github.com/kedartatwawadi/gtrac"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/atminter",
		URL:          []string{"https://github.com/csb5/atminter"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mltm",
		URL:          []string{"https://github.com/hsoleimani/mltm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/impre",
		URL:          []string{"https://github.com/zhangwei2015/impre"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/seeksv",
		URL:          []string{"https://github.com/qkl871118/seeksv"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ngs-bits",
		URL:          []string{"https://github.com/imgag/ngs-bits"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/swinger",
		URL:          []string{"https://github.com/yuma248/swinger"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/agin",
		URL:          []string{"https://github.com/hacone/agin"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vcfanno",
		URL:          []string{"https://github.com/brentp/vcfanno"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/glint",
		URL:          []string{"https://github.com/cozygene/glint"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/changlab",
		URL:          []string{"https://github.com/jefftc/changlab"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/seqkit",
		URL:          []string{"https://github.com/shenwei356/seqkit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dockq",
		URL:          []string{"https://github.com/bjornwallner/dockq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/discomark",
		URL:          []string{"https://github.com/hdetering/discomark"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/quin",
		URL:          []string{"https://github.com/ucarlab/quin"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pedla",
		URL:          []string{"https://github.com/wenjiegroup/pedla"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/qsimscan",
		URL:          []string{"https://github.com/abadona/qsimscan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/emdn",
		URL:          []string{"https://github.com/william0701/emdn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/debruijn-motif",
		URL:          []string{"https://github.com/hirvolt1/debruijn-motif"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metashot",
		URL:          []string{"https://github.com/bfosso/metashot"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/falco",
		URL:          []string{"https://github.com/vccri/falco"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mtgipick",
		URL:          []string{"https://github.com/bioinfo0706/mtgipick"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/microscope",
		URL:          []string{"https://github.com/bohdan-khomtchouk/microscope"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phyx",
		URL:          []string{"https://github.com/fephyfofum/phyx"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/admixture-graph",
		URL:          []string{"https://github.com/mailund/admixture_graph"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dmi",
		URL:          []string{"https://github.com/hongyu-miao/dmi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/horvathlab-ngs",
		URL:          []string{"https://github.com/horvathlab/ngs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deltabs",
		URL:          []string{"https://github.com/ucancompbio/deltabs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bidifuse",
		URL:          []string{"https://github.com/jandetrez/bidifuse"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nmfem",
		URL:          []string{"https://github.com/lanagarmire/nmfem"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/psynteract",
		URL:          []string{"https://github.com/psynteract"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/xmrf",
		URL:          []string{"https://github.com/zhandong/xmrf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lightassembler",
		URL:          []string{"https://github.com/sarael-metwally/lightassembler"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/getisdmr",
		URL:          []string{"https://github.com/dmu-lilab/getisdmr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sprint-race",
		URL:          []string{"https://github.com/watera427/sprint-race"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/beacon-network-inference",
		URL:          []string{"https://github.com/beaconprojectatvirginiatech/beacon_network_inference"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cytometry-clustering-comparison",
		URL:          []string{"https://github.com/lmweber/cytometry-clustering-comparison"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/proxl-web-app",
		URL:          []string{"https://github.com/yeastrc/proxl-web-app"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bionetgen",
		URL:          []string{"https://github.com/ruleworld/bionetgen"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lir",
		URL:          []string{"https://github.com/borisvassilev/lir"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/conformationalchange",
		URL:          []string{"https://github.com/bgsu-rna/conformationalchange"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/grabb",
		URL:          []string{"https://github.com/b-brankovics/grabb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/panget",
		URL:          []string{"https://github.com/pangetv1/panget"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/integrate-neo",
		URL:          []string{"https://github.com/chrismaherlab/integrate-neo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/osdb",
		URL:          []string{"https://github.com/stuchalk/osdb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hima",
		URL:          []string{"https://github.com/yinanzheng/hima"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/grail",
		URL:          []string{"https://github.com/theanswerisfortytwo/grail"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/worm-single-cell-data-and-codes",
		URL:          []string{"https://github.com/xthuang226/worm_single_cell_data_and_codes"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dysprog",
		URL:          []string{"https://github.com/bebeklab/dysprog"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nblda",
		URL:          []string{"https://github.com/yangchadam/nblda"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/detime",
		URL:          []string{"https://github.com/manchesterbioinference/detime"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cdpc",
		URL:          []string{"https://github.com/micheleallegra/cdpc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pyboolnet",
		URL:          []string{"https://github.com/hklarner/pyboolnet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/modict",
		URL:          []string{"https://github.com/ibrahimtanyalcin/modict"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/batch-ge",
		URL:          []string{"https://github.com/woutersteyaert/batch-ge"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gsp",
		URL:          []string{"https://github.com/bioinfogenome/gsp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lmmel-mir-miner",
		URL:          []string{"https://github.com/uomsystemsbiology/lmmel-mir-miner"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dnadu",
		URL:          []string{"https://github.com/cyinbox/dnadu"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nichesimulation",
		URL:          []string{"https://github.com/anwala/nichesimulation"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fait",
		URL:          []string{"https://github.com/godziklab/fait"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rnaifold",
		URL:          []string{"https://github.com/clotelab/rnaifold"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/non-parametric-algorithm-to-isolate-chunks-in-response-sequences",
		URL:          []string{"https://github.com/artipago/non-parametric-algorithm-to-isolate-chunks-in-response-sequences"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/brides",
		URL:          []string{"https://github.com/etiennelord/brides"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ispp",
		URL:          []string{"https://github.com/jacobhsu35/ispp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/isotopiclabelling",
		URL:          []string{"https://github.com/ruggeroferrazza/isotopiclabelling"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/estitics",
		URL:          []string{"https://github.com/michaellenz/estitics"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cpm-cytoscape",
		URL:          []string{"https://github.com/davcem/cpm-cytoscape"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/walt",
		URL:          []string{"https://github.com/smithlabcode/walt"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/circompara",
		URL:          []string{"https://github.com/egaffo/circompara"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mapcomp",
		URL:          []string{"https://github.com/enormandeau/mapcomp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/debga",
		URL:          []string{"https://github.com/hitbc/debga"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/startapp",
		URL:          []string{"https://github.com/jminnier/startapp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/selam",
		URL:          []string{"https://github.com/russcd/selam"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gc",
		URL:          []string{"https://github.com/biocryst/gc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/qpath",
		URL:          []string{"https://github.com/higex/qpath"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mcsc-decontamination",
		URL:          []string{"https://github.com/lafond-lapalmej/mcsc_decontamination"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/forwardgenomics",
		URL:          []string{"https://github.com/hillerlab/forwardgenomics"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/evoclust3d",
		URL:          []string{"https://github.com/doxeylab/evoclust3d"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/compound-eye-simulator",
		URL:          []string{"https://github.com/eeza-csic/compound-eye-simulator"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kcmbt-mt",
		URL:          []string{"https://github.com/abdullah009/kcmbt_mt"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/diseasediscovery",
		URL:          []string{"https://github.com/wrrzag/diseasediscovery"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fluff",
		URL:          []string{"https://github.com/simonvh/fluff"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lollipops",
		URL:          []string{"https://github.com/pbnjay/lollipops"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metpeak",
		URL:          []string{"https://github.com/compgenomics/metpeak"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sharpvisu",
		URL:          []string{"https://github.com/andronovl/sharpvisu"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pbrowse",
		URL:          []string{"https://github.com/vccri/pbrowse"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/semanticsco",
		URL:          []string{"https://github.com/usplssb/semanticsco"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/c-deva",
		URL:          []string{"https://github.com/cici333/c-deva"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sparkbwa",
		URL:          []string{"https://github.com/citiususc/sparkbwa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/asafe",
		URL:          []string{"https://github.com/biostatqian/asafe"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/imsa-a",
		URL:          []string{"https://github.com/jeremycoxbmi/imsa-a"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/varmatch",
		URL:          []string{"https://github.com/medvedevgroup/varmatch"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ensembler",
		URL:          []string{"https://github.com/choderalab/ensembler"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/motion-detection",
		URL:          []string{"https://github.com/thachnguyen/motion_detection"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/colormap",
		URL:          []string{"https://github.com/sfu-compbio/colormap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lattice-metage",
		URL:          []string{"https://github.com/lattclus/lattice-metage"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/spadevizr",
		URL:          []string{"https://github.com/tchitchek-lab/spadevizr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/graph2pro",
		URL:          []string{"https://github.com/col-iu/graph2pro"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/spectreok",
		URL:          []string{"https://github.com/mills-lab/spectreok"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/enscat",
		URL:          []string{"https://github.com/jlp2duke/enscat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tadtool",
		URL:          []string{"https://github.com/vaquerizaslab/tadtool"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gsalightning",
		URL:          []string{"https://github.com/billyhw/gsalightning"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/muxstep",
		URL:          []string{"https://github.com/petarv-/muxstep"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kinconform",
		URL:          []string{"https://github.com/esbg/kinconform"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/spiddor",
		URL:          []string{"https://github.com/spiddor/spiddor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pubmedportable",
		URL:          []string{"https://github.com/kerstendoering/pubmedportable"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/magellan",
		URL:          []string{"https://github.com/sristov/magellan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/wtchg-rg",
		URL:          []string{"https://github.com/popitsch/wtchg-rg"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/popstr",
		URL:          []string{"https://github.com/decodegenetics/popstr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mc-cbn",
		URL:          []string{"https://github.com/cbg-ethz/mc-cbn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/angular-ripleys-k",
		URL:          []string{"https://github.com/rubypeters/angular-ripleys-k"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/traitar",
		URL:          []string{"https://github.com/hzi-bifo/traitar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ymap",
		URL:          []string{"https://github.com/csb-kul/ymap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dacum",
		URL:          []string{"https://github.com/haddocking/dacum"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nanomark",
		URL:          []string{"https://github.com/kkrizanovic/nanomark"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/qc-dmet",
		URL:          []string{"https://github.com/sebwouters/qc-dmet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/argon",
		URL:          []string{"https://github.com/pierpal/argon"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metafast",
		URL:          []string{"https://github.com/ctlab/metafast"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/upho",
		URL:          []string{"https://github.com/ballesterus/upho"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fiqt",
		URL:          []string{"https://github.com/bacanusa/fiqt"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sparc",
		URL:          []string{"https://github.com/yechengxi/sparc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/epitracer",
		URL:          []string{"https://github.com/narmada26/epitracer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sdeap",
		URL:          []string{"https://github.com/ewyang089/sdeap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cosmomvpa",
		URL:          []string{"https://github.com/cosmomvpa/cosmomvpa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scphaser",
		URL:          []string{"https://github.com/edsgard/scphaser"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/optimum-rlda",
		URL:          []string{"https://github.com/danik0411/optimum-rlda"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/protasr",
		URL:          []string{"https://github.com/miguelarenas/protasr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/transmart-xnat-connector",
		URL:          []string{"https://github.com/sh107/transmart-xnat-connector"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/toxreporter",
		URL:          []string{"https://github.com/mgosink/toxreporter"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cogena",
		URL:          []string{"https://github.com/zhilongjia/cogena"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/goldilocks",
		URL:          []string{"https://github.com/samstudio8/goldilocks"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tw2",
		URL:          []string{"https://github.com/alekseyenko/tw2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mbirw",
		URL:          []string{"https://github.com//bioinfomaticscsu/mbirw"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/flexicoclusteringpackage",
		URL:          []string{"https://github.com/biotoolsleeds/flexicoclusteringpackage"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mhtrajectoryr",
		URL:          []string{"https://github.com/masedki/mhtrajectoryr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/azahar",
		URL:          []string{"https://github.com/agustinaarroyuelo/azahar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/apostl",
		URL:          []string{"https://github.com/bornea/apostl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vdjer",
		URL:          []string{"https://github.com/mozack/vdjer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/boostgapfill",
		URL:          []string{"https://github.com/tolutola/boostgapfill"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/daisy",
		URL:          []string{"https://github.com/ktrappe/daisy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/batchqc",
		URL:          []string{"https://github.com/mani2012/batchqc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/blend4php",
		URL:          []string{"https://github.com/galaxyproject/blend4php"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/case-bbb-prediction-data",
		URL:          []string{"https://github.com/bioinformatics-gao/case-bbb-prediction-data"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dt-all",
		URL:          []string{"https://github.com/nwpusy/dt_all"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gnuplot-iostream",
		URL:          []string{"https://github.com/dstahlke/gnuplot-iostream"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mimeanto",
		URL:          []string{"https://github.com/maureensmith/mimeanto"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/angsd-wrapper",
		URL:          []string{"https://github.com/mojaveazure/angsd-wrapper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/positionalmi",
		URL:          []string{"https://github.com/kassonlab/positionalmi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/jester",
		URL:          []string{"https://github.com/brielin/jester"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deuterater",
		URL:          []string{"https://github.com/jc-price/deuterater"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/conpair",
		URL:          []string{"https://github.com/nygenome/conpair"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/zibr",
		URL:          []string{"https://github.com/chvlyl/zibr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/medcrawler",
		URL:          []string{"https://github.com/a-hel/medcrawler"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/xgsa",
		URL:          []string{"https://github.com/vccri/xgsa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/network-snps",
		URL:          []string{"https://github.com/ncbi-hackathons/network_snps"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mptp",
		URL:          []string{"https://github.com/pas-kapli/mptp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sword",
		URL:          []string{"https://github.com/rvaser/sword"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/drugcombination",
		URL:          []string{"https://github.com/yiyiliu1/drugcombination"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/asap",
		URL:          []string{"https://github.com/ddofer/asap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/crowdnodnet",
		URL:          []string{"https://github.com/bioops/crowdnodnet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fqtools",
		URL:          []string{"https://github.com/alastair-droop/fqtools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dace",
		URL:          []string{"https://github.com/tinglab/dace"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metacrispr",
		URL:          []string{"https://github.com/hangelwen/metacrispr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lrssl",
		URL:          []string{"https://github.com/liangxujun/lrssl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biocemid",
		URL:          []string{"https://github.com/ferhtaydn/biocemid"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pong",
		URL:          []string{"https://github.com/abehr/pong"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/coala",
		URL:          []string{"https://github.com/statgenlmu/coala"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biocaddie2016mayodata",
		URL:          []string{"https://github.com/yanshanwang/biocaddie2016mayodata"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/egad",
		URL:          []string{"https://github.com/sarbal/egad"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rail-dbgap",
		URL:          []string{"https://github.com/nellore/rail-dbgap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/osfp",
		URL:          []string{"https://github.com/chaninn/osfp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kdetrees",
		URL:          []string{"https://github.com/grady/kdetrees"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/thapmix",
		URL:          []string{"https://github.com/illumina/thapmix"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/geneevolve",
		URL:          []string{"https://github.com/rtahmasbi/geneevolve"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tqec",
		URL:          []string{"https://github.com/alexandrupaler/tqec"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scoup",
		URL:          []string{"https://github.com/hmatsu1226/scoup"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ngsf-hmm",
		URL:          []string{"https://github.com/fgvieira/ngsf-hmm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metaspark",
		URL:          []string{"https://github.com/zhouweiyg/metaspark"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gptime",
		URL:          []string{"https://github.com/statisticalbiotechnology/gptime"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/twopaco",
		URL:          []string{"https://github.com/medvedevgroup/twopaco"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sepp",
		URL:          []string{"https://github.com/smirarab/sepp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pymod",
		URL:          []string{"https://github.com/pymodproject/pymod"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/haveyouswappedyoursamples",
		URL:          []string{"https://github.com/papenfusslab/haveyouswappedyoursamples"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lobstahs",
		URL:          []string{"https://github.com/vanmooylipidomics/lobstahs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phocos",
		URL:          []string{"https://github.com/altschulerwu-lab/phocos"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metabolicframework",
		URL:          []string{"https://github.com/ibalaur/metabolicframework"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/variancecomponenttest.jl",
		URL:          []string{"https://github.com/tao-hu/variancecomponenttest.jl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/treescaper",
		URL:          []string{"https://github.com/whuang08/treescaper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cloudphylo",
		URL:          []string{"https://github.com/xingjianxu/cloudphylo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lsdbs-mpi",
		URL:          []string{"https://github.com/turbo0628/lsdbs-mpi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/silkslider",
		URL:          []string{"https://github.com/wwood/silkslider"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rvtests",
		URL:          []string{"https://github.com/zhanxw/rvtests"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/grridgecodata",
		URL:          []string{"https://github.com/markvdwiel/grridgecodata"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sbvb0",
		URL:          []string{"https://github.com/mperezenciso/sbvb0"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/proper-extension",
		URL:          []string{"https://github.com/agaye/proper_extension"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/afresh",
		URL:          []string{"https://github.com/tparidae/afresh"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sofia",
		URL:          []string{"https://github.com/childsish/sofia"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/psmvc",
		URL:          []string{"https://github.com/oyl-cityu/psmvc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/debwt",
		URL:          []string{"https://github.com/hitbc/debwt"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/debwtcontact",
		URL:          []string{"https://github.com/dixianzhu/debwtcontact"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/issues",
		URL:          []string{"https://github.com/griffithlab/genvisr/issues"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/omblast",
		URL:          []string{"https://github.com/aldenleung/omblast"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mirmeta",
		URL:          []string{"https://github.com/xuelab/mirmeta"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pypanda",
		URL:          []string{"https://github.com/davidvi/pypanda"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/srbreak",
		URL:          []string{"https://github.com/hoangtn/srbreak"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nimfa",
		URL:          []string{"https://github.com/ccshao/nimfa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/keyregulatorygenes",
		URL:          []string{"https://github.com/maryamnazarieh/keyregulatorygenes"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/spcapp",
		URL:          []string{"https://github.com/ul-research-support/spcapp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tfit",
		URL:          []string{"https://github.com/azofeifa/tfit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/srf",
		URL:          []string{"https://github.com/rankmatrixfactorisation/srf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/video-database",
		URL:          []string{"https://github.com/remega/video_database"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/princess-opensource",
		URL:          []string{"https://github.com/achenfengb/princess_opensource"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metacycle",
		URL:          []string{"https://github.com/gangwug/metacycle"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/micado",
		URL:          []string{"https://github.com/cbib/micado"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ark",
		URL:          []string{"https://github.com/the-ark-informatics/ark"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fatslim",
		URL:          []string{"https://github.com/fatslim/fatslim"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bcalm",
		URL:          []string{"https://github.com/gatb/bcalm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sine-scan",
		URL:          []string{"https://github.com/maohlzj/sine_scan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/libcsam",
		URL:          []string{"https://github.com/rcanovas/libcsam"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mega",
		URL:          []string{"https://github.com/ciccalab/mega"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/w4cseq",
		URL:          []string{"https://github.com/wglab/w4cseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rentplus",
		URL:          []string{"https://github.com/sajadmirzaei/rentplus"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/chilin",
		URL:          []string{"https://github.com/cfce/chilin"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lacytools",
		URL:          []string{"https://github.com/tarskin/lacytools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rabifier",
		URL:          []string{"https://github.com/evocell/rabifier"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sctcrseq",
		URL:          []string{"https://github.com/elementolab/sctcrseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/spectraph",
		URL:          []string{"https://github.com/yakhinigroup/spectraph"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/physca",
		URL:          []string{"https://github.com/nluhmann/physca"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sbol-validator",
		URL:          []string{"https://github.com/synbiodex/sbol-validator"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fsda",
		URL:          []string{"https://github.com/compbio-uoft/fsda"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/paamia2015",
		URL:          []string{"https://github.com/fabkury/paamia2015"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mfsba",
		URL:          []string{"https://github.com/lsaravia/mfsba"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/d-gex",
		URL:          []string{"https://github.com/uci-cbcl/d-gex"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nplb",
		URL:          []string{"https://github.com/computationalbiology/nplb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/wyrm",
		URL:          []string{"https://github.com/bbci/wyrm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/popins",
		URL:          []string{"https://github.com/bkehr/popins"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/quandico",
		URL:          []string{"https://github.com/reineckef/quandico"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sowhat",
		URL:          []string{"https://github.com/josephryan/sowhat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rocketship",
		URL:          []string{"https://github.com/petmri/rocketship"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/infmod3dgen",
		URL:          []string{"https://github.com/wangsy11/infmod3dgen"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/roma",
		URL:          []string{"https://github.com/sysbio-curie/roma"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cssscl",
		URL:          []string{"https://github.com/oicr-ibc/cssscl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tangram",
		URL:          []string{"https://github.com/jiantao/tangram"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bstools",
		URL:          []string{"https://github.com/xfwang/bstools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pirsupplementary",
		URL:          []string{"https://github.com/ypriverol/pirsupplementary"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bigbwa",
		URL:          []string{"https://github.com/citiususc/bigbwa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/description-extractor",
		URL:          []string{"https://github.com/mutalyzer/description-extractor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pagal",
		URL:          []string{"https://github.com/sanchak/pagal"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/quasar",
		URL:          []string{"https://github.com/piquelab/quasar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/b-nem",
		URL:          []string{"https://github.com/martinfxp/b-nem"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/uniqtag",
		URL:          []string{"https://github.com/sjackman/uniqtag"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/uniqtag-paper",
		URL:          []string{"https://github.com/sjackman/uniqtag-paper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pathrings",
		URL:          []string{"https://github.com/ivcl/pathrings"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biomaj2galaxy",
		URL:          []string{"https://github.com/genouest/biomaj2galaxy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vt",
		URL:          []string{"https://github.com/atks/vt"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/seqbuster",
		URL:          []string{"https://github.com/lpantano/seqbuster"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lbeep",
		URL:          []string{"https://github.com/brsaran/lbeep"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/testsetbias",
		URL:          []string{"https://github.com/prpatil/testsetbias"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hitselect",
		URL:          []string{"https://github.com/diazlab/hitselect"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/neatfreq",
		URL:          []string{"https://github.com/bioh4x/neatfreq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pileup.js",
		URL:          []string{"https://github.com/hammerlab/pileup.js"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/neat",
		URL:          []string{"https://github.com/pschorderet/neat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kltepigenome",
		URL:          []string{"https://github.com/pmb59/kltepigenome"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/qvz",
		URL:          []string{"https://github.com/mikelhernaez/qvz"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ptw",
		URL:          []string{"https://github.com/rwehrens/ptw"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cryostage",
		URL:          []string{"https://github.com/pechano/cryostage"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/qreadselector",
		URL:          []string{"https://github.com/moneycat/qreadselector"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nxtrim",
		URL:          []string{"https://github.com/sequencing/nxtrim"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bgt",
		URL:          []string{"https://github.com/lh3/bgt"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cosmos",
		URL:          []string{"https://github.com/bioprocessdesignlab/cosmos"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/raser",
		URL:          []string{"https://github.com/jaegyoonahn/raser"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/comida",
		URL:          []string{"https://github.com/borenstein-lab/comida"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/massytools",
		URL:          []string{"https://github.com/tarskin/massytools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pidgin",
		URL:          []string{"https://github.com/lhm30/pidgin"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rmscca",
		URL:          []string{"https://github.com/hardin47/rmscca"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dalliance",
		URL:          []string{"https://github.com/dasmoth/dalliance"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ldserv",
		URL:          []string{"https://github.com/dasmoth/ldserv"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cometa",
		URL:          []string{"https://github.com/jkawulok/cometa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/starcode",
		URL:          []string{"https://github.com/gui11aume/starcode"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/magus",
		URL:          []string{"https://github.com/institut-de-genomique/magus"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mucor",
		URL:          []string{"https://github.com/blachlylab/mucor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/emringer",
		URL:          []string{"https://github.com/fraser-lab/emringer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cld",
		URL:          []string{"https://github.com/boutroslab/cld"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ionmf",
		URL:          []string{"https://github.com/mstrazar/ionmf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/proq-scripts",
		URL:          []string{"https://github.com/bjornwallner/proq_scripts"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/demix-q",
		URL:          []string{"https://github.com/userbz/demix-q"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/emsar",
		URL:          []string{"https://github.com/parklab/emsar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/regularized-manova",
		URL:          []string{"https://github.com/jaspere/regularized-manova"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/alview",
		URL:          []string{"https://github.com/ncip/alview"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/crowd-cid-relex",
		URL:          []string{"https://github.com/sulab/crowd_cid_relex"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/micro",
		URL:          []string{"https://github.com/carrineblank/micro"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/atma",
		URL:          []string{"https://github.com/rwalecki/atma"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/aksmooth",
		URL:          []string{"https://github.com/junfang/aksmooth"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/flaimapper",
		URL:          []string{"https://github.com/yhoogstrate/flaimapper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/arraylasso",
		URL:          []string{"https://github.com/adam-sam-brown/arraylasso"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dreamtools",
		URL:          []string{"https://github.com/dreamtools/dreamtools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/uproc",
		URL:          []string{"https://github.com/gobics/uproc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/django-chicp",
		URL:          []string{"https://github.com/d-i-l/django-chicp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/glycoseq",
		URL:          []string{"https://github.com/chpaul/glycoseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bcftools",
		URL:          []string{"https://github.com/samtools/bcftools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/reago",
		URL:          []string{"https://github.com/chengyuan/reago"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/flowsom",
		URL:          []string{"https://github.com/sofievg/flowsom"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/parastagonospora-nodorum-sn15",
		URL:          []string{"https://github.com/robsyme/parastagonospora_nodorum_sn15"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nextflu",
		URL:          []string{"https://github.com/blab/nextflu"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pimms",
		URL:          []string{"https://github.com/adac-uon/pimms"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/annocript",
		URL:          []string{"https://github.com/frankmusacchia/annocript"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/repeatsoaker",
		URL:          []string{"https://github.com/mdozmorov/repeatsoaker"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nanook",
		URL:          []string{"https://github.com/tgac/nanook"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fmrat",
		URL:          []string{"https://github.com/hggm-lim/fmrat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/oncotator",
		URL:          []string{"https://github.com/broadinstitute/oncotator"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/shesisplus",
		URL:          []string{"https://github.com/celaoforever/shesisplus"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/boa",
		URL:          []string{"https://github.com/idoerg/boa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pride-toolsuite",
		URL:          []string{"https://github.com/pride-toolsuite"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fermikit",
		URL:          []string{"https://github.com/lh3/fermikit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rampart",
		URL:          []string{"https://github.com/tgac/rampart"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mst-minor-allele-caller",
		URL:          []string{"https://github.com/zalmanv/mst_minor_allele_caller"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/d3m",
		URL:          []string{"https://github.com/ymatts/d3m"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scanindel",
		URL:          []string{"https://github.com/cauyrd/scanindel"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ledpred",
		URL:          []string{"https://github.com/aitgon/ledpred"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fusenet",
		URL:          []string{"https://github.com/marinkaz/fusenet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lb-impute",
		URL:          []string{"https://github.com/dellaporta-laboratory/lb-impute"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepnet-rbp",
		URL:          []string{"https://github.com/thucombio/deepnet-rbp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/openms",
		URL:          []string{"https://github.com/openms/openms"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/leap",
		URL:          []string{"https://github.com/omerwe/leap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/niftylink",
		URL:          []string{"https://github.com/niftk/niftylink"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/spladder",
		URL:          []string{"https://github.com/ratschlab/spladder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bio-tradis",
		URL:          []string{"https://github.com/sanger-pathogens/bio-tradis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sf-abc",
		URL:          []string{"https://github.com/stevenhwu/sf-abc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/juzhe1120-matlab-software",
		URL:          []string{"https://github.com/juzhe1120/matlab_software"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bc5cidtask",
		URL:          []string{"https://github.com/jhnlp/bc5cidtask"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hpg-pore",
		URL:          []string{"https://github.com/opencb/hpg-pore"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fbdenovo",
		URL:          []string{"https://github.com/strejcem/fbdenovo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ikap",
		URL:          []string{"https://github.com/marcel-mischnik/ikap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pycellerator",
		URL:          []string{"https://github.com/biomathman/pycellerator"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/datasink",
		URL:          []string{"https://github.com/shwhalen/datasink"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/filtus",
		URL:          []string{"https://github.com/magnusdv/filtus"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mendelian",
		URL:          []string{"https://github.com/bartbroeckx/mendelian"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rcount",
		URL:          []string{"https://github.com/mwschmid/rcount"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/manta",
		URL:          []string{"https://github.com/illumina/manta"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/triocnv",
		URL:          []string{"https://github.com/yongzhuang/triocnv"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/meshsim",
		URL:          []string{"https://github.com/jingzhou2015/meshsim"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/binning",
		URL:          []string{"https://github.com/smirarab/binning"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hbm",
		URL:          []string{"https://github.com/yolish/hbm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/spectrogene",
		URL:          []string{"https://github.com/fenderglass/spectrogene"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/npreader",
		URL:          []string{"https://github.com/mdcao/npreader"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kpal",
		URL:          []string{"https://github.com/lumc/kpal"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gips",
		URL:          []string{"https://github.com/synergy-zju/gips"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cloudforest",
		URL:          []string{"https://github.com/ilyalab/cloudforest"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fuma",
		URL:          []string{"https://github.com/erasmusmc-bioinformatics/fuma"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pez",
		URL:          []string{"https://github.com/willpearse/pez"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bam-abs",
		URL:          []string{"https://github.com/zhanglabvt/bam_abs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kipper",
		URL:          []string{"https://github.com/public-health-bioinformatics/kipper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/versioned-data",
		URL:          []string{"https://github.com/public-health-bioinformatics/versioned_data"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/george",
		URL:          []string{"https://github.com/jcapelladesto/george"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mpicbg-scicomp/snakemake-workflows",
		URL:          []string{"https://github.com/mpicbg-scicomp/snakemake-workflows"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/parties",
		URL:          []string{"https://github.com/oarnaiz/parties"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dive",
		URL:          []string{"https://github.com/ahmed-fakhry/dive"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/equation-generator",
		URL:          []string{"https://github.com/inemelckenbeeck/equation-generator"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/graphlet-naming",
		URL:          []string{"https://github.com/inemelckenbeeck/graphlet-naming"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/admixem",
		URL:          []string{"https://github.com/melop/admixem"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sputnik",
		URL:          []string{"https://github.com/sysbio-bioinf/sputnik"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/danq",
		URL:          []string{"https://github.com/uci-cbcl/danq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/jms",
		URL:          []string{"https://github.com/rubi-za/jms"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cmonkey2",
		URL:          []string{"https://github.com/baliga-lab/cmonkey2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hitime",
		URL:          []string{"https://github.com/bjpop/hitime"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/greenwoodlab",
		URL:          []string{"https://github.com/greenwoodlab"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/genefu",
		URL:          []string{"https://github.com/bhklab/genefu"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/shaoweinuaa",
		URL:          []string{"https://github.com/shaoweinuaa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pathwaymatrix",
		URL:          []string{"https://github.com/creativecodinglab/pathwaymatrix"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sisrs",
		URL:          []string{"https://github.com/rachelss/sisrs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gwas-scripts",
		URL:          []string{"https://github.com/jonicoleman/gwas_scripts"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pyvolve",
		URL:          []string{"https://github.com/sjspielman/pyvolve"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tcga-rnaseq-clinical",
		URL:          []string{"https://github.com/srp33/tcga_rnaseq_clinical"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/molti",
		URL:          []string{"https://github.com/gilles-didier/molti"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/binm",
		URL:          []string{"https://github.com/zhangxf-ccnu/binm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/aci",
		URL:          []string{"https://github.com/subangstrom/aci"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cudampf",
		URL:          []string{"https://github.com/super-hippo/cudampf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hitwalker2",
		URL:          []string{"https://github.com/biodev/hitwalker2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/domex",
		URL:          []string{"https://github.com/xuezhidong/domex"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/foreseqs",
		URL:          []string{"https://github.com/ddarriba/foreseqs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/imagebasedtranscriptomics",
		URL:          []string{"https://github.com/pelkmanslab/imagebasedtranscriptomics"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/matrixconverter",
		URL:          []string{"https://github.com/gburleigh/matrixconverter"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/harvest",
		URL:          []string{"https://github.com/marbl/harvest"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/red",
		URL:          []string{"https://github.com/redetector/red"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/oefinder",
		URL:          []string{"https://github.com/lengning/oefinder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fcs-point-correlator",
		URL:          []string{"https://github.com/dwaithe/fcs_point_correlator"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/locusexplorer",
		URL:          []string{"https://github.com/oncogenetics/locusexplorer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/adapt-mix",
		URL:          []string{"https://github.com/dpark27/adapt_mix"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gbs-snp-crop",
		URL:          []string{"https://github.com/halelab/gbs-snp-crop"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fizzy",
		URL:          []string{"https://github.com/eesi/fizzy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/defume",
		URL:          []string{"https://github.com/evdh0/defume"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mixclone",
		URL:          []string{"https://github.com/uci-cbcl/mixclone"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/andi",
		URL:          []string{"https://github.com/evolbioinf/andi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sumac",
		URL:          []string{"https://github.com/wf8/sumac"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mmr",
		URL:          []string{"https://github.com/ratschlab/mmr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/inmf",
		URL:          []string{"https://github.com/yangzi4/inmf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gene-block-evolution",
		URL:          []string{"https://github.com/reamdc1/gene_block_evolution"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sbb",
		URL:          []string{"https://github.com/flekschas/sbb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/confidence-information-ontology",
		URL:          []string{"https://github.com/bgeedb/confidence-information-ontology"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/confidence-information-ontology",
		URL:          []string{"https://raw.githubusercontent.com/bgeedb/confidence-information-ontology"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/optnetaligncpp",
		URL:          []string{"https://github.com/crclark/optnetaligncpp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mmseqs",
		URL:          []string{"https://github.com/soedinglab/mmseqs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/wgfast",
		URL:          []string{"https://github.com/jasonsahl/wgfast"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/limix",
		URL:          []string{"https://github.com/pmbio/limix"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/jcvi",
		URL:          []string{"https://github.com/tanghaibao/jcvi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/genome-bin-tools",
		URL:          []string{"https://github.com/kbseah/genome-bin-tools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/clusdca",
		URL:          []string{"https://github.com/wangshenguiuc/clusdca"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sifter-t",
		URL:          []string{"https://github.com/dcasbioinfo/sifter-t"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/minimap",
		URL:          []string{"https://github.com/lh3/minimap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/miniasm",
		URL:          []string{"https://github.com/lh3/miniasm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dycluster",
		URL:          []string{"https://github.com/emhanna/dycluster"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hedicim",
		URL:          []string{"https://github.com/fabkury/hedicim"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/clockstarg",
		URL:          []string{"https://github.com/sebastianduchene/clockstarg"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/qsubsec",
		URL:          []string{"https://github.com/alastair-droop/qsubsec"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/screenmasker",
		URL:          []string{"https://github.com/paulorlov/screenmasker"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vasp",
		URL:          []string{"https://github.com/mattmattmattmatt/vasp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/yoc",
		URL:          []string{"https://github.com/ruricaru/yoc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ghost-tree",
		URL:          []string{"https://github.com/jtfouquier/ghost-tree"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gqt",
		URL:          []string{"https://github.com/ryanlayer/gqt"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/epga2",
		URL:          []string{"https://github.com/bioinfomaticscsu/epga2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/genesthycan",
		URL:          []string{"https://github.com/chengkun-wu/genesthycan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ipo",
		URL:          []string{"https://github.com/glibiseller/ipo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/spaced-seeds-for-metagenomics",
		URL:          []string{"https://github.com/gregorykucherov/spaced-seeds-for-metagenomics"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mwfft2",
		URL:          []string{"https://github.com/zmzhang/mwfft2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/omicsintegrator",
		URL:          []string{"https://github.com/fraenkel-lab/omicsintegrator"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cesar",
		URL:          []string{"https://github.com/hillerlab/cesar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kema",
		URL:          []string{"https://github.com/dtuia/kema"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biopartsbuilder",
		URL:          []string{"https://github.com/baderzone/biopartsbuilder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dnn-hmm",
		URL:          []string{"https://github.com/wenjiegroup/dnn-hmm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/oclust",
		URL:          []string{"https://github.com/oscar-franzen/oclust"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/icma",
		URL:          []string{"https://github.com/abi-software-laboratory/icma"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/chippcr",
		URL:          []string{"https://github.com/michbur/chippcr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/moccasin",
		URL:          []string{"https://github.com/sbmlteam/moccasin"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/espresso-research",
		URL:          []string{"https://github.com/espresso-research"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/smallgenometools",
		URL:          []string{"https://github.com/drtrevorbell/smallgenometools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sampled-ancestors",
		URL:          []string{"https://github.com/compevol/sampled-ancestors"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/graphmap",
		URL:          []string{"https://github.com/isovic/graphmap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/chainrank",
		URL:          []string{"https://github.com/atenyi/chainrank"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/toggle",
		URL:          []string{"https://github.com/southgreenplatform/toggle"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/affylumcna",
		URL:          []string{"https://github.com/aplenchop/affylumcna"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/learningdl",
		URL:          []string{"https://github.com/alifahsyamsiyah/learningdl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cnvkit",
		URL:          []string{"https://github.com/etal/cnvkit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/splicing-analysis-pipeline",
		URL:          []string{"https://github.com/pulyakhina/splicing_analysis_pipeline"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ms2gs",
		URL:          []string{"https://github.com/mperezenciso/ms2gs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/networkregularisedcox",
		URL:          []string{"https://github.com/ssnhcom/networkregularisedcox"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/predrad",
		URL:          []string{"https://github.com/phrh/predrad"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/spectacle",
		URL:          []string{"https://github.com/jiminsong/spectacle"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/2kplus2",
		URL:          []string{"https://github.com/danmaclean/2kplus2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vip",
		URL:          []string{"https://github.com/keylabivdc/vip"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gempro",
		URL:          []string{"https://github.com/sbrg/gempro"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gnetlmm",
		URL:          []string{"https://github.com/pmbio/gnetlmm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hic-pro",
		URL:          []string{"https://github.com/nservant/hic-pro"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/som",
		URL:          []string{"https://github.com/bougui505/som"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ontologies",
		URL:          []string{"https://github.com/biointerchange/ontologies"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tailor",
		URL:          []string{"https://github.com/jhhung/tailor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/swarmsight",
		URL:          []string{"https://github.com/justasb/swarmsight"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/seqotron",
		URL:          []string{"https://github.com/4ment/seqotron"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/desir",
		URL:          []string{"https://github.com/stanlazic/desir"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/regshape",
		URL:          []string{"https://github.com/ramseylab/regshape"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/c-hmm",
		URL:          []string{"https://github.com/rslabncbs/c-hmm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metadiff",
		URL:          []string{"https://github.com/jiach/metadiff"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/agplus",
		URL:          []string{"https://github.com/kazumits/agplus"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/acgh-viewer",
		URL:          []string{"https://github.com/fredcommo/acgh_viewer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/maligner",
		URL:          []string{"https://github.com/leemendelowitz/maligner"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/coxboost",
		URL:          []string{"https://github.com/binderh/coxboost"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/causaltrail",
		URL:          []string{"https://github.com/dstoeckel/causaltrail"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bisque",
		URL:          []string{"https://github.com/hyulab/bisque"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/msreduce",
		URL:          []string{"https://github.com/pcdslab/msreduce"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pwtees",
		URL:          []string{"https://github.com/chengkun-wu/pwtees"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/stemcellsva",
		URL:          []string{"https://github.com/andrewejaffe/stemcellsva"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bpipe",
		URL:          []string{"https://github.com/ssadedin/bpipe"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ngs-qcbox",
		URL:          []string{"https://github.com/ceg-icrisat/ngs-qcbox"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/raspberry",
		URL:          []string{"https://github.com/ceg-icrisat/raspberry"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sealer-release",
		URL:          []string{"https://github.com/bcgsc/abyss/tree/sealer-release"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/optimal-seed-solver",
		URL:          []string{"https://github.com/cmu-safari/optimal-seed-solver"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nldmseq",
		URL:          []string{"https://github.com/pugea/nldmseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bin-passing-analyzer",
		URL:          []string{"https://github.com/1particle/bin-passing_analyzer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/abc",
		URL:          []string{"https://github.com/mlupien/abc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tophat-recondition",
		URL:          []string{"https://github.com/cbrueffer/tophat-recondition"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/infinity",
		URL:          []string{"https://github.com/bio-devel/infinity"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/amas",
		URL:          []string{"https://github.com/marekborowiec/amas"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dieterich-lab",
		URL:          []string{"https://github.com/dieterich-lab"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/clamms",
		URL:          []string{"https://github.com/rgcgithub/clamms"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/peak-grouping-alignment",
		URL:          []string{"https://github.com/joewandy/peak-grouping-alignment"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/chopbaicontact:birte.kehr@decode.is",
		URL:          []string{"https://github.com/decodegenetics/chopbaicontact:birte.kehr@decode.is"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cysbar",
		URL:          []string{"https://github.com/ts404/cysbar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/backclip",
		URL:          []string{"https://github.com/phrh/backclip"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bionmf-gpu",
		URL:          []string{"https://github.com/bioinfo-cnb/bionmf-gpu"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phylowgs",
		URL:          []string{"https://github.com/morrislab/phylowgs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gprotpred",
		URL:          []string{"https://github.com/vkostiou/gprotpred"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/geneticthesaurus",
		URL:          []string{"https://github.com/tkonopka/geneticthesaurus"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rgeneticthesaurus",
		URL:          []string{"https://github.com/tkonopka/rgeneticthesaurus"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fast",
		URL:          []string{"https://github.com/tlawrence3/fast"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gubbins",
		URL:          []string{"https://github.com/sanger-pathogens/gubbins"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/splicejumper",
		URL:          []string{"https://github.com/reedwarbler/splicejumper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/interactiverosetta",
		URL:          []string{"https://github.com/schenc3/interactiverosetta"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vizardous",
		URL:          []string{"https://github.com/modsim/vizardous"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bdc",
		URL:          []string{"https://github.com/sharlene/bdc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/orbkit",
		URL:          []string{"https://github.com/orbkit/orbkit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/funtoonorm",
		URL:          []string{"https://github.com/greenwoodlab/funtoonorm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/quantifly",
		URL:          []string{"https://github.com/dwaithe/quantifly"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/computel",
		URL:          []string{"https://github.com/lilit-nersisyan/computel"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/subclpred",
		URL:          []string{"https://github.com/tweirick/subclpred"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/glmvc",
		URL:          []string{"https://github.com/shengqh/glmvc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/simra",
		URL:          []string{"https://github.com/computationalgenomics/simra"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/xander-assembler",
		URL:          []string{"https://github.com/rdpstaff/xander_assembler"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/igb-mi-bundle",
		URL:          []string{"https://github.com/cruiit/igb-mi-bundle"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lsgkm",
		URL:          []string{"https://github.com/dongwon-lee/lsgkm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/xlsearch",
		URL:          []string{"https://github.com/col-iu/xlsearch"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/genomeartist",
		URL:          []string{"https://github.com/genomeartist/genomeartist"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nmrlipids",
		URL:          []string{"https://github.com/nmrlipids"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pyaudioanalysis",
		URL:          []string{"https://github.com/tyiannak/pyaudioanalysis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ttt",
		URL:          []string{"https://github.com/morganrcu/ttt"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tango-final",
		URL:          []string{"https://github.com/bsiranosian/tango_final"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/openhealth",
		URL:          []string{"https://github.com/mathbiol/openhealth"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/msproteomicstools",
		URL:          []string{"https://github.com/msproteomicstools/msproteomicstools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ptmapper",
		URL:          []string{"https://github.com/y-narushima/ptmapper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metica",
		URL:          []string{"https://github.com/daniellyz/metica"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ppsp",
		URL:          []string{"https://github.com/algbioi/ppsp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dispredict-v1.0",
		URL:          []string{"https://github.com/tamjidul/dispredict_v1.0"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rcorrector",
		URL:          []string{"https://github.com/mourisl/rcorrector"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sanefalcon",
		URL:          []string{"https://github.com/rstraver/sanefalcon"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phenominer",
		URL:          []string{"https://github.com/nhcollier/phenominer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/iterativeerrorcorrection",
		URL:          []string{"https://github.com/hillerlab/iterativeerrorcorrection"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ismb15",
		URL:          []string{"https://github.com/lsgh/ismb15"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ark",
		URL:          []string{"https://github.com/dkoslicki/ark"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/amap",
		URL:          []string{"https://github.com/laufercenter/amap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ipbt",
		URL:          []string{"https://github.com/benliemory/ipbt"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ms-data-core-api",
		URL:          []string{"https://github.com/pride-utilities/ms-data-core-api"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/perga",
		URL:          []string{"https://github.com/hitbio/perga"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/alpha-centauri",
		URL:          []string{"https://github.com/volkansevim/alpha-centauri"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/basics",
		URL:          []string{"https://github.com/catavallejos/basics"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/protaeljs",
		URL:          []string{"https://github.com/sanshu/protaeljs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/molgenis-imputation",
		URL:          []string{"https://github.com/molgenis/molgenis-imputation"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/loregic",
		URL:          []string{"https://github.com/gersteinlab/loregic"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hihmm",
		URL:          []string{"https://github.com/kasohn/hihmm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/redundans",
		URL:          []string{"https://github.com/gabaldonlab/redundans"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/stride",
		URL:          []string{"https://github.com/ythuang0522/stride"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/argyle",
		URL:          []string{"https://github.com/andrewparkermorgan/argyle"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hammock",
		URL:          []string{"https://github.com/hammock-dev/hammock"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gpga",
		URL:          []string{"https://github.com/chunbaozhou/gpga"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/superquantnode",
		URL:          []string{"https://github.com/caetera/superquantnode"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/crosslink-viewer",
		URL:          []string{"https://github.com/colin-combe/crosslink-viewer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/modistools",
		URL:          []string{"https://github.com/seantuck12/modistools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dgendoo",
		URL:          []string{"https://github.com/dgendoo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bhklab",
		URL:          []string{"https://github.com/bhklab"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/4pipe4",
		URL:          []string{"https://github.com/stuntspt/4pipe4"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rhat",
		URL:          []string{"https://github.com/hit-bioinformatics/rhat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/reactionflow",
		URL:          []string{"https://github.com/creativecodinglab/reactionflow"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gms",
		URL:          []string{"https://github.com/genome/gms"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hydra",
		URL:          []string{"https://github.com/arq5x/hydra"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/macsyfinder",
		URL:          []string{"https://github.com/gem-pasteur/macsyfinder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/parasail",
		URL:          []string{"https://github.com/jeffdaily/parasail"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/isosegmenter",
		URL:          []string{"https://github.com/bunop/isosegmenter"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/al3c",
		URL:          []string{"https://github.com/ahstram/al3c"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tvb-empirical-data-pipeline",
		URL:          []string{"https://github.com/brainmodes/tvb-empirical-data-pipeline"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/svaseq",
		URL:          []string{"https://github.com/jtleek/svaseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/simphy",
		URL:          []string{"https://github.com/adamallo/simphy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cgmisc",
		URL:          []string{"https://github.com/cgmisc-team/cgmisc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rsss",
		URL:          []string{"https://github.com/boboppie/rsss"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/assign",
		URL:          []string{"https://github.com/wevanjohnson/assign"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/htsint",
		URL:          []string{"https://github.com/ajrichards/htsint"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pypdb",
		URL:          []string{"https://github.com/williamgilpin/pypdb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bs-snper",
		URL:          []string{"https://github.com/hellbelly/bs-snper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/jamss",
		URL:          []string{"https://github.com/optimusmoose/jamss"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hodgkin-huxley-spde",
		URL:          []string{"https://github.com/deristnochda/hodgkin-huxley-spde"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tefma",
		URL:          []string{"https://github.com/mpgerstl/tefma"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/roulattice",
		URL:          []string{"https://github.com/roulattice/roulattice"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ksrepo",
		URL:          []string{"https://github.com/adam-sam-brown/ksrepo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pacemaker-clustering-methods",
		URL:          []string{"https://github.com/sebastianduchene/pacemaker_clustering_methods"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/electronfactors",
		URL:          []string{"https://github.com/simonbiggs/electronfactors"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pepr",
		URL:          []string{"https://github.com/usnistgov/pepr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mira",
		URL:          []string{"https://github.com/mhuttner/mira"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bayexer",
		URL:          []string{"https://github.com/haisiyi/bayexer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/synrio",
		URL:          []string{"https://github.com/nfmc/synrio"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/atsnp",
		URL:          []string{"https://github.com/keleslab/atsnp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phy-mer",
		URL:          []string{"https://github.com/danielnavarrogomez/phy-mer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mecors",
		URL:          []string{"https://github.com/metagenomics/mecors"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ncappc",
		URL:          []string{"https://github.com/cacha0227/ncappc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/grepsearch",
		URL:          []string{"https://github.com/meeibioinformaticscenter/grepsearch"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/carpools",
		URL:          []string{"https://github.com/boutroslab/carpools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/slimmer",
		URL:          []string{"https://github.com/enanomapper/slimmer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/trup",
		URL:          []string{"https://github.com/ruping/trup"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vdjtools",
		URL:          []string{"https://github.com/mikessh/vdjtools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deepdive-genegene-app",
		URL:          []string{"https://github.com/edoughty/deepdive_genegene_app"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/epicseg",
		URL:          []string{"https://github.com/lamortenera/epicseg"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sct",
		URL:          []string{"https://github.com/dww100/sct"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/degps",
		URL:          []string{"https://github.com/ll-lab-mcw/degps"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pathogenportal",
		URL:          []string{"https://github.com/cidvbi/pathogenportal"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/seaflow-cluster",
		URL:          []string{"https://github.com/jhyrkas/seaflow_cluster"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ims-quality",
		URL:          []string{"https://github.com/alexandrovteam/ims_quality"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tn-seqexplorer",
		URL:          []string{"https://github.com/sina-cb/tn-seqexplorer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mason",
		URL:          []string{"https://github.com/yeastrc/mason"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/disvis",
		URL:          []string{"https://github.com/haddocking/disvis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mirseqviewer",
		URL:          []string{"https://github.com/insoo078/mirseqviewer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/meristogram",
		URL:          []string{"https://github.com/waylandm/meristogram"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/multigems",
		URL:          []string{"https://github.com/cui-lab/multigems"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pypedia-server",
		URL:          []string{"https://github.com/kantale/pypedia_server"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pypedia",
		URL:          []string{"https://github.com/kantale/pypedia"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/networkrobustnesstoolbox",
		URL:          []string{"https://github.com/mpgerstl/networkrobustnesstoolbox"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ssusearch",
		URL:          []string{"https://github.com/dib-lab/ssusearch"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pvac-seq",
		URL:          []string{"https://github.com/griffithlab/pvac-seq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/coelho2015-netsdetermination",
		URL:          []string{"https://github.com/luispedro/coelho2015_netsdetermination"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cclasso",
		URL:          []string{"https://github.com/huayingfang/cclasso"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ucscnanopore-data",
		URL:          []string{"https://github.com/ucscnanopore/data"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mirmark",
		URL:          []string{"https://github.com/lanagarmire/mirmark"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kablammo",
		URL:          []string{"https://github.com/jwintersinger/kablammo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cellminercompanion",
		URL:          []string{"https://github.com/pepascuzzi/cellminercompanion"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tnseq",
		URL:          []string{"https://github.com/ffliu/tnseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/discordant",
		URL:          []string{"https://github.com/siskac/discordant"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lncrna-id",
		URL:          []string{"https://github.com/zhangy72/lncrna-id"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/enve",
		URL:          []string{"https://github.com/enve-tools/enve"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ost-pymodules",
		URL:          []string{"https://github.com/njohner/ost_pymodules"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/magic",
		URL:          []string{"https://github.com/chiuyc/magic"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/misfinder",
		URL:          []string{"https://github.com/hitbio/misfinder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/snpgenie",
		URL:          []string{"https://github.com/hugheslab/snpgenie"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/raptr-sv",
		URL:          []string{"https://github.com/njdbickhart/raptr-sv"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/snftool",
		URL:          []string{"https://github.com/maxconway/snftool"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/maracluster",
		URL:          []string{"https://github.com/statisticalbiotechnology/maracluster"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/exondel",
		URL:          []string{"https://github.com/slzhao/exondel"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/moccs",
		URL:          []string{"https://github.com/yuifu/moccs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pophelper",
		URL:          []string{"https://github.com/royfrancis/pophelper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rmne",
		URL:          []string{"https://github.com/fvnieuwe/rmne"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/frama",
		URL:          []string{"https://github.com/gengit/frama"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tcsbu",
		URL:          []string{"https://github.com/sairum/tcsbu"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/transit",
		URL:          []string{"https://github.com/mad-lab/transit"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ksp-puel",
		URL:          []string{"https://github.com/pengyiyang/ksp-puel"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rmfam",
		URL:          []string{"https://github.com/ppgardne/rmfam"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bitseq",
		URL:          []string{"https://github.com/bitseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bitseqvb-benchmarking",
		URL:          []string{"https://github.com/bitseq/bitseqvb_benchmarking"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dmfinder",
		URL:          []string{"https://github.com/mhayes20/dmfinder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bayesembler",
		URL:          []string{"https://github.com/bioinformatics-centre/bayesembler"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/baitfisher-package",
		URL:          []string{"https://github.com/cmayer/baitfisher-package"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sprites",
		URL:          []string{"https://github.com/zhangzhen/sprites"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bamhash",
		URL:          []string{"https://github.com/decodegenetics/bamhash"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/medline-collaboration-datasets",
		URL:          []string{"https://github.com/qingzhanggithub/medline-collaboration-datasets"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/team",
		URL:          []string{"https://github.com/zmzhang/team"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/qmlgalaxyportal",
		URL:          []string{"https://github.com/tarostar/qmlgalaxyportal"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rnaentropy",
		URL:          []string{"https://github.com/clotelab/rnaentropy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/drme",
		URL:          []string{"https://github.com/lzcyzm/drme"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/isc",
		URL:          []string{"https://github.com/mts-strathclyde/isc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/spotyping-v2.0",
		URL:          []string{"https://github.com/xiaeryu/spotyping-v2.0"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bvnlabscattome",
		URL:          []string{"https://github.com/bvnlabscattome"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/integrativeanalysis",
		URL:          []string{"https://github.com/mgt000/integrativeanalysis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cnv-prioritization",
		URL:          []string{"https://github.com/compbio-uoft/cnv-prioritization"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/screenbeam",
		URL:          []string{"https://github.com/jyyu/screenbeam"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/package-gfe",
		URL:          []string{"https://github.com/takahiro-maruki/package-gfe"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/papst",
		URL:          []string{"https://github.com/paulbible/papst"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hicplotter",
		URL:          []string{"https://github.com/kcakdemir/hicplotter"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/refeditor",
		URL:          []string{"https://github.com/superyuan/refeditor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/datarlatools",
		URL:          []string{"https://github.com/abdullah009/datarlatools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rlatools",
		URL:          []string{"https://github.com/abdullah009/rlatools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mast",
		URL:          []string{"https://github.com/rglab/mast"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pulmon",
		URL:          []string{"https://github.com/variani/pulmon"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cc-mds",
		URL:          []string{"https://github.com/zhangxf-ccnu/cc-mds"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/isoscm",
		URL:          []string{"https://github.com/shenkers/isoscm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/warpgroup",
		URL:          []string{"https://github.com/nathaniel-mahieu/warpgroup"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/queries",
		URL:          []string{"https://github.com/chenying2016/queries"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/libchebi",
		URL:          []string{"https://github.com/libchebi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gsca",
		URL:          []string{"https://github.com/zji90/gsca"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hlatyphon",
		URL:          []string{"https://github.com/lji-bioinformatics/hlatyphon"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/arraymaker",
		URL:          []string{"https://github.com/cw2014/arraymaker"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/princces",
		URL:          []string{"https://github.com/czirjakgabor/princces"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ims-project",
		URL:          []string{"https://github.com/frizfealer/ims_project"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/broad-peaks",
		URL:          []string{"https://github.com/zubekj/broad_peaks"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sms",
		URL:          []string{"https://github.com/alcu/sms"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bfc",
		URL:          []string{"https://github.com/lh3/bfc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/solver-for-the-jeffreys-type-equations-system",
		URL:          []string{"https://github.com/wswgg/solver-for-the-jeffreys-type-equations-system"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/profet",
		URL:          []string{"https://github.com/ddofer/profet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phylesystem-api",
		URL:          []string{"https://github.com/opentreeoflife/phylesystem-api"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phylesystem",
		URL:          []string{"https://github.com/opentreeoflife/phylesystem"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/opentree",
		URL:          []string{"https://github.com/opentreeoflife/opentree"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/databaseui",
		URL:          []string{"https://github.com/kelpforest-cameo/databaseui"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ncbi-hackathons",
		URL:          []string{"https://github.com/ncbi-hackathons"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pcetk",
		URL:          []string{"https://github.com/mfx9/pcetk"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/chilay",
		URL:          []string{"https://github.com/ivis-at-bilkent/chilay"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/maprepeat",
		URL:          []string{"https://github.com/dcbmariano/maprepeat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/genomed3plot",
		URL:          []string{"https://github.com/brinkmanlab/genomed3plot"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mme-apis",
		URL:          []string{"https://github.com/ga4gh/mme-apis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ublastx-stageone",
		URL:          []string{"https://github.com/biofuture/ublastx_stageone"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pmsignature",
		URL:          []string{"https://github.com/friend1ws/pmsignature"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pets",
		URL:          []string{"https://github.com/binxia/pets"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/resseq",
		URL:          []string{"https://github.com/hrbeubiocenter/resseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/motifbreakr",
		URL:          []string{"https://github.com/simon-coetzee/motifbreakr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/epigbs",
		URL:          []string{"https://github.com/thomasvangurp/epigbs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/close",
		URL:          []string{"https://github.com/xfwang/close"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/shape-component-analysis-matlab",
		URL:          []string{"https://github.com/ccdlcmu/shape_component_analysis_matlab"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/isotope-correction-toolbox",
		URL:          []string{"https://github.com/jungreuc/isotope_correction_toolbox"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rcpa.tools",
		URL:          []string{"https://github.com/shengqh/rcpa.tools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ptxqc",
		URL:          []string{"https://github.com/cbielow/ptxqc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gxbrowser",
		URL:          []string{"https://github.com/benaroyaresearch/gxbrowser"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/jaffa",
		URL:          []string{"https://github.com/oshlack/jaffa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/s4vdpca",
		URL:          []string{"https://github.com/mwsill/s4vdpca"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dreg",
		URL:          []string{"https://github.com/danko-lab/dreg"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/blastjs",
		URL:          []string{"https://github.com/teammaclean/blastjs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vp",
		URL:          []string{"https://github.com/virtualpharmacist/vp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pymbar",
		URL:          []string{"https://github.com/choderalab/pymbar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bbgp",
		URL:          []string{"https://github.com/handetopa/bbgp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/aggregategenefunctionprediction",
		URL:          []string{"https://github.com/wimverleyen/aggregategenefunctionprediction"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/panfp",
		URL:          []string{"https://github.com/srjun/panfp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/decon1d",
		URL:          []string{"https://github.com/hughests/decon1d"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/seed",
		URL:          []string{"https://github.com/danlbek/seed"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bio-scores",
		URL:          []string{"https://github.com/kilicogluh/bio-scores"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sv-bay",
		URL:          []string{"https://github.com/institutcurie/sv-bay"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/glp",
		URL:          []string{"https://github.com/axot/glp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metapathways2",
		URL:          []string{"https://github.com/hallamlab/metapathways2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hhcompare",
		URL:          []string{"https://github.com/rgacesa/hhcompare"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pfam-maps",
		URL:          []string{"https://github.com/chembl/pfam_maps"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pfam-map-loader",
		URL:          []string{"https://github.com/chembl/pfam_map_loader"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/optimize-it",
		URL:          []string{"https://github.com/topherconley/optimize-it"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bamchop",
		URL:          []string{"https://github.com/cbmi-big/bamchop"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hefpipe-rpos",
		URL:          []string{"https://github.com/atticus29/hefpipe_rpos"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tree-hmm",
		URL:          []string{"https://github.com/uci-cbcl/tree-hmm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/connjur-sandbox",
		URL:          []string{"https://github.com/connjur/connjur-sandbox"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/shinymethyl",
		URL:          []string{"https://github.com/jfortin1/shinymethyl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/genplay",
		URL:          []string{"https://github.com/julienlajugie/genplay"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/netpathminer",
		URL:          []string{"https://github.com/ahmohamed/netpathminer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cola",
		URL:          []string{"https://github.com/nedaz/cola"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/aligngraph",
		URL:          []string{"https://github.com/baoe/aligngraph"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ismags",
		URL:          []string{"https://github.com/mhoubraken/ismags"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/helo",
		URL:          []string{"https://github.com/larisa-soldatova/helo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mulrfrepo",
		URL:          []string{"https://github.com/ruchiherself/mulrfrepo"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bio-quinn2013",
		URL:          []string{"https://github.com/d-quinn/bio_quinn2013"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/risa",
		URL:          []string{"https://github.com/isa-tools/risa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kasimjs",
		URL:          []string{"https://github.com/nicolasoury/kasimjs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kasimcloud",
		URL:          []string{"https://github.com/mdpedersen/kasimcloud"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pulp",
		URL:          []string{"https://github.com/ahwagner/pulp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sap",
		URL:          []string{"https://github.com/davidsun/sap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/phylosub",
		URL:          []string{"https://github.com/morrislab/phylosub"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/workflows",
		URL:          []string{"https://github.com/misshie/workflows"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ngstools",
		URL:          []string{"https://github.com/mfumagalli/ngstools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/2matrix",
		URL:          []string{"https://github.com/nrsalinas/2matrix"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/subseq",
		URL:          []string{"https://github.com/storeylab/subseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bp-quant",
		URL:          []string{"https://github.com/pnnl-comp-mass-spec/bp-quant"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ode-rmhmc",
		URL:          []string{"https://github.com/a-kramer/ode_rmhmc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hot-scan",
		URL:          []string{"https://github.com/itojal/hot_scan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rvs",
		URL:          []string{"https://github.com/strug-lab/rvs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bioflosoftware",
		URL:          []string{"https://github.com/libourellab/bioflosoftware"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bamtools",
		URL:          []string{"https://github.com/pezmaster31/bamtools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mitoseek",
		URL:          []string{"https://github.com/riverlee/mitoseek"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/treepl",
		URL:          []string{"https://github.com/blackrim/treepl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/voronto",
		URL:          []string{"https://github.com/rodrigosantamaria/voronto"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hts-barcode-checker",
		URL:          []string{"https://github.com/naturalis/hts-barcode-checker"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sibjoin.git+",
		URL:          []string{"https://github.com/ddexter/sibjoin.git+"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/raxml-light-1.0.5",
		URL:          []string{"https://github.com/stamatak/raxml-light-1.0.5"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/breakpointer",
		URL:          []string{"https://github.com/ruping/breakpointer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bioruby-svgenes",
		URL:          []string{"https://github.com/danmaclean/bioruby-svgenes"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pviz",
		URL:          []string{"https://github.com/genentech/pviz"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/removepli",
		URL:          []string{"https://github.com/mrezak/removepli"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bam2ssj",
		URL:          []string{"https://github.com/pervouchine/bam2ssj"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mixmir",
		URL:          []string{"https://github.com/ldiao/mixmir"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/github.com:",
		URL:          []string{"https://github.com:"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hal",
		URL:          []string{"https://github.com/glennhickey/hal"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/to-mi",
		URL:          []string{"https://github.com/to-mi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sail",
		URL:          []string{"https://github.com/sail"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/macs",
		URL:          []string{"https://github.com/taoliu/macs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/symmetry",
		URL:          []string{"https://github.com/rcsb/symmetry"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nextclip",
		URL:          []string{"https://github.com/richardmleggett/nextclip"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tigar",
		URL:          []string{"https://github.com/nariai/tigar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/idepi",
		URL:          []string{"https://github.com/veg/idepi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/emmckmer",
		URL:          []string{"https://github.com/alibashir/emmckmer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cddapp",
		URL:          []string{"https://github.com/rbvi/cddapp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/snpsvm",
		URL:          []string{"https://github.com/brendanofallon/snpsvm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/methylcode",
		URL:          []string{"https://github.com/brentp/methylcode"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hgtector",
		URL:          []string{"https://github.com/dittmarlab/hgtector"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hsa",
		URL:          []string{"https://github.com/vlcc/hsa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/extreme",
		URL:          []string{"https://github.com/uci-cbcl/extreme"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/probhap",
		URL:          []string{"https://github.com/kuleshov/probhap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biosual",
		URL:          []string{"https://github.com/4ndr01d3/biosual"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/inmembrane",
		URL:          []string{"https://github.com/boscoh/inmembrane"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cima",
		URL:          []string{"https://github.com/hongjiezhu/cima"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/r-samstrt",
		URL:          []string{"https://github.com/shka/r-samstrt"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/illumina-utils",
		URL:          []string{"https://github.com/meren/illumina-utils"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cwords",
		URL:          []string{"https://github.com/simras/cwords"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/netgem",
		URL:          []string{"https://github.com/vjethava/netgem"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mir-prefer",
		URL:          []string{"https://github.com/hangelwen/mir-prefer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fourrussiansrnafolding",
		URL:          []string{"https://github.com/ijalabv/fourrussiansrnafolding"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gapest",
		URL:          []string{"https://github.com/scilifelab/gapest"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/circadb",
		URL:          []string{"https://github.com/itmat/circadb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/samtools",
		URL:          []string{"https://github.com/samtools/samtools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/plaac",
		URL:          []string{"https://github.com/whitehead/plaac"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fragbuilder",
		URL:          []string{"https://github.com/jensengroup/fragbuilder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/perpetually-updated-trees",
		URL:          []string{"https://github.com/fizquierdo/perpetually-updated-trees"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/error_correction",
		URL:          []string{"https://github.com/macmanes/error_correction"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/genome-maps",
		URL:          []string{"https://github.com/compbio-bigdata-viz/genome-maps"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/marker2sequence",
		URL:          []string{"https://github.com/pbr/marker2sequence"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/simsem",
		URL:          []string{"https://github.com/ninjin/simsem"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pywater",
		URL:          []string{"https://github.com/hiteshpatel379/pywater"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lumpy-sv",
		URL:          []string{"https://github.com/arq5x/lumpy-sv"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/solidmon",
		URL:          []string{"https://github.com/eriqande/solidmon"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nseq",
		URL:          []string{"https://github.com/songlab/nseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/codonhmm",
		URL:          []string{"https://github.com/somork/codonhmm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gengis",
		URL:          []string{"https://github.com/beiko-lab/gengis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/imglib",
		URL:          []string{"https://github.com/imagej/imglib"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/combined-pvalues",
		URL:          []string{"https://github.com/brentp/combined-pvalues"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ibiomes",
		URL:          []string{"https://github.com/jcvthibault/ibiomes"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hyb",
		URL:          []string{"https://github.com/gkudla/hyb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gembassy",
		URL:          []string{"https://github.com/celery-kotone/gembassy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/smith2012-insulin-signalling",
		URL:          []string{"https://github.com/graham1034/smith2012_insulin_signalling"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/chromozoom",
		URL:          []string{"https://github.com/rothlab/chromozoom"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bioclipse",
		URL:          []string{"https://github.com/bioclipse"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/jannovar",
		URL:          []string{"https://github.com/charite/jannovar"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/singlecellassay",
		URL:          []string{"https://github.com/rglab/singlecellassay"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/retroseq",
		URL:          []string{"https://github.com/tk2/retroseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biddsat",
		URL:          []string{"https://github.com/jotegui/biddsat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bioblend",
		URL:          []string{"https://github.com/afgane/bioblend"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/compressgv",
		URL:          []string{"https://github.com/aschlosberg/compressgv"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/temp",
		URL:          []string{"https://github.com/jialiumasswenglab/temp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/epicode",
		URL:          []string{"https://github.com/mcieslik-mctp/epicode"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/microbedb",
		URL:          []string{"https://github.com/mlangill/microbedb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bridges",
		URL:          []string{"https://github.com/rassis/bridges"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/orcacode",
		URL:          []string{"https://github.com/orcabioinfo/orcacode"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/star-genome-browser",
		URL:          []string{"https://github.com/angell1117/star-genome-browser"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/findis-db",
		URL:          []string{"https://github.com/findis-db"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/crdata",
		URL:          []string{"https://github.com/seerdata/crdata"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pinda",
		URL:          []string{"https://github.com/dgkontopoulos/pinda"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/downloads",
		URL:          []string{"https://github.com/piotrzakrzewski/meteval/downloads"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/spread",
		URL:          []string{"https://github.com/phylogeography/spread"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/subcloneseeker",
		URL:          []string{"https://github.com/yiq/subcloneseeker"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/afcluster",
		URL:          []string{"https://github.com/luscinius/afcluster"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/temt",
		URL:          []string{"https://github.com/uci-cbcl/temt"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/epof",
		URL:          []string{"https://github.com/gangchen/epof"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gobe",
		URL:          []string{"https://github.com/brentp/gobe"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tfbayes",
		URL:          []string{"https://github.com/pbenner/tfbayes"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/trap",
		URL:          []string{"https://github.com/bacak/trap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ubd",
		URL:          []string{"https://github.com/pelinakan/ubd"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ropebwt2",
		URL:          []string{"https://github.com/lh3/ropebwt2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scot",
		URL:          []string{"https://github.com/scot-dev/scot"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/msisensor",
		URL:          []string{"https://github.com/ding-lab/msisensor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cloudbrush",
		URL:          []string{"https://github.com/ice91/cloudbrush"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fermi",
		URL:          []string{"https://github.com/lh3/fermi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bioc",
		URL:          []string{"https://github.com/noname2020/bioc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dafga",
		URL:          []string{"https://github.com/outbig/dafga"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/blastgraph",
		URL:          []string{"https://github.com/bigwiv/blastgraph"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/statistical-agglomeration",
		URL:          []string{"https://github.com/optimusmoose/statistical_agglomeration"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mspire",
		URL:          []string{"https://github.com/princelab/mspire"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tgnet",
		URL:          []string{"https://github.com/ksanao/tgnet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/improver2013",
		URL:          []string{"https://github.com/uci-igb/improver2013"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/red",
		URL:          []string{"https://github.com/biolab/red"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fcnv",
		URL:          []string{"https://github.com/compbio-uoft/fcnv"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/express-d",
		URL:          []string{"https://github.com/adarob/express-d"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ontoquery",
		URL:          []string{"https://github.com/ilincatudose/ontoquery"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bide-2d",
		URL:          []string{"https://github.com/stevenhwu/bide-2d"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/quota-alignment",
		URL:          []string{"https://github.com/tanghaibao/quota-alignment"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/securema",
		URL:          []string{"https://github.com/xieconnect/securema"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/circuitservice",
		URL:          []string{"https://github.com/xieconnect/circuitservice"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/boolesim",
		URL:          []string{"https://github.com/matthiasbock/boolesim"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/herg-chembl-jcim",
		URL:          []string{"https://github.com/pzc/herg_chembl_jcim"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dnacc",
		URL:          []string{"https://github.com/patvarilly/dnacc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kbws",
		URL:          []string{"https://github.com/cory-ko/kbws"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sushi",
		URL:          []string{"https://github.com/dphansti/sushi"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/traxtile-public",
		URL:          []string{"https://github.com/braunb/traxtile-public"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/repeatseq",
		URL:          []string{"https://github.com/adaptivegenome/repeatseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metarep",
		URL:          []string{"https://github.com/jcvi/metarep"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/varianttoolchest",
		URL:          []string{"https://github.com/mebbert/varianttoolchest"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/geiger-v2",
		URL:          []string{"https://github.com/mwpennell/geiger-v2"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/snpr",
		URL:          []string{"https://github.com/gedankenstuecke/snpr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pipeliner",
		URL:          []string{"https://github.com/brunonevado/pipeliner"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mqap",
		URL:          []string{"https://github.com/sanchak/mqap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mitobim",
		URL:          []string{"https://github.com/chrishah/mitobim"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ontomaton",
		URL:          []string{"https://github.com/isa-tools/ontomaton"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nestly",
		URL:          []string{"https://github.com/fhcrc/nestly"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/msr",
		URL:          []string{"https://github.com/tknijnen/msr"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/chance",
		URL:          []string{"https://github.com/songlab/chance"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ensembl-rest",
		URL:          []string{"https://github.com/ensembl/ensembl-rest"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/biojs",
		URL:          []string{"https://github.com/biojs/biojs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rendertoolbox3",
		URL:          []string{"https://github.com/davidbrainard/rendertoolbox3"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/msmexplorer",
		URL:          []string{"https://github.com/simtk/msmexplorer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/flowpeaks",
		URL:          []string{"https://github.com/yongchao/flowpeaks"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dave-the-scientist",
		URL:          []string{"https://github.com/dave-the-scientist"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/galaxy",
		URL:          []string{"https://github.com/modencode-dcc/galaxy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mim",
		URL:          []string{"https://github.com/lucapinello/mim"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/franklin",
		URL:          []string{"https://github.com/joseblanca/franklin"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/progressivecactus",
		URL:          []string{"https://github.com/glennhickey/progressivecactus"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/affypipe",
		URL:          []string{"https://github.com/nicolazzie/affypipe"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fselector",
		URL:          []string{"https://github.com/need47/fselector"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/targetseqview",
		URL:          []string{"https://github.com/eitan177/targetseqview"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/arts",
		URL:          []string{"https://github.com/mmaiensc/arts"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/exonsuite",
		URL:          []string{"https://github.com/dilanustek/exonsuite"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tnrs",
		URL:          []string{"https://github.com/iplantcollaborativeopensource/tnrs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vseams",
		URL:          []string{"https://github.com/ollyburren/vseams"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/surpriseme",
		URL:          []string{"https://github.com/raldecoa/surpriseme"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/scct",
		URL:          []string{"https://github.com/wavefancy/scct"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/slamem",
		URL:          []string{"https://github.com/fjdf/slamem"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ale",
		URL:          []string{"https://github.com/ssolo/ale"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ebcall",
		URL:          []string{"https://github.com/friend1ws/ebcall"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nonlineargradientsui",
		URL:          []string{"https://github.com/statisticalbiotechnology/nonlineargradientsui"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mix",
		URL:          []string{"https://github.com/cbib/mix"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/seqpatch",
		URL:          []string{"https://github.com/zhixingfeng/seqpatch"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/questionnaire",
		URL:          []string{"https://github.com/apcdr/questionnaire"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/clj-biosequence",
		URL:          []string{"https://github.com/s312569/clj-biosequence"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/diffusivedynamics",
		URL:          []string{"https://github.com/lbf-ijs/diffusivedynamics"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cansnper",
		URL:          []string{"https://github.com/adrlar/cansnper"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ap11-samifier",
		URL:          []string{"https://github.com/intersectaustralia/ap11_samifier"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sbetoolbox",
		URL:          []string{"https://github.com/biocoder/sbetoolbox"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metamos",
		URL:          []string{"https://github.com/treangen/metamos"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/broccoli",
		URL:          []string{"https://github.com/wanderine/broccoli"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ngsplusplus",
		URL:          []string{"https://github.com/ngs-lib/ngsplusplus"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bits",
		URL:          []string{"https://github.com/arq5x/bits"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/methylmaps",
		URL:          []string{"https://github.com/epigenomics/methylmaps"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pangenome",
		URL:          []string{"https://github.com/rec3141/pangenome"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dcj-c",
		URL:          []string{"https://github.com/putnamlab/dcj-c"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/aryana-aligner",
		URL:          []string{"https://github.com/aryana-aligner"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/genomon-itdetector",
		URL:          []string{"https://github.com/ken0-1n/genomon-itdetector"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/samblaster",
		URL:          []string{"https://github.com/gregoryfaust/samblaster"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cordova",
		URL:          []string{"https://github.com/clcg/cordova"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dsm-framework",
		URL:          []string{"https://github.com/hiitmetagenomics/dsm-framework"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ennet",
		URL:          []string{"https://github.com/slawekj/ennet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/leaf",
		URL:          []string{"https://github.com/simoncb765/leaf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dendropy",
		URL:          []string{"https://github.com/jeetsukumaran/dendropy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fiddlercrab.info",
		URL:          []string{"https://github.com/msrosenberg/fiddlercrab.info"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/varcmp",
		URL:          []string{"https://github.com/lh3/varcmp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/qc3",
		URL:          []string{"https://github.com/slzhao/qc3"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/orderedpainting",
		URL:          []string{"https://github.com/bioprojects/orderedpainting"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bioruby-ucsc-api",
		URL:          []string{"https://github.com/misshie/bioruby-ucsc-api"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/randal",
		URL:          []string{"https://github.com/namsyvo/randal"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/magetbrain",
		URL:          []string{"https://github.com/pipitone/magetbrain"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gpmat",
		URL:          []string{"https://github.com/sheffieldml/gpmat"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ngsane",
		URL:          []string{"https://github.com/bauerlab/ngsane"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/emirge",
		URL:          []string{"https://github.com/csmiller/emirge"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/haploclique",
		URL:          []string{"https://github.com/armintoepfer/haploclique"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metriculator",
		URL:          []string{"https://github.com/princelab/metriculator"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rnie",
		URL:          []string{"https://github.com/ppgardne/rnie"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/canvasdb",
		URL:          []string{"https://github.com/uppsalagenomecenter/canvasdb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pyncs",
		URL:          []string{"https://github.com/inincs/pyncs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/abra",
		URL:          []string{"https://github.com/mozack/abra"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/patterncnv",
		URL:          []string{"https://github.com/topsoil/patterncnv"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mmseq",
		URL:          []string{"https://github.com/eturro/mmseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/peakzilla",
		URL:          []string{"https://github.com/steinmann/peakzilla"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fragalcode",
		URL:          []string{"https://github.com/sanchak/fragalcode"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/suns",
		URL:          []string{"https://github.com/godotgildor/suns"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/suns-cmd",
		URL:          []string{"https://github.com/gabriel439/suns-cmd"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/suns-search",
		URL:          []string{"https://github.com/gabriel439/suns-search"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/stochhmm",
		URL:          []string{"https://github.com/korflab/stochhmm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/derfinder",
		URL:          []string{"https://github.com/alyssafrazee/derfinder"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ftc",
		URL:          []string{"https://github.com/loopasam/ftc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/author-detection",
		URL:          []string{"https://github.com/matthewberryman/author-detection"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/grompy",
		URL:          []string{"https://github.com/grompy"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/rapidmic",
		URL:          []string{"https://github.com/helloworldcn/rapidmic"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/wavecnv",
		URL:          []string{"https://github.com/wavecnv"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/locings",
		URL:          []string{"https://github.com/shird/locings"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mfeprimer",
		URL:          []string{"https://github.com/quwubin/mfeprimer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/streamingtrim",
		URL:          []string{"https://github.com/gibacci/streamingtrim"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mynclist",
		URL:          []string{"https://github.com/bushlab/mynclist"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hpg-aligner",
		URL:          []string{"https://github.com/opencb/hpg-aligner"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hapmuc",
		URL:          []string{"https://github.com/usuyama/hapmuc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/blobology",
		URL:          []string{"https://github.com/blaxterlab/blobology"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/blobsplorer",
		URL:          []string{"https://github.com/mojones/blobsplorer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/swiftlink",
		URL:          []string{"https://github.com/ajm/swiftlink"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/fsmkl",
		URL:          []string{"https://github.com/jseoane/fsmkl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/chromatra",
		URL:          []string{"https://github.com/cmmt/chromatra"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/gobyweb2-plugins",
		URL:          []string{"https://github.com/campagnelaboratory/gobyweb2-plugins"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kbmtl",
		URL:          []string{"https://github.com/mehmetgonen/kbmtl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dag-viewer-biojs",
		URL:          []string{"https://github.com/alexkalderimis/dag-viewer-biojs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mychembl",
		URL:          []string{"https://github.com/rochoa85/mychembl"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vacceed",
		URL:          []string{"https://github.com/sgoodswe/vacceed"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nonpareil",
		URL:          []string{"https://github.com/lmrodriguezr/nonpareil"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/graph2tab",
		URL:          []string{"https://github.com/isa-tools/graph2tab"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/wholecellsimdb",
		URL:          []string{"https://github.com/covertlab/wholecellsimdb"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kraken",
		URL:          []string{"https://github.com/nedaz/kraken"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/epiq",
		URL:          []string{"https://github.com/yaarasegre/epiq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/odonata-traits",
		URL:          []string{"https://github.com/biologicalrecordscentre/odonata_traits"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/e-driver",
		URL:          []string{"https://github.com/eduardporta/e-driver"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pyloh",
		URL:          []string{"https://github.com/uci-cbcl/pyloh"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/metitree",
		URL:          []string{"https://github.com/netherlandsmetabolomicscentre/metitree"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mod-bio",
		URL:          []string{"https://github.com/lindenb/mod_bio"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/selscan",
		URL:          []string{"https://github.com/szpiech/selscan"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ragout",
		URL:          []string{"https://github.com/fenderglass/ragout"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/complete-striped-smith-waterman-library",
		URL:          []string{"https://github.com/mengyao/complete-striped-smith-waterman-library"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/snipviz",
		URL:          []string{"https://github.com/yeastrc/snipviz"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dyconet",
		URL:          []string{"https://github.com/juliemkauffman/dyconet"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ttree-source",
		URL:          []string{"https://github.com/0asa/ttree-source"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bluesnp",
		URL:          []string{"https://github.com/ibm-bioinformatics/bluesnp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "doc/visordoc",
		URL:          []string{"https://github.com/davidebolo1993/visordoc"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "doc/squigglekitdocs",
		URL:          []string{"https://github.com/psy-fer/squigglekitdocs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hyper-docs",
		URL:          []string{"https://github.com/montilab/hyper-docs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/maize-tes",
		URL:          []string{"https://github.com/mcstitzer/maize_tes"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/immunedeconv",
		URL:          []string{"https://github.com/grst/immunedeconv"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "doc/phatdocs",
		URL:          []string{"https://github.com/chgibb/phatdocs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "doc/pypdb-docs",
		URL:          []string{"https://github.com/williamgilpin/pypdb_docs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/allonkleinlab/spring",
		URL:          []string{"https://github.com/allonkleinlab/spring"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bimberlab/discvrseq",
		URL:          []string{"https://github.com/bimberlab/discvrseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/deplanckelab/asap",
		URL:          []string{"https://github.com/deplanckelab/asap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hysxe/acme",
		URL:          []string{"https://github.com/hysxe/acme"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/juliapalacios/phylodyn",
		URL:          []string{"https://github.com/juliapalacios/phylodyn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mathelab/ramp-db",
		URL:          []string{"https://github.com/mathelab/ramp-db"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pacificbiosciences/falcon",
		URL:          []string{"https://github.com/pacificbiosciences/falcon"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/shawn-xu/ppip",
		URL:          []string{"https://github.com/shawn-xu/ppip"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tf-chan-lab/omtools",
		URL:          []string{"https://github.com/tf-chan-lab/omtools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vccri/c3",
		URL:          []string{"https://github.com/vccri/c3"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vislab/eeg-blinks",
		URL:          []string{"https://github.com/vislab/eeg-blinks"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/abbyyan3/bhglm",
		URL:          []string{"https://github.com/abbyyan3/bhglm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/aditya-88/asap",
		URL:          []string{"https://github.com/aditya-88/asap"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/aldenleung/omtools",
		URL:          []string{"https://github.com/aldenleung/omtools"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bbarker/falcon",
		URL:          []string{"https://github.com/bbarker/falcon"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bigdatagenomics/mango",
		URL:          []string{"https://github.com/bigdatagenomics/mango"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/bimberlab/discvrseq",
		URL:          []string{"https://github.com/bimberlab/discvrseq"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/cabergh/ebdims",
		URL:          []string{"https://github.com/cabergh/ebdims"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/connor-lab/vapor",
		URL:          []string{"https://github.com/connor-lab/vapor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dib-lab/khmer",
		URL:          []string{"https://github.com/dib-lab/khmer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dphansti/coral",
		URL:          []string{"https://github.com/dphansti/coral"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/dphansti/mango",
		URL:          []string{"https://github.com/dphansti/mango"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/feizhe/pgs",
		URL:          []string{"https://github.com/feizhe/pgs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ged-lab/khmer",
		URL:          []string{"https://github.com/ged-lab/khmer"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/genenotebook",
		URL:          []string{"https://github.com/genenotebook/genenotebook"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/n1pas",
		URL:          []string{"https://github.com/grizant/n1pas"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/hr1912/treeexp",
		URL:          []string{"https://github.com/hr1912/treeexp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/innate2adaptive/decombinator",
		URL:          []string{"https://github.com/innate2adaptive/decombinator"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/isglobal-brge/lambda",
		URL:          []string{"https://github.com/isglobal-brge/lambda"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ishidalab-titech/3dcnn-mqa",
		URL:          []string{"https://github.com/ishidalab-titech/3dcnn_mqa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/iskana/pbwt-sec",
		URL:          []string{"https://github.com/iskana/pbwt-sec"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/jackhou2/c3",
		URL:          []string{"https://github.com/jackhou2/c3"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/jbudis/lambda",
		URL:          []string{"https://github.com/jbudis/lambda"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/jinchaofeng-code",
		URL:          []string{"https://github.com/jinchaofeng/code"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/jingwyang/treeexp",
		URL:          []string{"https://github.com/jingwyang/treeexp"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/johnbarton/ace",
		URL:          []string{"https://github.com/johnbarton/ace"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/jomayer/smurf",
		URL:          []string{"https://github.com/jomayer/smurf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kodicollins/pasta",
		URL:          []string{"https://github.com/kodicollins/pasta"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/krm15/acme",
		URL:          []string{"https://github.com/krm15/acme"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kubernetes/charts",
		URL:          []string{"https://github.com/kubernetes/charts"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/kyzhu/assembltrie",
		URL:          []string{"https://github.com/kyzhu/assembltrie"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/lamoureux-lab/3dcnn-mqa",
		URL:          []string{"https://github.com/lamoureux-lab/3dcnn_mqa"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/laura-orellana/ebdims",
		URL:          []string{"https://github.com/laura-orellana/ebdims"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ma-compbio/multires",
		URL:          []string{"https://github.com/ma-compbio/multires"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mathelab/ramp-db",
		URL:          []string{"https://github.com/mathelab/ramp-db"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mbeccuti/pgs",
		URL:          []string{"https://github.com/mbeccuti/pgs"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mdkarcher/phylodyn",
		URL:          []string{"https://github.com/mdkarcher/phylodyn"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/mills-lab/vapor",
		URL:          []string{"https://github.com/mills-lab/vapor"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/molgenis/molgenis",
		URL:          []string{"https://github.com/molgenis/molgenis"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nclaes/coral",
		URL:          []string{"https://github.com/nclaes/coral"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/nyiuab/bhglm",
		URL:          []string{"https://github.com/nyiuab/bhglm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/pachyderm/pachyderm",
		URL:          []string{"https://github.com/pachyderm/pachyderm"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ratschlab/pbwt-sec",
		URL:          []string{"https://github.com/ratschlab/pbwt-sec"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/raunaq-m/multires",
		URL:          []string{"https://github.com/raunaq-m/multires"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/shawn-xu/ppip",
		URL:          []string{"https://github.com/shawn-xu/ppip"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sheikhizadeh/ace",
		URL:          []string{"https://github.com/sheikhizadeh/ace"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/shubhamchandak94/spring",
		URL:          []string{"https://github.com/shubhamchandak94/spring"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/skandlab/smurf",
		URL:          []string{"https://github.com/skandlab/smurf"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/smirarab/pasta",
		URL:          []string{"https://github.com/smirarab/pasta"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/sysbiolux/falcon",
		URL:          []string{"https://github.com/sysbiolux/falcon"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/tgvaughan/icytree",
		URL:          []string{"https://github.com/tgvaughan/icytree"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/trtcrd/slim",
		URL:          []string{"https://github.com/trtcrd/slim"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/uclinfectionimmunity/decombinator",
		URL:          []string{"https://github.com/uclinfectionimmunity/decombinator"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vgteam/vg",
		URL:          []string{"https://github.com/vgteam/vg"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/vislab/eeg-blinks",
		URL:          []string{"https://github.com/vislab/eeg-blinks"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/xulabs/projects",
		URL:          []string{"https://github.com/xulabs/projects"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/yoann-dufresne/slim",
		URL:          []string{"https://github.com/yoann-dufresne/slim"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/yoheirosen/vg",
		URL:          []string{"https://github.com/yoheirosen/vg"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "github/ytzhong/projects",
		URL:          []string{"https://github.com/ytzhong/projects"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
}
