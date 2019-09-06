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
		Name:         "wkfl-encode-chipseq2",
		Tags:         []string{"pipeline", "chipseq"},
		URL:          []string{"https://github.com/ENCODE-DCC/chip-seq-pipeline2"},
		PostShellCmd: []string{"pip3 install caper || pip install caper"},
	},
}
