package spider

import (
	"fmt"
	"net/http"
	neturl "net/url"
	"path"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	cnet "github.com/openbiox/ligo/net"
	"github.com/openbiox/ligo/slice"
	"github.com/openbiox/ligo/stringo"
)

func UniVersalDoiSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "")
	urls = append(urls, directDoiSites(opt)...)
	if opt.FullText {
		UniVersalDoiSpiderListenPart1(c, opt, &urls)
		UniVersalDoiSpiderListenPart2(c, opt, &urls)
		UniVersalDoiSpiderListenPart3(c, opt, &urls)
		UniVersalDoiSpiderListenPart4(c, opt, &urls)
		staticUrl := static2pdf(opt)
		if staticUrl != "" {
			urls = append(urls, linkFilter(staticUrl, opt.URL))
		}
	}
	if opt.Supplementary {
		c.OnHTML("a.rewritten[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
		// https://www.microbiologyresearch.org/ specific
		c.OnHTML("#supplementary_data form.js-ft-download-form[action]", func(e *colly.HTMLElement) {
			link := e.Attr("action")
			link = stringo.StrReplaceAll(link, "[?].*", "")
			urls = append(urls, linkFilter(link, opt.URL))
		})
		c.OnHTML("#SuppDataIndexList .textoptionsFulltext ul.fulltext li a", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			urls = append(urls, linkFilter(link, opt.URL))
		})
	}
	if opt.FullText {
		visitFilter(c, opt)
	}
	if opt.Supplementary {
		supplPage := suppl2page(opt)
		if supplPage != "" {
			Visit(c, supplPage)
		}
	}
	postSpiderPrint(opt, &urls)
	return urls
}

func visitFilter(c *colly.Collector, opt *DoiSpiderOpt) {
	dxDoiSites := []string{"10.1561/", "10.15585/"}
	dxsites := false
	for _, v := range dxDoiSites {
		if strings.Contains(opt.Doi, v) {
			dxsites = true
		}
	}
	if strings.Contains(opt.Doi, "10.1049/htl") {
		Visit(c, fmt.Sprintf("https://digital-library.theiet.org/content/journals/%s", opt.Doi))
	} else if dxsites {
		Visit(c, fmt.Sprintf("https://dx.doi.org/%s", opt.Doi))
	} else {
		Visit(c, fmt.Sprintf("https://doi.org/%s", opt.Doi))
	}
}

func UniVersalDoiSpiderListenPart1(c *colly.Collector, opt *DoiSpiderOpt, urls *[]string) {
	c.OnHTML("meta[title='Full Text (PDF)']", func(e *colly.HTMLElement) {
		link := e.Attr("content")
		*urls = append(*urls, linkFilter(link, opt.URL))
	})
}

