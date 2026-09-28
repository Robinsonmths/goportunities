package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//inicializa o Router utilizando as configurações Default do gin
	router := gin.Default()
	// Definindo uma rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})
	// rodando a api
	router.Run(":8080") // padrão roda na porta 0.0.0.8080
}
