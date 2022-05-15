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
	router := gin.Default()

	goodsGroup := router.Group("/goods")
	{
		goodsGroup.GET("", goodsList)
		goodsGroup.GET("/:id/:action/add", goodsDetail) //获取商品id为1的详细信息 模式
		goodsGroup.POST("", createGoods)
	}

	router.Run(":8003")
}

func createGoods(c *gin.Context) {

}

func goodsDetail(c *gin.Context) {
	id := c.Param("id")
	action := c.Param("action")
	c.JSON(http.StatusOK, gin.H{
		"Router group": "goodsDetail",
		"id":           id,
		"action":       action,
	})
}

func goodsList(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"Router group": "goodsList",
	})
}
