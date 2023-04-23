package http

import (
	gocontext "context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Config struct {
	Port    string
	GinMode string
}

type RoutingSetup func(e *gin.Engine)

type Server struct {
	config Config

	engine     *gin.Engine
	httpServer *http.Server
}

func NewHTTPServer(config Config) *Server {
	return &Server{
		config: config,
	}
}

func (s *Server) Run(routing RoutingSetup) error {
	if s.config.Port == "" {
		return errors.New("port value is empty")
	}

	gin.SetMode(s.config.GinMode)
	s.engine = gin.Default()

	routing(s.engine)

	s.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%s", s.config.Port),
		Handler: s.engine,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Close(ctx gocontext.Context) error {
	return s.httpServer.Shutdown(ctx)
}