func UniVersalDoiSpiderListenPart2(c *colly.Collector, opt *DoiSpiderOpt, urls *[]string) {
	hrefPattern := []string{"#view-as-pdf", ".pdf_link a", "li a.format_pdf",
		"a.abstractFullText", "a.btn_descargaPDF",
		"#cpCuerpo_lnkDocumento", "#'DOWNLOAD PDFButton'", "li.calag-nav-pdf a",
		"#ctl00_ContentPlaceHolder_Right_download_pdf", "#columnRight table tr td:nth-child(3) a",
		".row .textoptionsFulltext ul.fulltext li a.pdf", "main div.row div.page-format a[target=_blank]",
		"#Toolbar a.article-pdfLink", ".article-head .btn-multi-block a.content-download",
		".d-flex main.syndicate .right-col-info div:nth-child(5) .card-body a:nth-child(5)[download]",
		"record-action-bar .list-group a.pdf", ".flexme-article .files .article_doc ul li:nth-child(3) a",
		".intent_pdf_link", ".icon-102_download_pdf a.download", ".download-and-link ul.additional_info a.triangle",
		".page-sidebar a.download-pdf"}
	for _, v := range hrefPattern {
		c.OnHTML(v, func(e *colly.HTMLElement) {
			link := e.Attr("href")
			*urls = append(*urls, linkFilter(link, opt.URL))
		})
	}
}
func UniVersalDoiSpiderListenPart3(c *colly.Collector, opt *DoiSpiderOpt, urls *[]string) {
	c.OnHTML("a.article-pdfLink[data-article-url]", func(e *colly.HTMLElement) {
		link := e.Attr("data-article-url")
		*urls = append(*urls, linkFilter(link, opt.URL))
	})
	c.OnHTML("iframe.pdf[data-src]", func(e *colly.HTMLElement) {
		link := e.Attr("data-src")
		*urls = append(*urls, linkFilter(link, opt.URL))
	})
	c.OnHTML("td b a[target='_blank']", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(link, "pdf.php") {
			*urls = append(*urls, linkFilter(link, opt.URL))
		}
	})
	c.OnHTML("ul.galleys_links li a.obj_galley_link", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		c.OnHTML("a.download", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			*urls = append(*urls, linkFilter(link, opt.URL))
		})
		Visit(c, linkFilter(link, opt.URL))
	})
	c.OnHTML(".article-content table td a:first-child", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(link, ".pdf") {
			*urls = append(*urls, linkFilter(link, opt.URL))
		}
	})
	c.OnHTML("#articleFullText a", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(link, "/download/") {
			*urls = append(*urls, linkFilter(link, opt.URL))
		}
	})
	c.OnHTML("a.pdf[data-popup]", func(e *colly.HTMLElement) {
		link := e.Attr("data-popup")
		if strings.Contains(link, "/search/") {
			*urls = append(*urls, linkFilter(link, opt.URL))
		}
	})
}

func UniVersalDoiSpiderListenPart4(c *colly.Collector, opt *DoiSpiderOpt, urls *[]string) {
	c.OnHTML("td.auto-style21 a.auto-style15", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		link = stringo.StrReplaceAll(link, "^../../", "/")
		*urls = append(*urls, linkFilter(link, opt.URL))
	})
	c.OnHTML("frameset frame:first-child", func(e *colly.HTMLElement) {
		link := e.Attr("src")
		if strings.Contains(link, "saje/article") {
			link = stringo.StrReplaceAll(link, "viewPDFInterstitial", "viewFile")
			*urls = append(*urls, linkFilter(link, opt.URL))
		}
	})
	c.OnHTML("#content p a.linkintext", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		link = stringo.StrReplaceAll(link, "^../../", "/")
		*urls = append(*urls, linkFilter(link, opt.URL))
	})
	c.OnHTML("div.pull-right a.btn-galley", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(link, "/view/") {
			link = strings.ReplaceAll(link, "/view/", "/download/")
			*urls = append(*urls, linkFilter(link, opt.URL))
		}
	})
	// https://www.microbiologyresearch.org/ specific
	c.OnHTML("form.ft-download-content__form--pdf[action]", func(e *colly.HTMLElement) {
		link := e.Attr("action")
		link = stringo.StrReplaceAll(link, "pdf[?].*", "pdf")
		*urls = append(*urls, linkFilter(link, opt.URL))
	})
	c.OnHTML("a[title='Article permanent link']", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		link = stringo.StrReplaceAll(link, "full", "")
		link = link + "pdf/" + path.Base(link) + ".pdf"
		*urls = append(*urls, linkFilter(link, opt.URL))
	})
}

func directDoiSites(opt *DoiSpiderOpt) (urls []string) {
	for _, v := range DirectJournalLinks {
		if strings.Contains(opt.Doi, v) {
			urls = append(urls, DirectSpider(opt)...)
		}
	}
	postSpiderPrint(opt, &urls)
	return urls
}

