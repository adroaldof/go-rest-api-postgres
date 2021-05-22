package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Database() *gorm.DB {
	db, error := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if error != nil {
		fmt.Println(error.Error())
		panic("Failed to connect to the database")
	}

	return db
}
