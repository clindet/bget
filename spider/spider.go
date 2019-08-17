package spider

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/JhuangLab/bget/chromedp"

	"github.com/JhuangLab/bget/log"
	"github.com/JhuangLab/bget/utils"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

// NatureComSpider access Nature.com files via spider
func NatureComSpider(doi string) (urls []string) {
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
		urls = append(urls, "https://nature.com"+link)
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
func ScienseComSpider(doi string) (urls []string) {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("doi.org", "advances.sciencemag.org", "immunology.sciencemag.org",
			"robotics.sciencemag.org", "stke.sciencemag.org", "stm.sciencemag.org", "secure.jbs.elsevierhealth.com",
			"id.elsevier.com", "science.sciencemag.org", "www.sciencemag.org"),
		colly.MaxDepth(2),
	)
	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	// Sciense advance
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

	c.OnHTML("div.panels-ajax-tab-wrap-jnl_sci_tab_pdf a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		urls = append(urls, link)
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
			"id.elsevier.com", "www.cancercell.org", "www.sciencedirect.com",
			"pdf.sciencedirectassets.com"),
		colly.MaxDepth(3),
	)
	extensions.RandomUserAgent(c)
	extensions.Referer(c)

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
		urls = append(urls, link)
	})

	c.OnHTML("span.article-attachment a.download-link[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
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
		link = "https://linkinghub.elsevier.com" + link
		c.Visit(link)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
		if utils.StrDetect(r.URL.String(), "^https://www.sciencedirect.com") {
			urls = append(urls, chromedp.Chrome2URLs(r.URL.String())...)
		}
	})

	// Start scraping on https://hackerspaces.org
	c.Visit(fmt.Sprintf("https://doi.org/%s", doi))
	return urls
}

// BloodJournalSpider access http://www.bloodjournal.org files via spider
func BloodJournalSpider(doi string) (urls []string) {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("doi.org", "www.bloodjournal.org", "signin.hematology.org"),
		colly.MaxDepth(1),
	)
	extensions.RandomUserAgent(c)

	c.OnHTML("a[data-panel-name=jnl_bloodjournal_tab_pdf]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		link = "http://www.bloodjournal.org" + link
		urls = append(urls, link)
	})
	c.OnHTML("a.rewritten[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		link = "http://www.bloodjournal.org" + link
		urls = append(urls, link)
	})
	c.OnHTML("a.[data-panel-name=jnl_bloodjournal_tab_data]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
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

// NejmSpider access http://www.nejm.org files via spider
func NejmSpider(doi string) (urls []string) {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("doi.org", "www.nejm.org"),
		colly.MaxDepth(1),
	)
	extensions.RandomUserAgent(c)

	c.OnHTML("a[data-tooltip='Download PDF']", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		link = "https://www.nejm.org" + link
		urls = append(urls, link)
	})
	c.OnHTML("a[data-interactionType=multimedia_download]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(link, "doi/suppl") {
			link = "https://www.nejm.org" + link
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
