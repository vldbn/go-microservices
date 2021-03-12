package service

import "users/internal/app/model/entities"

type UserService interface {
	Create(username string) (*entities.User, error)
	GetByID(id string) (*entities.User, error)
	GetList() ([]*entities.User, error)
	Update(id string, username string) (*entities.User, error)
	Delete(id string) error
}

type Services interface {
	User() UserService
}
