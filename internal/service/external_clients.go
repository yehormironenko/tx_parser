package service

import "github.com/yehormironenko/tx_parser/internal/client"

type ExternalClient struct {
	EthereumClient client.EthereumApiClient
}
