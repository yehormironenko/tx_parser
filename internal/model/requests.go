package model

type GetTransactionsRequest struct {
	Address string `json:"address,required"`
}

type SubscribeRequest struct {
	Action  string `json:"action,required"` // "subscribe" or "unsubscribe"
	Address string `json:"address,required"`
}
