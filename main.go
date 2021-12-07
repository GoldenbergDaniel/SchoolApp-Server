package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Month represents a month
type Month struct {
	Days []Day `json:"days"`
}

// Day
type Day struct {
	Day    int     `json:"day"`
	Events []Event `json:"events"`
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
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server started...")
	fmt.Println(" * Running on http://localhost:" + port)
	//fmt.Println(" * IP: localhost")
	fmt.Println(" * Port: " + port)

	r := mux.NewRouter()
	r.HandleFunc("/", fetchMonthResponse).Queries("month", "{month}", "year", "{year}").Methods("GET")

	http.ListenAndServe(":"+port, r)
}

func fetchMonthResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	getData(params["month"], params["year"], month)

	json.NewEncoder(w).Encode(month)
	fmt.Println("GET recieved from path " + r.URL.Path)
}
