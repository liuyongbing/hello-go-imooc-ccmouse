package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/get_params", getParams)
	router.POST("/post_params", postParams)
	router.POST("/get_and_post", getAndPost)

	router.Run(":8085")
}

// 获取参数： GET & POST
func getAndPost(c *gin.Context) {
	// GET Query
	id := c.Query("id")
	// Set default value
	page := c.DefaultQuery("page", "0")

	// POST Query
	name := c.PostForm("name")
	// Set default value
	message := c.DefaultPostForm("message", "信息")

	getParams := map[string]string{
		"get_id":   id,
		"get_page": page,
	}

	postParams := map[string]string{
		"name":    name,
		"message": message,
	}

	c.JSON(http.StatusOK, gin.H{
		"get_params":  getParams,
		"post_params": postParams,
		"time":        time.Now(),
	})
}

// 获取参数： POST
func postParams(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")
	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"nick":    nick,
	})
}

// 获取参数： GET
func getParams(c *gin.Context) {
	firstName := c.DefaultQuery("firstname", "bobby")
	lastName := c.DefaultQuery("lastname", "imooc")
	c.JSON(http.StatusOK, gin.H{
		"first_name": firstName,
		"last_name":  lastName,
	})
}
