package spider

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/openbiox/butils/log"
	bspider "github.com/openbiox/butils/spider"
	"github.com/openbiox/butils/stringo"
)

// NatureComSpider access Nature.com files via spider
func NatureComSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "www.nature.com", "idp.nature.com"),
		colly.MaxDepth(1),
	)
	bspider.SetSpiderProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.URL != nil {
		c.AllowedDomains = append(c.AllowedDomains, opt.URL.Host)
	}
	if opt.FullText {
		c.OnHTML("a.c-pdf-download__link[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, "https://nature.com"+link)
		})
	}
	if opt.Supplementary {
		c.OnHTML("a.print-link[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if !strings.HasPrefix(link, "http") {
				urls = append(urls, "https://nature.com"+link)
			} else {
				u, _ := url.Parse(link)
				linkTmp := strings.Split(u.Path, "/")
				if len(linkTmp) < 4 {
					return
				}
				linkTmp[2] = stringo.StrReplaceAll(linkTmp[2], "art:", "art%3A")
				newLink := append(linkTmp[0:2], strings.Join(linkTmp[2:4], "%2F"))
				newLink = append(newLink, linkTmp[4:len(linkTmp)]...)
				link = strings.Join(newLink, "/")
				link = u.Scheme + "://" + u.Host + link
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

// ScienseComSpider access sciencemag.org journal files via spider
func ScienseComSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "advances.sciencemag.org", "immunology.sciencemag.org",
			"robotics.sciencemag.org", "stke.sciencemag.org", "stm.sciencemag.org", "secure.jbs.elsevierhealth.com",
			"id.elsevier.com", "science.sciencemag.org", "www.sciencemag.org"),
		colly.MaxDepth(2),
	)
	bspider.SetSpiderProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	extensions.Referer(c)
	if opt.URL != nil {
		c.AllowedDomains = append(c.AllowedDomains, opt.URL.Host)
	}
	if opt.FullText {
		c.OnHTML("div.panels-ajax-tab-wrap-jnl_sci_tab_pdf a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, link)
		})
	}
	c.OnHTML("meta[name=citation_pdf_url]", func(e *colly.HTMLElement) {
		link := e.Attr("content")
		u, _ := url.Parse(link)
		link = u.Scheme + "://" + u.Host + u.Path
		if opt.FullText {
			urls = append(urls, link)
		}
		if opt.Supplementary {
			c.OnHTML("a.rewritten[href]", func(e *colly.HTMLElement) {
				link := e.Attr("href")
				u2, _ := url.Parse(link)
				link = u.Scheme + "://" + u.Host + u2.Path
				urls = append(urls, link)
			})
			c.Visit(strings.Replace(link, ".full.pdf", "", 1) + "/tab-figures-data")
		}
	})
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

// CellComSpider access cell.com journal files via spider
func CellComSpider(opt *DoiSpiderOpt) []string {
	urls := []string{}
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "www.cell.com", "cell.com", "linkinghub.elsevier.com", "secure.jbs.elsevierhealth.com",
			"id.elsevier.com", "www.cancercell.org", "www.sciencedirect.com",
			"pdf.sciencedirectassets.com", "www.thelancet.com", "www.gastrojournal.org"),
		colly.MaxDepth(1),
	)
	bspider.SetSpiderProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	extensions.Referer(c)
	if opt.URL != nil {
		c.AllowedDomains = append(c.AllowedDomains, opt.URL.Host)
	}
	if opt.FullText {
		c.OnHTML("meta[name=citation_pdf_url]", func(e *colly.HTMLElement) {
			link := e.Attr("content")
			urls = append(urls, link)
		})
		c.OnHTML("a.pdfLink[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if len(urls) == 0 {
				urls = append(urls, link)
			}
		})
		c.OnHTML("a.article-tools__item__displayStandardPdf[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if link != "#" && len(urls) == 0 {
				urls = append(urls, linkFilter(link, opt.URL))
			}
		})
		c.OnHTML("a.article-tools__item__displayExtendedPdf[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if link != "#" && len(urls) == 0 {
				urls = append(urls, linkFilter(link, opt.URL))
			}
		})
		c.OnHTML(".article-tools__pdf a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if link != "#" && len(urls) == 0 {
				urls = append(urls, linkFilter(link, opt.URL))
			}
		})
		c.OnHTML("div.PdfDownloadButton a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = "https://www.sciencedirect.com" + link
			if len(urls) == 0 {
				urls = append(urls, link)
			}
		})
	}
	c.OnHTML("#redirectURL", func(e *colly.HTMLElement) {
		link := e.Attr("value")
		u, _ := url.Parse(link)
		link, _ = url.QueryUnescape(u.Path)
		c.Visit(link)
	})
	if opt.Supplementary {
		c.OnHTML("#appsec1 a[target=new]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, link)
		})
		c.OnHTML("#appsec1 .externalFile a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, link)
		})
		c.OnHTML("span.article-attachment a.download-link[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, link)
		})
		c.OnHTML("a.supplemental-information__download[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if link != "#" {
				urls = append(urls, linkFilter(link, opt.URL))
			}
		})
	}
	c.OnHTML("meta[HTTP-EQUIV=REFRESH]", func(e *colly.HTMLElement) {
		link := e.Attr("content")
		link = stringo.StrReplaceAll(link, ".* url='", "")
		link = stringo.StrReplaceAll(link, "'$", "")
		link = "https://linkinghub.elsevier.com" + link
		c.Visit(link)
	})
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
		if stringo.StrDetect(r.URL.String(), "^https://www.sciencedirect.com") {
		}
	})
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

