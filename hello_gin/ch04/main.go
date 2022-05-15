package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	ID   int    `uri:"id" binding:"required"`
	Name string `uri:"name" binding:"required"`
}

func PersonDetail(c *gin.Context) {

	var person Person
	if err := c.ShouldBindUri(&person); err != nil {
		c.Status(404)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"router params": "/:name/:id",
		"name":          person.Name,
		"id":            person.ID,
		"time":          time.Now(),
	})
}

func main() {
	router := gin.Default()
	router.GET("/:name/:id", PersonDetail)

	router.Run(":8084")
}
