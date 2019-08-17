// Command visible is a chromedp example demonstrating how to wait until an
// element is visible.
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/chromedp"
)

var ok bool

func main() {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	var attbs []map[string]string

	// run task list
	err := chromedp.Run(ctx, visible("https://www.nejm.org/doi/full/10.1056/NEJMoa1902226", &attbs))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(attbs)
}

func visible(host string, attbs *[]map[string]string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(host),
		chromedp.WaitVisible(`#article_supplementary_material`, chromedp.ByID),
		chromedp.Click("#article_supplementary_material", chromedp.ByID),
		chromedp.ActionFunc(func(context.Context) error {
			log.Printf("waiting 3s for box to become visible")
			return nil
		}),
		chromedp.WaitVisible(`td a[data-download-type='Supplementary Protocol']`, chromedp.ByQuery),
		chromedp.ActionFunc(func(context.Context) error {
			log.Printf(">>>>>>>>>>>>>>>>>>>> td a[data-download-type='Supplementary Protocol'] IS VISIBLE")
			return nil
		}),
		chromedp.AttributesAll(".o-article-body__collapsible-content td a[data-interactiontype=multimedia_download]", attbs, chromedp.ByQueryAll),
		chromedp.ActionFunc(func(context.Context) error {
			log.Printf(">>>>>>>>>>>>>>>>>>>> get values")
			return nil
		}),
	}
}
