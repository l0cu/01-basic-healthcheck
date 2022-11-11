package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine
}

func New(host string, port uint) Server {
	httpServer := Server{
		httpAddr: fmt.Sprintf("#{host}:#{port}"),
		engine:   gin.New(),
	}

	httpServer.registerRoutes()
	return httpServer
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
}
