package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", HomePage).Methods("GET")
	router.HandleFunc("/articles", CreateArticle).Methods("POST")
	router.HandleFunc("/articles", GetArticles).Methods("GET")
	router.HandleFunc("/articles/{id}", DeleteArticle).Methods("DELETE")
	router.HandleFunc("/articles/{id}", GetArticle).Methods("GET")
	router.HandleFunc("/articles/{id}", PatchArticle).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func makeMigrations(db *gorm.DB) {
	db.AutoMigrate(&Article{})
}

func runSeeds(db *gorm.DB) {
	SeedArticles(db)
}

func main() {
	db := GetDBInstance()

	makeMigrations(db)
	runSeeds(db)

	handleRequests()
}
