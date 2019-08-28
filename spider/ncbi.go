package spider

import (
	"fmt"
	"strings"

	//"github.com/Miachol/bget/chromedp"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	butils "github.com/openbiox/butils"
	"github.com/openbiox/butils/log"
)

// GeoSpider access https://www.ncbi.nlm.nih.gov/geo files via spider
func GeoSpider(query, proxy string, timeout int) (gseURLs []string, gplURLs []string, sraLink string) {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("www.ncbi.nlm.nih.gov"),
		colly.MaxDepth(1),
	)
	setSpiderProxy(c, proxy, timeout)
	extensions.RandomUserAgent(c)

	c.OnHTML("table td a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(link, "/geo/download/?acc=GSE") {
			gseURLs = append(gseURLs, "https://www.ncbi.nlm.nih.gov"+link)
		} else if strings.Contains(link, "/geo/download/?acc=GPL") {
			gplURLs = append(gplURLs, "https://www.ncbi.nlm.nih.gov"+link)
		}
	})

	c.OnHTML("input[name=fulltable]", func(e *colly.HTMLElement) {
		link := e.Attr("onclick")
		if strings.Contains(link, "OpenLink") {
			link = "https://www.ncbi.nlm.nih.gov" + butils.StrReplaceAll(link, "(OpenLink[(])|(')", "")
			link = butils.StrReplaceAll(link, ",.*", "")
			gplURLs = append(gplURLs, link)
		}
	})

	c.OnHTML("table td a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(link, "geo/query/acc.cgi?acc=GPL") && !strings.Contains(link, "targ=self") {
			c.Visit("https://www.ncbi.nlm.nih.gov" + link)
		}
	})

	c.OnHTML("tr td a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(link, "/Traces/study/?acc=") {
			link = "https://www.ncbi.nlm.nih.gov" + e.Attr("href")
			sraLink = link
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit(fmt.Sprintf("https://www.ncbi.nlm.nih.gov/geo/query/acc.cgi?acc=%s", query))
	return gseURLs, gplURLs, sraLink
}
