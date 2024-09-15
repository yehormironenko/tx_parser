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
		mr.logger.Println("address is already in the list of subscribers")
		return false, nil
	}

	mr.subscribers[address] = blockNumber
	mr.logger.Printf("address: %v added to the list of subscribers with starting block number: %v", address, blockNumber)
	return true, nil
}

func (mr *MemoryRepo) RemoveSubscriber(address string) (bool, error) {
	mr.mu.Lock()
	defer mr.mu.Unlock()

	if _, exists := mr.subscribers[address]; !exists {
		mr.logger.Println("address is not in the list of subscribers")
		return false, nil
	}

	delete(mr.subscribers, address)
	mr.logger.Printf("address: %v removed from the list of subsribers", address)
	return true, nil
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
	log.Printf("Latest block has been updated for address: %v; new value is: %v", address, block)
	return true, nil
}
