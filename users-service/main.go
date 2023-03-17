package main

import (
	"github.com/Abdulrahman-Tayara/notes-app/users-service/initializers"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/prsentation/api"
	"log"
)

func loadConfig() (initializers.Config, error) {
	return initializers.LoadConfig(".", "app")
}

func init() {
	config, err := loadConfig()

	if err != nil {
		log.Fatal("error while loading the config: ", err)
	}

	err = initializers.ConnectToDB(&config)

	if err != nil {
		log.Fatal("error while connecting to the database: ", err)
	}
}

func main() {
	config, _ := loadConfig()

	server := api.NewHTTPServer(config)

	err := server.Run()

	if err != nil {
		log.Fatalf("error while starting the server: %v", err)
	}

}
