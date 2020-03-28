package spider

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	cnet "github.com/openbiox/ligo/net"
	stringo "github.com/openbiox/ligo/stringo"
)

// ScihupSpider access http://sci-hub.tw/ files via spider
func ScihupSpider(opt *DoiSpiderOpt) (urls []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("doi.org", "sci-hub.tw"),
		colly.MaxDepth(1),
	)
	cnet.SetCollyProxy(c, opt.Proxy, opt.Timeout)
	extensions.RandomUserAgent(c)
	if opt.FullText {
		c.OnHTML("#buttons a[onclick]", func(e *colly.HTMLElement) {
			link := e.Attr("onclick")
			link = stringo.StrExtract(link, "//.*", 1)[0]
			link = "http:" + link
			link = stringo.StrReplaceAll(link, "'$", "")
			urls = append(urls, link)
		})
	}
	c.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting %s", r.URL.String())
	})
	Visit(c, fmt.Sprintf("http://sci-hub.tw/%s", opt.Doi))
	return urls
}
