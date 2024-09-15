package repository

type SubscriberRepository interface {
	InsertNewSubscriber(address, blockNumber string) (bool, error)
	RemoveSubscriber(address string) (bool, error)
	IsSubscribed(address string) (bool, error)
	GetSubscribers() map[string]string
	UpdateValue(address, blockNumber string) (bool, error)
}
