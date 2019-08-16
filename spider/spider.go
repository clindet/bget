package spider

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/JhuangLab/bget/log"
	"github.com/JhuangLab/bget/utils"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

type doiSpiderType struct {
	doiOrg string
	spider map[string]interface{}
}

// DoiSpidersPool map doi to golang function
var DoiSpidersPool = map[string]func(doi string) []string{
	"10.5281": ZenodoSpider,
	"10.1038": NatureComSpider,
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
	extensions.RandomUserAgent(c)

	// On every a element which has href attribute call callback
	c.OnHTML("tbody a.filename[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(link, "?download=1") {
			u, _ := url.Parse(link)
			link = "https://zenodo.org" + u.Host + u.Path
			urls = append(urls, link)
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

// NatureComSpider access Zendo files via spider
func NatureComSpider(doi string) []string {
	urls := []string{}
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("doi.org", "www.nature.com", "idp.nature.com"),
		colly.MaxDepth(1),
	)
	extensions.RandomUserAgent(c)

	// On every a element which has href attribute call callback
	c.OnHTML("a.print-link[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if !strings.HasPrefix(link, "http") {
			urls = append(urls, "https://nature.com"+link)
		} else {
			u, _ := url.Parse(link)
			linkTmp := strings.Split(u.Path, "/")
			linkTmp[2] = utils.StrReplaceAll(linkTmp[2], "art:", "art%3A")
			newLink := append(linkTmp[0:2], strings.Join(linkTmp[2:4], "%2F"))
			newLink = append(newLink, linkTmp[4:len(linkTmp)]...)
			link = strings.Join(newLink, "/")
			link = u.Scheme + "://" + u.Host + link
			urls = append(urls, link)
		}
	})

	// On every a element which has href attribute call callback
	c.OnHTML("a.c-pdf-download__link[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if !strings.HasPrefix(link, "http") {
			urls = append(urls, "https://nature.com"+link)
		} else {
			u, _ := url.Parse(link)
			linkTmp := strings.Split(u.Path, "/")
			linkTmp[2] = utils.StrReplaceAll(linkTmp[2], "art:", "art%3A")
			newLink := append(linkTmp[0:1], strings.Join(linkTmp[2:3], "%2F"))
			newLink = append(newLink, linkTmp[4:len(linkTmp)]...)
			link = strings.Join(newLink, "/")
			link = u.Scheme + "://" + u.Host + link
			urls = append(urls, link)
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit(fmt.Sprintf("https://doi.org/%s", doi))
	return urls
}
