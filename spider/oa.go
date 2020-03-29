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
	c := initDoiColley(opt, "")
	c.AllowedDomains = append(c.AllowedDomains, []string{"zenodo.org"}...)
	if opt.FullText {
		c.OnHTML("tbody a.filename[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if strings.Contains(link, "?download=1") {
				urls = append(urls, linkFilter(link, opt.URL))
			}
		})
	}
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

// CshlpSpider access CshlpSpider files via spider
func CshlpSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "")
	c.AllowedDomains = append(c.AllowedDomains, []string{"biorxiv.org", "www.biorxiv.org",
		"genome.cshlp.org", "genesdev.cshlp.org", "www.medrxiv.org"}...)
	dataAvail := []string{}
	if opt.Supplementary {
		var supplVisited bool
		c.OnHTML("#supp-adjunct-data a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
		c.OnHTML("#mini-panel-biorxiv_art_tools .pane-highwire-variant-link a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if !stringo.StrDetect(link, "supplementary-material$|external-links$") {
				u, _ := url.Parse(link)
				link = linkFilter(u.Host+u.Path, opt.URL)
				urls = append(urls, link)
			}
		})
		c.OnHTML(".pane-biorxiv-supplementary-fragment a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if !stringo.StrDetect(link, "highwire/filestream") && !supplVisited {
				supplVisited = true
				Visit(c, linkFilter(link, opt.URL))
			}
		})
		c.OnHTML(".supplementary-material-expansion a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = stringo.StrReplaceAll(link, "[?]download=true$", "")
			urls = append(urls, linkFilter(link, opt.URL))
		})
		c.OnHTML("div.auto-clean a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
		c.OnHTML("a[rel=supplemental-data]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			Visit(c, linkFilter(link, opt.URL))
		})
		c.OnHTML(".data-availability p", func(e *colly.HTMLElement) {
			text := e.Text
			dataAvail = append(dataAvail, text)
		})
	}
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	if len(dataAvail) > 0 {
		dataAvailText := fmt.Sprintf(`Finding data availability from %s: ["`, opt.Doi)
		dataAvailText = dataAvailText + stringo.StrReplaceAll(strings.Join(dataAvail, `", "`), `", "$`, "")
		dataAvailText = dataAvailText + `"]`
		log.Infoln(dataAvailText)
	}
	postSpiderPrint(opt, &urls)
	return urls
}

