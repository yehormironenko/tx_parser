package repository

type SubscriberRepository interface {
	InsertNewSubscriber(address string) (bool, error)
	RemoveSubscriber(address string) (bool, error)
}
