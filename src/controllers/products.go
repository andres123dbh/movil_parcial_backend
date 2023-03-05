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

func GetProducts(c *gin.Context) {

	var mc *mongo.Client = configuration.MongoGetClient()
	coll := mc.Database("movil_parcial").Collection("products")

	cursor, err := coll.Find(context.TODO(), bson.D{})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}

	var arrayProducts []interfaces.Product

	for cursor.Next(context.TODO()) {
		var products interfaces.Product
		if err := cursor.Decode(&products); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
			return
		}
		arrayProducts = append(arrayProducts, products)
	}

	c.IndentedJSON(http.StatusOK, gin.H{"error": err,
		"products": arrayProducts})

}
