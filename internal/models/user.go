package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoleType string

const (
	RolePlayer RoleType = "player"
	RoleAdmin  RoleType = "admin"
)

type User struct {
	ID      primitive.ObjectID `bson:"_id" unique:"true" json:"id"`
	Code    string             `bson:"code" unique:"true"`
	Name    string             `bson:"name,omitempty" unique:"true"`
	Numbers []int              `bson:"numbers,omitempty"`
	Role    string             `bson:"role"`
}
