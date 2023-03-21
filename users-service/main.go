package main

import (
	"context"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/initializers"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/prsentation/api"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	go func() {
		err := server.Run()

		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("error while starting the server: %v", err)
		}
	}()

	handleCloseSignals(func(ctx context.Context) {
		if err := server.Close(ctx); err != nil {
			log.Fatal("Server forced to shutdown: ", err)
		}
	})
}

func handleCloseSignals(onClose func(ctx context.Context)) {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down the application...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	onClose(ctx)

	log.Println("Application exiting")
}
