package spider

import (
	"fmt"
	"strings"

	"github.com/JhuangLab/bget/log"
	"github.com/gocolly/colly"
)

// ZenodoSpider access Zendo files via spider
func ZenodoSpider(doi string) []string {
	urls := []string{}
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("doi.org", "zenodo.org"),
		colly.MaxDepth(1),
	)
	// On every a element which has href attribute call callback
	c.OnHTML("tbody a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(link, "?download=1") {
			urls = append(urls, "https://zenodo.org"+link)
		}
		// Print link
		//fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		//c.Visit(e.Request.AbsoluteURL(link))

	})
	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit(fmt.Sprintf("https://doi.org/%s", doi))
	return urls
}
