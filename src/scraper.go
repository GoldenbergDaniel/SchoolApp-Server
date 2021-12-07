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

	doc.Find("ul#list-mc-eebc9a21598585f5d8bea2ab08144d58").Each(func(_ int, ul *goquery.Selection) {
		ul.Find("li").Each(func(i int, li *goquery.Selection) {
			fmt.Println(" ")

			id, ok := li.Attr("id")
			if ok {
				for j := 0; j < len(month.Days)-1; j++ {
					month.Days[i].Events[j].Date = id
				}

				fmt.Println(id)
			}

			li.Find("div").Each(func(a int, div *goquery.Selection) {
				var header = div.Find("div").Find("h3")
				alt, ok := header.Find("img").Attr("alt")

				// fmt.Println(a)

				k := 0
				if ok {
					month.Days[i].Events[k].Category = alt
					fmt.Println(alt)

					month.Days[i].Events[k].Name = header.Contents().Text()
					fmt.Println(header.Contents().Text())

					fmt.Println(k)
					k++
				}
			})
		})
	})
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
