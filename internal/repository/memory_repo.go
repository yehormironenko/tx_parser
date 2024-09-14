package repository

import (
	"log"
	"sync"
)

type MemoryRepo struct {
	mu          sync.Mutex
	subscribers map[string]struct{}
	logger      *log.Logger
}

func NewMemoryRepo(subscribersMap map[string]struct{}, logger *log.Logger) *MemoryRepo {
	return &MemoryRepo{
		subscribers: subscribersMap,
		logger:      logger,
	}
}

func (mr *MemoryRepo) InsertNewSubscriber(address string) (bool, error) {
	mr.mu.Lock()
	defer mr.mu.Unlock()

	if _, exists := mr.subscribers[address]; exists {
		mr.logger.Println("address is already subscribed")
		return false, nil
	}

	mr.subscribers[address] = struct{}{}
	mr.logger.Printf("address subscribed %v", address)
	return true, nil
}

func (mr *MemoryRepo) RemoveSubscriber(address string) (bool, error) {
	mr.mu.Lock()
	defer mr.mu.Unlock()

	if _, exists := mr.subscribers[address]; !exists {
		mr.logger.Println("address is not subscribed")
		return false, nil
	}

	delete(mr.subscribers, address)
	mr.logger.Printf("address unsubscribed %v", address)
	return true, nil
}
