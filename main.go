package main

import (
	"log"
	"net/http"

	"github.com/sooonmitch/weather-api-demo/routes"
)

func main() {
	log.Println("Starting server...")

	appConfig := loadConfig()

	router := routes.SetupRouter()

	log.Printf("Server is running on %s", appConfig.ServerAddress)

	log.Fatal(http.ListenAndServe(appConfig.ServerAddress, router))
}
