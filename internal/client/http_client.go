package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"tx_parser/internal/client/model"
)

type EthereumApiClient interface {
	GetCurrentBlock() (int, error)
}

type ethereumApiClient struct {
	baseUrl    string
	httpClient http.Client
	logger     *log.Logger
}

func NewEthereumApiClient(baseUrl string, logger *log.Logger) EthereumApiClient {
	return &ethereumApiClient{
		baseUrl:    baseUrl,
		httpClient: http.Client{},
		logger:     logger,
	}
}

// GetCurrentBlock fetches the current block number from the Ethereum blockchain
func (e *ethereumApiClient) GetCurrentBlock() (int, error) {
	e.logger.Println("Getting current block in the Ethereum blockchain")

	// Define the request body
	// TODO change it and what is ID
	requestBody := []byte(`{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":83}`)

	req, err := http.NewRequest("POST", e.baseUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		return 0, fmt.Errorf("error creating request: %w", err)
	}

	e.logger.Printf("Sending request to %s", e.baseUrl)

	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := e.httpClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("error reading response body: %w", err)
	}

	// Unmarshal the response body
	var jsonResponse model.GetCurrentBlock
	err = json.Unmarshal(responseBody, &jsonResponse)
	if err != nil {
		return 0, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	// Extract the result and convert it to an integer
	hexValue := jsonResponse.Result
	if len(hexValue) > 2 && hexValue[:2] == "0x" {
		hexValue = hexValue[2:]
	}

	intValue, err := strconv.ParseInt(hexValue, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("error converting hex to int: %w", err)
	}
	log.Printf("Response from the extenal service: %v", jsonResponse)
	return int(intValue), nil
}
