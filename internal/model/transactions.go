package model

type Transactions []Transaction

type Transaction struct {
	Address          string `json:"address,omitempty"`
	Amount           string `json:"amount,omitempty"`
	BlockNumber      int64  `json:"blockNumber,omitempty"`
	TransactionIndex int64  `json:"transactionIndex,omitempty"`
	LogIndex         int64  `json:"logIndex,omitempty"`
	Removed          bool   `json:"removed,omitempty"`
}
