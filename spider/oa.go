package spider

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	cnet "github.com/openbiox/ligo/net"
	"github.com/openbiox/ligo/stringo"
)

// ZenodoSpider access Zendo files via spider
func ZenodoSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "zenodo.org"),
		colly.MaxDepth(1),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.FullText {
		c.OnHTML("tbody a.filename[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if strings.Contains(link, "?download=1") && opt.FullText {
				u, _ := url.Parse(link)
				link = "https://zenodo.org" + u.Host + u.Path
				urls = append(urls, link)
			}
		})
	}
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

// CshlpSpider access CshlpSpider files via spider
func CshlpSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "biorxiv.org", "www.biorxiv.org",
			"genome.cshlp.org", "genesdev.cshlp.org"),
		colly.MaxDepth(1),
	)
	if opt.URL != nil {
		c.AllowedDomains = append(c.AllowedDomains, opt.URL.Host)
	}
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.FullText {
		c.OnHTML("meta[name=citation_pdf_url]", func(e *colly.HTMLElement) {
			link := e.Attr("content")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	if opt.Supplementary {
		c.OnHTML(".pane-highwire-variant-link a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			u, _ := url.Parse(link)
			link = linkFilter(u.Host+u.Path, opt.URL)
			urls = append(urls, link)
		})
		c.OnHTML(".pane-biorxiv-supplementary-fragment a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			c.Visit(opt.URL.Host + link)
		})
		c.OnHTML(".supplementary-material-expansion a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = stringo.StrReplaceAll(link, "[?]download=true$", "")
			urls = append(urls, link)
		})
		c.OnHTML("div.auto-clean a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = linkFilter(link, opt.URL)
			urls = append(urls, link)
		})
		c.OnHTML("a[rel=supplemental-data]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			c.Visit(linkFilter(link, opt.URL))
		})
	}
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

