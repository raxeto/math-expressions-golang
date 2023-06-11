package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Config struct {
	ServiceAddress string `json:"service_address"`
}

func LoadConfig() Config {
	configFile, err := os.Open("config.json")
	if err != nil {
		log.Fatalf("Error opening config file: %s", err)
	}
	defer configFile.Close()

	byteValue, readErr := io.ReadAll(configFile)
	if readErr != nil {
		log.Fatalf("Error reading config file: %s", readErr)
	}

	var config Config
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Fatalf("Error decoding config file: %s", err)
	}

	return config
}