// BiomedcentralSpider access GenomeBiology files via spider
func BiomedcentralSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "")
	c.AllowedDomains = append(c.AllowedDomains, BiomedcentralJournalLinks...)
	if opt.FullText {
		c.OnHTML(".c-pdf-download a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
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
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

// PnasSpider access PnasSpider files via spider
func PnasSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "https://www.pnas.org")
	c.AllowedDomains = append(c.AllowedDomains, []string{"www.pnas.org"}...)
	if opt.FullText {
		c.OnHTML("a[data-trigger=tab-pdf]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	if opt.Supplementary {
		c.OnHTML("a['data-trigger'='tab-figures-data']", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			Visit(c, link)
		})
		c.OnHTML("a.rewritten[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
		c.OnResponse(func(r *colly.Response) {
			if !strings.HasSuffix(r.Request.URL.String(), "/tab-figures-data") {
				Visit(c, r.Request.URL.String()+"/tab-figures-data")
			}
		})
	}
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

// PlosSpider access PlosSpider files via spider
func PlosSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "")
	c.AllowedDomains = append(c.AllowedDomains, []string{"journals.plos.org", "dx.plos.org"}...)
	if opt.FullText {
		c.OnHTML("#downloadPdf", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	if opt.Supplementary {
		c.OnHTML(".supplementary-material a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if strings.Contains(link, "doi.org") {
				urls = append(urls, linkFilter(link, opt.URL))
			}
		})
	}
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

// FrontiersinSpider access Frontiers files via spider
func FrontiersinSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "https://www.frontiersin.org")
	c.AllowedDomains = append(c.AllowedDomains, []string{"www.frontiersin.org", "journal.frontiersin.org"}...)
	if opt.FullText {
		c.OnHTML("a.download-files-pdf", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	if opt.Supplementary {
		c.OnHTML("a.fs-download-button[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	c.OnResponse(func(r *colly.Response) {
		if !strings.HasSuffix(r.Request.URL.String(), "#supplementary-material") {
			Visit(c, r.Request.URL.String()+"#supplementary-material")
		}
	})
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

// PeerjSpider access Peerj files via spider
// supp not support now, need chromedp
func PeerjSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "https://peerj.com")
	c.AllowedDomains = append(c.AllowedDomains, []string{"peer.com"}...)
	if opt.FullText {
		c.OnHTML("a[data-format=PDF]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	if opt.Citations {
		c.OnHTML("a[data-format=BibText]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	if opt.Supplementary {
		c.OnHTML("a.article-supporting-download[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	c.OnResponse(func(r *colly.Response) {
		if !strings.HasSuffix(r.Request.URL.String(), "#supplementary-material") {
			Visit(c, r.Request.URL.String()+"#supplementary-material")
		}
	})
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

// OupComSpider access academic.oup.com files via spider
func OupComSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "https://academic.oup.com")
	c.AllowedDomains = append(c.AllowedDomains, []string{"academic.oup.com", "oup.silverchair-cdn.com"}...)
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
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

// EmbopressSpider access https://www.embopress.org files via spider
func EmbopressSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "")
	c.AllowedDomains = append(c.AllowedDomains, []string{"onlinelibrary.wiley.com", "www.embopress.org"}...)
	if opt.FullText {
		c.OnHTML("meta[name=citation_pdf_url]", func(e *colly.HTMLElement) {
			link := e.Attr("content")
			link = stringo.StrReplaceAll(link, "/doi/pdf/", "/doi/pdfdirect/")
			urls = append(urls, linkFilter(link, opt.URL))
		})
		c.OnHTML("div.article-action a[aria-label=PDF]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	if opt.Supplementary {
		c.OnHTML(".article-section__supporting a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

// AscopubsSpider access https://ascopubs.org/ files via spider
func AscopubsSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "http://ascopubs.org")
	c.AllowedDomains = append(c.AllowedDomains, []string{"ascopubs.org"}...)
	if opt.FullText {
		c.OnHTML(".pdfTools a[download]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	if opt.Supplementary {
		c.OnHTML("article.article ul li a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	c.OnResponse(func(r *colly.Response) {
		if r.Request.URL.String() == "https://ascopubs.org/doi/"+opt.Doi && opt.FullText {
			Visit(c, "https://ascopubs.org/doi/pdf/"+opt.Doi)
		}
		if r.Request.URL.String() == "https://ascopubs.org/doi/"+opt.Doi && opt.Supplementary {
			Visit(c, "https://ascopubs.org/doi/suppl/"+opt.Doi)
		}
	})
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

// HaematologicaSpider access https://ascopubs.org/ files via spider
func HaematologicaSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "http://www.haematologica.org")
	c.AllowedDomains = append(c.AllowedDomains, []string{"www.haematologica.org"}...)
	if opt.FullText {
		c.OnHTML(".pdfTools a[download]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	if opt.Supplementary {
		c.OnHTML("article.article ul li a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	c.OnResponse(func(r *colly.Response) {
		if !strings.HasSuffix(r.Request.URL.String(), ".figures-only") && opt.FullText {
			urls = append(urls, r.Request.URL.String()+".pdf")
		} else if opt.Supplementary {
			Visit(c, r.Request.URL.String()+".figures-only")
		}
	})
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

// WileyComSpider access https://onlinelibrary.wiley.com files via spider
func WileyComSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "")
	c.AllowedDomains = append(c.AllowedDomains, []string{"onlinelibrary.wiley.com", "doi.wiley.com",
		"aasldpubs.onlinelibrary.wiley.com", "currentprotocols.onlinelibrary.wiley.com",
		"bpspubs.onlinelibrary.wiley.com", "stemcellsjournals.onlinelibrary.wiley.com",
		"agupubs.onlinelibrary.wiley.com", "www.cochranelibrary.com"}...)
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
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

// ElifeSpider access https://elifesciences.org files via spider
func ElifeSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "https://elifesciences.org")
	c.AllowedDomains = append(c.AllowedDomains, []string{"elifesciences.org"}...)
	if opt.FullText {
		c.OnHTML("a[data-download-type='pdf-article']", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	if opt.Supplementary {
		c.OnHTML("a.additional-asset__link--download[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	c.OnResponse(func(r *colly.Response) {
		if strings.Contains(r.Request.URL.String(), "elifesciences.org") && !strings.HasSuffix(r.Request.URL.String(), "figures") && opt.Supplementary {
			Visit(c, r.Request.URL.String()+"/figures")
		}
	})
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

// JciSpider access https://www.jci.org files via spider
func JciSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "https://www.jci.org")
	c.AllowedDomains = append(c.AllowedDomains, []string{"www.jci.org"}...)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.FullText {
		c.OnHTML("h3 a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if strings.Contains(link, "cloudfront.net") {
				urls = append(urls, "http:"+link)
			} else {
				urls = append(urls, linkFilter(link, opt.URL))
			}
		})
	}
	if opt.Supplementary {
		c.OnHTML("#supplemental-material a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			Visit(c, linkFilter(link, opt.URL))
		})
	}
	c.OnResponse(func(r *colly.Response) {
		if strings.Contains(r.Request.URL.String(), "https://www.jci.org") && !strings.HasSuffix(r.Request.URL.String(), "pdf") && opt.FullText {
			Visit(c, r.Request.URL.String()+"/pdf")
		}
	})
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

