package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type RoleType string

const (
	PLAYER RoleType = "player"
	ADMIN  RoleType = "admin"
)

type User struct {
	Id      primitive.ObjectID `bson:"_id,omitempty"`
	Code    int                `bson:"code"`
	Name    string             `bson:"name,omitempty"`
	Numbers []int              `bson:"numbers,omitempty"`
	Role    string             `bson:"roleType"`
}
