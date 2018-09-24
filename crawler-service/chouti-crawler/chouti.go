// Package chouti 抽屉新热榜
package chouti

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

//ChoutiCrawler Crawler
type ChoutiCrawler struct{}

//Start 开始请求网页并解析
func (crawler *ChoutiCrawler) Start() {

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36"),
	)

	c.OnHTML("body", func(e *colly.HTMLElement) {
		e.DOM.Find(".content-list .item").Each(func(i int, selection *goquery.Selection) {
			picLink, _ := selection.Find(".news-pic img").Attr("original")
			detailLink, _ := selection.Find(".news-content .part1 a").First().Attr("href")
			title, _ := selection.Find(".news-content .part2").First().Attr("share-title")
			praiseSelection := selection.Find(".news-content .part2 .digg-a b").First()

			fmt.Printf("picLink->https:%v\n", picLink)
			fmt.Printf("detailLink->%v\n", detailLink)
			fmt.Printf("title->%v\n", title)
			fmt.Printf("praise->%v\n", praiseSelection.Text())
			fmt.Println("==================================")
		})
	})

	c.Visit("https://dig.chouti.com/all/hot/24hr/1")
}
