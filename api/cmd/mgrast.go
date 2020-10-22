package cmd

import (
	"github.com/clindet/bget/api/fetch"
	"github.com/clindet/bget/api/types"
	"github.com/spf13/cobra"
)

var MgRastEndp types.MgRastEndpoints
var MgRastCmd = &cobra.Command{
	Use:   "mgrast",
	Short: "Query mg-rast website APIs.",
	Long:  `Query mg-rast website APIs. Detail see http://api.mg-rast.org/api.html`,
	Run: func(cmd *cobra.Command, args []string) {
		MgRastCmdRunOptions(cmd, args)
	},
}

var MgRastAnnoCmd = &cobra.Command{
	Use:   "anno",
	Short: "Query http://api.mg-rast.org/anno website annotation APIs.",
	Long:  `Query http://api.mg-rast.org/anno website annotation APIs. Detail see http://api.mg-rast.org/api.html`,
	Run: func(cmd *cobra.Command, args []string) {
		MgRastEndp.Annotation = true
		MgRastCmdRunOptions(cmd, args)
	},
}

var MgRastComputeCmd = &cobra.Command{
	Use:   "compute",
	Short: "Query http://api.mg-rast.org/compute website compute APIs.",
	Long:  `Query http://api.mg-rast.org/compute website compute APIs. Detail see http://api.mg-rast.org/api.html`,
	Run: func(cmd *cobra.Command, args []string) {
		MgRastEndp.Compute = true
		MgRastEndp.ParamsCompute.Retry = bapiClis.Retries
		MgRastCmdRunOptions(cmd, args)
	},
}

var MgRastDarkMatterCmd = &cobra.Command{
	Use:   "darkmatter",
	Short: "Query http://api.mg-rast.org/darkmatter website compute APIs.",
	Long:  `Query http://api.mg-rast.org/darkmatter website compute APIs. Detail see http://api.mg-rast.org/api.html`,
	Run: func(cmd *cobra.Command, args []string) {
		MgRastEndp.DarkMatter = true
		MgRastCmdRunOptions(cmd, args)
	},
}

var MgRastDownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Query http://api.mg-rast.org/download website compute APIs.",
	Long:  `Query http://api.mg-rast.org/download website compute APIs. Detail see http://api.mg-rast.org/api.html`,
	Run: func(cmd *cobra.Command, args []string) {
		MgRastEndp.Download = true
		MgRastCmdRunOptions(cmd, args)
	},
}

var MgRastProjectCmd = &cobra.Command{
	Use:   "proj",
	Short: "Query http://api.mg-rast.org/project website compute APIs.",
	Long:  `Query http://api.mg-rast.org/project website compute APIs. Detail see http://api.mg-rast.org/api.html`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			MgRastEndp.Project = args[0]
		} else {
			MgRastEndp.Project = "nil"
		}
		MgRastCmdRunOptions(cmd, args)
	},
}

var MgRastLibraryCmd = &cobra.Command{
	Use:   "library [id]",
	Short: "Query http://api.mg-rast.org/library website compute APIs.",
	Long:  `Query http://api.mg-rast.org/library website compute APIs. Detail see http://api.mg-rast.org/api.html`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			MgRastEndp.Library = args[0]
		} else {
			MgRastEndp.Library = "nil"
		}
		MgRastCmdRunOptions(cmd, args)
	},
}

var MgRastSampleCmd = &cobra.Command{
	Use:   "sample [id]",
	Short: "Query http://api.mg-rast.org/sample website compute APIs.",
	Long:  `Query http://api.mg-rast.org/sample website compute APIs. Detail see http://api.mg-rast.org/api.html`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			MgRastEndp.Sample = args[0]
		} else {
			MgRastEndp.Sample = "nil"
		}
		MgRastCmdRunOptions(cmd, args)
	},
}

