package main

import (
	"log"

	bootstrap "github.com/l0cu/codely-hex-examples/04-test-course-handler-repository/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
