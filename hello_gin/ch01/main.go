package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func main() {
	r := gin.Default()

	r.GET("/ping", pong)

	r.Run(":8001")
}
