package repository

type SubscribeRepo interface {
	InsertNewSubscriber(address string) (bool, error)
	RemoveSubscriber(address string) (bool, error)
}