// JstatsoftSpider access https://www.jstatsoft.org files via spider
func JstatsoftSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "https://www.jstatsoft.org")
	c.AllowedDomains = append(c.AllowedDomains, []string{"www.jstatsoft.org"}...)
	if opt.FullText {
		c.OnHTML("a.file[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	if opt.Supplementary {
		c.OnHTML("a.action[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

// JciSpider access www.ejcrim.com files via spider
func EjcrimSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "https://www.ejcrim.com")
	c.AllowedDomains = append(c.AllowedDomains, []string{"www.ejcrim.com"}...)
	if opt.FullText {
		c.OnHTML("a.pdf[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = strings.ReplaceAll(link, "/view/", "/download/")
			urls = append(urls, link)
		})
	}
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

// KosuyoluheartjournalSpider access http://www.kosuyoluheartjournal.com/ files
func KosuyoluheartjournalSpider(opt *DoiSpiderOpt) (urls []string) {
	url := "https://doi.org/" + opt.Doi
	log.Infof("Visiting %s", url)
	req, _ := http.Get(url)
	urls = append(urls, req.Request.URL.String())
	postSpiderPrint(opt, &urls)
	return urls
}

// DovepressSpider access http://www.dovepress.com files via spider
func DovepressSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "https://www.dovepress.com")
	c.AllowedDomains = append(c.AllowedDomains, []string{"www.dovepress.com"}...)
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

// AutopsyandcasereportsSpider access https://autopsyandcasereports.org files via spider
func AutopsyandcasereportsSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "http://autopsyandcasereports.org")
	c.AllowedDomains = append(c.AllowedDomains, []string{"autopsyandcasereports.org"}...)
	if opt.FullText {
		c.OnHTML("a.pdfType1[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	Visit(c, fmt.Sprintf("https://autopsyandcasereports.org/article/doi/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

// FigshareSpider access https://figshare.com/ files via spider
func FigshareSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "https://figshare.com")
	c.AllowedDomains = append(c.AllowedDomains, []string{"figshare.com"}...)
	if opt.FullText {
		c.OnHTML("a.download-button[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

// PubsacsSpider access https://pubs.acs.org/ files via spider
func PubsacsSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "https://pubs.acs.org")
	c.AllowedDomains = append(c.AllowedDomains, []string{"pubs.acs.org"}...)
	if opt.FullText {
		link := fmt.Sprintf("/doi/pdf/%s", opt.Doi)
		urls = append(urls, linkFilter(link, opt.URL))
	}
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

// PubsRscSpider access https://pubs.rsc.org/ files via spider
func PubsRscSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "https://pubs.rsc.org")
	c.AllowedDomains = append(c.AllowedDomains, []string{"pubs.rsc.org", "xlink.rsc.org"}...)
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

// PubsRscSpider access https://www.annualreviews.org/ files via spider
func AnnualReviewsSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "https://www.annualreviews.org")
	c.AllowedDomains = append(c.AllowedDomains, []string{"www.annualreviews.org"}...)
	c.OnHTML(".tool-buttons a.icon-pdf[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		urls = append(urls, linkFilter(link, opt.URL))
	})
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

func MedknowSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "")
	c.AllowedDomains = append(c.AllowedDomains, MedknowJournalLinks...)
	if opt.FullText {
		c.OnHTML("td p a", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if strings.Contains(link, ".pdf") {
				urls = append(urls, linkFilter(link, opt.URL))
			}
		})
		var done bool
		var done2 bool
		c.OnHTML("meta[name=citation_pdf_url]", func(e *colly.HTMLElement) {
			link := e.Attr("content")
			if strings.Contains(link, "article.asp") {
				if !done {
					Visit(c, linkFilter(link, opt.URL))
				}
				done = true
			}
		})
		c.OnHTML("td a:nth-child(2)", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if strings.Contains(link, "type=2") {
				if !done2 {
					Visit(c, linkFilter(link, opt.URL))
				}
				done2 = true
			}
		})
	}
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

func EajmOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "https://www.eajm.org")
	c.AllowedDomains = append(c.AllowedDomains, []string{"www.eajm.org"}...)
	c.OnHTML("a.pdf-link", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(link, "/en/") {
			c.OnHTML("#mainFrame", func(e *colly.HTMLElement) {
				link := e.Attr("src")
				urls = append(urls, linkFilter(link, opt.URL))
			})
			Visit(c, linkFilter(link, opt.URL))
		}
	})
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

func AosisCoZaSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "")
	c.AllowedDomains = append(c.AllowedDomains, AosisCoZaJournalLinks...)
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

func KoreaMedSpider(opt *DoiSpiderOpt) (urls []string) {
	doi2host := make(map[string]string)
	for k := range KoreaMedJournalDois {
		doi2host[KoreaMedJournalDois[k]] = KoreaMedJournalLinks[k]
	}
	k := stringo.StrSplit(opt.Doi, "/", 2)[0]
	hostname := doi2host[k]
	c := initDoiColley(opt, "")
	c.AllowedDomains = append(c.AllowedDomains, KoreaMedJournalLinks...)
	if opt.FullText {
		c.OnHTML(".portlet-article-body-cell a", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if strings.Contains(link, "PDFData") {
				urls = append(urls, "https://"+hostname+link)
			}
		})
	}
	if opt.Supplementary {
		c.OnHTML(".portlet-article-body-cell-supplementary a", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			c.OnHTML(".supplementary-material-item a", func(e *colly.HTMLElement) {
				urls = append(urls, "https://"+hostname+"/"+e.Attr("href"))
			})
			Visit(c, fmt.Sprintf("%s/%s", hostname, link))
		})
	}
	Visit(c, fmt.Sprintf("https://"+hostname+"/DOIx.php?id=%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

func Bbk19Spider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "https://19.bbk.ac.uk")
	c.AllowedDomains = append(c.AllowedDomains, []string{"19.bbk.ac.uk"}...)
	if opt.FullText {
		c.OnHTML(".section ul li:nth-child(2) a:first-child", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if strings.Contains(link, "/download") {
				urls = append(urls, linkFilter(link, opt.URL))
			}
		})
	}
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

func TheietOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "https://digital-library.theiet.org")
	c.AllowedDomains = append(c.AllowedDomains, []string{"digital-library.theiet.org"}...)
	if opt.FullText {
		c.OnHTML(".headlinebox ul.fulltext li.pdf a:first-child", func(e *colly.HTMLElement) {
			link := "https://digital-library.theiet.org" + e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	Visit(c, fmt.Sprintf("https://digital-library.theiet.org/content/journals/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

func KargerComSpider(opt *DoiSpiderOpt) (urls []string) {
	if opt.FullText {
		urls = append(urls, strings.ReplaceAll(opt.URL.String(), "/FullText/", "/Pdf/"))
	}
	postSpiderPrint(opt, &urls)
	return urls
}

func AappublicationsOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "https://hosppeds.aappublications.org")
	c.AllowedDomains = append(c.AllowedDomains, []string{"hosppeds.aappublications.org", "www.nfaap.org"}...)
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}
