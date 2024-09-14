package model

type GetTransactionsRequest struct {
	Address string `json:"address,required"`
}
