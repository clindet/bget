package spider

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/openbiox/butils/log"
	bspider "github.com/openbiox/butils/spider"
	"github.com/openbiox/butils/stringo"
)

// GeoSpider access https://www.ncbi.nlm.nih.gov/geo files via spider
func GeoSpider(opt *QuerySpiderOpt) (gseURLs []string, gplURLs []string, sraLink string) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.ncbi.nlm.nih.gov"),
		colly.MaxDepth(1),
	)
	bspider.SetSpiderProxy(c, opt.Proxy, opt.Timeout)
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
			link = "https://www.ncbi.nlm.nih.gov" + stringo.StrReplaceAll(link, "(OpenLink[(])|(')", "")
			link = stringo.StrReplaceAll(link, ",.*", "")
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
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	c.Visit(fmt.Sprintf("https://www.ncbi.nlm.nih.gov/geo/query/acc.cgi?acc=%s", opt.Query))
	return gseURLs, gplURLs, sraLink
}
func PmcSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.ncbi.nlm.nih.gov"),
		colly.MaxDepth(1),
	)
	bspider.SetSpiderProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.FullText {
		c.OnHTML(fmt.Sprintf(".doi b:contains('%s')", opt.Doi), func(e *colly.HTMLElement) {
			e.DOM.Parents().Filter(".rslt").Find(".aux .links a.view").Each(func(i int, s *goquery.Selection) {
				link, _ := s.Attr("href")
				if strings.Contains(link, "pdf") {
					link = "https://www.ncbi.nlm.nih.gov" + link
					urls = append(urls, link)
				}
			})
		})
	}
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	c.Visit(fmt.Sprintf("https://www.ncbi.nlm.nih.gov/pmc/?term=%s", opt.Doi))
	return urls
}
