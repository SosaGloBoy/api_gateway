package middleware

import "github.com/gin-gonic/gin"

func RemoveXFrameOptions() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// После того, как прокси/обработчики вернули ответ
		c.Writer.Header().Del("X-Frame-Options")
		c.Writer.Header().Set("Content-Security-Policy", "frame-ancestors *")
	}
}
