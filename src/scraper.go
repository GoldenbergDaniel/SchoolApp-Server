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

	month.Name = doc.Find("ul").Text()

	fmt.Println(month.Name)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
