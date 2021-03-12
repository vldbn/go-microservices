package impl

import (
	"address/internal/app/model/entities"
	"address/internal/app/model/messages"
	"address/internal/app/store"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type AddressService struct {
	store store.Store
}

func (a *AddressService) GetList() ([]*entities.Address, error) {
	addresses, err := a.store.Address().GetAddresses()

	if err != nil {
		return nil, err
	}

	return addresses, nil
}

func (a *AddressService) GetByID(id string) (*entities.Address, error) {
	idP, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	adr, err := a.store.Address().GetByAddressID(idP)

	if err != nil {
		return nil, err
	}

	return adr, nil
}

func (a *AddressService) Update(id string, country string, city string, address string) (*entities.Address, error) {
	idP, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	adr, err := a.store.Address().GetByAddressID(idP)

	if err != nil {
		return nil, err
	}

	adr.Country = country
	adr.City = city
	adr.Address = address

	if err := a.store.Address().UpdateAddress(adr); err != nil {
		return nil, err
	}

	return adr, nil
}

func (a *AddressService) UpdateUsername(body []byte) {
	log.Println("[INFO] UPDATE - ", string(body))

	usr := messages.UserMsg{}

	if err := json.Unmarshal(body, &usr); err != nil {
		log.Println("Failed to unmarshal message - ", err)
		return
	}

	uid, err := primitive.ObjectIDFromHex(usr.ID)

	if err != nil {
		log.Println(err)
		return
	}

	if err := a.store.Address().UpdateUsername(uid, usr.Username); err != nil {
		log.Println("Failed to update address in db - ", err)
		return
	}
}

func (a *AddressService) Delete(body []byte) {
	log.Println("[INFO] DELETE - ", string(body))

	usr := messages.UserMsg{}

	if err := json.Unmarshal(body, &usr); err != nil {
		log.Println("Failed to unmarshal message - ", err)
		return
	}

	uid, err := primitive.ObjectIDFromHex(usr.ID)

	if err != nil {
		log.Println(err)
		return
	}

	if err := a.store.Address().Delete(uid); err != nil {
		log.Println("Failed to delete address in db - ", err)
		return
	}
}

func (a *AddressService) Create(body []byte) {
	log.Println("[INFO] CREATE - ", string(body))
	usr := messages.UserMsg{}

	if err := json.Unmarshal(body, &usr); err != nil {
		log.Println("Failed to unmarshal message - ", err)
		return
	}

	uid, err := primitive.ObjectIDFromHex(usr.ID)

	if err != nil {
		log.Println(err)
		return
	}

	adr := entities.Address{
		UID:      uid,
		Username: usr.Username,
	}

	if err := a.store.Address().CreateAddress(&adr); err != nil {
		log.Println("Failed to create address in db - ", err)
		return
	}

}
