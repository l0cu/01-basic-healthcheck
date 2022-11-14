package main

import (
	"log"

	"github.com/l0cu/codely-hex-examples/02-01-architectured-gin-health/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
