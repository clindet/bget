package urlpool

var wkflFiles = []bgetFilesURLType{}

var wdlWkflFiles = []bgetFilesURLType{
	{
		Name:         "wkfl/encode-chipseq2",
		Tags:         []string{"pipeline", "chipseq"},
		URL:          []string{"https://github.com/ENCODE-DCC/chip-seq-pipeline2"},
		PostShellCmd: []string{"pip3 install caper || pip install caper"},
	},
	{
		Name: "wkfl/antonkulaga/antonkulaga",
		URL:  []string{"https://github.com/antonkulaga/antonkulaga"},
	},
	{
		Name: "wkfl/bgm-cwg/Whole-Exome-Sequencing-Workflow",
		URL:  []string{"https://github.com/bgm-cwg/Whole-Exome-Sequencing-Workflow"},
	},
	{
		Name: "wkfl/bgm-cwg/Whole-Genome-Sequencing-Workflow",
		URL:  []string{"https://github.com/bgm-cwg/Whole-Genome-Sequencing-Workflow"},
	},
	{
		Name: "wkfl/biowdl/aligning",
		URL:  []string{"https://github.com/biowdl/aligning"},
	},
	{
		Name: "wkfl/biowdl/bam-to-gvcf",
		URL:  []string{"https://github.com/biowdl/bam-to-gvcf"},
	},
	{
		Name: "wkfl/biowdl/ChIP-seq",
		URL:  []string{"https://github.com/biowdl/ChIP-seq"},
	},
	{
		Name: "wkfl/biowdl/gams",
		URL:  []string{"https://github.com/biowdl/gams"},
	},
	{
		Name: "wkfl/biowdl/germline-DNA",
		URL:  []string{"https://github.com/biowdl/germline-DNA"},
	},
	{
		Name: "wkfl/biowdl/jointgenotyping",
		URL:  []string{"https://github.com/biowdl/jointgenotyping"},
	},
	{
		Name: "wkfl/biowdl/pipeline-template",
		URL:  []string{"https://github.com/biowdl/pipeline-template"},
	},
	{
		Name: "wkfl/biowdl/QC",
		URL:  []string{"https://github.com/biowdl/QC"},
	},
	{
		Name: "wkfl/biowdl/RNA-seq",
		URL:  []string{"https://github.com/biowdl/RNA-seq"},
	},
	{
		Name: "wkfl/biowdl/somatic-variantcalling",
		URL:  []string{"https://github.com/biowdl/somatic-variantcalling"},
	},
	{
		Name: "wkfl/biowdl/tasks",
		URL:  []string{"https://github.com/biowdl/tasks"},
	},
	{
		Name: "wkfl/biowdl/virus-assembly",
		URL:  []string{"https://github.com/biowdl/virus-assembly"},
	},
	{
		Name: "wkfl/broadinstitute/TSCA_firecloud_WDLS",
		URL:  []string{"https://github.com/broadinstitute/TSCA_firecloud_WDLS"},
	},
	{
		Name: "wkfl/camccowan/PacBio-WDL-Tools",
		URL:  []string{"https://github.com/camccowan/PacBio-WDL-Tools"},
	},
	{
		Name: "wkfl/cancerit/BRASS-wdl",
		URL:  []string{"https://github.com/cancerit/BRASS-wdl"},
	},
	{
		Name: "wkfl/dnanexus-rnd/DeepVariant-GLnexus-WDL",
		URL:  []string{"https://github.com/dnanexus-rnd/DeepVariant-GLnexus-WDL"},
	},
	{
		Name: "wkfl/edawson/firecloud-tools",
		URL:  []string{"https://github.com/edawson/firecloud-tools"},
	},
	{
		Name: "wkfl/gatk/broad-prod-wgs-germline-snps-indels",
		URL:  []string{"https://github.com/gatk-workflows/broad-prod-wgs-germline-snps-indels"},
	},
	{
		Name: "wkfl/gatk/five-dollar-genome-analysis-pipeline",
		URL:  []string{"https://github.com/gatk-workflows/five-dollar-genome-analysis-pipeline"},
	},
	{
		Name: "wkfl/gatk/gatk3-4-rnaseq-germline-snps-indels",
		URL:  []string{"https://github.com/gatk-workflows/gatk3-4-rnaseq-germline-snps-indels"},
	},
	{
		Name: "wkfl/gatk/gatk3-data-processing",
		URL:  []string{"https://github.com/gatk-workflows/gatk3-data-processing"},
	},
	{
		Name: "wkfl/gatk/gatk4-data-processing",
		URL:  []string{"https://github.com/gatk-workflows/gatk4-data-processing"},
	},
	{
		Name: "wkfl/gatk/gatk4-germline-snps-indels",
		URL:  []string{"https://github.com/gatk-workflows/gatk4-germline-snps-indels"},
	},
	{
		Name: "wkfl/gatk/gatk4-somatic-cnvs",
		URL:  []string{"https://github.com/gatk-workflows/gatk4-somatic-cnvs"},
	},
	{
		Name: "wkfl/gatk/gatk4-somatic-snvs-indels",
		URL:  []string{"https://github.com/gatk-workflows/gatk4-somatic-snvs-indels"},
	},
	{
		Name: "wkfl/gatk/gatk4-somatic-with-preprocessing",
		URL:  []string{"https://github.com/gatk-workflows/gatk4-somatic-with-preprocessing"},
	},
	{
		Name: "wkfl/gatk/intel-faster-alternative-samples",
		URL:  []string{"https://github.com/gatk-workflows/intel-faster-alternative-samples"},
	},
	{
		Name: "wkfl/gatk/intel-gatk3-4-germline-snps-indels",
		URL:  []string{"https://github.com/gatk-workflows/intel-gatk3-4-germline-snps-indels"},
	},
	{
		Name: "wkfl/gatk/intel-gatk3-germline-snps-indels",
		URL:  []string{"https://github.com/gatk-workflows/intel-gatk3-germline-snps-indels"},
	},
	{
		Name: "wkfl/gatk/intel-gatk4-germline-snps-indels",
		URL:  []string{"https://github.com/gatk-workflows/intel-gatk4-germline-snps-indels"},
	},
	{
		Name: "wkfl/gatk/intel-gatk4-somatic-with-preprocessing",
		URL:  []string{"https://github.com/gatk-workflows/intel-gatk4-somatic-with-preprocessing"},
	},
	{
		Name: "wkfl/gatk/seq-format-conversion",
		URL:  []string{"https://github.com/gatk-workflows/seq-format-conversion"},
	},
	{
		Name: "wkfl/gatk/seq-format-validation",
		URL:  []string{"https://github.com/gatk-workflows/seq-format-validation"},
	},
	{
		Name: "wkfl/gatk/genomicsITER-developers/wdl",
		URL:  []string{"https://github.com/genomicsITER-developers/wdl"},
	},
	{
		Name: "wkfl/hongiiv/gatk-workflows",
		URL:  []string{"https://github.com/hongiiv/gatk-workflows"},
	},
	{
		Name: "wkfl/jimmybgammyknee/uofaBioWDL",
		URL:  []string{"https://github.com/jimmybgammyknee/uofaBioWDL"},
	},
	{
		Name: "wkfl/jimmybgammyknee/workflows",
		URL:  []string{"https://github.com/jimmybgammyknee/workflows"},
	},
	{
		Name: "wkfl/johnmous/ChIPseq",
		URL:  []string{"https://github.com/johnmous/ChIPseq"},
	},
	{
		Name: "wkfl/karlrl/mgs_sop",
		URL:  []string{"https://github.com/karlrl/mgs_sop"},
	},
	{
		Name: "wkfl/Kurt-Hetrick/WDL",
		URL:  []string{"https://github.com/Kurt-Hetrick/WDL"},
	},
	{
		Name: "wkfl/manning-lab/saigeWdl",
		URL:  []string{"https://github.com/manning-lab/saigeWdl"},
	},
	{
		Name: "wkfl/manning-lab/wgsaParsrWdl",
		URL:  []string{"https://github.com/manning-lab/wgsaParsrWdl"},
	},
	{
		Name: "wkfl/mkim55/deepvar-wdl",
		URL:  []string{"https://github.com/mkim55/deepvar-wdl"},
	},
	{
		Name: "wkfl/ncsa/MayomicsVC",
		URL:  []string{"https://github.com/ncsa/MayomicsVC"},
	},
	{
		Name: "wkfl/omixer/metagenophile",
		URL:  []string{"https://github.com/omixer/metagenophile"},
	},
	{
		Name: "wkfl/oskarvid/wdl_germline_pipeline",
		URL:  []string{"https://github.com/oskarvid/wdl_germline_pipeline"},
	},
	{
		Name: "wkfl/ottojolanki/kallisto-wdl",
		URL:  []string{"https://github.com/ottojolanki/kallisto-wdl"},
	},
	{
		Name: "wkfl/raisok/RNAseq_WDL",
		URL:  []string{"https://github.com/raisok/RNAseq_WDL"},
	},
	{
		Name: "wkfl/SanderEST/cromwell_wdl",
		URL:  []string{"https://github.com/SanderEST/cromwell_wdl"},
	},
	{
		Name: "wkfl/Seonwhee-Genome/GATK_WDL",
		URL:  []string{"https://github.com/Seonwhee-Genome/GATK_WDL"},
	},
}

func init() {
	wkflFiles = append(wkflFiles, wdlWkflFiles...)
}
