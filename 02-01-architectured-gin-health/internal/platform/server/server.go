package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/l0cu/codely-hex-examples/02-01-architectured-gin-health/internal/platform/server/handler/health"
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
