package middlewares

import (
	"github.com/andres123dbh/movil_parcial_backend/configuration"

	"fmt"
    "net/http"

    "github.com/gin-gonic/gin"
)

// ensure there is a mongo connection
func CheckMongoConnection() gin.HandlerFunc {
	return func(c *gin.Context){

		err := configuration.MongoCheckConnection()
		if (err != nil){
			fmt.Println(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"message": "Internal server error"})
			return
		}

	}
}