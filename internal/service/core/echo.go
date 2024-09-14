package core

import (
	"log"
)

type Echo struct {
	logger *log.Logger
}

func NewEcho(logger *log.Logger) *Echo {
	return &Echo{logger: logger}
}

func (s *Echo) Echo() string {
	s.logger.Println("success from Echo service")
	return "Success from Echo!"
}
