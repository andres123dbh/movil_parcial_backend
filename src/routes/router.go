package routes

import (
	"github.com/andres123dbh/movil_parcial_backend/controllers"
	"github.com/andres123dbh/movil_parcial_backend/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(engine *gin.Engine) {
	// Test
	engine.GET("/test", middlewares.CheckMongoConnection(), controllers.Test)
}