package spider

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/openanno/bget/chrome"
	glog "github.com/openbiox/ligo/log"
	cnet "github.com/openbiox/ligo/net"
	"github.com/openbiox/ligo/stringo"
)

var log = glog.Logger

// NatureComSpider access Nature.com files via spider
func NatureComSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "")
	c.AllowedDomains = append(c.AllowedDomains, []string{"www.nature.com", "idp.nature.com"}...)
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
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

// ScienseComSpider access sciencemag.org journal files via spider
func ScienseComSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "")
	c.AllowedDomains = append(c.AllowedDomains, ScienseComJournalLinks...)
	if opt.FullText {
		c.OnHTML("div.panels-ajax-tab-wrap-jnl_sci_tab_pdf a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	if opt.Supplementary {
		c.OnHTML("a.rewritten[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
		Visit(c, opt.URL.String()+"/tab-figures-data")
	}
	postSpiderPrint(opt, &urls)
	return urls
}

// CellComSpider access cell.com journal files via spider
func CellComSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "")
	c.AllowedDomains = append(c.AllowedDomains, CellComJournalLinks...)
	if opt.FullText {
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
		Visit(c, link)
	})
	c.OnHTML("meta[HTTP-EQUIV=REFRESH]", func(e *colly.HTMLElement) {
		link := e.Attr("content")
		link = stringo.StrReplaceAll(link, ".* url='", "")
		link = stringo.StrReplaceAll(link, "'$", "")
		link = "https://linkinghub.elsevier.com" + link
		Visit(c, link)
	})
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
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
	postSpiderPrint(opt, &urls)
	return urls
}

// BloodJournalSpider access http://www.bloodjournal.org files via spider
func BloodJournalSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "")
	c.AllowedDomains = append(c.AllowedDomains, []string{"signin.hematology.org", "www.bloodjournal.org", "ashpublications.org"}...)
	if opt.Supplementary {
		c.OnHTML("a.[data-panel-name=jnl_bloodjournal_tab_data]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			Visit(c, link)
		})
		c.OnHTML("a.rewritten[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = linkFilter(link, opt.URL)
			urls = append(urls, link)
		})
	}
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

// NejmSpider access http://www.nejm.org files via spider
func NejmSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "https://www.nejm.org")
	c.AllowedDomains = append(c.AllowedDomains, []string{"www.nejm.org"}...)
	if opt.FullText {
		c.OnHTML("a[data-tooltip='Download PDF']", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	if opt.Supplementary {
		c.OnHTML("a[data-interactionType=multimedia_download]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if strings.Contains(link, "doi/suppl") {
				urls = append(urls, linkFilter(link, opt.URL))
			}
		})
	}
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

// AhajournalsSpider access https://www.ahajournals.org files via spider
func AhajournalsSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "https://www.ahajournals.org")
	c.AllowedDomains = append(c.AllowedDomains, []string{"www.ahajournals.org"}...)
	if opt.FullText {
		urls = append(urls, fmt.Sprintf("https://www.ahajournals.org/doi/pdf/%s?download=true", opt.Doi))
	}
	if opt.Supplementary {
		c.OnHTML("li.supplemental-material__item a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, link)
		})
		Visit(c, fmt.Sprintf("https://www.ahajournals.org/doi/suppl/%s", opt.Doi))
	}
	postSpiderPrint(opt, &urls)
	return urls
}

