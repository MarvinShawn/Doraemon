package juejin

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

//JuejinCrawler Crawler
type JuejinCrawler struct{}

//Start 开始请求网页并解析
func Start() {

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36"),
	)

	c.OnHTML("body", func(e *colly.HTMLElement) {
		e.DOM.Find("li .content-box").Each(func(i int, selection *goquery.Selection) {
			titleSelection := selection.Find(".info-box .title").First()
			link, ok := titleSelection.Attr("href")
			countSelection := selection.Find(".info-box .count").First()
			picSelection, picOk := selection.Find(".info-box+div").First().Attr("data-src")
			if ok {
				fmt.Printf("title->%v\n", titleSelection.Text())
				fmt.Printf("link->https://juejin.im%v\n", link)
				fmt.Printf("praise->%v\n", countSelection.Text())
				if picOk {
					fmt.Printf("picUrl->%v\n", picSelection)

				}
				fmt.Println("====================================")
			}
		})

	})

	c.Visit("https://juejin.im/")

}
