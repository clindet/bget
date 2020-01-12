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
