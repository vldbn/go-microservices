package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Address struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	UID      primitive.ObjectID `json:"uid" bson:"uid"`
	Username string             `json:"username" bson:"username"`
	Country  string             `json:"country" bson:"country"`
	City     string             `json:"city" bson:"city"`
	Address  string             `json:"address" bson:"address"`
}
