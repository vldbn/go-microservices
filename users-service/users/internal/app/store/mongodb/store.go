package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
	"users/internal/app/store"
)

type MongoStore struct {
	client         *mongo.Client
	database       string
	userRepository *UserRepository
}

func (m *MongoStore) User() store.UserRepository {
	if m.userRepository != nil {
		return m.userRepository
	}

	m.userRepository = &UserRepository{
		collection: m.client.Database(m.database).Collection("users"),
	}
	return m.userRepository
}

func NewMongoStore(client *mongo.Client, database string) store.Store {
	return &MongoStore{
		client:   client,
		database: database,
	}
}
