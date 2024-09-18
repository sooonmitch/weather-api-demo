package main

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Application configuration
type configuration struct {
	ServerAddress string `envconfig:"SERVER_ADDRESS" required:"true" default:":8080"`
}

// loadConfig will load the configuration and exit the app if it can't.
func loadConfig() configuration {
	var appConfig configuration

	err := envconfig.Process("APP", &appConfig)
	if err != nil {
		log.Fatalf("Failed to load app config: %v", err)
	}

	return appConfig
}
