package repository

import (
	"io"
	"log"
	"testing"
)

func TestInsertNewSubscriber_Success(t *testing.T) {
	logger := log.New(io.Discard, "", log.LstdFlags)
	repo := NewMemoryRepo(make(map[string]string), logger)

	success, err := repo.InsertNewSubscriber("0x123", "100")
	if err != nil {
		t.Fatalf("InsertNewSubscriber() error = %v", err)
	}
	if !success {
		t.Fatal("Expected InsertNewSubscriber to return true, got false")
	}
	if blockNumber := repo.GetSubscribers()["0x123"]; blockNumber != "100" {
		t.Errorf("Expected blockNumber to be '100', got '%v'", blockNumber)
	}
}

func TestInsertNewSubscriber_Duplicate(t *testing.T) {
	logger := log.New(io.Discard, "", log.LstdFlags)
	repo := NewMemoryRepo(map[string]string{"0x123": "100"}, logger)

	success, err := repo.InsertNewSubscriber("0x123", "101")
	if err != nil {
		t.Fatalf("InsertNewSubscriber() error = %v", err)
	}
	if success {
		t.Fatal("Expected InsertNewSubscriber to return false, got true")
	}
}

func TestRemoveSubscriber_Success(t *testing.T) {
	logger := log.New(io.Discard, "", log.LstdFlags)
	repo := NewMemoryRepo(map[string]string{"0x123": "100"}, logger)

	success, err := repo.RemoveSubscriber("0x123")
	if err != nil {
		t.Fatalf("RemoveSubscriber() error = %v", err)
	}
	if !success {
		t.Fatal("Expected RemoveSubscriber to return true, got false")
	}
	if _, exists := repo.GetSubscribers()["0x123"]; exists {
		t.Fatal("Expected subscriber to be removed")
	}
}

func TestRemoveSubscriber_NotFound(t *testing.T) {
	logger := log.New(io.Discard, "", log.LstdFlags)
	repo := NewMemoryRepo(make(map[string]string), logger)

	success, err := repo.RemoveSubscriber("0x123")
	if err != nil {
		t.Fatalf("RemoveSubscriber() error = %v", err)
	}
	if success {
		t.Fatal("Expected RemoveSubscriber to return false, got true")
	}
}

func TestUpdateValue_Success(t *testing.T) {
	logger := log.New(io.Discard, "", log.LstdFlags)
	repo := NewMemoryRepo(map[string]string{"0x123": "100"}, logger)

	success, err := repo.UpdateValue("0x123", "200")
	if err != nil {
		t.Fatalf("UpdateValue() error = %v", err)
	}
	if !success {
		t.Fatal("Expected UpdateValue to return true, got false")
	}
	if blockNumber := repo.GetSubscribers()["0x123"]; blockNumber != "200" {
		t.Errorf("Expected blockNumber to be '200', got '%v'", blockNumber)
	}
}

func TestGetSubscribers(t *testing.T) {
	logger := log.New(io.Discard, "", log.LstdFlags)
	subscribersMap := map[string]string{"0x123": "100", "0x456": "200"}
	repo := NewMemoryRepo(subscribersMap, logger)

	subscribers := repo.GetSubscribers()
	if len(subscribers) != 2 {
		t.Fatalf("Expected 2 subscribers, got %v", len(subscribers))
	}
	if subscribers["0x123"] != "100" || subscribers["0x456"] != "200" {
		t.Errorf("Unexpected subscribers: %v", subscribers)
	}
}
