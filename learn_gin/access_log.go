package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	//取消控制台中日志的字体颜色
	gin.DisableConsoleColor()

	//创建一个日志文件
	access_log,_ := os.Create("access_log.log")
	gin.DefaultWriter = io.MultiWriter(access_log)

	router := gin.Default()
	router.GET("/log", func(c *gin.Context) {
		c.JSON(200,gin.H{"message":"hello world"})
	})
	router.Run()
}