// BiomedcentralSpider access GenomeBiology files via spider
func BiomedcentralSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "aacijournal.biomedcentral.com", "actaneurocomms.biomedcentral.com", "actavetscand.biomedcentral.com", "advancesinrheumatology.biomedcentral.com", "advancesinsimulation.biomedcentral.com", "aepi.biomedcentral.com", "agricultureandfoodsecurity.biomedcentral.com", "aidsrestherapy.biomedcentral.com", "almob.biomedcentral.com", "alzres.biomedcentral.com", "animalbiotelemetry.biomedcentral.com", "animalmicrobiome.biomedcentral.com", "annals-general-psychiatry.biomedcentral.com", "ann-clinmicrob.biomedcentral.com", "appliedcr.biomedcentral.com", "appliedvolc.biomedcentral.com", "archivesphysiotherapy.biomedcentral.com", "archpublichealth.biomedcentral.com", "aricjournal.biomedcentral.com", "arrhythmia.biomedcentral.com", "arthritis-research.biomedcentral.com", "arthroplasty.biomedcentral.com", "ascpjournal.biomedcentral.com", "asthmarp.biomedcentral.com", "autoimmunhighlights.biomedcentral.com", "avianres.biomedcentral.com", "bacandrology.biomedcentral.com", "bdataanalytics.biomedcentral.com", "behavioralandbrainfunctions.biomedcentral.com", "biodatamining.biomedcentral.com", "bioelecmed.biomedcentral.com", "biologicalproceduresonline.biomedcentral.com", "biologydirect.biomedcentral.com", "biolres.biomedcentral.com", "biomarkerres.biomedcentral.com", "biomaterialsres.biomedcentral.com", "biomeddermatol.biomedcentral.com", "biomedical-engineering-online.biomedcentral.com", "biosignaling.biomedcentral.com", "biotechnologyforbiofuels.biomedcentral.com", "bmcanesthesiol.biomedcentral.com", "bmcbiochem.biomedcentral.com", "bmcbioinformatics.biomedcentral.com", "bmcbiol.biomedcentral.com", "bmcbiomedeng.biomedcentral.com", "bmcbiophys.biomedcentral.com", "bmcbiotechnol.biomedcentral.com", "bmccancer.biomedcentral.com", "bmccardiovascdisord.biomedcentral.com", "bmcchem.biomedcentral.com", "bmcchemeng.biomedcentral.com", "bmcclinpathol.biomedcentral.com", "bmccomplementalternmed.biomedcentral.com", "bmcdermatol.biomedcentral.com", "bmcdevbiol.biomedcentral.com", "bmcearnosethroatdisord.biomedcentral.com", "bmcecol.biomedcentral.com", "bmcemergmed.biomedcentral.com", "bmcendocrdisord.biomedcentral.com", "bmcenergy.biomedcentral.com", "bmcevolbiol.biomedcentral.com", "bmcfampract.biomedcentral.com", "bmcgastroenterol.biomedcentral.com", "bmcgenet.biomedcentral.com", "bmcgenomics.biomedcentral.com", "bmcgeriatr.biomedcentral.com", "bmchealthservres.biomedcentral.com", "bmchematol.biomedcentral.com", "bmcimmunol.biomedcentral.com", "bmcinfectdis.biomedcentral.com", "bmcinthealthhumrights.biomedcentral.com", "bmcmaterials.biomedcentral.com", "bmcmecheng.biomedcentral.com", "bmcmededuc.biomedcentral.com", "bmcmedethics.biomedcentral.com", "bmcmedgenet.biomedcentral.com", "bmcmedgenomics.biomedcentral.com", "bmcmedicine.biomedcentral.com", "bmcmedimaging.biomedcentral.com", "bmcmedinformdecismak.biomedcentral.com", "bmcmedresmethodol.biomedcentral.com", "bmcmicrobiol.biomedcentral.com", "bmcmolbiol.biomedcentral.com", "bmcmolcellbiol.biomedcentral.com", "bmcmusculoskeletdisord.biomedcentral.com", "bmcnephrol.biomedcentral.com", "bmcneurol.biomedcentral.com", "bmcneurosci.biomedcentral.com", "bmcnurs.biomedcentral.com", "bmcnutr.biomedcentral.com", "bmcobes.biomedcentral.com", "bmcophthalmol.biomedcentral.com", "bmcoralhealth.biomedcentral.com", "bmcpalliatcare.biomedcentral.com", "bmcpediatr.biomedcentral.com", "bmcpharmacoltoxicol.biomedcentral.com", "bmcphysiol.biomedcentral.com", "bmcplantbiol.biomedcentral.com", "bmcpregnancychildbirth.biomedcentral.com", "bmcproc.biomedcentral.com", "bmcpsychiatry.biomedcentral.com", "bmcpsychology.biomedcentral.com", "bmcpublichealth.biomedcentral.com", "bmcpulmmed.biomedcentral.com", "bmcresnotes.biomedcentral.com", "bmcrheumatol.biomedcentral.com", "bmcsportsscimedrehabil.biomedcentral.com", "bmcstructbiol.biomedcentral.com", "bmcsurg.biomedcentral.com", "bmcsystbiol.biomedcentral.com", "bmcurol.biomedcentral.com", "bmcvetres.biomedcentral.com", "bmcwomenshealth.biomedcentral.com", "bmczool.biomedcentral.com", "bpded.biomedcentral.com", "bpsmedicine.biomedcentral.com", "breast-cancer-research.biomedcentral.com", "bsd.biomedcentral.com", "burnstrauma.biomedcentral.com", "cabiagbio.biomedcentral.com", "cancerandmetabolism.biomedcentral.com", "cancerci.biomedcentral.com", "cancercommun.biomedcentral.com", "cancerconvergence.biomedcentral.com", "cancerimagingjournal.biomedcentral.com", "cancer-nano.biomedcentral.com", "cancersheadneck.biomedcentral.com", "capmh.biomedcentral.com", "cardiab.biomedcentral.com", "cardiooncologyjournal.biomedcentral.com", "cardiothoracicsurgery.biomedcentral.com", "cardiovascularultrasound.biomedcentral.com", "cbmjournal.biomedcentral.com", "ccforum.biomedcentral.com", "cellandbioscience.biomedcentral.com", "celldiv.biomedcentral.com", "cerebellumandataxias.biomedcentral.com", "cgejournal.biomedcentral.com", "chiromt.biomedcentral.com", "ciliajournal.biomedcentral.com", "clindiabetesendo.biomedcentral.com", "clinicalepigeneticsjournal.biomedcentral.com", "clinicalhypertension.biomedcentral.com", "clinicalmolecularallergy.biomedcentral.com", "clinicalmovementdisorders.biomedcentral.com", "clinicalproteomicsjournal.biomedcentral.com", "clinicalsarcomaresearch.biomedcentral.com", "cmbl.biomedcentral.com", "cmjournal.biomedcentral.com", "cnjournal.biomedcentral.com", "conflictandhealth.biomedcentral.com", "contraceptionmedicine.biomedcentral.com", "crimesciencejournal.biomedcentral.com", "ctajournal.biomedcentral.com", "diagnosticpathology.biomedcentral.com", "diagnprognres.biomedcentral.com", "dmsjournal.biomedcentral.com", "eandv.biomedcentral.com", "edintegrity.biomedcentral.com", "ehjournal.biomedcentral.com", "ehoonline.biomedcentral.com", "energsustainsoc.biomedcentral.com", "environhealthprevmed.biomedcentral.com", "environmentalevidencejournal.biomedcentral.com", "environmentalmicrobiome.biomedcentral.com", "epigeneticsandchromatin.biomedcentral.com", "equityhealthj.biomedcentral.com", "ete-online.biomedcentral.com", "ethnobiomed.biomedcentral.com", "eurapa.biomedcentral.com", "eurjmedres.biomedcentral.com", "evodevojournal.biomedcentral.com", "evolution-outreach.biomedcentral.com", "exrna.biomedcentral.com", "fas.biomedcentral.com", "fertilityresearchandpractice.biomedcentral.com", "fluidsbarrierscns.biomedcentral.com", "foodcontaminationjournal.biomedcentral.com", "fppn.biomedcentral.com", "frontiersinzoology.biomedcentral.com", "fungalbiolbiotech.biomedcentral.com", "genesandnutrition.biomedcentral.com", "genesenvironment.biomedcentral.com", "genomebiology.biomedcentral.com", "genomemedicine.biomedcentral.com", "geochemicaltransactions.biomedcentral.com", "ghrp.biomedcentral.com", "globalizationandhealth.biomedcentral.com", "gsejournal.biomedcentral.com", "gutpathogens.biomedcentral.com", "harmreductionjournal.biomedcentral.com", "hccpjournal.biomedcentral.com", "head-face-med.biomedcentral.com", "healthandjusticejournal.biomedcentral.com", "healtheconomicsreview.biomedcentral.com", "health-policy-systems.biomedcentral.com", "hereditasjournal.biomedcentral.com", "hmr.biomedcentral.com", "hqlo.biomedcentral.com", "human-resources-health.biomedcentral.com", "humgenomics.biomedcentral.com", "idpjournal.biomedcentral.com", "ijbnpa.biomedcentral.com", "ij-healthgeographics.biomedcentral.com", "ijhpr.biomedcentral.com", "ijmhs.biomedcentral.com", "ijpeonline.biomedcentral.com", "ijponline.biomedcentral.com", "imafungus.biomedcentral.com", "immunityageing.biomedcentral.com", "implementationscience.biomedcentral.com", "implementationsciencecomms.biomedcentral.com", "infectagentscancer.biomedcentral.com", "inflammregen.biomedcentral.com", "injepijournal.biomedcentral.com", "innovationeducation.biomedcentral.com", "internationalbreastfeedingjournal.biomedcentral.com", "intjem.biomedcentral.com", "irishvetjournal.biomedcentral.com", "jasbsci.biomedcentral.com", "jbioleng.biomedcentral.com", "jbiolres.biomedcentral.com", "jbiomedsci.biomedcentral.com", "jbiomedsem.biomedcentral.com", "jcannabisresearch.biomedcentral.com", "jcheminf.biomedcentral.com", "jcmr-online.biomedcentral.com", "jcongenitalcardiology.biomedcentral.com", "jcottonres.biomedcentral.com", "jeatdisord.biomedcentral.com", "jeccr.biomedcentral.com", "jecoenv.biomedcentral.com", "jfootankleres.biomedcentral.com", "jhoonline.biomedcentral.com", "jhpn.biomedcentral.com", "jintensivecare.biomedcentral.com", "jissn.biomedcentral.com", "jitc.biomedcentral.com", "jmedicalcasereports.biomedcentral.com", "jnanobiotechnology.biomedcentral.com", "jneurodevdisorders.biomedcentral.com", "jneuroengrehab.biomedcentral.com", "jneuroinflammation.biomedcentral.com", "joppp.biomedcentral.com", "josr-online.biomedcentral.com", "journal-inflammation.biomedcentral.com", "journalofethnicfoods.biomedcentral.com", "journalotohns.biomedcentral.com", "journalretinavitreous.biomedcentral.com", "jphcs.biomedcentral.com", "jphysiolanthropol.biomedcentral.com", "jps.biomedcentral.com", "kneesurgrelatres.biomedcentral.com", "labanimres.biomedcentral.com", "lipidworld.biomedcentral.com", "lsspjournal.biomedcentral.com", "malariajournal.biomedcentral.com", "mbr.biomedcentral.com", "measurementinstrumentssocialscience.biomedcentral.com", "mhnpjournal.biomedcentral.com", "microbialcellfactories.biomedcentral.com", "microbiomejournal.biomedcentral.com", "mmrjournal.biomedcentral.com", "mobilednajournal.biomedcentral.com", "molecularautism.biomedcentral.com", "molecularbrain.biomedcentral.com", "molecular-cancer.biomedcentral.com", "molecularcytogenetics.biomedcentral.com", "molecularneurodegeneration.biomedcentral.com", "molmed.biomedcentral.com", "movementecologyjournal.biomedcentral.com", "mrmjournal.biomedcentral.com", "msddjournal.biomedcentral.com", "neuraldevelopment.biomedcentral.com", "neurocommons.biomedcentral.com", "neurolrespract.biomedcentral.com", "nutritionandmetabolism.biomedcentral.com", "nutritionj.biomedcentral.com", "occup-med.biomedcentral.com", "ojrd.biomedcentral.com", "onehealthoutlook.biomedcentral.com", "ovarianresearch.biomedcentral.com", "parasitesandvectors.biomedcentral.com", "particleandfibretoxicology.biomedcentral.com", "ped-rheum.biomedcentral.com", "peh-med.biomedcentral.com", "perioperativemedicinejournal.biomedcentral.com", "phytopatholres.biomedcentral.com", "pilotfeasibilitystudies.biomedcentral.com", "plantmethods.biomedcentral.com", "pneumonia.biomedcentral.com", "pophealthmetrics.biomedcentral.com", "porcinehealthmanagement.biomedcentral.com", "proteomesci.biomedcentral.com", "pssjournal.biomedcentral.com", "publichealthreviews.biomedcentral.com", "rbej.biomedcentral.com", "reproductive-health-journal.biomedcentral.com", "researchintegrityjournal.biomedcentral.com", "researchinvolvement.biomedcentral.com", "resource-allocation.biomedcentral.com", "respiratory-research.biomedcentral.com", "retrovirology.biomedcentral.com", "revchilhistnat.biomedcentral.com", "ro-journal.biomedcentral.com", "rrtjournal.biomedcentral.com", "scfbm.biomedcentral.com", "signals.biomedcentral.com", "sjtrem.biomedcentral.com", "skeletalmusclejournal.biomedcentral.com", "sleep.biomedcentral.com", "stemcellres.biomedcentral.com", "substanceabusepolicy.biomedcentral.com", "surgexppathol.biomedcentral.com", "sustainableearth.biomedcentral.com", "sustainenvironres.biomedcentral.com", "systematicreviewsjournal.biomedcentral.com", "tbiomed.biomedcentral.com", "tdtmvjournal.biomedcentral.com", "thejournalofheadacheandpain.biomedcentral.com", "threedmedprint.biomedcentral.com", "thrombosisjournal.biomedcentral.com", "thyroidresearchjournal.biomedcentral.com", "translational-medicine.biomedcentral.com", "translationalneurodegeneration.biomedcentral.com", "transmedcomms.biomedcentral.com", "trialsjournal.biomedcentral.com", "tropmedhealth.biomedcentral.com", "urbantransformations.biomedcentral.com", "veterinaryresearch.biomedcentral.com", "virologyj.biomedcentral.com", "wjes.biomedcentral.com", "wjso.biomedcentral.com", "womensmidlifehealthjournal.biomedcentral.com", "www.biomedcentral.com/journals#top)", "zoologicalletters.biomedcentral.com"),
		colly.MaxDepth(1),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.URL != nil {
		c.AllowedDomains = append(c.AllowedDomains, opt.URL.Host)
	}
	if opt.FullText {
		c.OnHTML(".c-pdf-download a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			u, _ := url.Parse(link)
			link = "https://" + u.Host + u.Path
			urls = append(urls, link)
		})
	}
	if opt.Supplementary {
		c.OnHTML(".c-article-supplementary__item a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if !stringo.StrDetect(link, "^/articles/") {
				link = stringo.StrReplaceAll(link, "[?]download=true$", "")
				urls = append(urls, link)
			}
		})
	}
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

