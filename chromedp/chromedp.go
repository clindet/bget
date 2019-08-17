package chromedp

import (
	"context"

	"github.com/JhuangLab/bget/log"

	"github.com/chromedp/chromedp"
)

var ok bool

func Chrome2URLs(url string) []string {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	var attbs []map[string]string
	urls := []string{}

	// run task list
	//err := chromedp.Run(ctx, visibleNejm("https://www.nejm.org/doi/full/10.1056/NEJMoa1902226", &attbs))
	err := chromedp.Run(ctx, visibleScienceDirect(url, &attbs))

	if err != nil {
		log.Fatal(err)
	}
	for i := range attbs {
		for k, v := range attbs[i] {
			if k == "href" {
				urls = append(urls, v)
			}
		}
	}
	return urls
}
func visibleScienceDirect(host string, attbs *[]map[string]string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(host),
		chromedp.WaitVisible(`.show-toc-button`, chromedp.ByQuery),
		chromedp.Click(`.show-toc-button`, chromedp.ByQuery),
		chromedp.WaitVisible(`a[href="#app2"]`, chromedp.ByQuery),
		chromedp.Click(`a[href="#app2"]`, chromedp.ByQuery),
		chromedp.WaitVisible(`#app2`, chromedp.ByQuery),
		chromedp.AttributesAll(".Appendices a.icon-link[href]", attbs, chromedp.ByQueryAll),
	}
}
func visibleNejm(host string, attbs *[]map[string]string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(host),
		chromedp.WaitVisible(`#article_supplementary_material`, chromedp.ByID),
		chromedp.Click("#article_supplementary_material", chromedp.ByID),
		chromedp.ActionFunc(func(context.Context) error {
			return nil
		}),
		chromedp.WaitVisible(`td a[data-download-type='Supplementary Protocol']`, chromedp.ByQuery),
		chromedp.ActionFunc(func(context.Context) error {
			return nil
		}),
		chromedp.AttributesAll(".o-article-body__collapsible-content td a[data-interactiontype=multimedia_download]", attbs, chromedp.ByQueryAll),
		chromedp.ActionFunc(func(context.Context) error {
			return nil
		}),
	}
}
