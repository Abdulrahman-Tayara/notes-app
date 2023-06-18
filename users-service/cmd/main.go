package main

import (
	"context"
	"fmt"
	"github.com/Abdulrahman-Tayara/notes-app/pkg/http"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/api"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/configs"
	grpc2 "github.com/Abdulrahman-Tayara/notes-app/users-service/grpc"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/initializers"
	"google.golang.org/grpc"
	"log"
	"net"
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

func httpServerSetup(config *configs.Config) *http.Server {
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

	return server
}

func grpcServerSetup(config *configs.Config) *grpc.Server {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	grpc2.RegisterAuthenticationServer(s)

	log.Printf("grpc server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return s
}

func main() {
	config, _ := loadConfig()
	configs.AppConfig = &config

	httpServer := httpServerSetup(&config)
	grpcServer := grpcServerSetup(&config)

	handleCloseSignals(func(ctx context.Context) {
		grpcServer.Stop()

		if err := httpServer.Close(ctx); err != nil {
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
