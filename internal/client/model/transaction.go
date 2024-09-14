package model

type EthGetLogsParams struct {
	//FromBlock string `json:"from_block"`
	//ToBlock   string `json:"to_block"`
	Address string `json:"address"`
}

type EthLogResult struct {
	Result []EthLog `json:"result"`
}

type EthLog struct {
	Address          string `json:"address,omitempty"`
	BlockNumber      string `json:"blockNumber,omitempty"`
	Data             string `json:"data,omitempty"`
	TransactionIndex string `json:"transactionIndex,omitempty"`
	LogIndex         string `json:"logIndex,omitempty"`
	Removed          bool   `json:"removed,omitempty"`
}
