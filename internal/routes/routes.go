package routes

import (
	"api_gateway/internal/handlers"
	"api_gateway/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.Use(middleware.CORS())
	router.Use(middleware.AuthMiddleware())

	router.POST("/register", handlers.AuthRegisterHandler)
	router.POST("/login", handlers.AuthLoginHandler)

}
