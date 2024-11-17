package actions

import (
	"errors"
	"log"
	"os"
	"testing"

	"github.com/yehormironenko/tx_parser/internal/client/model"
	"github.com/yehormironenko/tx_parser/internal/service"
	mock "github.com/yehormironenko/tx_parser/internal/service/actions/testing"
)

// TestGetCurrentBlock_Success tests the successful retrieval of the current block
func TestGetCurrentBlock_Success(t *testing.T) {
	logger := log.New(os.Stdout, "", log.LstdFlags) // Use nil for output
	mockClient := &mock.MockEthereumApiClient{
		GetCurrentBlockFunc: func() (*model.GetCurrentBlock, error) {
			return &model.GetCurrentBlock{Result: "0x13cbc26"}, nil
		},
	}
	serviceMock := NewGetCurrentBlockService(service.ExternalClient{EthereumClient: mockClient}, logger)

	expectedBlockNumber := 20757542

	blockNumber, err := serviceMock.GetCurrentBlock()

	if err != nil {
		t.Errorf("GetCurrentBlock() error = %v", err)
		return
	}
	if blockNumber != int(expectedBlockNumber) {
		t.Errorf("GetCurrentBlock() = %d; want %d", blockNumber, expectedBlockNumber)
	}
}

// TestGetCurrentBlock_ErrorFromClient tests handling errors from the external client
func TestGetCurrentBlock_ErrorFromClient(t *testing.T) {
	logger := log.New(os.Stdout, "", log.LstdFlags) // Use nil for output
	mockClient := &mock.MockEthereumApiClient{
		GetCurrentBlockFunc: func() (*model.GetCurrentBlock, error) {
			return nil, errors.New("client error")
		},
	}
	serviceMock := NewGetCurrentBlockService(service.ExternalClient{EthereumClient: mockClient}, logger)

	t.Cleanup(func() {
		mockClient = nil
	})

	blockNumber, err := serviceMock.GetCurrentBlock()

	if err == nil {
		t.Errorf("GetCurrentBlock() error = nil; want error")
	}
	if blockNumber != 0 {
		t.Errorf("GetCurrentBlock() = %d; want 0", blockNumber)
	}
}

// TestGetCurrentBlock_ErrorParsing tests handling errors from the ConvertHexToInt function
func TestGetCurrentBlock_ErrorParsing(t *testing.T) {
	logger := log.New(os.Stdout, "", log.LstdFlags) // Use nil for output
	mockClient := &mock.MockEthereumApiClient{
		GetCurrentBlockFunc: func() (*model.GetCurrentBlock, error) {
			return &model.GetCurrentBlock{Result: "0xG"}, nil // Invalid hex value
		},
	}
	serviceMock := NewGetCurrentBlockService(service.ExternalClient{EthereumClient: mockClient}, logger)

	blockNumber, err := serviceMock.GetCurrentBlock()

	if err == nil {
		t.Errorf("GetCurrentBlock() error = nil; want error")
	}
	if blockNumber != 0 {
		t.Errorf("GetCurrentBlock() = %d; want 0", blockNumber)
	}
}
