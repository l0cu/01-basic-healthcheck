package bootstrap

const (
	host = "localhost"
	port = 8080
)

func Run() error {
	httpServer := server.New(host, port)
	return httpServer.Run()
}
