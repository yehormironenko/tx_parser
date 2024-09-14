package actions

import (
	"log"
	"tx_parser/internal/model"
	"tx_parser/internal/repository"
)

type NotificationServiceImpl struct {
	repository repository.SubscribeRepo
	logger     *log.Logger
}

func NewNotificationService(repo repository.SubscribeRepo, logger *log.Logger) *NotificationServiceImpl {
	return &NotificationServiceImpl{
		repository: repo,
		logger:     logger,
	}
}

func (s *NotificationServiceImpl) Subscribe(address string) (bool, error) {
	s.logger.Println("Adding address to the subscribers")
	isAdded, err := s.repository.InsertNewSubscriber(address)
	if err != nil {
		return isAdded, err
	}
	return isAdded, err
}

func (s *NotificationServiceImpl) Unsubscribe(address string) (bool, error) {
	s.logger.Println("Deleting user from subscriber list")
	isDeleted, err := s.repository.RemoveSubscriber(address)
	if err != nil {
		return isDeleted, err
	}
	return isDeleted, err
}

func (s *NotificationServiceImpl) Notify(address string, transaction model.Transaction) {

}
