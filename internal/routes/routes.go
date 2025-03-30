package routes

import (
	"api_gateway/internal/handlers"
	"api_gateway/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.Use(middleware.CORS())

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", handlers.AuthRegisterHandler)
		authGroup.POST("/login", handlers.AuthLoginHandler)
	}

	labGroup := router.Group("/labs")
	{
		labGroup.POST("", handlers.LabCreateHandler)                             // Создание лаборатории
		labGroup.PUT("/:id", handlers.LabUpdateHandler)                          // Обновление лаборатории
		labGroup.DELETE("/:id", handlers.LabDeleteHandler)                       // Удаление лаборатории
		labGroup.GET("/:id", handlers.LabGetHandler)                             // Получение лаборатории
		labGroup.POST("/:id/start", handlers.LabStartHandler)                    // Запуск лаборатории
		labGroup.POST("/:id/stop", handlers.LabStopHandler)                      // Остановка лаборатории
		labGroup.POST("/:id/execute-command", handlers.LabExecuteCommandHandler) // Выполнение команды
	}
	taskGroup := router.Group("/tasks")
	{
		taskGroup.POST("", handlers.TaskCreateHandler)
		taskGroup.GET("/:id", handlers.TaskGetByIDHandler)
		taskGroup.GET("", handlers.TaskGetAllHandler)
		taskGroup.PUT("/:id", handlers.TaskUpdateHandler)
		taskGroup.DELETE("/:id", handlers.TaskDeleteHandler)
	}
}
