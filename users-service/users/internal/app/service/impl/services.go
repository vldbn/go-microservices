package impl

import (
	"github.com/streadway/amqp"
	"users/internal/app/service"
	"users/internal/app/store"
)

type ServicesImpl struct {
	userService *UserService
	store       store.Store
	ch          *amqp.Channel
}

func (s *ServicesImpl) User() service.UserService {
	if s.userService != nil {
		return s.userService
	}

	s.userService = &UserService{
		store: s.store,
		ch:    s.ch,
	}
	return s.userService
}

func NewServices(store store.Store, ch *amqp.Channel) service.Services {
	return &ServicesImpl{
		store: store,
		ch:    ch,
	}
}