// JamaNetworkSpider access https://jamanetwork.com files via spider
func JamaNetworkSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "https://jamanetwork.com")
	c.AllowedDomains = append(c.AllowedDomains, []string{"jamanetwork.com"}...)
	if opt.FullText {
		c.OnHTML("#contents-tab a.toolbar-pdf[data-article-url]", func(e *colly.HTMLElement) {
			link := e.Attr("data-article-url")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	if opt.Supplementary {
		c.OnHTML(".supplement a.supplement-download[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, link)
		})
	}
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

// AacrJournalsSpider access aacrjournals.org files via spider
func AacrJournalsSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "")
	c.AllowedDomains = append(c.AllowedDomains, AacrJournalLinks...)
	if opt.Supplementary {
		c.OnHTML("a.rewritten[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	c.OnResponse(func(r *colly.Response) {
		cd1 := !stringo.StrDetect(r.Request.URL.String(), ".figures-only$")
		cd2 := !stringo.StrDetect(r.Request.URL.String(), ".full-text.pdf$")
		if strings.Contains(r.Request.URL.String(), "aacrjournals.org") && cd1 && cd2 && opt.Supplementary {
			Visit(c, r.Request.URL.String()+".figures-only")
		}
		if strings.Contains(r.Request.URL.String(), "aacrjournals.org") && cd1 && cd2 && opt.FullText {
			urls = append(urls, r.Request.URL.String()+".full-text.pdf")
		}
	})
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

// TandfonlineSpider access https://www.tandfonline.com files via spider
// not support now, need chromedp
func TandfonlineSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "www.tandfonline.com")
	c.AllowedDomains = append(c.AllowedDomains, []string{"www.tandfonline.com"}...)
	if opt.FullText {
		c.OnHTML("a[title='Download all']", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
		link := fmt.Sprintf("https://www.tandfonline.com/doi/pdf/%s?needAccess=true", opt.Doi)
		urls = append(urls, link)
	}
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	if opt.Supplementary {
		c.OnHTML("a.show-pdf[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
		c.OnHTML("#supplementaryPanel a", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
		Visit(c, fmt.Sprintf("https://www.tandfonline.com/doi/suppl/%s", opt.Doi))
	}
	postSpiderPrint(opt, &urls)
	return urls
}

// BmjComSpider access www.bmj.com files via spider
func BmjComSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "")
	c.AllowedDomains = append(c.AllowedDomains, BmjComJournalLinks...)
	fulltextUrl := ""
	if opt.URL.Hostname() == "www.bmj.com" {
		c.OnHTML("a.pdf-link[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			fulltextUrl = "https://" + opt.URL.Hostname() + link
			if opt.FullText {
				urls = append(urls, fulltextUrl)
			}
			if opt.Supplementary {
				Visit(c, stringo.StrReplaceAll(fulltextUrl, ".full.pdf", "/related"))
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
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

// JournalsApsSpider access https://journals.aps.org/ files via spider
func JournalsApsSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "http://journals.aps.org")
	c.AllowedDomains = append(c.AllowedDomains, []string{"journals.aps.org", "link.aps.org"}...)
	if opt.FullText {
		c.OnHTML(".article-nav-actions a", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if strings.Contains(link, "/pdf/") {
				urls = append(urls, linkFilter(link, opt.URL))
			}
		})
	}
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

func CellimageLibrarySpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "http://www.cellimagelibrary.org")
	c.AllowedDomains = append(c.AllowedDomains, []string{"www.cellimagelibrary.org", "cellimagelibrary.org"}...)
	if opt.FullText {
		c.OnHTML("a.download_menu_anchor", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if strings.Contains(link, ".zip") {
				urls = append(urls, linkFilter(link, opt.URL))
			}
		})
	}
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

func IeeexploreSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "https://ieeexplore.ieee.org")
	c.AllowedDomains = append(c.AllowedDomains, []string{"ieeexplore.ieee.org"}...)
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
		Visit(c, linkFilter(link, opt.URL))
	})
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

func SagepubComSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "")
	c.AllowedDomains = append(c.AllowedDomains, []string{"journals.sagepub.com"}...)
	if opt.FullText {
		c.OnHTML(".pdf-access a", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if strings.Contains(link, "/pdf/") {
				urls = append(urls, linkFilter(link, opt.URL))
			}
		})
	}
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	if len(urls) == 0 {
		urls = append(urls, "http://sage.cnpereading.com/paragraph/download/"+opt.Doi)
	}
	postSpiderPrint(opt, &urls)
	return urls
}

func LwwComSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "")
	c.AllowedDomains = append(c.AllowedDomains, []string{"journals.lww.com",
		"links.lww.com", "download.lww.com"}...)
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
	Visit(c, fmt.Sprintf("https://journals.lww.com/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

func LiebertpubSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "https://www.liebertpub.com")
	c.AllowedDomains = append(c.AllowedDomains, []string{"www.liebertpub.com"}...)
	if opt.FullText {
		urls = append(urls, linkFilter("/doi/pdfplus/"+opt.Doi, opt.URL))
	}
	if opt.Supplementary {
		c.OnHTML("a.ext-link", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	Visit(c, fmt.Sprintf("https://www.liebertpub.com/doi/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

func PhysiologyOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	if opt.FullText {
		urls = append(urls, linkFilter("/doi/pdfplus/"+opt.Doi, opt.URL))
	}
	postSpiderPrint(opt, &urls)
	return urls
}

func RoyalsocietypublishingOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	if opt.FullText {
		urls = append(urls, linkFilter("/doi/pdf/"+opt.Doi, opt.URL))
	}
	postSpiderPrint(opt, &urls)
	return urls
}

func AmetsocOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	if opt.FullText {
		urls = append(urls, linkFilter("/doi/pdf/"+opt.Doi, opt.URL))
	}
	postSpiderPrint(opt, &urls)
	return urls
}

func AmegroupsSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "http://tlcr.amegroups.com")
	c.AllowedDomains = append(c.AllowedDomains, []string{"tlcr.amegroups.com"}...)
	if opt.FullText {
		c.OnHTML("li a.pdf", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = strings.ReplaceAll(link, "view", "download")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

func JmirOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "")
	c.AllowedDomains = append(c.AllowedDomains, []string{"jmir.org", "mhealth.jmir.org"}...)
	if opt.FullText {
		link := opt.URL.String()
		urls = append(urls, link+"/pdf")
	}
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

func ThiemeConnectDeSpider(opt *DoiSpiderOpt) (urls []string) {
	if opt.FullText {
		link := opt.URL.String()
		link = strings.ReplaceAll(link, "/abstract/", "/pdf/")
		urls = append(urls, link+".pdf")
	}
	postSpiderPrint(opt, &urls)
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
	err := Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	if err != nil {
		fmt.Println(err)
	}
	postSpiderPrint(opt, &urls)
return urls
}*/

func ThnoOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "")
	c.AllowedDomains = append(c.AllowedDomains, []string{"www.thno.org", "www.jcancer.org", "www.ntno.org",
		"www.ijbs.com", "www.medsci.org", "www.jgenomics.com", "www.jbji.net"}...)
	if opt.FullText {
		c.OnHTML("a.textbutton", func(e *colly.HTMLElement) {
			link := "/" + e.Attr("href")
			if strings.Contains(link, ".pdf") {
				urls = append(urls, linkFilter(link, opt.URL))
			}
		})
	}
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

func GeochemicalperspectivesOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "http://www.geochemicalperspectives.org")
	c.AllowedDomains = append(c.AllowedDomains, []string{"www.geochemicalperspectives.org"}...)
	if opt.FullText {
		c.OnHTML(".entry-content p a", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if strings.Contains(link, ".pdf") {
				urls = append(urls, linkFilter(link, opt.URL))
			}
		})
		c.OnHTML("#GPLpdf tr td a:first-child", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if strings.Contains(link, ".pdf") {
				urls = append(urls, linkFilter(link, opt.URL))
			}
		})
	}
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

func IospressComSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "https://content.iospress.com")
	c.AllowedDomains = append(c.AllowedDomains, []string{"www.medra.org", "content.iospress.com", "content.iospress.com:443"}...)
	Visit(c, fmt.Sprintf("https://content.iospress.com/doi/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

func IucrOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "http://journals.iucr.org")
	c.AllowedDomains = append(c.AllowedDomains, []string{"journals.iucr.org", "scripts.iucr.org"}...)
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
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

func GeoscienceworldOrg(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "https://pubs.geoscienceworld.org")
	c.AllowedDomains = append(c.AllowedDomains, []string{"pubs.geoscienceworld.org"}...)
	if opt.FullText {
		c.OnHTML(".article-pdfLink", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, "https://pubs.geoscienceworld.org"+link)
		})
	}
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

func AeawebOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "https://www.aeaweb.org")
	c.AllowedDomains = append(c.AllowedDomains, []string{"pubs.aeaweb.org", "www.aeaweb.org"}...)
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
	Visit(c, fmt.Sprintf("https://www.aeaweb.org/articles?id=%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

func InformsOrgSPider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "https://pubsonline.informs.org")
	c.AllowedDomains = append(c.AllowedDomains, []string{"pubsonline.informs.org"}...)
	urls = AddPdfSpider(opt)
	if opt.Supplementary {
		c.OnHTML(".article-section__content p a", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
		supplHost := stringo.StrReplaceAll(opt.URL.String(), "/doi/", "/doi/suppl/")
		Visit(c, supplHost)
	}
	postSpiderPrint(opt, &urls)
	return urls
}

func AsnjournalsOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "http://jasn.asnjournals.org")
	c.AllowedDomains = append(c.AllowedDomains, []string{"jasn.asnjournals.org", "www.jasn.org"}...)
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
		Visit(c, opt.URL.String()+"/tab-figures-data")
	}
	postSpiderPrint(opt, &urls)
	return urls
}

