package main

import (
	"context"
	"github.com/Abdulrahman-Tayara/notes-app/shared/http"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/configs"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/initializers"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/prsentation/api"
	"log"
	nethttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func loadConfig() (configs.Config, error) {
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
	configs.AppConfig = &config

	server := http.NewHTTPServer(http.Config{
		Port:    config.Port,
		GinMode: config.GinMode,
	})

	go func() {
		err := server.Run(api.SetupRouters)

		if err != nil && err != nethttp.ErrServerClosed {
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
