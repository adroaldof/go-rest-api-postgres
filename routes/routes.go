package routes

import (
	"log"
	"net/http"

	"github.com/adroaldof/go-rest-api-postgres/api/articles"
	"github.com/adroaldof/go-rest-api-postgres/api/home"
	"github.com/gorilla/mux"
)

func GetRoutes() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", home.HomePage).Methods("GET")

	router.HandleFunc("/articles", articles.CreateArticle).Methods("POST")
	router.HandleFunc("/articles", articles.GetArticles).Methods("GET")
	router.HandleFunc("/articles/{id}", articles.DeleteArticle).Methods("DELETE")
	router.HandleFunc("/articles/{id}", articles.GetArticle).Methods("GET")
	router.HandleFunc("/articles/{id}", articles.PatchArticle).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":8000", router))
}
