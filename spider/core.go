package spider

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/openbiox/bget/chrome"
	glog "github.com/openbiox/ligo/log"
	cnet "github.com/openbiox/ligo/net"
	"github.com/openbiox/ligo/stringo"
)

var log = glog.Logger

// NatureComSpider access Nature.com files via spider
func NatureComSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "www.nature.com", "idp.nature.com"),
		colly.MaxDepth(1),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.URL != nil {
		c.AllowedDomains = append(c.AllowedDomains, opt.URL.Host)
	}
	if opt.FullText {
		c.OnHTML("a.c-pdf-download__link[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	if opt.Supplementary {
		c.OnHTML("a.print-link[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if !strings.Contains(link, "/figures/") {
				if !strings.HasPrefix(link, "http") {
					urls = append(urls, linkFilter(link, opt.URL))
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
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	extensions.Referer(c)
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	if opt.URL != nil {
		c.AllowedDomains = append(c.AllowedDomains, opt.URL.Host)
	}
	if opt.FullText {
		c.OnHTML("div.panels-ajax-tab-wrap-jnl_sci_tab_pdf a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
		c.OnHTML("meta[name=citation_pdf_url]", func(e *colly.HTMLElement) {
			link := e.Attr("content")
			urls = append(urls, linkFilter(link, opt.URL))
		})
		c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	}
	if opt.Supplementary {
		c.OnHTML("a.rewritten[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
		c.Visit(opt.URL.String() + "/tab-figures-data")
	}
	return urls
}

// CellComSpider access cell.com journal files via spider
func CellComSpider(opt *DoiSpiderOpt) []string {
	urls := []string{}
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "www.cell.com", "cell.com", "linkinghub.elsevier.com", "secure.jbs.elsevierhealth.com",
			"id.elsevier.com", "www.cancercell.org", "www.sciencedirect.com",
			"pdf.sciencedirectassets.com", "www.thelancet.com", "www.gastrojournal.org",
			"www.clinicalkey.com"),
		colly.MaxDepth(1),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
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
	if opt.Supplementary {
		urls = append(urls, chrome.DoiSupplURLs(fmt.Sprintf("https://doi.org/%s", opt.Doi),
			time.Duration(opt.Timeout)*time.Second, opt.Proxy)...)
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
	return urls
}

// BloodJournalSpider access http://www.bloodjournal.org files via spider
func BloodJournalSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "signin.hematology.org", "www.bloodjournal.org", "ashpublications.org"),
		colly.MaxDepth(1),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
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
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
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
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
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
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
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
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
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
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
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
		colly.AllowedDomains("doi.org", "ard.bmj.com", "adc.bmj.com", "casereports.bmj.com", "ebm.bmj.com", "gh.bmj.com", "informatics.bmj.com", "innovations.bmj.com", "bmjleader.bmj.com", "militaryhealth.bmj.com", "neurologyopen.bmj.com", "nutrition.bmj.com", "bmjopen.bmj.com", "drc.bmj.com", "bmjopengastro.bmj.com", "bmjophth.bmj.com", "qir.bmj.com", "bmjopenrespres.bmj.com", "openscience.bmj.com", "bmjopensem.bmj.com", "qualitysafety.bmj.com", "bmjpaedsopen.bmj.com", "srh.bmj.com", "stel.bmj.com", "spcare.bmj.com", "sit.bmj.com", "bjo.bmj.com", "bjsm.bmj.com", "considerations.bmj.com", "dtb.bmj.com", "ep.bmj.com", "emj.bmj.com", "esmoopen.bmj.com", "ejhp.bmj.com", "ebmh.bmj.com", "ebn.bmj.com", "fmch.bmj.com", "fn.bmj.com", "fg.bmj.com", "gpsych.bmj.com", "gut.bmj.com", "heart.bmj.com", "heartasia.bmj.com", "injuryprevention.bmj.com", "inpractice.bmj.com", "ihj.bmj.com", "ijgc.bmj.com", "jitc.bmj.com", "jcp.bmj.com", "jech.bmj.com", "jim.bmj.com", "jisakos.bmj.com", "jme.bmj.com", "jmg.bmj.com", "jnnp.bmj.com", "jnis.bmj.com", "lupus.bmj.com", "mh.bmj.com", "oem.bmj.com", "openheart.bmj.com", "pmj.bmj.com", "pn.bmj.com", "rapm.bmj.com", "rmdopen.bmj.com", "sti.bmj.com", "svn.bmj.com", "www.bmj.com", "thorax.bmj.com", "tobaccocontrol.bmj.com", "tsaco.bmj.com", "veterinaryrecord.bmj.com", "vetrecordcasereports.bmj.com", "vetrecordopen.bmj.com", "wjps.bmj.com"),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.URL != nil {
		c.AllowedDomains = append(c.AllowedDomains, opt.URL.Host)
	}
	fulltextUrl := ""
	if opt.URL.Hostname() != "www.bmj.com" {
		if opt.FullText {
			c.OnHTML("meta[name=citation_pdf_url]", func(e *colly.HTMLElement) {
				urls = append(urls, e.Attr("content"))
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
	}
	if opt.Supplementary {
		c.OnHTML(".supplementary-material a[href]", func(e *colly.HTMLElement) {
			urls = append(urls, e.Attr("href"))
		})
		c.OnHTML("a.rewritten[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = "https://" + opt.URL.Hostname() + link
			urls = append(urls, link)
		})
	}
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

// JournalsApsSpider access https://journals.aps.org/ files via spider
func JournalsApsSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "journals.aps.org", "link.aps.org"),
		colly.MaxDepth(1),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.FullText {
		c.OnHTML(".article-nav-actions a", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if strings.Contains(link, "/pdf/") {
				urls = append(urls, linkFilter(link, opt.URL))
			}
		})
	}
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

func CellimageLibrarySpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "www.cellimagelibrary.org"),
		colly.MaxDepth(2),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.FullText {
		c.OnHTML("a.download_menu_anchor", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if strings.Contains(link, ".zip") {
				urls = append(urls, linkFilter(link, opt.URL))
			}
		})
	}
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

func IeeexploreSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "ieeexplore.ieee.org"),
		colly.MaxDepth(2),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	if opt.FullText {
		c.OnHTML("a.download_menu_anchor", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if strings.Contains(link, ".zip") {
				urls = append(urls, linkFilter(link, opt.URL))
			}
		})
	}
	var done bool
	c.OnResponse(func(r *colly.Response) {
		if done {
			return
		}
		done = true
		link := fmt.Sprintf("/stamp/stamp.jsp?arnumber=%s", path.Base(r.Request.URL.String()))
		c.OnHTML("iframe", func(e *colly.HTMLElement) {
			link := e.Attr("src")
			urls = append(urls, linkFilter(link, opt.URL))
		})
		c.Visit(linkFilter(link, opt.URL))
	})
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

func SagepubComSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "journals.sagepub.com"),
		colly.MaxDepth(2),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	if opt.FullText {
		c.OnHTML(".pdf-access a", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if strings.Contains(link, "/pdf/") {
				urls = append(urls, linkFilter(link, opt.URL))
			}
		})
	}
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	if len(urls) == 0 {
		urls = append(urls, "http://sage.cnpereading.com/paragraph/download/"+opt.Doi)
	}
	return urls
}

func LwwComSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "journals.lww.com", "links.lww.com", "download.lww.com"),
		colly.MaxDepth(2),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	if opt.FullText {
		c.OnHTML("div.ejp-article-wrapper #js-ejp-article-tools", func(e *colly.HTMLElement) {
			link := e.Attr("data-pdf-url")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	if opt.Supplementary {
		c.OnHTML("#ej-article-sam-container a", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			client := cnet.NewHTTPClient(opt.Timeout, opt.Proxy)
			req, _ := http.NewRequest("HEAD", link, nil)
			resp, _ := client.Do(req)
			urls = append(urls, resp.Request.URL.String())
		})
	}
	c.Visit(fmt.Sprintf("https://journals.lww.com/%s", opt.Doi))
	return urls
}

func LiebertpubSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "www.liebertpub.com"),
		colly.MaxDepth(2),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	if opt.FullText {
		urls = append(urls, linkFilter("/doi/pdfplus/"+opt.Doi, opt.URL))
	}
	if opt.Supplementary {
		c.OnHTML("a.ext-link", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	c.Visit(fmt.Sprintf("https://www.liebertpub.com/doi/%s", opt.Doi))
	return urls
}

func PhysiologyOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	if opt.FullText {
		urls = append(urls, linkFilter("/doi/pdfplus/"+opt.Doi, opt.URL))
	}
	return urls
}

func RoyalsocietypublishingOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	if opt.FullText {
		urls = append(urls, linkFilter("/doi/pdf/"+opt.Doi, opt.URL))
	}
	return urls
}

func AmetsocOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	if opt.FullText {
		urls = append(urls, linkFilter("/doi/pdf/"+opt.Doi, opt.URL))
	}
	return urls
}

func AmegroupsSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "tlcr.amegroups.com"),
		colly.MaxDepth(2),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	if opt.FullText {
		c.OnHTML("li a.pdf", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = strings.ReplaceAll(link, "view", "download")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

func JmirOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "tlcr.amegroups.com"),
		colly.MaxDepth(2),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	if opt.FullText {
		link := opt.URL.String()
		urls = append(urls, link+"/pdf")
	}
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

func ThiemeConnectDeSpider(opt *DoiSpiderOpt) (urls []string) {
	if opt.FullText {
		link := opt.URL.String()
		link = strings.ReplaceAll(link, "/abstract/", "/pdf/")
		urls = append(urls, link+".pdf")
	}
	return urls
}

/*func DegruyterComSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "www.degruyter.com", "www.degruyter.com:443", "46.51.207.106",
			"108.128.99.123", "46.51.207.106:443"),
		colly.MaxDepth(5),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	if opt.FullText {
		c.OnHTML("meta[name=citation_pdf_url]", func(e *colly.HTMLElement) {
			link := e.Attr("content")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	err := c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	if err != nil {
		fmt.Println(err)
	}
	return urls
}*/

func ThnoOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "www.thno.org"),
		colly.MaxDepth(5),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	if opt.FullText {
		c.OnHTML("a.textbutton", func(e *colly.HTMLElement) {
			link := "/" + e.Attr("href")
			if strings.Contains(link, ".pdf") {
				urls = append(urls, linkFilter(link, opt.URL))
			}
		})
	}
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

func GeochemicalperspectivesOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "www.geochemicalperspectives.org"),
		colly.MaxDepth(5),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	if opt.FullText {
		c.OnHTML(".entry-content p a", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if strings.Contains(link, ".pdf") {
				urls = append(urls, linkFilter(link, opt.URL))
			}
		})
	}
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

func IospressComSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "www.medra.org", "content.iospress.com", "content.iospress.com:443"),
		colly.MaxDepth(5),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	if opt.FullText {
		c.OnHTML("meta[name='citation_pdf_url']", func(e *colly.HTMLElement) {
			link := e.Attr("content")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	c.Visit(fmt.Sprintf("https://content.iospress.com/doi/%s", opt.Doi))
	return urls
}

func IucrOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "journals.iucr.org", "scripts.iucr.org"),
		colly.MaxDepth(5),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	if opt.FullText {
		c.OnHTML(".bubbleInfo .sidebutton a[title=PDF]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	if opt.Supplementary {
		c.OnHTML(".file_links_other p a", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

func GeoscienceworldOrg(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "pubs.geoscienceworld.org"),
		colly.MaxDepth(5),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	if opt.FullText {
		c.OnHTML(".article-pdfLink", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, "https://pubs.geoscienceworld.org"+link)
		})
	}
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

func AeawebOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "pubs.aeaweb.org", "www.aeaweb.org"),
		colly.MaxDepth(5),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	if opt.FullText {
		c.OnHTML(".download a", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	if opt.Supplementary {
		c.OnHTML("#additionalMaterials li a", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	c.Visit(fmt.Sprintf("https://www.aeaweb.org/articles?id=%s", opt.Doi))
	return urls
}

func InformsOrgSPider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "pubsonline.informs.org"),
		colly.MaxDepth(5),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	urls = AddPdfSpider(opt)
	if opt.Supplementary {
		c.OnHTML(".article-section__content p a", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
		supplHost := stringo.StrReplaceAll(opt.URL.String(), "/doi/", "/doi/suppl/")
		c.Visit(supplHost)
	}
	return urls
}

func AsnjournalsOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "jasn.asnjournals.org", "www.jasn.org"),
		colly.MaxDepth(5),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	if opt.FullText {
		c.OnHTML("a[data-panel-name=jnl_asnjnls_tab_pdf]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	if opt.Supplementary {
		c.OnHTML("a.rewritten", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
		c.Visit(opt.URL.String() + "/tab-figures-data")
	}
	return urls
}

func CogitatiopressComSpider(opt *DoiSpiderOpt) (urls []string) {
	if opt.FullText {
		urls = append(urls, stringo.StrReplaceAll(opt.URL.String(),
			"/view/", "/download/")+"/"+path.Base(opt.URL.String()))
	}
	return urls
}

func AdiccionesEsSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "www.adicciones.es"),
		colly.MaxDepth(5),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	if opt.FullText {
		c.OnHTML("#articleFullText a:nth-child(3)", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = strings.ReplaceAll(link, "/view/", "/download/")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

func EurosurveillanceOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "www.eurosurveillance.org"),
		colly.MaxDepth(5),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	if opt.FullText {
		c.OnHTML(".pdfItem a.pdf", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if !strings.Contains(link, "suppdata") {
				link, _ = RetriveRedirectLink(linkFilter(link, opt.URL), opt.Timeout, opt.Proxy)
				urls = append(urls, link)
			}
		})
	}
	if opt.Supplementary {
		c.OnHTML(".pdfItem a.pdf", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if strings.Contains(link, "suppdata") {
				link, _ = RetriveRedirectLink(linkFilter(link, opt.URL), opt.Timeout, opt.Proxy)
				urls = append(urls, link)
			}
		})
	}
	c.Visit(fmt.Sprintf("https://www.eurosurveillance.org/content/%s#html_fulltext", opt.Doi))
	return urls
}

func AerzteblattDeSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "www.aerzteblatt.de"),
		colly.MaxDepth(5),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	if opt.FullText {
		c.OnHTML("a.pdfLink", func(e *colly.HTMLElement) {
			c.OnHTML("div.save a", func(e *colly.HTMLElement) {
				link := e.Attr("href")
				link = stringo.StrReplaceAll(link, "[?].*", "")
				urls = append(urls, linkFilter(link, opt.URL))
			})
			link := e.Attr("href")
			c.Visit(linkFilter(link, opt.URL))
		})
	}
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

func KjronlineOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	return KoreaMedSpider(opt, "kjronline.org")
}

func ImmunenetworkOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	return KoreaMedSpider(opt, "immunenetwork.org")
}

func TosOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "tos.org"),
		colly.MaxDepth(5),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	if opt.FullText {
		c.OnHTML(".large-links-blue a", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if strings.Contains(link, "docs") {
				urls = append(urls, linkFilter(link, opt.URL))
			}
		})
	}
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

func JstrokeOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "www.j-stroke.org", "j-stroke.org"),
		colly.MaxDepth(5),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	if opt.FullText {
		c.OnHTML("meta[name=citation_pdf_url]", func(e *colly.HTMLElement) {
			link := e.Attr("content")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

func AnnalsOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "annals.org"),
		colly.MaxDepth(5),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	if opt.FullText {
		c.OnHTML("#tagmasterPDF", func(e *colly.HTMLElement) {
			link := e.Attr("data-article-url")
			urls = append(urls, "https://annals.org"+link)
		})
	}
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

func PortlandpressComSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "portlandpress.com"),
		colly.MaxDepth(5),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	if opt.FullText {
		c.OnHTML("a.article-pdfLink", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, "https://portlandpress.com"+link)
		})
	}
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

func GeoscienceworldOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "pubs.geoscienceworld.org"),
		colly.MaxDepth(5),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	if opt.FullText {
		c.OnHTML("a.article-pdfLink", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, "https://pubs.geoscienceworld.org"+link)
		})
	}
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}
