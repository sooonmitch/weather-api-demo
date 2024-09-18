package routes

import (
	"github.com/gorilla/mux"
	v1 "github.com/sooonmitch/weather-api-demo/routes/v1"
)

const (
	pathAPIV1 string = "/api/v1"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	api := router.PathPrefix(pathAPIV1).Subrouter()

	v1.RegisterRoutes(api)

	return router
}
