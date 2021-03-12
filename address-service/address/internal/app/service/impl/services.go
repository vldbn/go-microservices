package impl

import (
	"address/internal/app/service"
	"address/internal/app/store"
)

type Services struct {
	addressService *AddressService
	store          store.Store
}

func (s *Services) Address() service.AddressService {
	if s.addressService != nil {
		return s.addressService
	}

	s.addressService = &AddressService{
		store: s.store,
	}
	return s.addressService
}

func NewServices(store store.Store) service.Services {
	return &Services{
		store: store,
	}
}
