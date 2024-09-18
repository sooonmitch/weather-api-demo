package v1

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/sooonmitch/weather-api-demo/middlewares"
	"github.com/sooonmitch/weather-api-demo/routes/v1/handlers"
)

func TestRouter(t *testing.T) {
	weather := &handlers.Weather{}

	router := mux.NewRouter()
	router.Use(middlewares.JSONResponse)
	router.Use(middlewares.Logger)
	router.Use(middlewares.Metrics)

	router.HandleFunc("/weather/{latitude},{longitude}", weather.Get).Methods("GET")

	tests := []struct {
		method     string
		target     string
		statusCode int
	}{
		{"GET", "/weather/40,74", http.StatusOK},
		{"GET", "/weather/e,74", http.StatusBadRequest},
		{"POST", "/weather/40,74", http.StatusMethodNotAllowed},
	}

	for _, tt := range tests {
		req, err := http.NewRequest(tt.method, tt.target, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != tt.statusCode {
			t.Errorf("handler returned wrong status code for %s %s: got %v want %v",
				tt.method, tt.target, status, tt.statusCode)
		}
	}
}