func CogitatiopressComSpider(opt *DoiSpiderOpt) (urls []string) {
	if opt.FullText {
		urls = append(urls, stringo.StrReplaceAll(opt.URL.String(),
			"/view/", "/download/")+"/"+path.Base(opt.URL.String()))
	}
	postSpiderPrint(opt, &urls)
	return urls
}

func AdiccionesEsSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "http://www.adicciones.es")
	c.AllowedDomains = append(c.AllowedDomains, []string{"www.adicciones.es"}...)
	if opt.FullText {
		c.OnHTML("#articleFullText a:nth-child(3)", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = strings.ReplaceAll(link, "/view/", "/download/")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

func EurosurveillanceOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "https://www.eurosurveillance.org")
	c.AllowedDomains = append(c.AllowedDomains, []string{"www.eurosurveillance.org"}...)
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
	Visit(c, fmt.Sprintf("https://www.eurosurveillance.org/content/%s#html_fulltext", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

func AerzteblattDeSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "https://www.aerzteblatt.de")
	c.AllowedDomains = append(c.AllowedDomains, []string{"www.aerzteblatt.de"}...)
	if opt.FullText {
		c.OnHTML("a.pdfLink", func(e *colly.HTMLElement) {
			c.OnHTML("div.save a", func(e *colly.HTMLElement) {
				link := e.Attr("href")
				link = stringo.StrReplaceAll(link, "[?].*", "")
				urls = append(urls, linkFilter(link, opt.URL))
			})
			link := e.Attr("href")
			Visit(c, linkFilter(link, opt.URL))
		})
	}
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

func TosOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "http://tos.org")
	c.AllowedDomains = append(c.AllowedDomains, []string{"tos.org"}...)
	if opt.FullText {
		c.OnHTML(".large-links-blue a", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if strings.Contains(link, "docs") {
				urls = append(urls, linkFilter(link, opt.URL))
			}
		})
	}
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

func JstrokeOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "http://www.j-stroke.org")
	c.AllowedDomains = append(c.AllowedDomains, []string{"www.j-stroke.org", "j-stroke.org"}...)
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

func AnnalsOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "http://annals.org")
	c.AllowedDomains = append(c.AllowedDomains, []string{"annals.org"}...)
	if opt.FullText {
		c.OnHTML("#tagmasterPDF", func(e *colly.HTMLElement) {
			link := e.Attr("data-article-url")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

func PortlandpressComSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "https://portlandpress.com")
	c.AllowedDomains = append(c.AllowedDomains, []string{"portlandpress.com"}...)
	if opt.FullText {
		c.OnHTML("a.article-pdfLink", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

func GeoscienceworldOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "https://pubs.geoscienceworld.org")
	c.AllowedDomains = append(c.AllowedDomains, []string{"pubs.geoscienceworld.org"}...)
	if opt.FullText {
		c.OnHTML("a.article-pdfLink", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

func SpringerComSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "https://link.springer.com")
	c.AllowedDomains = append(c.AllowedDomains, []string{"link.springer.com", "idp.springer.com"}...)
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

func AomOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "https://journals.aom.org")
	c.AllowedDomains = append(c.AllowedDomains, []string{"journals.aom.org"}...)
	if opt.FullText {
		if urls = AddPdfSpider(opt); len(urls) > 0 {
			urls[0] = urls[0] + "?download=true"
		}
	}
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

func AltexOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "https://www.altex.org")
	c.AllowedDomains = append(c.AllowedDomains, []string{"www.altex.org"}...)
	if opt.FullText {
		c.OnHTML(".article-sidebar div.download a.pdf:nth-child(1)", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = strings.ReplaceAll(link, "/view/", "/download/")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	if opt.Supplementary {
		c.OnHTML(".article-sidebar div.download a.pdf:nth-child(2)", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = strings.ReplaceAll(link, "/view/", "/download/")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	postSpiderPrint(opt, &urls)
	return urls
}

func IopOrgSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "https://iopscience.iop.org")
	c.AllowedDomains = append(c.AllowedDomains, []string{"iopscience.iop.org"}...)
	if opt.FullText {
		urls = append(urls, fmt.Sprintf("https://iopscience.iop.org/article/%s/pdf", opt.Doi))
	}
	postSpiderPrint(opt, &urls)
	return urls
}
