package actions

import (
	"errors"
	"github.com/yehormironenko/tx_parser/internal/client/model"
	txn "github.com/yehormironenko/tx_parser/internal/model"
	client "github.com/yehormironenko/tx_parser/internal/service"
	mock "github.com/yehormironenko/tx_parser/internal/service/actions/testing"
	"log"
	"os"
	"testing"
	"time"
)

func TestStartPolling_Success(t *testing.T) {
	mockRepo := &mock.MockSubscriberRepository{
		GetSubscribersFunc: func() map[string]string {
			return map[string]string{
				"0x123": "0x1",
			}
		},
		UpdateValueFunc: func(address, block string) (bool, error) {
			return true, nil
		},
	}

	mockClient := &mock.MockEthereumApiClient{
		GetCurrentBlockFunc: func() (*model.GetCurrentBlock, error) {
			return &model.GetCurrentBlock{Result: "0x10"}, nil // 0x10 - 16
		},
		GetTransactionsFunc: func(address, fromBlock, toBlock *string) (*model.EthLogResult, error) {
			return &model.EthLogResult{
				Result: []model.EthLog{
					{
						BlockNumber:      "0x2",
						TransactionIndex: "0x1",
						LogIndex:         "0x1",
						Address:          "0x123",
						Data:             "100",
						Removed:          false,
					},
				},
			}, nil
		},
	}

	logger := log.New(os.Stdout, "", log.LstdFlags)
	service := NewNotificationService(mockRepo, client.ExternalClient{EthereumClient: mockClient}, logger)

	go service.StartPolling()

	time.Sleep(1 * time.Second)

	// Assertions
	if service.currentBlock != 16 {
		t.Errorf("StartPolling() currentBlock = %v, want %v", service.currentBlock, 9)
	}
}

func TestStartPolling_ErrorFetchingBlock(t *testing.T) {
	mockRepo := &mock.MockSubscriberRepository{}

	mockClient := &mock.MockEthereumApiClient{
		GetCurrentBlockFunc: func() (*model.GetCurrentBlock, error) {
			return nil, errors.New("error fetching block")
		},
	}

	logger := log.New(os.Stdout, "", log.LstdFlags)
	service := NewNotificationService(mockRepo, client.ExternalClient{EthereumClient: mockClient}, logger)

	// Running in a separate goroutine to simulate polling
	go service.StartPolling()

	time.Sleep(1 * time.Second)

	if service.currentBlock != 0 {
		t.Errorf("StartPolling() currentBlock = %v, want %v", service.currentBlock, 0)
	}
}

func TestProcessNotifications_Success(t *testing.T) {
	mockRepo := &mock.MockSubscriberRepository{}
	mockClient := &mock.MockEthereumApiClient{}

	logger := log.New(os.Stdout, "", log.LstdFlags)
	service := NewNotificationService(mockRepo, client.ExternalClient{EthereumClient: mockClient}, logger)

	tx := txn.Transaction{
		Address:          "0x123",
		Amount:           "100",
		BlockNumber:      1,
		LogIndex:         1,
		TransactionIndex: 1,
		Removed:          false,
	}

	go func() {
		service.notifications <- tx
		close(service.notifications)
	}()

	service.ProcessNotifications()

	// Check if the notification was processed (logging output would indicate that)
	// Since there's no output, we just ensure no panics occurred
}

func TestFetchLatestBlock_Error(t *testing.T) {
	mockClient := &mock.MockEthereumApiClient{
		GetCurrentBlockFunc: func() (*model.GetCurrentBlock, error) {
			return nil, errors.New("error fetching block")
		},
	}

	logger := log.New(os.Stdout, "", log.LstdFlags)
	service := NewNotificationService(nil, client.ExternalClient{EthereumClient: mockClient}, logger)

	_, err := service.fetchLatestBlock()
	if err == nil {
		t.Error("fetchLatestBlock() error = nil, want error")
	}
}

func TestFetchLatestBlock_Success(t *testing.T) {
	mockClient := &mock.MockEthereumApiClient{
		GetCurrentBlockFunc: func() (*model.GetCurrentBlock, error) {
			return &model.GetCurrentBlock{Result: "0xA"}, nil
		},
	}

	logger := log.New(os.Stdout, "", log.LstdFlags)
	service := NewNotificationService(nil, client.ExternalClient{EthereumClient: mockClient}, logger)

	block, err := service.fetchLatestBlock()
	if err != nil {
		t.Fatalf("fetchLatestBlock() error = %v", err)
	}
	if block != 10 {
		t.Errorf("fetchLatestBlock() block = %v, want %v", block, 10)
	}
}
