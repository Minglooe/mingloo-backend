package routes

import (
	"mingloo/web-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

// @title Sample API
// @version 1.0
// @description This is a sample API with Swagger documentation.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@example.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api/v1
func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.GET("", handlers.GetUsers)
			users.POST("", handlers.CreateUser)
		}

		events := api.Group("/events")
		{
			events.GET("", handlers.GetEvents)
			events.POST("", handlers.CreateEvent)
			events.PATCH("", handlers.UpdateEvent)
			events.DELETE("", handlers.DeleteEvent)
		}
	}
}
