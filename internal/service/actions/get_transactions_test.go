package actions

import (
	"errors"
	"github.com/yehormironenko/tx_parser/internal/client/model"
	txn "github.com/yehormironenko/tx_parser/internal/model"
	"github.com/yehormironenko/tx_parser/internal/service"
	mock "github.com/yehormironenko/tx_parser/internal/service/actions/testing"
	"log"
	"os"
	"testing"
)

/*// MockEthereumApiClientTransactions is a mock for the EthereumApiClient
type MockEthereumApiClientTransactions struct {
	GetTransactionsFunc func(address, fromBlock, toBlock *string) (*model.EthLogResult, error)
}

func (m *MockEthereumApiClientTransactions) GetCurrentBlock() (*model.GetCurrentBlock, error) {
	return nil, nil
}

func (m *MockEthereumApiClientTransactions) GetTransactions(address, fromBlock, toBlock *string) (*model.EthLogResult, error) {
	return m.GetTransactionsFunc(address, fromBlock, toBlock)
}*/

// TestGetTransactions_Success tests the successful retrieval of txn
func TestGetTransactions_Success(t *testing.T) {
	logger := log.New(os.Stdout, "", log.LstdFlags) // Use nil for output
	mockClient := &mock.MockEthereumApiClient{
		GetTransactionsFunc: func(address, fromBlock, toBlock *string) (*model.EthLogResult, error) {
			return &model.EthLogResult{
				Result: []model.EthLog{
					{
						Address:          "0x123",
						BlockNumber:      "0x1",
						TransactionIndex: "0x1",
						LogIndex:         "0x1",
						Data:             "0xabc",
						Removed:          false,
					},
				},
			}, nil
		},
	}
	serviceMock := NewGetTransactionsService(service.ExternalClient{EthereumClient: mockClient}, logger)

	expectedTransactions := txn.Transactions{
		{
			Address:          "0x123",
			Amount:           "0xabc",
			BlockNumber:      1,
			LogIndex:         1,
			TransactionIndex: 1,
			Removed:          false,
		},
	}

	transactions, err := serviceMock.GetTransactions("0x123")

	if err != nil {
		t.Errorf("GetTransactions() error = %v", err)
		return
	}
	if len(transactions) != len(expectedTransactions) {
		t.Errorf("GetTransactions() = %v; want %v", transactions, expectedTransactions)
		return
	}
	for i, tx := range transactions {
		if tx != expectedTransactions[i] {
			t.Errorf("Transaction %d = %v; want %v", i, tx, expectedTransactions[i])
		}
	}
}

// TestGetTransactions_ErrorFromClient tests handling errors from the external client
func TestGetTransactions_ErrorFromClient(t *testing.T) {
	logger := log.New(os.Stdout, "", log.LstdFlags)

	mockClient := &mock.MockEthereumApiClient{
		GetTransactionsFunc: func(address, fromBlock, toBlock *string) (*model.EthLogResult, error) {
			return nil, errors.New("client error")
		},
	}
	serviceMock := NewGetTransactionsService(service.ExternalClient{EthereumClient: mockClient}, logger)

	transactions, err := serviceMock.GetTransactions("0x123")

	if err == nil {
		t.Errorf("GetTransactions() error = nil; want error")
	}
	if len(transactions) != 0 {
		t.Errorf("GetTransactions() = %v; want empty slice", transactions)
	}
}

// TestGetTransactions_ErrorParsing tests handling errors from the ConvertHexToInt function
func TestGetTransactions_ErrorParsing(t *testing.T) {
	logger := log.New(os.Stdout, "", log.LstdFlags) // Use nil for output
	mockClient := &mock.MockEthereumApiClient{
		GetTransactionsFunc: func(address, fromBlock, toBlock *string) (*model.EthLogResult, error) {
			return &model.EthLogResult{
				Result: []model.EthLog{
					{
						Address:          "0x123",
						BlockNumber:      "0xG", // Invalid hex value
						TransactionIndex: "0x1",
						LogIndex:         "0x1",
						Data:             "0xabc",
						Removed:          false,
					},
				},
			}, nil
		},
	}
	serviceMock := NewGetTransactionsService(service.ExternalClient{EthereumClient: mockClient}, logger)

	transactions, err := serviceMock.GetTransactions("0x123")

	if err == nil {
		t.Errorf("GetTransactions() error = nil; want error")
	}
	if len(transactions) != 0 {
		t.Errorf("GetTransactions() = %v; want empty slice", transactions)
	}
}
