package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func getPostgresConnectionString() string {
	databaseUrlPattern := "host=%s port=%s user=%s dbname=%s password=%s sslmode=disable"

	databaseUrl := fmt.Sprintf(databaseUrlPattern,
		GetEnvironmentVariable("POSTGRES_HOST", "0.0.0.0"),
		GetEnvironmentVariable("POSTGRES_PORT", "5432"),
		GetEnvironmentVariable("POSTGRES_USER", "user"),
		GetEnvironmentVariable("POSTGRES_NAME", "go_api_db"),
		GetEnvironmentVariable("POSTGRES_PASSWORD", "password"))

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