// PnasSpider access PnasSpider files via spider
func PnasSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "www.pnas.org"),
		colly.MaxDepth(1),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.URL != nil {
		c.AllowedDomains = append(c.AllowedDomains, opt.URL.Host)
	}
	if opt.FullText {
		c.OnHTML("a[data-trigger=tab-pdf]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = "https://www.pnas.org" + link
			urls = append(urls, link)
		})
	}
	if opt.Supplementary {
		c.OnHTML("a['data-trigger'='tab-figures-data']", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			c.Visit(link)
		})
		c.OnHTML("a.rewritten[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = "https://www.pnas.org" + link
			urls = append(urls, link)
		})
		c.OnResponse(func(r *colly.Response) {
			if !strings.HasSuffix(r.Request.URL.String(), "/tab-figures-data") {
				c.Visit(r.Request.URL.String() + "/tab-figures-data")
			}
		})
	}
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

// PlosSpider access PlosSpider files via spider
func PlosSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "journals.plos.org", "dx.plos.org"),
		colly.MaxDepth(1),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.URL != nil {
		c.AllowedDomains = append(c.AllowedDomains, opt.URL.Host)
	}
	if opt.FullText {
		c.OnHTML("#downloadPdf", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = "https://journals.plos.org" + link
			urls = append(urls, link)
		})
	}
	if opt.Supplementary {
		c.OnHTML(".supplementary-material a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if strings.Contains(link, "doi.org") {
				urls = append(urls, link)
			}
		})
	}
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

