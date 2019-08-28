package spider

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	butils "github.com/openbiox/butils"
	"github.com/openbiox/butils/log"
)

// ScihubSpider access http://sci-hub.tw/ files via spider
func ScihubSpider(doi, proxy string, timeout int) (urls []string) {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("doi.org", "sci-hub.tw"),
		colly.MaxDepth(1),
	)
	setSpiderProxy(c, proxy, timeout)
	extensions.RandomUserAgent(c)

	c.OnHTML("#buttons a[onclick]", func(e *colly.HTMLElement) {
		link := e.Attr("onclick")
		link = butils.StrExtract(link, "//.*", 1)[0]
		link = "http:" + link
		link = butils.StrReplaceAll(link, "'$", "")
		urls = append(urls, link)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit(fmt.Sprintf("http://sci-hub.tw/%s", doi))
	return urls
}
