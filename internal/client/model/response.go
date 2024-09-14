package model

type GetCurrentBlock struct {
	JSONRPC string `json:"jsonrpc"`
	Result  string `json:"result"`
	ID      int64  `json:"id"`
}