// FrontiersinSpider access Frontiers files via spider
func FrontiersinSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "www.frontiersin.org", "journal.frontiersin.org"),
		colly.MaxDepth(1),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.URL != nil {
		c.AllowedDomains = append(c.AllowedDomains, opt.URL.Host)
	}
	if opt.FullText {
		c.OnHTML("a.download-files-pdf", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = "https://www.frontiersin.org" + link
			urls = append(urls, link)
		})
	}
	if opt.Supplementary {
		c.OnHTML("a.fs-download-button[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, link)
		})
	}
	c.OnResponse(func(r *colly.Response) {
		if !strings.HasSuffix(r.Request.URL.String(), "#supplementary-material") {
			c.Visit(r.Request.URL.String() + "#supplementary-material")
		}
	})
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

// PeerjSpider access Peerj files via spider
// supp not support now, need chromedp
func PeerjSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "peerj.com"),
		colly.MaxDepth(1),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.URL != nil {
		c.AllowedDomains = append(c.AllowedDomains, opt.URL.Host)
	}
	if opt.FullText {
		c.OnHTML("a[data-format=PDF]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = "https://peerj.com" + link
			urls = append(urls, link)
		})
	}
	if opt.Citations {
		c.OnHTML("a[data-format=BibText]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = "https://peerj.com" + link
			urls = append(urls, link)
		})
	}
	if opt.Supplementary {
		c.OnHTML("a.article-supporting-download[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, link)
		})
	}
	c.OnResponse(func(r *colly.Response) {
		if !strings.HasSuffix(r.Request.URL.String(), "#supplementary-material") {
			c.Visit(r.Request.URL.String() + "#supplementary-material")
		}
	})
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

