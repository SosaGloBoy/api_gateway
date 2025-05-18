package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

const knowledgeServiceURL = "http://localhost:8091"

func CreateCategoryHandler(c *gin.Context) {
	ProxyRequest(knowledgeServiceURL, "/knowledges/categories/", c)
}

func GetAllCategoriesHandler(c *gin.Context) {
	ProxyRequest(knowledgeServiceURL, "/knowledges/categories/", c)
}

func GetCategoryByIDHandler(c *gin.Context) {
	categoryID := c.Param("id")
	ProxyRequest(knowledgeServiceURL, fmt.Sprintf("/knowledges/categories/%s", categoryID), c)
}

func DeleteCategoryHandler(c *gin.Context) {
	categoryID := c.Param("id")
	ProxyRequest(knowledgeServiceURL, fmt.Sprintf("/knowledges/categories/%s", categoryID), c)
}

func CreateArticleHandler(c *gin.Context) {
	ProxyRequest(knowledgeServiceURL, "/knowledges/articles/", c)
}

func GetAllArticlesHandler(c *gin.Context) {
	ProxyRequest(knowledgeServiceURL, "/knowledges/articles/", c)
}

func GetArticleByIDHandler(c *gin.Context) {
	articleID := c.Param("id")
	ProxyRequest(knowledgeServiceURL, fmt.Sprintf("/knowledges/articles/%s", articleID), c)
}

func DeleteArticleHandler(c *gin.Context) {
	articleID := c.Param("id")
	ProxyRequest(knowledgeServiceURL, fmt.Sprintf("/knowledges/articles/%s", articleID), c)
}