func MgRastCmdRunOptions(cmd *cobra.Command, args []string) {
	if fetch.MgRast(&MgRastEndp, &bapiClis, func() { initCmd(cmd, args) }, nil) {
		bapiClis.HelpFlags = false
	}
	if bapiClis.HelpFlags {
		cmd.Help()
	}
}

func init() {
	setGlobalFlag(MgRastCmd, &bapiClis)

	// annotation module
	MgRastAnnoCmd.Flags().BoolVarP(&MgRastEndp.Similarity, "sim", "", false, `annotate the similarity`)
	MgRastAnnoCmd.Flags().StringVarP(&MgRastEndp.ParamsAnno.ID, "id", "", "", `unique metagenome identifier`)
	MgRastAnnoCmd.Flags().StringVarP(&MgRastEndp.ParamsAnno.Type, "type", "", "", `can be chosen from the following: organism, function, ontology, feature, all`)
	MgRastAnnoCmd.Flags().StringVarP(&MgRastEndp.ParamsAnno.Source, "source", "", "", `can be chosen from the following: RefSeq, GenBank, IMG, SEED, TrEMBL, SwissProt, PATRIC, KEGG, eggNOG, RDP, Greengenes, LSU, SSU, ITS, Subsystems, NOG, COG, KO`)
	MgRastAnnoCmd.Flags().StringVarP(&MgRastEndp.ParamsAnno.Filter, "filter", "", "", `text string to filter annotations by: only return those that contain text`)
	MgRastAnnoCmd.Flags().StringVarP(&MgRastEndp.ParamsAnno.FilterLevel, "filter-level", "", "", `hierarchal level to filter annotations by, for organism or ontology only`)
	MgRastAnnoCmd.Flags().IntVarP(&MgRastEndp.ParamsAnno.Identity, "identity", "", 60, `percent value for minimum % identity cutoff: default is 60`)
	MgRastAnnoCmd.Flags().IntVarP(&MgRastEndp.ParamsAnno.Length, "length", "", 15, `value for minimum alignment length cutoff: default is 15`)
	MgRastAnnoCmd.Flags().BoolVarP(&MgRastEndp.ParamsAnno.NoCutoffs, "no-cutoffs", "", false, `do not use any cutoffs. default is to use default cutoffs`)
	MgRastAnnoCmd.Flags().IntVarP(&MgRastEndp.ParamsAnno.Version, "vers", "", 1, `M5NR version, default is 1`)
	MgRastAnnoCmd.Flags().StringVarP(&MgRastEndp.Md5s, "md5s", "", "", `given md5sums (M5NR ID)`)
	MgRastAnnoCmd.Flags().IntVarP(&MgRastEndp.ParamsAnno.Evalue, "evalue", "", 5, `negative exponent value for maximum e-value cutoff`)
	MgRastAnnoCmd.Flags().StringVarP(&MgRastEndp.ParamsAnno.Format, "format", "", "tab", `tab or fasta`)

	// compute module
	MgRastComputeCmd.Flags().BoolVarP(&MgRastEndp.ComputeAlphadiversity, "alphadiversity", "", false, `compute alphadiversity`)
	MgRastComputeCmd.Flags().BoolVarP(&MgRastEndp.ComputeRarefaction, "rarefaction", "", false, `compute rarefaction`)
	MgRastComputeCmd.Flags().BoolVarP(&MgRastEndp.ComputeBlast, "blast", "", false, `compute blast`)
	MgRastComputeCmd.Flags().BoolVarP(&MgRastEndp.ComputeNormalize, "normalize", "", false, `compute normalize`)
	MgRastComputeCmd.Flags().BoolVarP(&MgRastEndp.ComputeDistance, "distance", "", false, `compute distance`)
	MgRastComputeCmd.Flags().BoolVarP(&MgRastEndp.ComputeHeatmap, "heatmap", "", false, `compute heatmap`)
	MgRastComputeCmd.Flags().BoolVarP(&MgRastEndp.ComputePcoa, "pcoa", "", false, `compute pcoa`)

	MgRastComputeCmd.Flags().BoolVarP(&MgRastEndp.ParamsCompute.Asynchronous, "async", "", false, `if true return process id to query status resource for results, default is false`)
	MgRastComputeCmd.Flags().BoolVarP(&MgRastEndp.ParamsCompute.Alpha, "return-alpha", "", false, `if true also return alphadiversity, default is false`)
	MgRastComputeCmd.Flags().BoolVarP(&MgRastEndp.ParamsCompute.Rna, "rna", "", false, `if true input md5sum is RNA feature, default is false (md5sum is protein)`)
	MgRastComputeCmd.Flags().IntVarP(&MgRastEndp.ParamsCompute.Raw, "raw", "", 0, `option to use raw data (not normalize) [0 or 1]`)
	MgRastComputeCmd.Flags().IntVarP(&MgRastEndp.ParamsCompute.AnnVer, "ann-ver", "", 1, `M5NR annotation version, default 1`)
	MgRastComputeCmd.Flags().StringVarP(&MgRastEndp.ParamsCompute.Level, "level", "", "strain", `can be chosen from the following (the first being default): strain, species, genus, family, order, class, phylum, domain`)
	MgRastComputeCmd.Flags().IntVarP(&MgRastEndp.ParamsCompute.SeqNum, "seq-num", "", 0, `number of sequences in metagenome`)
	MgRastComputeCmd.Flags().StringVarP(&MgRastEndp.ParamsCompute.Md5, "md5", "", "", `md5sum of M5NR feature to search against`)
	MgRastComputeCmd.Flags().StringVarP(&MgRastEndp.Columns, "columns", "", "", `column id`)
	MgRastComputeCmd.Flags().StringVarP(&MgRastEndp.Data, "data", "", "", `raw or normalized value`)
	MgRastComputeCmd.Flags().StringVarP(&MgRastEndp.ParamsCompute.Norm, "norm", "", "DESeq_blind", `can be chosen from the following (the first being default): DESeq_blind, standardize, quantile, DESeq_per_condition, DESeq_pooled, DESeq_pooled_CR`)
	MgRastComputeCmd.Flags().StringVarP(&MgRastEndp.Rows, "rows", "", "", `row id`)
	MgRastComputeCmd.Flags().StringVarP(&MgRastEndp.ParamsCompute.Distance, "dist", "", "bray-curtis", ` can be chosen from the following (the first being default): bray-curtis, euclidean, maximum, manhattan, canberra, minkowski, difference`)
	MgRastComputeCmd.Flags().StringVarP(&MgRastEndp.ParamsCompute.Cluster, "cluster", "", "ward", `can be chosen from the following (the first being default): ward, single, complete, mcquitty, median, centroid`)
	MgRastComputeCmd.Flags().IntVarP(&MgRastEndp.ParamsCompute.Evalue, "evalue", "", 5, `negative exponent value for maximum e-value cutoff`)

	MgRastDarkMatterCmd.Flags().StringVarP(&MgRastEndp.ID, "id", "", "", `unique metagenome identifier`)

	MgRastDownloadCmd.Flags().BoolVarP(&MgRastEndp.DownloadHistory, "history", "", false, `summary of MG-RAST analysis-pipeline workflow and commands.`)
	MgRastDownloadCmd.Flags().StringVarP(&MgRastEndp.ParamsDownload.File, "file", "", "", `file name or identifier`)
	MgRastDownloadCmd.Flags().StringVarP(&MgRastEndp.ParamsDownload.AweID, "awe-id", "", "", `optional: AWE ID of MG-RAST metagenome`)
	MgRastDownloadCmd.Flags().StringVarP(&MgRastEndp.ParamsDownload.Stage, "stage", "", "", `stage name or identifier`)
	MgRastDownloadCmd.Flags().BoolVarP(&MgRastEndp.ParamsDownload.Link, "link", "", false, `if true return one time link for download and not file stream`)
	MgRastDownloadCmd.Flags().BoolVarP(&MgRastEndp.ParamsDownload.Delete, "delete", "", false, `if true (and user is admin) delete original document from AWE on completion.`)
	MgRastDownloadCmd.Flags().BoolVarP(&MgRastEndp.ParamsDownload.Force, "force", "", false, `if true, recreate document in Shock from AWE.`)
	MgRastDownloadCmd.Flags().StringVarP(&MgRastEndp.ID, "id", "", "", `unique metagenome identifier`)

	MgRastProjectCmd.Flags().IntVarP(&MgRastEndp.ParamsProjOrLibrary.Limit, "size", "", 20, `maximum number of items requested`)
	MgRastProjectCmd.Flags().IntVarP(&MgRastEndp.ParamsProjOrLibrary.Offset, "from", "", 0, `zero based index of the first data object to be returned`)
	MgRastProjectCmd.Flags().StringVarP(&MgRastEndp.ParamsProjOrLibrary.Order, "order", "", "id", `This parameter value can be chosen from the following (the first being default): id, name`)
	MgRastProjectCmd.Flags().StringVarP(&MgRastEndp.ParamsProjOrLibrary.Verbosity, "verbosity", "", "minimal", `This parameter value can be chosen from the following (the first being default): minimal, verbose, full`)

	MgRastLibraryCmd.Flags().IntVarP(&MgRastEndp.ParamsProjOrLibrary.Limit, "size", "", 20, `maximum number of items requested`)
	MgRastLibraryCmd.Flags().IntVarP(&MgRastEndp.ParamsProjOrLibrary.Offset, "from", "", 0, `zero based index of the first data object to be returned`)
	MgRastLibraryCmd.Flags().StringVarP(&MgRastEndp.ParamsProjOrLibrary.Order, "order", "", "id", `This parameter value can be chosen from the following (the first being default): id, name`)
	MgRastLibraryCmd.Flags().StringVarP(&MgRastEndp.ParamsProjOrLibrary.Verbosity, "verbosity", "", "minimal", `This parameter value can be chosen from the following (the first being default): minimal, verbose, full`)

	//MgRastCmd.Flags().BoolVarP(&MgRastEndp.Download, "download", "", false, ``)
	//MgRastCmd.Flags().BoolVarP(&MgRastEndp.Inbox, "inbox", "", false, ``)
	//MgRastCmd.Flags().BoolVarP(&MgRastEndp.Library, "lib", "", false, ``)
	//MgRastCmd.Flags().BoolVarP(&MgRastEndp.M5nr, "m5nr", "", false, ``)
	//MgRastCmd.Flags().BoolVarP(&MgRastEndp.Matrix, "mat", "", false, ``)
	//MgRastCmd.Flags().BoolVarP(&MgRastEndp.MetaData, "meta", "", false, ``)
	//MgRastCmd.Flags().BoolVarP(&MgRastEndp.MetaGenome, "gmeta", "", false, ``)
	//MgRastCmd.Flags().BoolVarP(&MgRastEndp.Mixs, "mixs", "", false, ``)
	//MgRastCmd.Flags().BoolVarP(&MgRastEndp.Profile, "profile", "", false, ``)
	//MgRastCmd.Flags().BoolVarP(&MgRastEndp.Project, "project", "", false, ``)
	//MgRastCmd.Flags().BoolVarP(&MgRastEndp.ResearchObject, "res-obj", "", false, ``)
	//MgRastCmd.Flags().BoolVarP(&MgRastEndp.Sample, "sample", "", false, ``)
	//MgRastCmd.Flags().BoolVarP(&MgRastEndp.Search, "search", "", false, ``)
	//MgRastCmd.Flags().BoolVarP(&MgRastEndp.Submission, "submit", "", false, ``)
	//MgRastCmd.Flags().BoolVarP(&MgRastEndp.Validation, "valid", "", false, ``)

	MgRastCmd.PersistentFlags().StringVarP(&MgRastEndp.Sequence, "seq", "", "", `set sequence`)
	MgRastCmd.PersistentFlags().BoolVarP(&MgRastEndp.Info, "info", "", false, `returns description of parameters and attributes.`)
	MgRastCmd.PersistentFlags().StringVarP(&MgRastEndp.Auth, "auth", "", "", `auth key`)

	MgRastCmd.AddCommand(MgRastAnnoCmd)
	MgRastCmd.AddCommand(MgRastComputeCmd)
	MgRastCmd.AddCommand(MgRastDarkMatterCmd)
	MgRastCmd.AddCommand(MgRastDownloadCmd)
	MgRastCmd.AddCommand(MgRastProjectCmd)
	MgRastCmd.AddCommand(MgRastLibraryCmd)
	MgRastCmd.AddCommand(MgRastSampleCmd)

	MgRastAnnoCmd.Example = `  bget api mgrast anno --info
  # retrieval of SwissProt taxonomy annotations with a cut-off 10^100 for dataset mgm4447943.3
  bget api mgrast anno --evalue 100 --type organism --source SwissProt --seq mgm4447943.3
  bget api mgrast anno --filter Immunoreactive --type function --source SEED --seq mgm4447943.3
  # annotated read sequences from mgm4447943.3 with hits in SwissProt organisms for given md5s
  bget api mgrast anno --type organism --source RefSeq --seq mgm4447943.3 --md5s 000a8d74068603c9e8674bff9970f367,0001c08aa276d154b7696f9758839786
  bget api mgrast anno --sim --type organism --source RefSeq --seq mgm4447943.3 --md5s 000a8d74068603c9e8674bff9970f367,0001c08aa276d154b7696f9758839786
  bget api mgrast anno --sim --identity 80 --type function --source KO --seq mgm4447943.3`

	MgRastComputeCmd.Example = `  bget api mgrast compute --info
  bget api mgrast compute --alphadiversity --seq mgm4447943.3 --level order
  bget api mgrast compute --rarefaction --seq mgm4447943.3 --level order
  bget api mgrast compute --blast --seq mgm4447943.3 --md5 0001c08aa276d154b7696f9758839786
  bget api mgrast compute --pcoa --raw 1 --dist euclidean --columns mgm4441619.3,mgm4441656.4,mgm4441680.3,mgm4441681.3 --rows Eukaryota,Bacteria,Archaea --data '[135,410,848,1243],[4397,6529,71423,204413],[1422,2156,874,1138]'
  bget api mgrast compute --normalize --columns mgm4441619.3,mgm4441656.4,mgm4441680.3,mgm4441681.3 --rows Eukaryota,Bacteria,Archaea --data '[135,410,848,1243],[4397,6529,71423,204413],[1422,2156,874,1138]'
  bget api mgrast compute --distance --dist euclidean --columns mgm4441619.3,mgm4441656.4,mgm4441680.3,mgm4441681.3 --rows Eukaryota,Bacteria,Archaea --data '[135,410,848,1243],[4397,6529,71423,204413],[1422,2156,874,1138]'
  bget api mgrast compute --heatmap --cluster mcquitty --columns mgm4441619.3,mgm4441656.4,mgm4441680.3,mgm4441681.3 --rows Eukaryota,Bacteria,Archaea --data '[135,410,848,1243],[4397,6529,71423,204413],[1422,2156,874,1138]' --raw 0`

	MgRastDarkMatterCmd.Example = `  bget api mgrast darkmatter --id mgm4447943.3`

	MgRastDownloadCmd.Example = `  bget api mgrast download --id mgm4447943.3 --file 350.1
  bget api mgrast download --history --id mgm4447943.3
	bget api mgrast download --id mgm4447943.3 --stage 650`

	MgRastProjectCmd.Example = `  bget api mgrast proj
	bget api mgrast proj mgp128 --verbosity full`

	MgRastLibraryCmd.Example = `  bget api mgrast library 
	bget api mgrast library mgl52924 --verbosity full`

	MgRastSampleCmd.Example = `  bget api mgrast sample 
  bget api mgrast sample mgs25823 --verbosity full`
}
