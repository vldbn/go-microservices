package store

import (
	"address/internal/app/model/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AddressRepository interface {
	GetAddresses() ([]*entities.Address, error)
	GetByAddressID(id primitive.ObjectID) (*entities.Address, error)
	UpdateAddress(address *entities.Address) error
	CreateAddress(address *entities.Address) error
	UpdateUsername(uid primitive.ObjectID, username string) error
	Delete(id primitive.ObjectID) error
}
