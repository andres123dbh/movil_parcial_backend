package routes

import (
	"github.com/andres123dbh/movil_parcial_backend/controllers"
	"github.com/andres123dbh/movil_parcial_backend/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(engine *gin.Engine) {
	// Test
	engine.GET("/test", middlewares.CheckMongoConnection(), controllers.Test)
	engine.GET("/token", middlewares.ProvideAccessToken(), controllers.Token)

	//Session
	engine.POST("/login", middlewares.CheckMongoConnection(), controllers.Login)

	//Products
	engine.GET("/products/get", middlewares.ProvideAccessToken(), middlewares.CheckMongoConnection(), controllers.GetProducts)
}
