package actions

import (
	"errors"
	"github.com/yehormironenko/tx_parser/internal/client/model"
	client "github.com/yehormironenko/tx_parser/internal/service"
	mock "github.com/yehormironenko/tx_parser/internal/service/actions/testing"
	"log"
	"os"
	"testing"
)

func TestSubscribe_Success(t *testing.T) {
	mockRepo := &mock.MockSubscriberRepository{
		InsertNewSubscriberFunc: func(address, blockNumber string) (bool, error) {
			if address == "0x123" && blockNumber == "0x10" {
				return true, nil
			}
			return false, nil
		},
	}
	mockClient := &mock.MockEthereumApiClient{
		GetCurrentBlockFunc: func() (*model.GetCurrentBlock, error) {
			return &model.GetCurrentBlock{Result: "0x10"}, nil
		},
	}
	logger := log.New(os.Stdout, "", log.LstdFlags)
	service := NewSubscriptionsService(mockRepo, client.ExternalClient{EthereumClient: mockClient}, logger)

	t.Cleanup(func() {
		mockClient = nil
		mockRepo = nil
	})

	success, err := service.Subscribe("0x123")
	if err != nil {
		t.Errorf("Subscribe() error = %v", err)
	}
	if !success {
		t.Error("Subscribe() success = false, want true")
	}
}

func TestSubscribe_ErrorFromClient(t *testing.T) {
	mockRepo := &mock.MockSubscriberRepository{}
	mockClient := &mock.MockEthereumApiClient{
		GetCurrentBlockFunc: func() (*model.GetCurrentBlock, error) {
			return nil, errors.New("client error")
		},
	}
	logger := log.New(os.Stdout, "", log.LstdFlags)
	service := NewSubscriptionsService(mockRepo, client.ExternalClient{EthereumClient: mockClient}, logger)

	success, err := service.Subscribe("0x123")
	if err == nil {
		t.Error("Subscribe() error = nil, want error")
	}
	if success {
		t.Error("Subscribe() success = true, want false")
	}
}

func TestUnsubscribe_Success(t *testing.T) {
	mockRepo := &mock.MockSubscriberRepository{
		RemoveSubscriberFunc: func(address string) (bool, error) {
			if address == "0x123" {
				return true, nil
			}
			return false, nil
		},
	}

	logger := log.New(os.Stdout, "", log.LstdFlags)
	service := NewSubscriptionsService(mockRepo, client.ExternalClient{}, logger)
	t.Cleanup(func() {
		mockRepo = nil
	})
	success, err := service.Unsubscribe("0x123")
	if err != nil {
		t.Errorf("Unsubscribe() error = %v", err)
	}
	if !success {
		t.Error("Unsubscribe() success = false, want true")
	}
}

func TestUnsubscribe_ErrorFromRepo(t *testing.T) {
	mockRepo := &mock.MockSubscriberRepository{
		RemoveSubscriberFunc: func(address string) (bool, error) {
			return false, errors.New("repository error")
		},
	}
	mockClient := &mock.MockEthereumApiClient{}
	logger := log.New(os.Stdout, "", log.LstdFlags)
	service := NewSubscriptionsService(mockRepo, client.ExternalClient{EthereumClient: mockClient}, logger)

	success, err := service.Unsubscribe("0x123")
	if err == nil {
		t.Error("Unsubscribe() error = nil, want error")
	}
	if success {
		t.Error("Unsubscribe() success = true, want false")
	}
}
