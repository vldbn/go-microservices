package impl

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"users/internal/app/model/entities"
	"users/internal/app/model/messages"
	"users/internal/app/service/broker"
	"users/internal/app/store"
)

type UserService struct {
	store store.Store
	ch    *amqp.Channel
}

func (u *UserService) Create(username string) (*entities.User, error) {
	usr := entities.User{
		Username: username,
	}

	if err := u.store.User().CreateUser(&usr); err != nil {
		return nil, err
	}

	go func() {

		usrJ, err := json.Marshal(&usr)

		if err != nil {
			log.Println(err)
		}

		if err := broker.SendMessage(u.ch, "users.create", usrJ); err != nil {
			log.Println(err)
		}
	}()

	return &usr, nil
}

func (u *UserService) GetByID(id string) (*entities.User, error) {
	idP, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	usr, err := u.store.User().GetUserByID(idP)

	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (u *UserService) GetList() ([]*entities.User, error) {
	users, err := u.store.User().GetUsers()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserService) Update(id string, username string) (*entities.User, error) {
	idP, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	usr, err := u.store.User().GetUserByID(idP)

	if err != nil {
		return nil, err
	}

	usr.Username = username

	err = u.store.User().UpdateUser(usr)

	if err != nil {
		return nil, err
	}

	go func() {
		usrJ, err := json.Marshal(&usr)

		if err != nil {
			log.Println(err)
		}

		if err := broker.SendMessage(u.ch, "users.update", usrJ); err != nil {
			log.Println(err)
		}
	}()

	return usr, nil
}

func (u *UserService) Delete(id string) error {
	idP, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	if err := u.store.User().DeleteUser(idP); err != nil {
		return err
	}

	go func() {

		usrDelMsg := messages.UserDeleteMsg{
			ID: id,
		}

		msgJ, err := json.Marshal(usrDelMsg)

		if err != nil {
			log.Println(err)
		}

		if err := broker.SendMessage(u.ch, "users.delete", msgJ); err != nil {
			log.Println(err)
		}
	}()

	return nil
}
