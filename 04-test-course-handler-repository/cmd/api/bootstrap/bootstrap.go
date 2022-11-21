package bootstrap

import (
	"database/sql"
	"fmt"

	server "github.com/l0cu/codely-hex-examples/04-test-course-handler-repository/internal/platform/server"
)

const (
	host = "localhost"
	port = 8080

	dbUser = "codely"
	dbPass = "codely"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "codely"
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	courseRepository := mysql.NewCourseRepository(db)

	httpServer := server.New(host, port, courseRepository)
	return httpServer.Run()
}
