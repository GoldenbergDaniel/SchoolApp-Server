package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const url string = "https://www.cavsconnect.com/calendar/"
const cid string = "mc-eebc9a21598585f5d8bea2ab08144d58"

func getData(month_ string, year_ string, month *Month) {
	res, err := http.Get(url + "?cid=" + cid + "&format=list" + "&month=" + string(month_) + "&yr=" + string(year_))
	checkError(err)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkError(err)

	doc.Find("ul#list-mc-eebc9a21598585f5d8bea2ab08144d58").Each(func(_ int, ul *goquery.Selection) {
		ul.Find("li").Each(func(i int, li *goquery.Selection) {
			var day = new(Day)
			fmt.Println(" ")
			id, ok := li.Attr("id")
			if ok {
				date := string(year_) + "-" + string(month_) + "-" + id[len(id)-2:]
				day.Day = date
				fmt.Println(date)
			}

			li.Find("div").Each(func(a int, div *goquery.Selection) {

				var event = new(Event)

				var header = div.Find("div").Find("h3")
				alt, ok := header.Find("img").Attr("alt")

				// fmt.Println(a)

				if header.Contents().Text() != "" && header.Contents().Text() != "Even" && header.Contents().Text() != "Odd" {
					if ok {
						event.Category = alt
						fmt.Println(alt)

						event.Name = header.Contents().Text()
						fmt.Println(header.Contents().Text())
					}
					day.Events = append(day.Events, *event)
				}
			})
			month.Days = append(month.Days, *day)
		})
	})
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
