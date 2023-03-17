package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func ConnectToDB(dsn string) (err error) {
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	fmt.Println("Connected successfully to the database")

	return
}

func DBInstance() *gorm.DB {
	return db
}
