package service

type Echo interface {
	Echo() string
}

type GetCurrentBlock interface {
	GetCurrentBlock() (int, error)
}
