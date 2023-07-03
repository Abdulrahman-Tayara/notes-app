package main

import (
	"fmt"
	"github.com/Abdulrahman-Tayara/notes-app/notes-service/configs"
	"github.com/Abdulrahman-Tayara/notes-app/notes-service/domain"
	"github.com/Abdulrahman-Tayara/notes-app/notes-service/infrastructure/db"
	"log"
)

func init() {
	config, err := configs.LoadConfig(".", "app")

	if err != nil {
		log.Fatal("error while loading the config", err)
	}

	err = db.ConnectToDB(config.DbDSN)

	if err != nil {
		log.Fatal("error while connecting to the database", err)
	}
}

func main() {
	err := db.Instance().AutoMigrate(&domain.Note{})

	if err != nil {
		log.Fatal("[error]: ", err)
	}

	fmt.Println("Migration complete")
}
