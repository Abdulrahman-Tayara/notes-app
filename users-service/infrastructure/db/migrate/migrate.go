package main

import (
	"fmt"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/configs"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/auth"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/domain/entity"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/infrastructure/db"
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
	err := db.Instance().AutoMigrate(&entity.User{}, &auth.RefreshToken{})

	if err != nil {
		log.Fatal("[error]: ", err)
	}

	fmt.Println("Migration complete")
}
