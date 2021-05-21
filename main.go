package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Article struct {
	ID     string `json:"ID"`
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
	json.NewEncoder(response).Encode(Articles)
}

func getArticle(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	for _, article := range Articles {
		if article.ID == id {
			json.NewEncoder(response).Encode(article)
		}
	}

}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/articles", getArticles)
	router.HandleFunc("/articles/{id}", getArticle)

	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	Articles = []Article{
		{
			ID:     "1",
			Title:  "Python Intermediate and Advanced 101",
			Author: "Arkaprabha Majumdar",
			Link:   "https://www.amazon.com/dp/B089KVK23P",
		},
		{
			ID:     "2",
			Title:  "R programming Advanced",
			Author: "Arkaprabha Majumdar",
			Link:   "https://www.amazon.com/dp/B089WH12CR",
		},
		{
			ID:     "3",
			Title:  "R programming Fundamentals",
			Author: "Arkaprabha Majumdar",
			Link:   "https://www.amazon.com/dp/B089S58WWG",
		},
	}

	handleRequests()
}
