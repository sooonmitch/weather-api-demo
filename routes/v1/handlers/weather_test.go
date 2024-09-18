package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetWeather(t *testing.T) {
	weather := &Weather{}

	req, err := http.NewRequest("GET", "/weather/40,74", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/weather/{latitude},{longitude}", weather.Get)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response Weather
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("could not decode response: %v", err)
	}

	if response.Latitude != 40 || response.Longitude != 74 {
		t.Errorf("handler returned unexpected lat/lon: got %v/%v want %v/%v",
			response.Latitude, response.Longitude, 40, 74)
	}

	if len(response.Forecasts) == 0 {
		t.Errorf("expected forecasts, got none")
	}

	temp := response.Forecasts[0].Temperature
	if temp < -20 || temp > 100 {
		t.Errorf("temperature out of expected range: got %v", temp)
	}
}
