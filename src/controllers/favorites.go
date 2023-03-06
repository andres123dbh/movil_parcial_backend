package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/andres123dbh/movil_parcial_backend/configuration"
	"github.com/andres123dbh/movil_parcial_backend/interfaces"
)

func UpdateFavorites(c *gin.Context) {

	var mc *mongo.Client = configuration.MongoGetClient()
	coll := mc.Database("movil_parcial").Collection("users")
	userid, _ := c.Get("userId")

	bodyFavs := interfaces.BodyFavorites{}

	if err := c.BindJSON(&bodyFavs); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	// organizing an array of ObjectIDs
	var user interfaces.User
	fav := user.Favorites
	for _, favId := range bodyFavs.BodyFavorites {
		object_id, _ := primitive.ObjectIDFromHex(favId.ID)
		fav = append(fav, object_id)
	}

	// updating favorites
	id, _ := primitive.ObjectIDFromHex(userid.(string))
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"favorites", fav}}}}
	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}
	// returns # of changes
	c.IndentedJSON(http.StatusOK, gin.H{"message": result})

}
