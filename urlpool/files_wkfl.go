package urlpool

var wkflFiles = []bgetFilesURLType{}

var snakemakeWkflFiles = []bgetFilesURLType{
	{
		Name:         "snakemake/dna-seq-gatk-variant-calling",
		URL:          []string{"https://github.com/snakemake-workflows/dna-seq-gatk-variant-calling"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
	{
		Name:         "snakemake/sequana",
		URL:          []string{"https://github.com/sequana/sequana"},
		PostShellCmd: []string{"cd {{dest}} && git checkout {{version}}"},
	},
}

var wdlWkflFiles = []bgetFilesURLType{
	{
		Name:         "wdl/encode-chipseq2",
		Tags:         []string{"pipeline", "chipseq"},
		URL:          []string{"https://github.com/ENCODE-DCC/chip-seq-pipeline2"},
		PostShellCmd: []string{"pip3 install caper || pip install caper"},
	},
	{
		Name: "wdl/antonkulaga",
		URL:  []string{"https://github.com/antonkulaga/antonkulaga"},
	},
	{
		Name: "wdl/bgm-cwg-whole-exome-sequencing",
		URL:  []string{"https://github.com/bgm-cwg/Whole-Exome-Sequencing-Workflow"},
	},
	{
		Name: "wdl/bgm-cwg-whole-genome-sequencing",
		URL:  []string{"https://github.com/bgm-cwg/Whole-Genome-Sequencing-Workflow"},
	},
	{
		Name: "wdl/biowdl-aligning",
		URL:  []string{"https://github.com/biowdl/aligning"},
	},
	{
		Name: "wdl/biowdl-bam-to-gvcf",
		URL:  []string{"https://github.com/biowdl/bam-to-gvcf"},
	},
	{
		Name: "wdl/biowdl-chip-seq",
		URL:  []string{"https://github.com/biowdl/ChIP-seq"},
	},
	{
		Name: "wdl/biowdl-gams",
		URL:  []string{"https://github.com/biowdl/gams"},
	},
	{
		Name: "wdl/biowdl-germline-dna",
		URL:  []string{"https://github.com/biowdl/germline-DNA"},
	},
	{
		Name: "wdl/biowdl-jointgenotyping",
		URL:  []string{"https://github.com/biowdl/jointgenotyping"},
	},
	{
		Name: "wdl/biowdl-pipeline-template",
		URL:  []string{"https://github.com/biowdl/pipeline-template"},
	},
	{
		Name: "wdl/biowdl-qc",
		URL:  []string{"https://github.com/biowdl/QC"},
	},
	{
		Name: "wdl/biowdl-rna-seq",
		URL:  []string{"https://github.com/biowdl/RNA-seq"},
	},
	{
		Name: "wdl/biowdl-somatic-variantcalling",
		URL:  []string{"https://github.com/biowdl/somatic-variantcalling"},
	},
	{
		Name: "wdl/biowdl-tasks",
		URL:  []string{"https://github.com/biowdl/tasks"},
	},
	{
		Name: "wdl/biowdl-virus-assembly",
		URL:  []string{"https://github.com/biowdl/virus-assembly"},
	},
	{
		Name: "wdl/tsca-firecloud",
		URL:  []string{"https://github.com/broadinstitute/TSCA_firecloud_WDLS"},
	},
	{
		Name: "wdl/pacbio-tools",
		URL:  []string{"https://github.com/camccowan/PacBio-WDL-Tools"},
	},
	{
		Name: "wdl/brass",
		URL:  []string{"https://github.com/cancerit/BRASS-wdl"},
	},
	{
		Name: "wdl/deepvariant-glnexus",
		URL:  []string{"https://github.com/dnanexus-rnd/DeepVariant-GLnexus-WDL"},
	},
	{
		Name: "wdl/firecloud-tools",
		URL:  []string{"https://github.com/edawson/firecloud-tools"},
	},
	{
		Name: "wdl/gatk-wgs-germline-snps-indels",
		URL:  []string{"https://github.com/gatk-workflows/broad-prod-wgs-germline-snps-indels"},
	},
	{
		Name: "wdl/gatk-five-dollar-genome-analysis",
		URL:  []string{"https://github.com/gatk-workflows/five-dollar-genome-analysis-pipeline"},
	},
	{
		Name: "wdl/gatk-3-4-rnaseq-germline-snps-indels",
		URL:  []string{"https://github.com/gatk-workflows/gatk3-4-rnaseq-germline-snps-indels"},
	},
	{
		Name: "wdl/gatk-3-data-processing",
		URL:  []string{"https://github.com/gatk-workflows/gatk3-data-processing"},
	},
	{
		Name: "wdl/gatk-4-data-processing",
		URL:  []string{"https://github.com/gatk-workflows/gatk4-data-processing"},
	},
	{
		Name: "wdl/gatk-4-germline-snps-indels",
		URL:  []string{"https://github.com/gatk-workflows/gatk4-germline-snps-indels"},
	},
	{
		Name: "wdl/gatk-4-somatic-cnvs",
		URL:  []string{"https://github.com/gatk-workflows/gatk4-somatic-cnvs"},
	},
	{
		Name: "wdl/gatk-4-somatic-snvs-indels",
		URL:  []string{"https://github.com/gatk-workflows/gatk4-somatic-snvs-indels"},
	},
	{
		Name: "wdl/gatk-4-somatic-with-preprocessing",
		URL:  []string{"https://github.com/gatk-workflows/gatk4-somatic-with-preprocessing"},
	},
	{
		Name: "wdl/gatk-intel-faster-alternative-samples",
		URL:  []string{"https://github.com/gatk-workflows/intel-faster-alternative-samples"},
	},
	{
		Name: "wdl/gatk-3-4-intel-germline-snps-indels",
		URL:  []string{"https://github.com/gatk-workflows/intel-gatk3-4-germline-snps-indels"},
	},
	{
		Name: "wdl/gatk-3-intel-germline-snps-indels",
		URL:  []string{"https://github.com/gatk-workflows/intel-gatk3-germline-snps-indels"},
	},
	{
		Name: "wdl/gatk-4-intel-germline-snps-indels",
		URL:  []string{"https://github.com/gatk-workflows/intel-gatk4-germline-snps-indels"},
	},
	{
		Name: "wdl/gatk-4-intel-somatic-with-preprocessing",
		URL:  []string{"https://github.com/gatk-workflows/intel-gatk4-somatic-with-preprocessing"},
	},
	{
		Name: "wdl/gatk-seq-format-conversion",
		URL:  []string{"https://github.com/gatk-workflows/seq-format-conversion"},
	},
	{
		Name: "wdl/gatk-seq-format-validation",
		URL:  []string{"https://github.com/gatk-workflows/seq-format-validation"},
	},
	{
		Name: "wdl/genomics-iter-developers-wgs-wes-germline",
		URL:  []string{"https://github.com/genomicsITER-developers/wdl"},
	},
	{
		Name: "wdl/hongiiv-gatk-workflows",
		URL:  []string{"https://github.com/hongiiv/gatk-workflows"},
	},
	{
		Name: "wdl/uofabio",
		URL:  []string{"https://github.com/jimmybgammyknee/uofaBioWDL"},
	},
	{
		Name: "wdl/jimmybgammyknee-workflows",
		URL:  []string{"https://github.com/jimmybgammyknee/workflows"},
	},
	{
		Name: "wdl/johnmous-chip-seq",
		URL:  []string{"https://github.com/johnmous/ChIPseq"},
	},
	{
		Name: "wdl/mgs-sop",
		URL:  []string{"https://github.com/karlrl/mgs_sop"},
	},
	{
		Name: "wdl/kurt-hetrick-variants",
		URL:  []string{"https://github.com/Kurt-Hetrick/WDL"},
	},
	{
		Name: "wdl/saige",
		URL:  []string{"https://github.com/manning-lab/saigeWdl"},
	},
	{
		Name: "wdl/wgsa-parsr",
		URL:  []string{"https://github.com/manning-lab/wgsaParsrWdl"},
	},
	{
		Name: "wdl/deepvar",
		URL:  []string{"https://github.com/mkim55/deepvar-wdl"},
	},
	{
		Name: "wdl/mayomics-vc",
		URL:  []string{"https://github.com/ncsa/MayomicsVC"},
	},
	{
		Name: "wdl/metagenophile",
		URL:  []string{"https://github.com/omixer/metagenophile"},
	},
	{
		Name: "wdl/oskarvid-germline-pipeline",
		URL:  []string{"https://github.com/oskarvid/wdl_germline_pipeline"},
	},
	{
		Name: "wdl/ottojolanki-kallisto",
		URL:  []string{"https://github.com/ottojolanki/kallisto-wdl"},
	},
	{
		Name: "wdl/raisok-rnaseq",
		URL:  []string{"https://github.com/raisok/RNAseq_WDL"},
	},
	{
		Name: "wdl/sanderest-cromwell",
		URL:  []string{"https://github.com/SanderEST/cromwell_wdl"},
	},
	{
		Name: "wdl/seonwhee-genome-gatk",
		URL:  []string{"https://github.com/Seonwhee-Genome/GATK_WDL"},
	},
}

func init() {
	wkflFiles = append(wkflFiles, wdlWkflFiles...)
	wkflFiles = append(wkflFiles, snakemakeWkflFiles...)
}
