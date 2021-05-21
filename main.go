package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Article struct {
	// ID  string `json:"ID"`
	Title  string `json:"Title"`
	Author string `json:"Author"`
	Link   string `json:"Link"`
}

var Articles []Article

type Home struct {
	Now string `json:"Now"`
}

func homePage(response http.ResponseWriter, request *http.Request) {
	now := Home{
		Now: time.Now().String(),
	}

	json.NewEncoder(response).Encode(now)
}

func getArticles(response http.ResponseWriter, request *http.Request) {
	Articles = []Article{
		{
			Title:  "Python Intermediate and Advanced 101",
			Author: "Arkaprabha Majumdar",
			Link:   "https://www.amazon.com/dp/B089KVK23P",
		},
		{
			Title:  "R programming Advanced",
			Author: "Arkaprabha Majumdar",
			Link:   "https://www.amazon.com/dp/B089WH12CR",
		},
		{
			Title:  "R programming Fundamentals",
			Author: "Arkaprabha Majumdar",
			Link:   "https://www.amazon.com/dp/B089S58WWG",
		},
	}

	json.NewEncoder(response).Encode(Articles)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/articles", getArticles)

	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	handleRequests()
}
