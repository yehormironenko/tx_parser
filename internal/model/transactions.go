package model

type Transactions []Transaction

type Transaction struct {
	Address          string `json:"address,omitempty"`
	Amount           string `json:"amount,omitempty"`
	BlockNumber      int    `json:"blockNumber,omitempty"`
	TransactionIndex int    `json:"transactionIndex,omitempty"`
	LogIndex         int    `json:"logIndex,omitempty"`
	Removed          bool   `json:"removed,omitempty"`
}
