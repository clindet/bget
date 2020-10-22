package spider

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	stringo "github.com/openbiox/ligo/stringo"
)

// ScihupSpider access http://sci-hub.tw/ files via spider
func ScihupSpider(opt *DoiSpiderOpt) (urls []string) {
	c := initDoiColley(opt, "")
	c.AllowedDomains = append(c.AllowedDomains, []string{"sci-hub.tw"}...)
	if opt.FullText {
		c.OnHTML("#buttons a[onclick]", func(e *colly.HTMLElement) {
			link := e.Attr("onclick")
			link = stringo.StrExtract(link, "//.*", 1)[0]
			link = "http:" + link
			link = stringo.StrReplaceAll(link, "'$", "")
			urls = append(urls, link)
		})
	}
	Visit(c, fmt.Sprintf("http://sci-hub.tw/%s", opt.Doi))
	return urls
}
