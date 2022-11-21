package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	health "github.com/l0cu/codely-hex-examples/04-test-course-handler-repository/internal/platform/server/handler/health"
	courses "github.com/l0cu/codely-hex-examples/04-test-course-handler-repository/internal/platform\server/handler/courses"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine
}

func New(host string, port uint) Server {
	httpServer := Server{
		httpAddr: fmt.Sprintf("%s:%d", host, port),
		engine:   gin.New(),
	}

	httpServer.registerRoutes()
	return httpServer
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/courses", courses.CreateHandler())
}
