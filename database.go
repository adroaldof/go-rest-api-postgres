package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	DBUser     = "user"
	DBPassword = "password"
	DBName     = "go_api_db"
	DBHost     = "0.0.0.0"
	DBPort     = "5432"
	DBType     = "postgres"
)

func getPostgresConnectionString() string {
	dataBase := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		DBHost,
		DBPort,
		DBUser,
		DBName,
		DBPassword)

	return dataBase
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
