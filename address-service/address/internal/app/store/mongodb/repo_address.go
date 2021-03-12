package mongodb

import (
	"address/internal/app/model/entities"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type AddressRepository struct {
	collection *mongo.Collection
}

func (a *AddressRepository) GetAddresses() ([]*entities.Address, error) {
	var addresses []*entities.Address = make([]*entities.Address, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := a.collection.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		adr := entities.Address{}
		if err := cur.Decode(&adr); err != nil {
			return nil, err
		}
		addresses = append(addresses, &adr)
	}

	return addresses, nil
}

func (a *AddressRepository) GetByAddressID(id primitive.ObjectID) (*entities.Address, error) {
	var adr *entities.Address

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	if err := a.collection.FindOne(ctx, filter).Decode(&adr); err != nil {
		return nil, err
	}

	return adr, nil
}

func (a *AddressRepository) UpdateAddress(address *entities.Address) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	filter := bson.M{"_id": address.ID}

	upd := bson.M{
		"country": address.Country,
		"city":    address.City,
		"address": address.Address,
	}

	res, err := a.collection.UpdateOne(ctx, filter, bson.D{primitive.E{Key: "$set", Value: upd}})

	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		e := errors.New("address not found")
		return e
	}

	return nil
}

func (a *AddressRepository) CreateAddress(address *entities.Address) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := a.collection.InsertOne(ctx, address)

	if err != nil {
		return err
	}

	return nil
}

func (a *AddressRepository) UpdateUsername(uid primitive.ObjectID, username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	filter := bson.M{"uid": uid}
	upd := bson.M{
		"username": username,
	}

	res, err := a.collection.UpdateOne(ctx, filter, bson.D{primitive.E{Key: "$set", Value: upd}})

	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		e := errors.New("address not found")
		return e
	}

	return nil
}

func (a *AddressRepository) Delete(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	filter := bson.M{"uid": id}

	res, err := a.collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		e := errors.New("address not found")
		return e
	}

	return nil
}
