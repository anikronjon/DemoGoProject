package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ok":   true,
			"app":  "railway-gin-test",
			"port": os.Getenv("PORT"),
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "healthy")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Railway expects your app to listen on 0.0.0.0:<PORT>
	_ = r.Run("0.0.0.0:" + port)
}