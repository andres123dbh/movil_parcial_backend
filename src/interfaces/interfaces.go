package interfaces

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID   `json:"_id"       bson:"_id"`
	Firstname string               `json:"firstname" bson:"firstname"`
	Email     string               `json:"email"     bson:"email"`
	Password  string               `json:"password"  bson:"password"`
	Favorites []primitive.ObjectID `json:"favorites" bson:"favorites"`
}

type LoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
