package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initializer Router
	router := gin.Default()

	// Initialize Routes
	initializeRoutes(router)
	// Run the server
	router.Run(":8080") // padrão roda na porta 0.0.0.8080
}
