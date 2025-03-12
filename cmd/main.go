package main

import (
	"api_gateway/internal/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	// Регистрируем маршруты API Gateway
	routes.SetupRoutes(router)

	// Запуск API Gateway
	log.Println("API Gateway running on port 8083")
	if err := router.Run(":8083"); err != nil {
		log.Fatal(err)
	}
}
