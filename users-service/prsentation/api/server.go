package api

import (
	gocontext "context"
	"errors"
	"fmt"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/configs"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/prsentation/api/context"
	"github.com/Abdulrahman-Tayara/notes-app/users-service/prsentation/api/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HTTPServer struct {
	config configs.Config

	engine     *gin.Engine
	httpServer *http.Server
}

func NewHTTPServer(config configs.Config) *HTTPServer {
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

	s.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%s", s.config.Port),
		Handler: s.engine,
	}

	return s.httpServer.ListenAndServe()
}

func (s *HTTPServer) Close(ctx gocontext.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *HTTPServer) setupRouters() {
	apiGroup := s.engine.Group("api/")

	apiGroup.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "I'm good thanks for asking",
		})
	})

	apiGroup.POST("/signup", context.GinWrapper(controllers.SignUpController))
	apiGroup.POST("/login", context.GinWrapper(controllers.LoginController))

	apiGroup.POST("/refresh", context.GinWrapper(controllers.RefreshAccessToken))
}
