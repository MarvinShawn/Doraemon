package main

import (
	"fmt"

	"github.com/uniplaces/carbon"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

//Aritical 文章
type Aritical struct {
	CreateTime int64
	Summary    string
	ImageURL   string
	AriticalID int64
	DetailURL  string
}

func main() {

	startParsing()
}

func startParsing() {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36"),
	)
	c.OnHTML("body", func(e *colly.HTMLElement) {

		artical := Aritical{}
		dateSelection := e.DOM.Find(".daily small").First()
		if len(dateSelection.Text()) <= 0 {
			fmt.Println("没有日期")
			artical.CreateTime = carbon.Now().Unix()
		} else {
			car, err := carbon.CreateFromFormat(carbon.DateFormat, dateSelection.Text(), "Local")
			if err != nil {
				fmt.Printf("格式化日期出错%v", err)
			}
			artical.CreateTime = car.Unix()
		}
		tmpArr := make([]Aritical, 0)
		e.DOM.Find(".content a").Each(func(i int, selection *goquery.Selection) {
			title := selection.Text()
			articalLink, ok := selection.Attr("href")
			if len(title) > 0 && ok {
				artical.DetailURL = fmt.Sprintf("https://toutiao.io%v", articalLink)
				artical.Summary = title
				tmpArr = append(tmpArr, artical)
			}
		})

		for _, atl := range tmpArr {
			fmt.Printf("%+v\n", atl)
		}
	})
	c.OnError(func(res *colly.Response, err error) {
		fmt.Print(err)
	})
	c.Visit("https://toutiao.io/")
}
