package actions

import (
	"log"

	"github.com/yehormironenko/tx_parser/internal/repository"
	"github.com/yehormironenko/tx_parser/internal/service"
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
	s.logger.Printf("Adding address %v to the subscribers", address)

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
	s.logger.Printf("Deleting address %v from subscriber list", address)
	isDeleted, err := s.repository.RemoveSubscriber(address)
	if err != nil {
		return isDeleted, err
	}
	return isDeleted, nil
}
