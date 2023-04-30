package main

import (
	"log"

	"food/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := routes.SetupRouter()

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
