package config

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	Name         string   `json:"name"`
	IP           string   `json:"ip"`
	Port         string   `json:"port"`
	Url          string   `json:"url"`
	AllowOrigins []string `json:"allowOrigins"`
}

func ConfigFile() (Config, error) {
	// Open the JSON file
	file, err := os.Open("config.json")
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	// Read the file contents into a byte slice
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return Config{}, err
	}

	// Unmarshal the JSON data into a Person struct
	var configFile Config
	err = json.Unmarshal(fileBytes, &configFile)
	if err != nil {
		return Config{}, err
	}

	return configFile, nil
}