// OupComSpider access academic.oup.com files via spider
func OupComSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "academic.oup.com", "oup.silverchair-cdn.com"),
		colly.MaxDepth(1),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.URL != nil {
		c.AllowedDomains = append(c.AllowedDomains, opt.URL.Host)
	}
	if opt.FullText {
		c.OnHTML("a.article-pdfLink", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = "https://academic.oup.com" + link
			urls = append(urls, link)
		})
	}
	if opt.Supplementary {
		c.OnHTML(".dataSuppLink a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, link)
		})
	}
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

// EmbopressSpider access https://www.embopress.org files via spider
func EmbopressSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "onlinelibrary.wiley.com", "www.embopress.org"),
		colly.MaxDepth(1),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.URL != nil {
		c.AllowedDomains = append(c.AllowedDomains, opt.URL.Host)
	}
	if opt.FullText {
		c.OnHTML("div.article-action a[aria-label=PDF]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = "https://www.embopress.org" + link
			urls = append(urls, link)
		})
	}
	if opt.Supplementary {
		c.OnHTML(".article-section__supporting a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = "https://www.embopress.org" + link
			urls = append(urls, link)
		})
	}
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

// AscopubsSpider access https://ascopubs.org/ files via spider
func AscopubsSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "ascopubs.org"),
		colly.MaxDepth(1),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.URL != nil {
		c.AllowedDomains = append(c.AllowedDomains, opt.URL.Host)
	}
	if opt.FullText {
		c.OnHTML(".pdfTools a[download]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = "https://ascopubs.org" + link
			urls = append(urls, link)
		})
	}
	if opt.Supplementary {
		c.OnHTML("article.article ul li a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = "https://ascopubs.org" + link
			urls = append(urls, link)
		})
	}
	c.OnResponse(func(r *colly.Response) {
		if r.Request.URL.String() == "https://ascopubs.org/doi/"+opt.Doi && opt.FullText {
			c.Visit("https://ascopubs.org/doi/pdf/" + opt.Doi)
		}
		if r.Request.URL.String() == "https://ascopubs.org/doi/"+opt.Doi && opt.Supplementary {
			c.Visit("https://ascopubs.org/doi/suppl/" + opt.Doi)
		}
	})
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

