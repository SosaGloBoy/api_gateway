package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const wettyURL = "http://localhost:61829"

func WettyProxyHandler(c *gin.Context) {
	targetURL, err := url.Parse(wettyURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid target URL"})
		return
	}

	// Используем оригинальный путь без изменений
	c.Request.Host = targetURL.Host

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
