package testing

import "github.com/yehormironenko/tx_parser/internal/client/model"

type MockSubscriberRepository struct {
	GetSubscribersFunc      func() map[string]string
	UpdateValueFunc         func(address, block string) (bool, error)
	InsertNewSubscriberFunc func(address, blockNumber string) (bool, error)
	RemoveSubscriberFunc    func(address string) (bool, error)
}

func (m *MockSubscriberRepository) GetSubscribers() map[string]string {
	return m.GetSubscribersFunc()
}

func (m *MockSubscriberRepository) UpdateValue(address, block string) (bool, error) {
	return m.UpdateValueFunc(address, block)
}

func (m *MockSubscriberRepository) InsertNewSubscriber(address, blockNumber string) (bool, error) {
	return m.InsertNewSubscriberFunc(address, blockNumber)
}

func (m *MockSubscriberRepository) RemoveSubscriber(address string) (bool, error) {
	return m.RemoveSubscriberFunc(address)
}

// Mock for ApiClient
type MockEthereumApiClient struct {
	GetCurrentBlockFunc func() (*model.GetCurrentBlock, error)
	GetTransactionsFunc func(address, fromBlock, toBlock *string) (*model.EthLogResult, error)
}

func (m *MockEthereumApiClient) GetCurrentBlock() (*model.GetCurrentBlock, error) {
	return m.GetCurrentBlockFunc()
}

func (m *MockEthereumApiClient) GetTransactions(address, fromBlock, toBlock *string) (*model.EthLogResult, error) {
	return m.GetTransactionsFunc(address, fromBlock, toBlock)
}