// HaematologicaSpider access https://ascopubs.org/ files via spider
func HaematologicaSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "www.haematologica.org"),
		colly.MaxDepth(1),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.URL != nil {
		c.AllowedDomains = append(c.AllowedDomains, opt.URL.Host)
	}
	if opt.FullText {
		c.OnHTML(".pdfTools a[download]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = "https://ascopubs.org" + link
			urls = append(urls, link)
		})
	}
	if opt.Supplementary {
		c.OnHTML("article.article ul li a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = "https://ascopubs.org" + link
			urls = append(urls, link)
		})
	}
	c.OnResponse(func(r *colly.Response) {
		if !strings.HasSuffix(r.Request.URL.String(), ".figures-only") && opt.FullText {
			urls = append(urls, r.Request.URL.String()+".pdf")
		} else if opt.Supplementary {
			c.Visit(r.Request.URL.String() + ".figures-only")
		}
	})
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

// WileyComSpider access https://onlinelibrary.wiley.com files via spider
func WileyComSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "onlinelibrary.wiley.com", "doi.wiley.com",
			"aasldpubs.onlinelibrary.wiley.com", "currentprotocols.onlinelibrary.wiley.com",
			"bpspubs.onlinelibrary.wiley.com", "stemcellsjournals.onlinelibrary.wiley.com",
			"agupubs.onlinelibrary.wiley.com"),
		colly.MaxDepth(1),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.URL != nil {
		c.AllowedDomains = append(c.AllowedDomains, opt.URL.Host)
	}
	if opt.FullText {
		c.OnHTML("meta[name=citation_pdf_url]", func(e *colly.HTMLElement) {
			link := e.Attr("content")
			link = stringo.StrReplaceAll(link, "/doi/pdf/", "/doi/pdfdirect/")
			urls = append(urls, linkFilter(link, opt.URL))
		})
		if strings.Contains(opt.Doi, "10.3982/") {
			urls = append(urls, "https://onlinelibrary.wiley.com/doi/pdfdirect/"+opt.Doi)
		}
	}
	if opt.Supplementary {
		c.OnHTML(".support-info__table td a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
		c.OnHTML("a[title='Download full book']", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

// ElifeSpider access https://elifesciences.org files via spider
func ElifeSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "elifesciences.org"),
		colly.MaxDepth(1),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.FullText {
		c.OnHTML("a[data-download-type='pdf-article']", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, link)
		})
	}
	if opt.Supplementary {
		c.OnHTML("a.additional-asset__link--download[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, link)
		})
	}
	c.OnResponse(func(r *colly.Response) {
		if strings.Contains(r.Request.URL.String(), "elifesciences.org") && !strings.HasSuffix(r.Request.URL.String(), "figures") && opt.Supplementary {
			c.Visit(r.Request.URL.String() + "/figures")
		}
	})
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

