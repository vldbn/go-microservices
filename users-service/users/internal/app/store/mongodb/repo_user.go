package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"users/internal/app/model/entities"
)

type UserRepository struct {
	collection *mongo.Collection
}

func (u *UserRepository) CreateUser(user *entities.User) error {
	user.ID = primitive.NewObjectID()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := u.collection.InsertOne(ctx, user)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) GetUserByID(id primitive.ObjectID) (*entities.User, error) {
	var user *entities.User

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	if err := u.collection.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) GetUsers() ([]*entities.User, error) {
	var users []*entities.User

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := u.collection.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		usr := entities.User{}
		if err := cur.Decode(&usr); err != nil {
			return nil, err
		}
		users = append(users, &usr)
	}

	return users, nil
}

func (u *UserRepository) UpdateUser(user *entities.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	filter := bson.M{"_id": user.ID}

	upd := bson.M{
		"username": user.Username,
	}

	res, err := u.collection.UpdateOne(ctx, filter, bson.D{primitive.E{Key: "$set", Value: upd}})

	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		e := errors.New("user not found")
		return e
	}

	return nil
}

func (u *UserRepository) DeleteUser(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	filter := bson.M{"_id": id}

	res, err := u.collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		e := errors.New("user not found")
		return e
	}

	return nil
}
