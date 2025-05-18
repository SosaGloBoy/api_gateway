package routes

import (
	"api_gateway/internal/handlers"
	"api_gateway/internal/middleware"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.Use(middleware.CORS())
	router.Use(middleware.RemoveXFrameOptions())
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Wetty Proxy is running. Use /terminal to access the terminal.")
	})

	router.Any("/terminal/*path", func(c *gin.Context) {
		target, _ := url.Parse("http://localhost:61829")

		proxy := httputil.NewSingleHostReverseProxy(target)

		// Модифицируем запрос
		proxy.Director = func(req *http.Request) {
			req.URL.Scheme = target.Scheme
			req.URL.Host = target.Host
			req.URL.Path = c.Param("path") // Перенаправляем путь
			req.Host = target.Host

			// Важно для WebSockets
			if req.Header.Get("Upgrade") == "websocket" {
				req.Header.Set("Connection", "upgrade")
			}
		}

		// Модифицируем ответ
		proxy.ModifyResponse = func(resp *http.Response) error {
			// Удаляем заголовки, блокирующие iframe
			resp.Header.Del("X-Frame-Options")
			resp.Header.Del("Content-Security-Policy")

			// Разрешаем встраивание с любого источника
			resp.Header.Set("Access-Control-Allow-Origin", "*")
			return nil
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	})
	authGroup := router.Group("/auth")

	{
		authGroup.POST("/register", handlers.AuthRegisterHandler)
		authGroup.POST("/login", handlers.AuthLoginHandler)
	}

	labGroup := router.Group("/labs")
	{
		labGroup.POST("", handlers.LabCreateHandler)                             // Создание лаборатории
		labGroup.PUT("/:id", handlers.LabUpdateHandler)                          // Обновление лаборатории
		labGroup.POST("/:id/delete", handlers.LabDeleteHandler)                  // Удаление лаборатории
		labGroup.GET("/:id", handlers.LabGetHandler)                             // Получение лаборатории
		labGroup.POST("/:id/start", handlers.LabStartHandler)                    // Запуск лаборатории
		labGroup.POST("/:id/stop", handlers.LabStopHandler)                      // Остановка лаборатории
		labGroup.POST("/:id/execute-command", handlers.LabExecuteCommandHandler) // Выполнение команды
		labGroup.POST("/:id/commit", handlers.LabCommitLabHandler)
	}
	taskGroup := router.Group("/tasks")
	{
		taskGroup.POST("", handlers.TaskCreateHandler)
		taskGroup.GET("/:id", handlers.TaskGetByIDHandler)
		taskGroup.GET("", handlers.TaskGetAllHandler)
		taskGroup.PUT("/:id", handlers.TaskUpdateHandler)
		taskGroup.DELETE("/:id", handlers.TaskDeleteHandler)
	}
	sessionGroup := router.Group("/sessions")
	{
		sessionGroup.POST("", handlers.CreateSessionHandler) // Для создания сессии
		sessionGroup.GET("/user/:user_id", handlers.GetSessionByUserIDHandler)
		sessionGroup.GET("/:session_id", handlers.GetSessionByIDHandler)
		sessionGroup.POST("/:session_id/steps/:step_id/validate", handlers.ValidateStepHandler)
		sessionGroup.GET("/:session_id/steps/:step_id", handlers.GetStepStatusHandler)
		sessionGroup.POST("/:session_id/suspend", handlers.SuspendSessionHandler)
		sessionGroup.POST("/:session_id/complete", handlers.CompleteSessionHandler)
	}
	knowledgeGroup := router.Group("/knowledges")

	category := knowledgeGroup.Group("/categories")
	{
		category.POST("", handlers.CreateCategoryHandler)
		category.GET("/all", handlers.GetAllCategoriesHandler)
		category.GET("/:id", handlers.GetCategoryByIDHandler)
		category.DELETE("/:id", handlers.DeleteCategoryHandler)
	}

	article := knowledgeGroup.Group("/articles")
	{
		article.POST("", handlers.CreateArticleHandler)
		article.GET("", handlers.GetAllArticlesHandler)
		article.GET("/:id", handlers.GetArticleByIDHandler)
		article.DELETE("/:id", handlers.DeleteArticleHandler)
	}

}
