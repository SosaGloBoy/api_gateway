package handlers

import "github.com/gin-gonic/gin"

const authServiceURL = "http://localhost:8081"

func AuthRegisterHandler(c *gin.Context) {
	ProxyRequest(authServiceURL, "/register", c)
}

func AuthLoginHandler(c *gin.Context) {
	ProxyRequest(authServiceURL, "/login", c)
}
