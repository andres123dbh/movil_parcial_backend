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

type Product struct {
	Id     primitive.ObjectID `json:"_id,omitempty"       bson:"_id"`
	ID     int                `json:"id" bson:"id"`
	Image  string             `json:"image" bson:"image" `
	Title  string             `json:"title" bson:"title" `
	Rating float64            `json:"rating" bson:"rating"`
	Seller string             `json:"seller" bson:"seller"`
}
