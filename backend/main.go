package main

import (
	"foodnetwork/routes"

	"github.com/gin-contrib/cors" // Import the CORS middleware
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Add the CORS middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	router.Use(cors.New(config))

	router.Use(cors.New(config)) // Add the CORS middleware to your Gin engine

	routes.SetupRouter(router)

	router.Run(":8080")
}
