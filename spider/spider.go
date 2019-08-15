package spider

import (
	"fmt"
	"strings"

	"github.com/JhuangLab/bget/log"
	"github.com/gocolly/colly"
)

type doiSpiderType struct {
	doiOrg string
	spider map[string]interface{}
}

// DoiSpidersPool map doi to golang function
var DoiSpidersPool = map[string]func(doi string) []string{
	"10.5281": ZenodoSpider,
}

// ZenodoSpider access Zendo files via spider
func ZenodoSpider(doi string) []string {
	urls := []string{}
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("doi.org", "zenodo.org"),
		colly.MaxDepth(1),
	)
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.108 Safari/537.36"

	// On every a element which has href attribute call callback
	c.OnHTML("tbody a.filename[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(link, "?download=1") {
			urls = append(urls, "https://zenodo.org"+link)
		}
	})
	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Connection", "keep-alive")
		log.Infof("Visiting %s", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit(fmt.Sprintf("https://doi.org/%s", doi))
	return urls
}
