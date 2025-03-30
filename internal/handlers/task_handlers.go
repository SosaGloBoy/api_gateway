package handlers

import "github.com/gin-gonic/gin"

const taskServiceURL = "http://localhost:8086"

func TaskCreateHandler(c *gin.Context) {
	ProxyRequest(taskServiceURL, "/tasks", c)
}

// TaskGetByIDHandler для получения задания по ID
func TaskGetByIDHandler(c *gin.Context) {
	id := c.Param("id")
	ProxyRequest(taskServiceURL, "/tasks/"+id, c)
}

// TaskGetAllHandler для получения всех заданий
func TaskGetAllHandler(c *gin.Context) {
	ProxyRequest(taskServiceURL, "/tasks", c)
}

// TaskUpdateHandler для обновления задания
func TaskUpdateHandler(c *gin.Context) {
	id := c.Param("id")
	ProxyRequest(taskServiceURL, "/tasks/"+id, c)
}

// TaskDeleteHandler для удаления задания
func TaskDeleteHandler(c *gin.Context) {
	id := c.Param("id")
	ProxyRequest(taskServiceURL, "/tasks/"+id, c)
}
