package actions

import (
	"log"
	"tx_parser/internal/repository"
	"tx_parser/internal/service"
)

type SubscriptionsService struct {
	repository      repository.SubscriberRepository
	externalClients service.ExternalClient
	logger          *log.Logger
}

func NewSubscriptionsService(repo repository.SubscriberRepository, clients service.ExternalClient, logger *log.Logger) *SubscriptionsService {
	return &SubscriptionsService{
		repository:      repo,
		externalClients: clients,
		logger:          logger,
	}
}

func (s *SubscriptionsService) Subscribe(address string) (bool, error) {
	s.logger.Println("Adding address to the subscribers")

	currentBlock, err := s.externalClients.EthereumClient.GetCurrentBlock()
	if err != nil {
		s.logger.Printf("Unsucsessfull subscribtion: %v", err)
		return false, err
	}

	isAdded, err := s.repository.InsertNewSubscriber(address, currentBlock.Result)
	if err != nil {
		return isAdded, err
	}
	return isAdded, err
}

func (s *SubscriptionsService) Unsubscribe(address string) (bool, error) {
	s.logger.Println("Deleting user from subscriber list")
	isDeleted, err := s.repository.RemoveSubscriber(address)
	if err != nil {
		return isDeleted, err
	}
	return isDeleted, err
}
