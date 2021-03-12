package responses

import (
	"users/internal/app/model/entities"
)

type UserRes struct {
	User  *entities.User `json:"user,omitempty"`
	Error string         `json:"error,omitempty"`
}

type UsersRes struct {
	Users []*entities.User `json:"users,omitempty"`
	Error string           `json:"error,omitempty"`
}