// JciSpider access https://www.jci.org files via spider
func JciSpider(opt *DoiSpiderOpt) (urls []string) {
	host := "https://www.jci.org"
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "www.jci.org"),
		colly.MaxDepth(1),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.FullText {
		c.OnHTML("h3 a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if strings.Contains(link, "cloudfront.net") {
				urls = append(urls, "http:"+link)
			} else {
				urls = append(urls, host+link)
			}
		})
	}
	if opt.Supplementary {
		c.OnHTML("#supplemental-material a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			c.Visit(host + link)
		})
	}
	c.OnResponse(func(r *colly.Response) {
		if strings.Contains(r.Request.URL.String(), "https://www.jci.org") && !strings.HasSuffix(r.Request.URL.String(), "pdf") && opt.FullText {
			c.Visit(r.Request.URL.String() + "/pdf")
		}
	})
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

// JstatsoftSpider access https://www.jstatsoft.org files via spider
func JstatsoftSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "www.jstatsoft.org"),
		colly.MaxDepth(1),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.FullText {
		c.OnHTML("a.file[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, link)
		})
	}
	if opt.Supplementary {
		c.OnHTML("a.action[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, link)
		})
	}
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

// JciSpider access www.ejcrim.com files via spider
func EjcrimSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "www.ejcrim.com"),
		colly.MaxDepth(1),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.FullText {
		c.OnHTML("a.pdf[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = strings.ReplaceAll(link, "/view/", "/download/")
			urls = append(urls, link)
		})
	}
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

// KosuyoluheartjournalSpider access http://www.kosuyoluheartjournal.com/ files
func KosuyoluheartjournalSpider(opt *DoiSpiderOpt) (urls []string) {
	url := "https://doi.org/" + opt.Doi
	log.Infof("Visiting %s", url)
	req, _ := http.Get(url)
	urls = append(urls, req.Request.URL.String())
	return urls
}

// DovepressSpider access http://www.dovepress.com files via spider
func DovepressSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "www.dovepress.com"),
		colly.MaxDepth(1),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.FullText {
		c.OnHTML("meta[name=citation_pdf_url]", func(e *colly.HTMLElement) {
			link := e.Attr("content")
			urls = append(urls, link)
		})
	}
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

// AutopsyandcasereportsSpider access https://autopsyandcasereports.org files via spider
func AutopsyandcasereportsSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "autopsyandcasereports.org"),
		colly.MaxDepth(1),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.FullText {
		c.OnHTML("a.pdfType1[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, "https://autopsyandcasereports.org"+link)
		})
	}
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

// FigshareSpider access https://figshare.com/ files via spider
func FigshareSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "figshare.com"),
		colly.MaxDepth(1),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.FullText {
		c.OnHTML("a.download-button[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, link)
		})
	}
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

// PubsacsSpider access https://pubs.acs.org/ files via spider
func PubsacsSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "pubs.acs.org"),
		colly.MaxDepth(1),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.FullText {
		link := fmt.Sprintf("/doi/pdf/%s", opt.Doi)
		urls = append(urls, linkFilter(link, opt.URL))
	}
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

// PubsRscSpider access https://pubs.rsc.org/ files via spider
func PubsRscSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "pubs.rsc.org", "xlink.rsc.org"),
		colly.MaxDepth(1),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	c.OnHTML("meta[name=citation_pdf_url]", func(e *colly.HTMLElement) {
		link := e.Attr("content")
		urls = append(urls, linkFilter(link, opt.URL))
	})
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

// PubsRscSpider access https://www.annualreviews.org/ files via spider
func AnnualReviewsSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "www.annualreviews.org"),
		colly.MaxDepth(1),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	c.OnHTML(".tool-buttons a.icon-pdf[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		urls = append(urls, linkFilter(link, opt.URL))
	})
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}
