package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ProxyRequest(target, path string, c *gin.Context) {
	targetURL, err := url.Parse(target)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid target URL"})
		return
	}
	c.Request.URL.Path = path

	fmt.Println("Proxying request to:", targetURL.String()+path)

	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	fmt.Printf("Proxying: %s%s → %s%s\n",
		c.Request.Host, c.Request.URL.Path,
		target, path)

	proxy.ServeHTTP(c.Writer, c.Request)
}
