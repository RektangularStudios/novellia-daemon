package main

import (
	"fmt"
	"net/http"

	novellia_api "github.com/RektangularStudios/novellia/generated/novellia-api"
)

const (
	server_port = 8080
)

func main() {
	fmt.Printf("Starting server on port %d", server_port)

	DefaultApiService := novellia_api.NewDefaultApiService()
	DefaultApiController := novellia_api.NewDefaultApiController(DefaultApiService)

	router := novellia_api.NewRouter(DefaultApiController)

	err := http.ListenAndServe(fmt.Sprintf(":%d", server_port), router)
	if err != nil {
		fmt.Printf("Error starting server")
	}
}
