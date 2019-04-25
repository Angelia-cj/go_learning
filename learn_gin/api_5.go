package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/*
	映射为querystring或postform参数

	POST ：/post?ids[a]=1234&ids[b]=hello HTTP/1.1
	请求头：Content-Type: application/x-www-form-urlencoded
	form表单参数: names[first]=thinkerou&names[second]=tianou
*/

func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		fmt.Printf("ids: %v ; names: %v ",ids,names)
	})
	router.Run(":8080")
}