package spider

import neturl "net/url"

type doiSpiderType struct {
	doiOrg string
	spider map[string]interface{}
}
type DoiSpiderOpt struct {
	Doi           string
	Proxy         string
	Timeout       int
	FullText      bool
	Supplementary bool
	Citations     bool
	URL           *neturl.URL
}
type QuerySpiderOpt struct {
	Query   string
	Proxy   string
	Timeout int
}

// DoiSpidersPool map doi to golang function
var DoiSpidersPool = map[string]func(opt *DoiSpiderOpt) []string{
	"10.1001":  JamaNetworkSpider,
	"10.1002":  WileyComSpider,
	"10.1016":  CellComSpider,
	"10.1038":  NatureComSpider,
	"10.1053":  CellComSpider,
	"10.1056":  NejmSpider,
	"10.1073":  PnasSpider,
	"10.1080":  TandfonlineSpider,
	"10.1093":  OupComSpider,
	"10.1101":  CshlpSpider,
	"10.1111":  WileyComSpider,
	"10.1164":  AtsjournalsOrgSpider,
	"10.1126":  ScienseComSpider,
	"10.1136":  BmjComSpider,
	"10.1158":  AacrJournalsSpider,
	"10.1161":  AhajournalsSpider,
	"10.1172":  JciSpider,
	"10.1182":  BloodJournalSpider,
	"10.1186":  BiomedcentralSpider,
	"10.1200":  AscopubsSpider,
	"10.1371":  PlosSpider,
	"10.2147":  DovepressSpider,
	"10.3322":  WileyComSpider,
	"10.3324":  HaematologicaSpider,
	"10.3389":  FrontiersinSpider,
	"10.4322":  AutopsyandcasereportsSpider,
	"10.5281":  ZenodoSpider,
	"10.5578":  KosuyoluheartjournalSpider,
	"10.5665":  OupComSpider,
	"10.7554":  ElifeSpider,
	"10.7717":  PeerjSpider,
	"10.12890": EjcrimSpider,
	"10.15252": EmbopressSpider,
	"10.18637": JstatsoftSpider,
}
