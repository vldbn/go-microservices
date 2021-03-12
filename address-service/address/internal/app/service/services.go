package service

import (
	"address/internal/app/model/entities"
)

type AddressService interface {
	GetList() ([]*entities.Address, error)
	GetByID(id string) (*entities.Address, error)
	Update(id string, country string, city string, address string) (*entities.Address, error)
	Create(body []byte)
	UpdateUsername(body []byte)
	Delete(body []byte)
}

type Services interface {
	Address() AddressService
}
