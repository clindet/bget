package spider

import (
	"fmt"
	neturl "net/url"
	"strings"

	"github.com/gocolly/colly"
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

func linkFilter(link string, u *neturl.URL) string {
	if !strings.Contains(link, "http") {
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
