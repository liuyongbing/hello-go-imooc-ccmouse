package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"time":    time.Now(),
	})
}

func main() {
	r := gin.Default()

	r.GET("/ping", pong)

	r.Run(":8001")
}
