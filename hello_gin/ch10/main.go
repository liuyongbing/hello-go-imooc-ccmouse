package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func Graceful(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Graceful restart or stop",
		"time":    time.Now(),
	})
}

func main() {
	router := gin.Default()

	router.GET("graceful", Graceful)

	go func() {
		router.Run(":8090")
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<--quit

	fmt.Println("Server关闭中。。。")
	fmt.Println("服务注销中。。。")
}
