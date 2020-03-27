package types

// CligovEndpoints is http://clinicaltrials.gov/ website endpoints
type CligovEndpoints struct {
	InfoDataVrs       bool
	InfoAPIVrs        bool
	InfoAPIDefs       bool
	InfoStuStru       bool
	InfoStuFieldsList bool
	InfoStuStat       bool
	InfoSearchArea    bool
	FullStudies       bool
	StuFields         bool
	FieldValues       bool
	Field             string
	Fields            string
}

// Datasets2toolsEndpoints is Datasets2tools website endpoints
type Datasets2toolsEndpoints struct {
	ObjectType              string
	PageSize                int
	DatasetAccession        string
	CannedAnalysisAccession string
	Query                   string
	ToolName                string
	DiseaseName             string
	Gneset                  string
}

// BioToolsEndpoints is bio.tools website endpoints
type BioToolsEndpoints struct {
	Tool         string
	ID           string
	Name         string
	Topic        string
	DataType     string
	DataFormat   string
	OutputFormat string
	Publication  string
}

type MgRastAnnoEndpoints struct {
	Evalue      int      `json:"evalue"`
	Filter      string   `json:"filter"`
	FilterLevel string   `json:"filter_level"`
	Format      string   `json:"format"`
	Identity    int      `json:"identity"`
	Length      int      `json:"length"`
	NoCutoffs   bool     `json:"no_cutoffs"`
	Source      string   `json:"source"`
	Type        string   `json:"type"`
	Version     int      `json:"version"`
	ID          string   `json:"id"`
	Md5s        []string `json:"md5s"`
}

type MgRastComputeEndpoints struct {
	Asynchronous bool        `json:"asynchronous"`
	Evalue       int         `json:"evalue"`
	Alpha        bool        `json:"alpha"`
	Rna          bool        `json:"rna"`
	Raw          int         `json:"raw"`
	AnnVer       int         `json:"ann_ver"`
	Level        string      `json:"level"`
	SeqNum       int         `json:"seq_num"`
	Md5          string      `json:"md5"`
	Columns      []string    `json:"columns"`
	Data         [][]float64 `json:"data"`
	Norm         string      `json:"norm"`
	Rows         []string    `json:"rows"`
	Distance     string      `json:"distance"`
	Cluster      string      `json:"cluster"`
	ID           string      `json:"id"`
	Retry        int         `json:"retry"`
}

type MgRastDownloadEndpoints struct {
	File   string `json:"file"`
	Link   bool   `json:"link"`
	AweID  string `json:"awe_id"`
	Delete bool   `json:"delete"`
	Force  bool   `json:"force"`
	Stage  string `json:"stage"`
}

type MgRastProjLibEndpoints struct {
	Limit     int    `json:"limit"`
	Offset    int    `json:"offset"`
	Order     string `json:"order"`
	Verbosity string `json:"Verbosity"`
}

type MgRastEndpoints struct {
	Info bool

	Annotation     bool   `json:"annotation"`
	Similarity     bool   `json:"similarity"`
	Sequence       string `json:"sequence"`
	Project        string
	ResearchObject string
	Library        string
	Sample         string

	ParamsAnno          MgRastAnnoEndpoints
	ParamsCompute       MgRastComputeEndpoints
	ParamsDownload      MgRastDownloadEndpoints
	ParamsProjOrLibrary MgRastProjLibEndpoints

	Compute               bool
	ComputeAlphadiversity bool
	ComputeRarefaction    bool
	ComputeBlast          bool
	ComputeNormalize      bool
	ComputeDistance       bool
	ComputeHeatmap        bool
	ComputePcoa           bool
	Rows                  string
	Data                  string
	Md5s                  string
	Columns               string
	ID                    string

	DarkMatter bool

	Download        bool
	DownloadHistory bool

	Inbox      bool
	M5nr       bool
	Matrix     bool
	MetaData   bool
	MetaGenome bool
	Mixs       bool
	Profile    bool
	Search     bool
	Submission bool
	Validation bool

	Auth string
}
