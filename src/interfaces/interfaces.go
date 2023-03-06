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

// todo delete id string, is just for dev purposes
type Product struct {
	Id     primitive.ObjectID `json:"_id,omitempty"       bson:"_id"`
	ID     string             `json:"id" bson:"id"`
	Title  string             `json:"title" bson:"title" `
	Seller string             `json:"seller" bson:"seller"`
	Rating float64            `json:"rating" bson:"rating"`
	Image  string             `json:"image" bson:"image" `
}

type BodyFavorites struct {
	BodyFavorites []Product `json:"favorites"`
}

type Favorites struct {
	Favorites primitive.ObjectID `json:"favorites" bson:"favorites"`
}
