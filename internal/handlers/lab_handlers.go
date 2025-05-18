package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

const labServiceURL = "http://localhost:8082" // URL вашего lab_service

func LabCreateHandler(c *gin.Context) {
	ProxyRequest(labServiceURL, "/labs", c)
}

func LabUpdateHandler(c *gin.Context) {
	ProxyRequest(labServiceURL, fmt.Sprintf("/labs/%s", c.Param("id")), c)
}

func LabDeleteHandler(c *gin.Context) {
	ProxyRequest(labServiceURL, fmt.Sprintf("/labs/%s/delete", c.Param("id")), c)
}

func LabGetHandler(c *gin.Context) {
	ProxyRequest(labServiceURL, fmt.Sprintf("/labs/%s", c.Param("id")), c)
}

func LabStartHandler(c *gin.Context) {
	ProxyRequest(labServiceURL, fmt.Sprintf("/labs/%s/start", c.Param("id")), c)
}

func LabStopHandler(c *gin.Context) {
	ProxyRequest(labServiceURL, fmt.Sprintf("/labs/%s/stop", c.Param("id")), c)
}

func LabExecuteCommandHandler(c *gin.Context) {
	ProxyRequest(labServiceURL, fmt.Sprintf("/labs/%s/execute-command", c.Param("id")), c)
}
func LabCommitLabHandler(c *gin.Context) {
	ProxyRequest(labServiceURL, fmt.Sprintf("/labs/%s/commit", c.Param("id")), c)
}
