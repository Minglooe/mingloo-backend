package routes

import (
	"github.com/gin-gonic/gin"
	"mingloo/web-api/internal/handlers"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/users", handlers.GetUsers)
		api.POST("/users", handlers.CreateUser)
	}
}
