package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

const sessionServiceURL = "http://localhost:8087"

func CreateSessionHandler(c *gin.Context) {
	ProxyRequest(sessionServiceURL, "/sessions", c)
}

func CompleteSessionHandler(c *gin.Context) {
	sessionID := c.Param("session_id")
	ProxyRequest(sessionServiceURL, fmt.Sprintf("/sessions/%s/complete", sessionID), c)
}

func GetSessionByUserIDHandler(c *gin.Context) {
	userID := c.Param("user_id")
	ProxyRequest(sessionServiceURL, fmt.Sprintf("/sessions/user/%s", userID), c)
}
func GetSessionByIDHandler(c *gin.Context) {
	sessionID := c.Param("session_id")
	ProxyRequest(sessionServiceURL, fmt.Sprintf("/sessions/%s", sessionID), c)

}
func ValidateStepHandler(c *gin.Context) {
	sessionID := c.Param("session_id")
	stepID := c.Param("step_id")
	ProxyRequest(sessionServiceURL, fmt.Sprintf("/sessions/%s/steps/%s/validate", sessionID, stepID), c)
}

func GetStepStatusHandler(c *gin.Context) {
	sessionID := c.Param("session_id")
	stepID := c.Param("step_id")
	ProxyRequest(sessionServiceURL, fmt.Sprintf("/sessions/%s/steps/%s", sessionID, stepID), c)
}
func SuspendSessionHandler(c *gin.Context) {
	sessionID := c.Param("session_id")
	ProxyRequest(sessionServiceURL, fmt.Sprintf("/sessions/%s/suspend", sessionID), c)
}
