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
	"10.1021":  PubsacsSpider,
	"10.1029":  WileyComSpider,
	"10.1038":  NatureComSpider,
	"10.1039":  PubsRscSpider,
	"10.1042":  PortlandpressComSpider,
	"10.1053":  CellComSpider,
	"10.1055":  ThiemeConnectDeSpider,
	"10.1056":  NejmSpider,
	"10.1073":  PnasSpider,
	"10.1080":  TandfonlineSpider,
	"10.1086":  AddPdfplusSpider,
	"10.1089":  LiebertpubSpider,
	"10.1093":  OupComSpider,
	"10.1094":  AddPdfplusSpider,
	"10.1096":  AddPdfSpider,
	"10.1097":  LwwComSpider,
	"10.1098":  RoyalsocietypublishingOrgSpider,
	"10.1101":  CshlpSpider,
	"10.1103":  JournalsApsSpider,
	"10.1107":  IucrOrgSpider,
	"10.1109":  IeeexploreSpider,
	"10.1111":  WileyComSpider,
	"10.1126":  ScienseComSpider,
	"10.1130":  GeoscienceworldOrg,
	"10.1137":  AddPdfSpider,
	"10.1136":  BmjComSpider,
	"10.1145":  AddPdfSpider,
	"10.1146":  AnnualReviewsSpider,
	"10.1152":  PhysiologyOrgSpider,
	"10.1158":  AacrJournalsSpider,
	"10.1161":  AhajournalsSpider,
	"10.1164":  AddPdfWithSupplSpider,
	"10.1172":  JciSpider,
	"10.1175":  AmetsocOrgSpider,
	"10.1176":  AddPdfplusSpider,
	"10.1177":  SagepubComSpider,
	"10.1182":  BloodJournalSpider,
	"10.1186":  BiomedcentralSpider,
	"10.1200":  AscopubsSpider,
	"10.1210":  OupComSpider,
	"10.1246":  AddPdfSpider,
	"10.1257":  AeawebOrgSpider,
	"10.1287":  InformsOrgSPider,
	"10.1289":  AddPdfSpider,
	"10.1371":  PlosSpider,
	"10.1373":  OupComSpider,
	"10.1513":  AddPdfWithSupplSpider,
	"10.1556":  AddPdfSpider,
	"10.1681":  AsnjournalsOrgSpider,
	"10.2105":  AddPdfplusWithSupplSpider,
	"10.2113":  GeoscienceworldOrgSpider,
	"10.2136":  WileyComSpider,
	"10.2147":  DovepressSpider,
	"10.2196":  JmirOrgSpider,
	"10.2217":  AddPdfplusSpider,
	"10.2340":  AddDownloadSpider,
	"10.2471":  DirectSpider,
	"10.2807":  EurosurveillanceOrgSpider,
	"10.3102":  SagepubComSpider,
	"10.3168":  CellComSpider,
	"10.3233":  IospressComSpider,
	"10.3238":  AerzteblattDeSpider,
	"10.3322":  WileyComSpider,
	"10.3324":  HaematologicaSpider,
	"10.3348":  KjronlineOrgSpider,
	"10.3386":  DirectSpider,
	"10.3389":  FrontiersinSpider,
	"10.3835":  WileyComSpider,
	"10.3982":  WileyComSpider,
	"10.4110":  ImmunenetworkOrgSpider,
	"10.4155":  AddPdfplusSpider,
	"10.4158":  AddPdfSpider,
	"10.4322":  AutopsyandcasereportsSpider,
	"10.5281":  ZenodoSpider,
	"10.5578":  KosuyoluheartjournalSpider,
	"10.5598":  BiomedcentralSpider,
	"10.5670":  TosOrgSpider,
	"10.5694":  WileyComSpider,
	"10.5665":  OupComSpider,
	"10.5853":  JstrokeOrgSpider,
	"10.6084":  FigshareSpider,
	"10.7150":  ThnoOrgSpider,
	"10.7185":  GeochemicalperspectivesOrgSpider,
	"10.7189":  DirectSpider,
	"10.7295":  CellimageLibrarySpider,
	"10.7326":  AnnalsOrgSpider,
	"10.7554":  ElifeSpider,
	"10.7717":  PeerjSpider,
	"10.12890": EjcrimSpider,
	"10.14309": LwwComSpider,
	"10.15252": EmbopressSpider,
	"10.17645": CogitatiopressComSpider,
	"10.17712": DirectSpider,
	"10.18637": JstatsoftSpider,
	"10.20882": AdiccionesEsSpider,
	"10.21037": AmegroupsSpider,
	"10.35946": ReplaceHtmlSpider,
}
