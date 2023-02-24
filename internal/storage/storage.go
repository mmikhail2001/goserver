package storage

import (
	"github.com/mmikhail2001/goserver/model"
)

type Storage interface {
	CreateUser(user model.User) (int, error)
}

type storage struct {
}

func (s *storage) CreateUser(user model.User) (int, error) {
	return 0, nil
}

func NewStorage() Storage {
	return &storage{}
}
