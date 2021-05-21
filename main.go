package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Article struct {
	Title  string `json:"Title"`
	Author string `json:"Author"`
	Link   string `json:"Link"`
}

var Articles []Article

type Home struct {
	Now string `json:"Now"`
}

func homePage(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Endpoint home page")

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

	fmt.Println("Endpoint get articles")
	json.NewEncoder(response).Encode(Articles)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", getArticles)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	handleRequests()
}