func AddPdfSpider(opt *DoiSpiderOpt) (urls []string) {
	if opt.FullText {
		link := opt.URL.String()
		link = stringo.StrReplaceAll(link, "/doi/", "/doi/pdf/")
		urls = append(urls, link)
	}
	postSpiderPrint(opt, &urls)
	return urls
}

func AddPdfSuffixSpider(opt *DoiSpiderOpt) (urls []string) {
	if opt.FullText {
		link := opt.URL.String()
		urls = append(urls, link+".pdf")
	}
	postSpiderPrint(opt, &urls)
	return urls
}

func AddTextPdfSuffixSpider(opt *DoiSpiderOpt) (urls []string) {
	if opt.FullText {
		link := opt.URL.String()
		if strings.Contains(link, "nn.neurology.org") {
			link = link + ".full-text.pdf"
			urls = append(urls, strings.ReplaceAll(link, "/content/", "/content/nnn/"))
		} else {
			link = link + "-text.pdf"
			urls = append(urls, link)
		}
	}
	postSpiderPrint(opt, &urls)
	return urls
}

func AddPdfplusSpider(opt *DoiSpiderOpt) (urls []string) {
	if opt.FullText {
		link := opt.URL.String()
		link = stringo.StrReplaceAll(link, "/doi/", "/doi/pdfplus/")
		urls = append(urls, link)
	}
	postSpiderPrint(opt, &urls)
	return urls
}

func DirectSpider(opt *DoiSpiderOpt) (urls []string) {
	if opt.FullText {
		urls = append(urls, opt.URL.String())
	}
	postSpiderPrint(opt, &urls)
	return urls
}

func AddPdfplusWithSupplSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "")
	urls = AddPdfplusSpider(opt)
	if opt.Supplementary {
		c.OnHTML(".suppl_list a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = "https://" + opt.URL.Hostname() + link
			urls = append(urls, link)
		})
		Visit(c, stringo.StrReplaceAll(opt.URL.String(), "/doi/", "/doi/suppl/"))
	}
	postSpiderPrint(opt, &urls)
	return urls
}

func AddPdfWithSupplSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "")
	urls = AddPdfSpider(opt)
	if opt.Supplementary {
		c.OnHTML(".suppl_list a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = "https://" + opt.URL.Hostname() + link
			urls = append(urls, link)
		})
		Visit(c, stringo.StrReplaceAll(opt.URL.String(), "/doi/", "/doi/suppl/"))
	}
	postSpiderPrint(opt, &urls)
	return urls
}

func AddLastPdfSpider(opt *DoiSpiderOpt) (urls []string) {
	urls = append(urls, opt.URL.String()+"/pdf")
	postSpiderPrint(opt, &urls)
	return urls
}

func AddDownloadSpider(opt *DoiSpiderOpt) (urls []string) {
	if opt.FullText {
		link := stringo.StrReplaceAll(opt.URL.String(), "/content/abstract/", "/download/")
		link, err := RetriveRedirectLink(link, opt.Timeout, opt.Proxy)
		if err != nil {
			log.Warnln(err)
		}
		urls = append(urls, link)
	}
	postSpiderPrint(opt, &urls)
	return urls
}

func ReplaceHtmlSpider(opt *DoiSpiderOpt) (urls []string) {
	if opt.FullText {
		link := stringo.StrReplaceAll(opt.URL.String(), ".htm$|.html$", ".pdf")
		urls = append(urls, link)
	}
	postSpiderPrint(opt, &urls)
	return urls
}

