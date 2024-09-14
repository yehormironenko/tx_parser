package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
	"tx_parser/internal/client/model"
)

type EthereumApiClient interface {
	GetCurrentBlock() (*model.GetCurrentBlock, error)
	GetTransactions(address string) (*model.EthLogResult, error)
}

type ethereumApiClient struct {
	baseUrl    string
	httpClient http.Client
	logger     *log.Logger
}

func NewEthereumApiClient(baseUrl string, logger *log.Logger) EthereumApiClient {
	return &ethereumApiClient{
		baseUrl:    baseUrl,
		httpClient: http.Client{Timeout: 2 * time.Second},
		logger:     logger,
	}
}

// GetCurrentBlock fetches the current block number from the Ethereum blockchain
func (e *ethereumApiClient) GetCurrentBlock() (*model.GetCurrentBlock, error) {
	e.logger.Println("Getting current block in the Ethereum blockchain")

	// Define the request body
	// TODO change it and what is ID
	requestBody := []byte(`{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}`)

	req, err := http.NewRequest("POST", e.baseUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	e.logger.Printf("Sending request to %s", e.baseUrl)

	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := e.httpClient.Do(req)
	if err != nil {
		//e.logger.Fatalf("error sending request: %v", err)
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// Unmarshal the response body
	var jsonResponse model.GetCurrentBlock
	err = json.Unmarshal(responseBody, &jsonResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %w", err)
	}
	e.logger.Printf("Response from %v : %v", e.baseUrl, jsonResponse)

	return &jsonResponse, nil
}

func (e *ethereumApiClient) GetTransactions(address string) (*model.EthLogResult, error) {
	e.logger.Printf("Getting transactions for address %v", address)

	params := model.EthGetLogsParams{
		Address: address,
	}

	// Prepare the JSON request body
	requestBody, err := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_getLogs",
		"params":  []interface{}{params},
		"id":      1,
	})
	if err != nil {
		return nil, fmt.Errorf("error marshalling request: %w", err)
	}

	req, err := http.NewRequest("POST", e.baseUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	e.logger.Printf("Sending request to %s", e.baseUrl)

	resp, err := e.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// Parse the JSON response
	var response model.EthLogResult
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	e.logger.Printf("Response from %v : %v", e.baseUrl, response)

	return &response, nil

}
