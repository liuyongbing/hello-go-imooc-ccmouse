package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func getting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "someGet",
		"time":    time.Now(),
	})
}

func main() {
	// 使用默认中间件创建一个gin路由器
	// logger and recovery (crash-free) 中间件
	router := gin.Default()

	//restful 的开发中
	router.GET("/someGet", getting)
	// router.POST("/somePost", posting)
	// router.PUT("/somePut", putting)
	// router.DELETE("/someDelete", deleting)
	// router.PATCH("/somePatch", patching)
	// router.HEAD("/someHead", head)
	// router.OPTIONS("/someOptions", options)

	// 默认启动的是 8080端口，也可以自己定义启动端口
	router.Run(":8002")
	// router.Run(":3000") for a hard coded port
}
