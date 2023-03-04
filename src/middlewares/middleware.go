package middlewares

import (
	"github.com/andres123dbh/movil_parcial_backend/configuration"
	"github.com/andres123dbh/movil_parcial_backend/utils"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ensure there is a mongo connection
func CheckMongoConnection() gin.HandlerFunc {
	return func(c *gin.Context) {

		err := configuration.MongoCheckConnection()
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{"message": "Internal server error"})
			return
		}

	}
}

func ProvideAccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get access token from header
		accessToken := c.GetHeader("accessToken")

		if accessToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Access token is required"})
			return
		}

		// Check if access token is valid
		id, error := utils.ValidateAccessToken(accessToken)
		if error != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": error.Error()})
			return
		}

		// Set user id to context
		c.Set("userId", id)
	}
}
