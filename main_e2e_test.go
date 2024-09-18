package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sooonmitch/weather-api-demo/middlewares"
	"github.com/sooonmitch/weather-api-demo/routes/v1/handlers"

	"github.com/gorilla/mux"
)

func TestGetWeatherE2E(t *testing.T) {
	weather := &handlers.Weather{}

	router := mux.NewRouter()

	router.Use(middlewares.JSONResponse)
	router.Use(middlewares.Logger)
	router.Use(middlewares.Metrics)

	router.HandleFunc("/weather/{latitude},{longitude}", weather.Get)

	req, err := http.NewRequest("GET", "/weather/40,174", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
