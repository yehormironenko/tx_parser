package actions

import (
	"log"
	"tx_parser/internal/model"
	"tx_parser/internal/repository"
)

type SubscriptionsService struct {
	repository repository.SubscriberRepository
	logger     *log.Logger
}

func NewNotificationService(repo repository.SubscriberRepository, logger *log.Logger) *SubscriptionsService {
	return &SubscriptionsService{
		repository: repo,
		logger:     logger,
	}
}

func (s *SubscriptionsService) Subscribe(address string) (bool, error) {
	s.logger.Println("Adding address to the subscribers")
	isAdded, err := s.repository.InsertNewSubscriber(address)
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

func (s *SubscriptionsService) Notify(address string, transaction model.Transaction) {

}
