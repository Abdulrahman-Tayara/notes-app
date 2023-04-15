package main

import (
	"fmt"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/application/auth"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/core/domain/entity"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/infrastructure/db"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/initializers"
	"log"
)

func init() {
	config, err := initializers.LoadConfig(".", "app")

	if err != nil {
		log.Fatal("error while loading the config", err)
	}

	err = initializers.ConnectToDB(&config)

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
