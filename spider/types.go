package spider

import neturl "net/url"

type DoiSpiderOpt struct {
	Doi               string
	Proxy             string
	Timeout           int
	FullText          bool
	Supplementary     bool
	Citations         bool
	PrintSiteMeta     bool
	PrintCrossRefMeta bool
	CitationMeta      *map[string]string
	URL               *neturl.URL
}
type QuerySpiderOpt struct {
	Query   string
	Proxy   string
	Timeout int
}