// BloodJournalSpider access http://www.bloodjournal.org files via spider
func BloodJournalSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "signin.hematology.org", "www.bloodjournal.org"),
		colly.MaxDepth(1),
	)
	bspider.SetSpiderProxy(c, opt.Proxy, opt.Timeout)
	if opt.URL != nil {
		c.AllowedDomains = append(c.AllowedDomains, opt.URL.Host)
	}
	if opt.FullText {
		c.OnHTML("meta[name=citation_pdf_url]", func(e *colly.HTMLElement) {
			link := e.Attr("content")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	if opt.Supplementary {
		c.OnHTML("a.[data-panel-name=jnl_bloodjournal_tab_data]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			c.Visit(link)
		})
		c.OnHTML("a.rewritten[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = linkFilter(link, opt.URL)
			urls = append(urls, link)
		})
	}

	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

// NejmSpider access http://www.nejm.org files via spider
func NejmSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "www.nejm.org"),
		colly.MaxDepth(1),
	)
	bspider.SetSpiderProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.URL != nil {
		c.AllowedDomains = append(c.AllowedDomains, opt.URL.Host)
	}
	if opt.FullText {
		c.OnHTML("a[data-tooltip='Download PDF']", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = "https://www.nejm.org" + link
			urls = append(urls, link)
		})
	}
	if opt.Supplementary {
		c.OnHTML("a[data-interactionType=multimedia_download]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if strings.Contains(link, "doi/suppl") {
				link = "https://www.nejm.org" + link
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

// AhajournalsSpider access https://www.ahajournals.org files via spider
func AhajournalsSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "www.ahajournals.org"),
		colly.MaxDepth(1),
	)
	bspider.SetSpiderProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.URL != nil {
		c.AllowedDomains = append(c.AllowedDomains, opt.URL.Host)
	}
	if opt.FullText {
		c.OnHTML(".citation__access__actions a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = "https://www.ahajournals.org" + link
			urls = append(urls, link)
			c.Visit(stringo.StrReplaceAll(link, "/doi/pdf/", "/doi/suppl/"))
		})
	}
	if opt.Supplementary {
		c.OnHTML(".supplemental-material__item a.green-text-color[href]", func(e *colly.HTMLElement) {
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

// JamaNetworkSpider access https://jamanetwork.com files via spider
func JamaNetworkSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "jamanetwork.com"),
		colly.MaxDepth(1),
	)
	bspider.SetSpiderProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.URL != nil {
		c.AllowedDomains = append(c.AllowedDomains, opt.URL.Host)
	}
	if opt.FullText {
		c.OnHTML("#contents-tab a.toolbar-pdf[data-article-url]", func(e *colly.HTMLElement) {
			link := e.Attr("data-article-url")
			urls = append(urls, "https://jamanetwork.com"+link)
		})
	}
	if opt.Supplementary {
		c.OnHTML(".supplement a.supplement-download[href]", func(e *colly.HTMLElement) {
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

// AacrJournalsSpider access aacrjournals.org files via spider
func AacrJournalsSpider(opt *DoiSpiderOpt) (urls []string) {
	host := ""
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "aacrjournals.org", "cancerdiscovery.aacrjournals.org",
			"clincancerres.aacrjournals.org", "cancerimmunolres.aacrjournals.org"),
		colly.MaxDepth(1),
	)
	bspider.SetSpiderProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.URL != nil {
		c.AllowedDomains = append(c.AllowedDomains, opt.URL.Host)
	}
	if opt.Supplementary {
		c.OnHTML("a.rewritten[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = "https://" + host + link
			urls = append(urls, link)
		})
	}
	c.OnResponse(func(r *colly.Response) {
		host = r.Request.URL.Hostname()
		cd1 := !stringo.StrDetect(r.Request.URL.String(), ".figures-only$")
		cd2 := !stringo.StrDetect(r.Request.URL.String(), ".full-text.pdf$")
		if strings.Contains(r.Request.URL.String(), "aacrjournals.org") && cd1 && cd2 && opt.Supplementary {
			c.Visit(r.Request.URL.String() + ".figures-only")
		}
		if strings.Contains(r.Request.URL.String(), "aacrjournals.org") && cd1 && cd2 && opt.FullText {
			urls = append(urls, r.Request.URL.String()+".full-text.pdf")
		}
	})
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

// TandfonlineSpider access https://www.tandfonline.com files via spider
// not support now, need chromedp
func TandfonlineSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "www.tandfonline.com"),
		colly.MaxDepth(1),
	)
	bspider.SetSpiderProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.FullText {
		c.OnHTML("a[title='Download all']", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
		c.OnHTML("li.pdf-tab", func(e *colly.HTMLElement) {
			link := fmt.Sprintf("https://www.tandfonline.com/doi/pdf/%s?needAccess=true", opt.Doi)
			urls = append(urls, link)
		})
	}
	if opt.Supplementary {
		c.OnHTML("a.show-pdf[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
		c.OnHTML("#supplementaryPanel a", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
		c.Visit(fmt.Sprintf("https://www.tandfonline.com/doi/suppl/%s", opt.Doi))
	}
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

// BmjComSpider access www.bmj.com files via spider
func BmjComSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "www.bmj.com", "gut.bmj.com"),
		colly.MaxDepth(1),
	)
	bspider.SetSpiderProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.URL != nil {
		c.AllowedDomains = append(c.AllowedDomains, opt.URL.Host)
	}
	fulltextUrl := ""
	if opt.URL.Hostname() == "gut.bmj.com" {
		if opt.FullText {
			c.OnHTML("meta[name=citation_pdf_url]", func(e *colly.HTMLElement) {
				urls = append(urls, e.Attr("content"))
			})
		}
		if opt.Supplementary {
			c.OnHTML(".supplementary-material a[href]", func(e *colly.HTMLElement) {
				urls = append(urls, e.Attr("href"))
			})
		}
	} else {
		c.OnHTML("a.pdf-link[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			fulltextUrl = "https://" + opt.URL.Hostname() + link
			if opt.FullText {
				urls = append(urls, fulltextUrl)
			}
			if opt.Supplementary {
				c.Visit(stringo.StrReplaceAll(fulltextUrl, ".full.pdf", "/related"))
			}
		})
		if opt.Supplementary {
			c.OnHTML("a.rewritten[href]", func(e *colly.HTMLElement) {
				link := e.Attr("href")
				link = "https://" + opt.URL.Hostname() + link
				urls = append(urls, link)
			})
		}
	}
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

// AtsjournalsOrgSpider access www.atsjournals.org files via spider
func AtsjournalsOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "www.atsjournals.org"),
		colly.MaxDepth(2),
	)
	bspider.SetSpiderProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.URL != nil {
		c.AllowedDomains = append(c.AllowedDomains, opt.URL.Host)
	}
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	log.Infof("Visiting https://www.atsjournals.org/doi/full/%s", opt.Doi)
	if opt.FullText {
		urls = append(urls, "https://www.atsjournals.org/doi/pdf/"+opt.Doi)
	}
	if opt.Supplementary {
		c.OnHTML(".suppl_list a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = "https://" + opt.URL.Hostname() + link
			urls = append(urls, link)
		})
		c.Visit("https://www.atsjournals.org/doi/suppl/" + opt.Doi)
	}
	return urls
}

// JournalsApsSpider access https://journals.aps.org/ files via spider
func JournalsApsSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "journals.aps.org", "link.aps.org"),
		colly.MaxDepth(1),
	)
	bspider.SetSpiderProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	c.OnHTML(".article-nav-actions a", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(link, "/pdf/") {
			urls = append(urls, linkFilter(link, opt.URL))
		}
	})
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}
