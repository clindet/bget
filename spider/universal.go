package spider

import (
	"fmt"
	"net/http"
	neturl "net/url"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	cnet "github.com/openbiox/ligo/net"
	"github.com/openbiox/ligo/stringo"
)

func UniVersalDoiSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org"),
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
		c.OnHTML("meta[title='Full Text (PDF)']", func(e *colly.HTMLElement) {
			link := e.Attr("content")
			urls = append(urls, linkFilter(link, opt.URL))
		})
		// https://www.microbiologyresearch.org/ specific
		c.OnHTML("form.ft-download-content__form--pdf[action]", func(e *colly.HTMLElement) {
			link := e.Attr("action")
			link = stringo.StrReplaceAll(link, "pdf[?].*", "pdf")
			urls = append(urls, linkFilter(link, opt.URL))
		})
		c.OnHTML("a.article-pdfLink[data-article-url]", func(e *colly.HTMLElement) {
			link := e.Attr("data-article-url")
			urls = append(urls, linkFilter(link, opt.URL))
		})
		staticUrl := static2pdf(opt)
		if staticUrl != "" {
			urls = append(urls, linkFilter(staticUrl, opt.URL))
		}
	}
	if opt.Supplementary {
		supplPage := suppl2page(opt)
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

		c.Visit(supplPage)
	}
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	c.Visit(fmt.Sprintf("https://doi.org/%s", opt.Doi))
	return urls
}

func AddPdfSpider(opt *DoiSpiderOpt) (urls []string) {
	if opt.FullText {
		link := opt.URL.String()
		link = stringo.StrReplaceAll(link, "/doi/", "/doi/pdf/")
		urls = append(urls, link)
	}
	return urls
}

func AddPdfplusSpider(opt *DoiSpiderOpt) (urls []string) {
	if opt.FullText {
		link := opt.URL.String()
		link = stringo.StrReplaceAll(link, "/doi/", "/doi/pdfplus/")
		urls = append(urls, link)
	}
	return urls
}

func DirectSpider(opt *DoiSpiderOpt) (urls []string) {
	if opt.FullText {
		urls = append(urls, opt.URL.String())
	}
	return urls
}

func AddPdfplusWithSupplSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org"),
		colly.MaxDepth(2),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.URL != nil {
		c.AllowedDomains = append(c.AllowedDomains, opt.URL.Host)
	}
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	urls = AddPdfplusSpider(opt)
	if opt.Supplementary {
		c.OnHTML(".suppl_list a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = "https://" + opt.URL.Hostname() + link
			urls = append(urls, link)
		})
		c.Visit(stringo.StrReplaceAll(opt.URL.String(), "/doi/", "/doi/suppl/"))
	}
	return urls
}

func AddPdfWithSupplSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org"),
		colly.MaxDepth(2),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.URL != nil {
		c.AllowedDomains = append(c.AllowedDomains, opt.URL.Host)
	}
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	urls = AddPdfSpider(opt)
	if opt.Supplementary {
		c.OnHTML(".suppl_list a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			link = "https://" + opt.URL.Hostname() + link
			urls = append(urls, link)
		})
		c.Visit(stringo.StrReplaceAll(opt.URL.String(), "/doi/", "/doi/suppl/"))
	}
	return urls
}

func AddLastPdfSpider(opt *DoiSpiderOpt) (urls []string) {
	return []string{opt.URL.String() + "/pdf"}
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
	return urls
}

func KoreaMedSpider(opt *DoiSpiderOpt, hostname string) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", hostname),
		colly.MaxDepth(5),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
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
			c.Visit(fmt.Sprintf("%s/%s", hostname, link))
		})
	}
	c.Visit(fmt.Sprintf("https://"+hostname+"/DOIx.php?id=%s", opt.Doi))
	return urls
}

func ReplaceHtmlSpider(opt *DoiSpiderOpt) (urls []string) {
	if opt.FullText {
		link := stringo.StrReplaceAll(opt.URL.String(), ".htm$|.html$", ".pdf")
		urls = append(urls, link)
	}
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
	if !strings.Contains(link, "http") && u != nil {
		link = u.Scheme + "://" + u.Host + link
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
