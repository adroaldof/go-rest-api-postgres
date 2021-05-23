package articles

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/adroaldof/go-rest-api-postgres/database"
)

type Article struct {
	ID     string `json:"ID"`
	Title  string `json:"Title"`
	Author string `json:"Author"`
	Link   string `json:"Link"`
}

var Articles []Article

func GetArticles(response http.ResponseWriter, request *http.Request) {
	var articles []Article

	db := database.GetDBInstance()
	db.Find(&articles)

	json.NewEncoder(response).Encode(articles)
}

func GetArticle(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	var article Article

	db := database.GetDBInstance()
	db.Find(&article, id)

	json.NewEncoder(response).Encode(article)
}

func CreateArticle(response http.ResponseWriter, request *http.Request) {
	body, _ := ioutil.ReadAll(request.Body)

	var article Article

	json.Unmarshal(body, &article)

	db := database.GetDBInstance()
	db.Create(article)

	json.NewEncoder(response).Encode(article)
}

func PatchArticle(response http.ResponseWriter, request *http.Request) {
	body, _ := ioutil.ReadAll(request.Body)
	vars := mux.Vars(request)
	id := vars["id"]

	var payload Article

	json.Unmarshal(body, &payload)

	var article Article

	db := database.GetDBInstance()
	db.Find(&article, id)

	article.Title = payload.Title
	article.Author = payload.Author
	article.Link = payload.Link

	db.Save(&article)

	json.NewEncoder(response).Encode(article)
}

func DeleteArticle(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	var article Article

	db := database.GetDBInstance()
	db.Find(&article, id)
	db.Delete(&article)

	json.NewEncoder(response)
}

func SeedArticles(db *gorm.DB) {
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

	db.CreateInBatches(&Articles, 3)
}
