package service

import (
	"github.com/mmikhail2001/goserver/internal/storage"
	"github.com/mmikhail2001/goserver/model"
)

type Service interface {
	CreateUser(user model.User) (int, error)
}

type service struct {
	storage storage.Storage
}

func (s *service) CreateUser(user model.User) (int, error) {
	return s.storage.CreateUser(user)
}

func NewService(storage storage.Storage) Service {
	return &service{
		storage: storage,
	}
}
