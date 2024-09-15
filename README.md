# Tx Parser

## Description
`Tx Parser` is an Ethereum blockchain parser that allows querying transactions for subscribed addresses.

## Installation
1. Clone the repository:
    ```bash
    git clone https://github.com/yehormironenko/tx_parser
    cd tx_parser
    ```

## Usage

### Using Docker

To build and run the application using Docker:

1. Build the Docker image:
    ```bash
    docker build -t tx_parser .
    ```

2. Run the Docker container:
    ```bash
    docker run -p 8092:8092 tx_parser
    ```

### Running the Go Application

You can also run the application directly from the terminal using Go:

```bash
go run ./cmd/main.go
```
Application will be available on the 127.0.0.1:8092 address

## List of Operations

- **Echo** - Check availability of the service
    - **Request**:
      ```http
      GET: localhost:8092/echo
      ```
    - **Response**:
      ```plaintext
      Success from Echo!
      ```

- **GetCurrentBlock** - Returns the last parsed block
    - **Request**:
      ```http
      GET: localhost:8092/current-block
      ```
    - **Response**:
      ```json
      {
        "blockNumber": 20757286
      }
      ```

- **GetTransactions** - Returns a list of inbound or outbound transactions for an address
    - **Request**:
      ```http
      POST: localhost:8092/transactions
      ```
    - **Response**:
      ```json
      [
        {
          "address": "0xbe206379252ed32b85cf8d1f53195c6daac75801",
          "amount": "0xffffffffffffffffffffffffffffffffffffffffffffffffffffce33e0fa0b720000000000000000000000000000000000000000000000000b7ba40800000000000000000000000000000000000000000000007a602c49392af18f1d66d55bcc0000000000000000000000000000000000000000000000005e8713ff23e7403a0000000000000000000000000000000000000000000000000000000000017792",
          "blockNumber": 20757973,
          "transactionIndex": 5,
          "logIndex": 32
        }
      ]
      ```

- **Subscribe** - Add an address to observers. Two available actions: subscribe/unsubscribe
    - **Example**:
        - **Request**:
          ```http
          POST: localhost:8092/subscribe
          ```
          ```json
          {
            "action": "subscribe",
            "address": "0xbe206379252ed32b85cf8d1f53195c6daac75801"
          }
          ```
        - **Response**:
          ```plaintext
          Subscribed to address 0xbe206379252ed32b85cf8d1f53195c6daac75801
          ```

        - **Request**:
          ```json
          {
            "action": "unsubscribe",
            "address": "0xbe206379252ed32b85cf8d1f53195c6daac75801"
          }
          ```
        - **Response**:
          ```plaintext
          Unsubscribed from address 0xbe206379252ed32b85cf8d1f53195c6daac75801
          ```
## Testing
To run tests, use the following command:
```bash
go test ./...
```

## External Client
All information is received from [Ethereum RPC Public Node](https://ethereum-rpc.publicnode.com)

## References
- [Ethereum JSON RPC Interface](https://ethereum.org/en/developers/docs/apis/json-rpc/)

## Limitations
This project uses only internal libraries of Go.

## Contact Information
For support or questions, please contact yegor.mironenko@gmail.com

## Changelog
- **v1.0.0**: Initial release