package mongodb

import (
	"address/internal/app/store"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoStore struct {
	client            *mongo.Client
	database          string
	addressRepository *AddressRepository
}

func (m *MongoStore) Address() store.AddressRepository {
	if m.addressRepository != nil {
		return m.addressRepository
	}

	m.addressRepository = &AddressRepository{
		collection: m.client.Database(m.database).Collection("addresses"),
	}
	return m.addressRepository
}

func NewMongoStore(client *mongo.Client, database string) store.Store {
	return &MongoStore{
		client:   client,
		database: database,
	}
}
