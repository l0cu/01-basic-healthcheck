package bootstrap

import "github.com/l0cu/codely-hex-examples/03-post-course-handler/internal/platform/server"

const (
	host = "localhost"
	port = 8080
)

func Run() error {
	httpServer := server.New(host, port)
	return httpServer.Run()
}
