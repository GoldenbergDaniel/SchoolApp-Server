package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Month represents a month
type Month struct {
	Days [31]Day `json:"days"`
}

// Day
type Day struct {
	Day    int       `json:"day"`
	Events [30]Event `json:"events"`
}

// Event represents an event in a month
type Event struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Date        string `json:"date"`
}

var month *Month = new(Month)

func main() {
	fmt.Println("Server started...")
	fmt.Println(" * Running on http://127.0.0.1:8080/")
	fmt.Println(" * IP: localhost")
	fmt.Println(" * Port: 8080")

	for i := 0; i < len(month.Days); i++ {
		month.Days[i].Day = i + 1
	}

	r := mux.NewRouter()
	r.HandleFunc("/", fetchMonthResponse).Queries("month", "{month}", "year", "{year}").Methods("GET")

	http.ListenAndServe(":8080", r)
}

func fetchMonthResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	getData(params["month"], params["year"])

	json.NewEncoder(w).Encode(month)
	fmt.Println("GET recieved from path " + r.URL.Path)
}
