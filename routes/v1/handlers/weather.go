package handlers

import (
	"encoding/json"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var ForecastFriendlyName = map[int]string{
	0:  "Sunny",
	1:  "Partly Cloudy",
	2:  "Cloudy",
	3:  "Overcast",
	4:  "Rainy",
	5:  "Stormy",
	6:  "Snowy",
	7:  "Windy",
	8:  "Foggy",
	10: "Thunderstorms",
	12: "Showers",
}

// Weather reports and history of a given location
type Weather struct {
	Latitude  int        `json:"latitude"`  // What latitude the report comes from
	Longitude int        `json:"longitude"` // What longitude the report comes from
	Forecasts []Forecast `json:"forecasts"` // History of forecasts
	Hazards   []Hazard   `json:"hazards"`   // Any active weather alerts
}

// Hazard/Alert/Special Consideration in the weather that has been reported
type Hazard struct {
	Code         int    // Standardized code of the event
	FriendlyName string // User friendly name of the event
	Description  string // Description of the event
	ReportTime   int    // Unix time of when the hazard was made
	ExpireTime   int    // Unix time of when the hazard should expire
}

// The Forecast report
type Forecast struct {
	Code                    int     `json:"code"`                    // Standardized forecast code
	FriendlyName            string  `json:"friendly_name"`           // User friendly name of the forecast
	Description             string  `json:"description"`             // Description of the forecast
	Temperature             float32 `json:"temperature"`             // Floating value of temperature
	TemperatureUnit         string  `json:"temperature_unit"`        // Temperature unit, "C", "F", "K"
	FriendlyTemperatureName string  `json:"friendly_tempature_name"` // User friendly temperature name "hot", "cold, "moderate"
	ReportTime              int     `json:"report_time"`             // Unix time of when the report was made
}

func (weather *Weather) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	latStr := vars["latitude"]
	lonStr := vars["longitude"]

	lat, err := strconv.Atoi(latStr)
	if err != nil {
		http.Error(w, "Invalid latitude", http.StatusBadRequest)
		return
	}

	lon, err := strconv.Atoi(lonStr)
	if err != nil {
		http.Error(w, "Invalid longitude", http.StatusBadRequest)
		return
	}

	// Generate random temperature in Fahrenheit
	ra := rand.New(rand.NewSource(time.Now().UnixNano()))
	temp := ra.Float32()*120 - 20 // Range from -20F to 100F

	// Clean up and round the temp to tenth
	temp = float32(math.Round(float64(temp)*10) / 10)

	var friendlyTempName string

	// Get the friendly temp name
	switch {
	case temp < 50:
		friendlyTempName = "cold"
	case temp < 70:
		friendlyTempName = "moderate"
	default:
		friendlyTempName = "warm"
	}

	// Dummy data for the response
	response := Weather{
		Latitude:  lat,
		Longitude: lon,
		Forecasts: []Forecast{
			{
				Code:                    0,
				FriendlyName:            ForecastFriendlyName[0],
				Description:             "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
				Temperature:             temp,
				TemperatureUnit:         "F",
				FriendlyTemperatureName: friendlyTempName,
				ReportTime:              int(time.Now().Unix()),
			},
		},
	}

	json.NewEncoder(w).Encode(response)
}

func (weather *Weather) Put(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(weather) // Out of scope for project, should return what was modified with the correct status.
}

func (weather *Weather) Update(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(weather) // Out of scope for project, should return what was modified with the correct status.
}

func (weather *Weather) Delete(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(weather) // Out of scope for project, should return what was modified with the correct status.
}
