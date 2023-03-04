package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/andres123dbh/movil_parcial_backend/configuration"
	"github.com/andres123dbh/movil_parcial_backend/interfaces"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Login(c *gin.Context) {
	var err error
	var form interfaces.LoginForm

	if err := c.BindJSON(&form); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	// find user with email

	if form.Email == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Email cannot be empty"})
		return
	}

	if form.Password == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Password cannot be empty"})
		return
	}

	var mc *mongo.Client = configuration.MongoGetClient()
	var user interfaces.User
	coll := mc.Database("movil_parcial").Collection("users")

	err = coll.FindOne(
		context.TODO(),
		bson.D{{Key: "email", Value: form.Email}},
	).Decode(&user)

	if err != nil {
		// ErrNoDocuments: the filter did not match any documents
		if err == mongo.ErrNoDocuments {
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"message": "Wrong user/password"})
			return
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{"message": "Internal server error"})
			return
		}
	}

	// check password

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password)); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"message": "Wrong user/password"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully logged in"})
}
