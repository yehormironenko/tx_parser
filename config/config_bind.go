package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Config struct {
	Server Server
	Client Client
}

func LoadConfig(filePath string, logger *log.Logger) (*Config, error) {
	logger.Println("Opening config file:", filePath)

	configFile, err := os.Open(filePath)
	if err != nil {
		logger.Printf("Error opening config file: %v\n", err)
		return nil, fmt.Errorf("error opening config file: %v", err)
	}
	defer configFile.Close()

	logger.Println("Config file opened successfully")

	byteValue, err := io.ReadAll(configFile)
	if err != nil {
		logger.Printf("Error reading config file: %v\n", err)
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	logger.Println("Config file read successfully")

	var config Config
	if err = json.Unmarshal(byteValue, &config); err != nil {
		logger.Printf("Error unmarshaling config file: %v\n", err)
		return nil, fmt.Errorf("error unmarshaling config file: %v", err)
	}

	logger.Println("Config file parsed successfully")

	return &config, nil
}
