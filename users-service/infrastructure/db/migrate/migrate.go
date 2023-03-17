package main

import (
	"fmt"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/infrastructure"
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
	err := initializers.DB.AutoMigrate(&infrastructure.User{})

	if err != nil {
		log.Fatal("[error]: ", err)
	}

	fmt.Println("Migration complete")
}
