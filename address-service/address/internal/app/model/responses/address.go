package responses

import "address/internal/app/model/entities"

type AddressRes struct {
	Address *entities.Address `json:"address,omitempty"`
	Error   string            `json:"error,omitempty"`
}

type AddressesRes struct {
	Addresses []*entities.Address `json:"addresses,omitempty"`
	Error     string              `json:"error,omitempty"`
}
