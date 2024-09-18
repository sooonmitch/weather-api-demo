package v1

import (
	"fmt"

	"github.com/sooonmitch/weather-api-demo/middlewares"
	"github.com/sooonmitch/weather-api-demo/routes/v1/handlers"

	"github.com/gorilla/mux"
)

const (
	pathWeather string = "/weather"
)

func RegisterRoutes(router *mux.Router) {
	// Middlewares
	router.Use(middlewares.JSONResponse)
	router.Use(middlewares.Logger)
	router.Use(middlewares.Metrics)

	// Handlers
	weather := &handlers.Weather{}

	router.HandleFunc(fmt.Sprintf("%s/{latitude},{longitude}", pathWeather), weather.Get).Methods("GET")
	router.HandleFunc(pathWeather, weather.Put).Methods("PUT")
	router.HandleFunc(pathWeather, weather.Update).Methods("POST")
	router.HandleFunc(pathWeather, weather.Delete).Methods("DELETE")
}
