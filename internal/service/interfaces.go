package service

import (
	"tx_parser/internal/model"
)

type Echo interface {
	Echo() string
}

type GetCurrentBlock interface {
	GetCurrentBlock() (int, error)
}

type GetTransactions interface {
	GetTransactions(address string) (model.Transactions, error)
}

type Notification interface {
	Subscribe(address string) (bool, error)
	Unsubscribe(address string) (bool, error)
	Notify(address string, transaction model.Transaction)
}
