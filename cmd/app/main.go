package main

import (
	"todoapi/pkg/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Initialize routes
	routes.SetupTodoRoutes(router)

	// Start the server
	router.Run("localhost:3333")
}
