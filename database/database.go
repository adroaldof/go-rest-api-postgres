package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/adroaldof/go-rest-api-postgres/handlers"
)

func getPostgresConnectionString() string {
	databaseUrlPattern := "host=%s port=%s user=%s dbname=%s password=%s sslmode=disable"

	databaseUrl := fmt.Sprintf(databaseUrlPattern,
		handlers.GetEnvironmentVariable("POSTGRES_HOST", "0.0.0.0"),
		handlers.GetEnvironmentVariable("POSTGRES_PORT", "5432"),
		handlers.GetEnvironmentVariable("POSTGRES_USER", "user"),
		handlers.GetEnvironmentVariable("POSTGRES_NAME", "go_api_db"),
		handlers.GetEnvironmentVariable("POSTGRES_PASSWORD", "password"))

	return databaseUrl
}

func GetDBInstance() *gorm.DB {
	connectionString := getPostgresConnectionString()

	db, error := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connectionString,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if error != nil {
		fmt.Println(error.Error())
		panic("Failed to connect to the database")
	}

	return db
}
