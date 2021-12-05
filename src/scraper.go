package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const url string = "https://www.cavsconnect.com/calendar/"
const cid string = "mc-eebc9a21598585f5d8bea2ab08144d58"

func getData(month_ string, year_ string) {
	res, err := http.Get(url + "?cid=" + cid + "&format=list" + "&month=" + string(month_) + "&yr=" + string(year_))
	checkError(err)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkError(err)

	// This is where you scrape
	//month.Name = doc.Find("ul#list-mc-eebc9a21598585f5d8bea2ab08144d58").Text()

	doc.Find("ul#list-mc-eebc9a21598585f5d8bea2ab08144d58").Each(func(i int, ul *goquery.Selection) {
		ul.Find("li").Each(func(i int, li *goquery.Selection) {
			id, ok := li.Attr("id")
			if ok {
				println(id)
			}
			li.Find("div").Each(func(i int, div *goquery.Selection) {
				var header = (div.Find("div").Find("h3"))
				alt, ok := header.Find("img").Attr("alt")
				if ok {
					println(alt)
				}
				println(header.Contents().Text())
			})
		})
	})

	fmt.Println(month.Name)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
