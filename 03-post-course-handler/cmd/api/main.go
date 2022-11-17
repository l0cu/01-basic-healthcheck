package main

import (
	"log"

	"github.com/l0cu/codely-hex-examples/03-post-course-handler/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
