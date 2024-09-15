package repository

import (
	"log"
	"sync"
)

type MemoryRepo struct {
	mu          sync.Mutex
	subscribers map[string]string
	logger      *log.Logger
}

func NewMemoryRepo(subscribersMap map[string]string, logger *log.Logger) *MemoryRepo {
	return &MemoryRepo{
		subscribers: subscribersMap,
		logger:      logger,
	}
}

func (mr *MemoryRepo) InsertNewSubscriber(address, blockNumber string) (bool, error) {
	mr.mu.Lock()
	defer mr.mu.Unlock()

	if _, exists := mr.subscribers[address]; exists {
		mr.logger.Println("address is already subscribed")
		return false, nil
	}

	mr.subscribers[address] = blockNumber
	mr.logger.Printf("address added to subscribers %v with starting block number: %v", address, blockNumber)
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

func (mr *MemoryRepo) IsSubscribed(address string) (bool, error) {
	mr.mu.Lock()
	defer mr.mu.Unlock()

	if _, exists := mr.subscribers[address]; exists {
		mr.logger.Println("address is already subscribed")
		return true, nil
	}

	return false, nil
}

func (mr *MemoryRepo) GetSubscribers() map[string]string {
	mr.mu.Lock()
	defer mr.mu.Unlock()

	return mr.subscribers
}

func (mr *MemoryRepo) UpdateValue(address, block string) (bool, error) {
	mr.mu.Lock()
	defer mr.mu.Unlock()
	mr.subscribers[address] = block
	return true, nil
}
