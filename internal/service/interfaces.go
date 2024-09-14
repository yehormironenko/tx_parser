package service

import (
	"tx_parser/internal/model"
)

type Echo interface {
	Echo() string
}

type GetCurrentBlock interface {
	GetCurrentBlock() (int64, error)
}

type GetTransactions interface {
	GetTransactions(address string) (model.Transactions, error)
}
