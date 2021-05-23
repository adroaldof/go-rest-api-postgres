package main

import (
	"gorm.io/gorm"

	"github.com/adroaldof/go-rest-api-postgres/api/articles"
	"github.com/adroaldof/go-rest-api-postgres/database"
	"github.com/adroaldof/go-rest-api-postgres/routes"
)

func makeMigrations(db *gorm.DB) {
	db.AutoMigrate(&articles.Article{})
}

func runSeeds(db *gorm.DB) {
	articles.SeedArticles(db)
}

func main() {
	db := database.GetDBInstance()

	makeMigrations(db)
	runSeeds(db)

	routes.GetRoutes()
}
