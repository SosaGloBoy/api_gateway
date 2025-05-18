package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func WettyRequest(target string, c *gin.Context) {
	targetURL, err := url.Parse(target)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid target URL"})
		return
	}

	// Подменяем путь (важно для wetty, если ты обращаешься к /wetty)
	c.Request.URL.Path = c.Param("path")
	c.Request.Host = targetURL.Host

	// Создаём reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	// Удаляем заголовки, мешающие iframe
	proxy.ModifyResponse = func(resp *http.Response) error {
		resp.Header.Del("X-Frame-Options")
		resp.Header.Del("Content-Security-Policy")
		resp.Header.Set("Content-Security-Policy", "frame-ancestors *")
		return nil
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}
