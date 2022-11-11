package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const httpAddr = ":8080"

func main() {
	fmt.Println("Server running on", httpAddr)

	server := gin.New()
	server.GET("/health", healthHandler)

	log.Fatal(server.Run(httpAddr))
}

func healthHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "everything is ok!")
}
