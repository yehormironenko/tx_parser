package service

import (
	"tx_parser/internal/model"
)

type Echoer interface {
	Echo() string
}

type BlockRetriever interface {
	GetCurrentBlock() (int, error)
}

type TransactionFetcher interface {
	GetTransactions(address string) (model.Transactions, error)
}

type SubscriptionManager interface {
	Subscribe(address string) (bool, error)
	Unsubscribe(address string) (bool, error)
}

type Notifier interface {
	Notify(address string, transaction model.Transaction)
}
