package spider

type doiSpiderType struct {
	doiOrg string
	spider map[string]interface{}
}

// DoiSpidersPool map doi to golang function
var DoiSpidersPool = map[string]func(doi string) []string{
	"10.1001":    JamaNetworkSpider,
	"10.1002":    WileyComSpider,
	"10.1016":    CellComSpider,
	"10.1038":    NatureComSpider,
	"10.1053":    CellComSpider,
	"10.1056":    NejmSpider,
	"10.1073":    PnasSpider,
	"10.1093":    OupComSpider,
	"10.1101":    BiorxivSpider,
	"10.1101.gr": GenomeResSpider,
	"10.1111":    WileyComSpider,
	"10.1126":    ScienseComSpider,
	"10.1161":    AhajournalsSpider,
	"10.1182":    BloodJournalSpider,
	"10.1186":    BiomedcentralSpider,
	"10.1200":    AscopubsSpider,
	"10.1371":    PlosSpider,
	"10.15252":   EmbopressSpider,
	"10.1158":    AacrJournalsSpider,
	"10.3322":    WileyComSpider,
	"10.3324":    HaematologicaSpider,
	"10.3389":    FrontiersinSpider,
	"10.5281":    ZenodoSpider,
	"10.7717":    PeerjSpider,
}
