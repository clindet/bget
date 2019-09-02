package spider

import (
	"fmt"
	neturl "net/url"
	"strings"

	"github.com/gocolly/colly"
	"github.com/openbiox/butils/log"
	bspider "github.com/openbiox/butils/spider"
)

func UniVersalDoiSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org"),
		colly.MaxDepth(1),
	)
	bspider.SetSpiderProxy(c, opt.Proxy, opt.Timeout)
	c.AllowedDomains = append(c.AllowedDomains, opt.URL.Host)
	if opt.FullText {
		c.OnHTML("meta[name=citation_pdf_url]", func(e *colly.HTMLElement) {
			link := e.Attr("content")
			urls = append(urls, linkFilter(link, opt.URL))
		})
		c.OnHTML("meta[title='Full Text (PDF)']", func(e *colly.HTMLElement) {
			link := e.Attr("content")
			urls = append(urls, linkFilter(link, opt.URL))
		})
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
