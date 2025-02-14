package main

import (
	"fmt"
	"log"

	"mingloo/web-api/config"
	"mingloo/web-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Set Gin mode
	gin.SetMode(config.AppConfig.Server.Mode)

	// Initialize router
	r := gin.Default()

	// TODO: Modificar para settar o CORS do config
	r.SetTrustedProxies([]string{"127.0.0.1"})

	r.Use(gin.Logger(), gin.Recovery())

	// Register routes
	routes.RegisterRoutes(r)

	// Start the server
	port := config.AppConfig.Server.Port
	log.Printf("Server running on port %s", port)
	r.Run(fmt.Sprintf(":%s", port))
}
