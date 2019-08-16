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
	"10.1126": ScienseComSpider,
	"10.1016": CellComSpider,
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

// NatureComSpider access Nature.com files via spider
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

// ScienseComSpider access sciencemag.org journal files via spider
func ScienseComSpider(doi string) []string {
	urls := []string{}
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("doi.org", "advances.sciencemag.org", "immunology.sciencemag.org",
			"robotics.sciencemag.org", "stke.sciencemag.org", "stm.sciencemag.org", "secure.jbs.elsevierhealth.com",
			"id.elsevier.com"),
		colly.MaxDepth(2),
	)
	extensions.RandomUserAgent(c)

	c.OnHTML("meta[name=citation_pdf_url]", func(e *colly.HTMLElement) {
		link := e.Attr("content")
		u, _ := url.Parse(link)
		link = u.Scheme + "://" + u.Host + u.Path
		urls = append(urls, link)
		c.OnHTML("a.rewritten[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			u2, _ := url.Parse(link)
			link = u.Scheme + "://" + u.Host + u2.Path
			urls = append(urls, link)
		})
		c.Visit(strings.Replace(link, ".full.pdf", "", 1) + "/tab-figures-data")
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit(fmt.Sprintf("https://doi.org/%s", doi))
	return urls
}

// CellComSpider access cell.com journal files via spider
func CellComSpider(doi string) []string {
	urls := []string{}
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("doi.org", "www.cell.com", "cell.com", "linkinghub.elsevier.com", "secure.jbs.elsevierhealth.com",
			"id.elsevier.com", "www.cancercell.org", "www.sciencedirect.com"),
		colly.MaxDepth(2),
	)
	extensions.RandomUserAgent(c)

	c.OnHTML("#redirectURL", func(e *colly.HTMLElement) {
		link := e.Attr("value")
		u, _ := url.Parse(link)
		link, _ = url.QueryUnescape(u.Path)
		c.Visit(link)
	})

	c.OnHTML("a.article-tools__item__displayStandardPdf[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if link != "#" {
			link = "https://www.cell.com" + link
			urls = append(urls, link)
		}
	})

	c.OnHTML("a.article-tools__item__displayExtendedPdf[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if link != "#" {
			link = "https://www.cell.com" + link
			urls = append(urls, link)
		}
	})

	c.OnHTML(".article-tools__pdf a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if link != "#" {
			link = "https://www.cell.com" + link
			urls = append(urls, link)
		}
	})

	// cell stem cell start
	c.OnHTML("div.PdfDownloadButton a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		link = "https://www.sciencedirect.com" + link
		c.Visit(link)
	})

	c.OnHTML("span.article-attachment a.download-link[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println(link)
		urls = append(urls, link)
	})
	// cell stem cell end

	c.OnHTML("a.supplemental-information__download[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if link != "#" {
			link = "https://www.cell.com" + link
			urls = append(urls, link)
		}

	})

	c.OnHTML("meta[HTTP-EQUIV=REFRESH]", func(e *colly.HTMLElement) {
		link := e.Attr("content")
		link = utils.StrReplaceAll(link, ".* url='", "")
		link = utils.StrReplaceAll(link, "'$", "")
		link = "https://linkinghub.elsevier.com/" + link
		c.Visit(link)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit(fmt.Sprintf("https://doi.org/%s", doi))
	return urls
}
