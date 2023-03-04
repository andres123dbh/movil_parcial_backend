package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/andres123dbh/movil_parcial_backend/configuration"
	"github.com/andres123dbh/movil_parcial_backend/interfaces"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Test(c *gin.Context) {

	var mc *mongo.Client = configuration.MongoGetClient()
	var user interfaces.User
	coll := mc.Database("movil_parcial").Collection("users")

	err := coll.FindOne(
		context.TODO(),
		bson.D{{Key: "email", Value: "andres@gmail.co"}},
	).Decode(&user)

	c.IndentedJSON(http.StatusOK, gin.H{"error": err,
		"message": user})

}

func Token(c *gin.Context) {

	userid, _ := c.Get("userId")

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Successfully retrieved user",
		"user":    userid,
	})

}
