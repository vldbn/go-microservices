package store

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"users/internal/app/model/entities"
)

type UserRepository interface {
	CreateUser(user *entities.User) error
	GetUserByID(id primitive.ObjectID) (*entities.User, error)
	GetUsers() ([]*entities.User, error)
	UpdateUser(user *entities.User) error
	DeleteUser(id primitive.ObjectID) error
}
