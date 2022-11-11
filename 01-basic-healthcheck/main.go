package main

import (
	"fmt"
	"log"
	http "net/http"
)

const httpAddr = ":8080"

func main() {
	fmt.Println("Server running on", httpAddr)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)

	log.Fatal(http.ListenAndServe(httpAddr, mux))
}

func healthHandler(responseWriter http.ResponseWriter, _ *http.Request) {
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write([]byte("everything is ok!"))
}