func RetriveRedirectLink(url string, timeout int, proxy string) (string, error) {
	client := cnet.NewHTTPClient(timeout, proxy)
	req, _ := http.NewRequest("HEAD", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	return resp.Request.URL.String(), nil
}

func linkFilter(link string, u *neturl.URL) string {
	link = siteLinkFilter(link)
	if !strings.Contains(link, "http") && u != nil {
		if !stringo.StrDetect(link, "^/") {
			link = "/" + link
		}
		link = u.Scheme + "://" + u.Host + link
	}
	return link
}

func siteLinkFilter(link string) string {
	twoSlashPattern := []string{"^//calag.ucanr.edu", "^//prc.springeropen.com", "//aricjournal.biomedcentral.com"}
	for _, v := range twoSlashPattern {
		if stringo.StrDetect(link, v) {
			link = stringo.StrReplaceAll(link, v, "")
		}
	}
	twoSlashPattern = []string{"//journals.iucr.org"}
	for _, v := range twoSlashPattern {
		if stringo.StrDetect(link, v) {
			link = "http:" + link
		}
	}
	return link
}

func suppl2page(opt *DoiSpiderOpt) (link string) {
	baseLink := opt.URL.Scheme + "://" + opt.URL.Host +
		opt.URL.EscapedPath()
	if stringo.StrDetect(opt.Doi, "10.2337/|10.1242/") {
		link = baseLink + ".supplemental"
	}
	return link
}

func static2pdf(opt *DoiSpiderOpt) (url string) {
	return url
}

func initDoiColley(opt *DoiSpiderOpt, host string) *colly.Collector {
	c := initBasicColley(opt.Timeout, opt.Proxy)
	if opt.URL == nil && host != "" {
		u, _ := neturl.Parse(host)
		opt.URL = u
	}
	if opt.URL != nil {
		c.AllowedDomains = append(c.AllowedDomains, opt.URL.Host)
	}
	c.OnHTML("meta[title='Full Text (PDF)']", func(e *colly.HTMLElement) {
		link := e.Attr("content")
		(*opt.CitationMeta)["citation_pdf_url"] = linkFilter(link, opt.URL)
	})
	c.OnHTML("meta[name]", func(e *colly.HTMLElement) {
		content := e.Attr("content")
		name := e.Attr("name")
		if _, ok := (*opt.CitationMeta)[name]; ok && content != "" && (*opt.CitationMeta)[name] != "" &&
			(*opt.CitationMeta)[name] != content {
			(*opt.CitationMeta)[name] = (*opt.CitationMeta)[name] + "; " + content
		} else if content != "" {
			(*opt.CitationMeta)[name] = content
		}
	})
	return c
}

func initQueryColley(opt *QuerySpiderOpt, host string) *colly.Collector {
	c := initBasicColley(opt.Timeout, opt.Proxy)
	return c
}

func initBasicColley(timeout int, proxy string) *colly.Collector {
	c := colly.NewCollector(
		colly.MaxDepth(5),
	)
	c.AllowedDomains = append(c.AllowedDomains, UniversalJournalLinks...)
	cnet.SetCollyProxy(c, proxy, timeout)
	extensions.RandomUserAgent(c)
	extensions.Referer(c)
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	return c
}

func Visit(c *colly.Collector, url string) error {
	err := c.Visit(url)
	if err != nil && strings.Contains(err.Error(), "Not following redirect to") {
		link := stringo.StrSplit(err.Error(), "Not following redirect to ", 2)[1]
		link = stringo.StrReplaceAll(link, " because its not in AllowedDomains", "")
		log.Infof("Adding %s to AllowedDomains domain and retring.....", link)
		c.AllowedDomains = append(c.AllowedDomains, link)
		c.AllowURLRevisit = true
		err = c.Visit(url)
		c.AllowURLRevisit = false
	}
	if err != nil {
		log.Warnln(err)
	}
	return err
}

func postSpiderPrint(opt *DoiSpiderOpt, urls *[]string) {
	if len(*urls) == 0 && opt.FullText {
		for k, v := range *opt.CitationMeta {
			if k == "citation_pdf_url" && v != "" {
				*urls = append(*urls, v)
			}
		}
	}
	if len(*urls) > 0 {
		*urls = slice.DropSliceDup(*urls)
		log.Infof(`Finding total %d links from %s: ["%s"]`,
			len(*urls), opt.Doi, strings.Join(*urls, `", "`))
	}
}
