package config

import (
	"io"
	"log"
	"os"
	"testing"
)

func TestLoadConfig_Success(t *testing.T) {
	logger := log.New(io.Discard, "", log.LstdFlags)

	// Create a temporary file with valid JSON content
	tempFile, err := os.CreateTemp("", "config*.json")
	if err != nil {
		t.Fatalf("Error creating temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	configData := `{"server": {"host": "localhost", "port": 8080}}`
	if _, err := tempFile.WriteString(configData); err != nil {
		t.Fatalf("Error writing to temp file: %v", err)
	}
	tempFile.Close()

	config, err := LoadConfig(tempFile.Name(), logger)
	if err != nil {
		t.Fatalf("LoadConfig() error = %v", err)
	}

	if config.Server.Host != "localhost" || config.Server.Port != 8080 {
		t.Errorf("Unexpected config values: %+v", config)
	}
}

func TestLoadConfig_FileNotFound(t *testing.T) {
	logger := log.New(io.Discard, "", log.LstdFlags)

	_, err := LoadConfig("nonexistent.json", logger)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

func TestLoadConfig_ReadError(t *testing.T) {
	logger := log.New(io.Discard, "", log.LstdFlags)

	// Create a temporary file and set its permissions to simulate a read error
	tempFile, err := os.CreateTemp("", "config*.json")
	if err != nil {
		t.Fatalf("Error creating temp file: %v", err)
	}
	tempFile.Close()
	os.Chmod(tempFile.Name(), 0000) // No read permissions
	defer os.Remove(tempFile.Name())

	_, err = LoadConfig(tempFile.Name(), logger)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

func TestLoadConfig_UnmarshalError(t *testing.T) {
	logger := log.New(io.Discard, "", log.LstdFlags)

	// Create a temporary file with invalid JSON content
	tempFile, err := os.CreateTemp("", "config*.json")
	if err != nil {
		t.Fatalf("Error creating temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	invalidConfigData := `{"server": {"host": "localhost", "port": "invalid"}}` // Port should be an integer
	if _, err := tempFile.WriteString(invalidConfigData); err != nil {
		t.Fatalf("Error writing to temp file: %v", err)
	}
	tempFile.Close()

	_, err = LoadConfig(tempFile.Name(), logger)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}
