package api

import (
	"errors"
	"fmt"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/initializers"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/prsentation/api/context"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/prsentation/api/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HTTPServer struct {
	config initializers.Config

	engine *gin.Engine
}

func NewHTTPServer(config initializers.Config) *HTTPServer {
	return &HTTPServer{
		config: config,
	}
}

func (s *HTTPServer) Run() error {
	if s.config.Port == "" {
		return errors.New("port value is empty")
	}

	gin.SetMode(s.config.GinMode)
	s.engine = gin.Default()

	s.setupRouters()

	return s.engine.Run(fmt.Sprintf(":%s", s.config.Port))
}

func (s *HTTPServer) setupRouters() {
	apiGroup := s.engine.Group("api/")

	apiGroup.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "I'm good thanks for asking",
		})
	})

	apiGroup.POST("/signup", context.GinWrapper(controllers.SignUpController))
}
