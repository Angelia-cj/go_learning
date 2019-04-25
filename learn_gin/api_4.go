package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
/*
    POST   /post/test?id=1234&page=1  HTTP/1.1
    请求头:  Content-Type: application/x-www-form-urlencoded
    form表单参数:  name=manu&message=this_is_great
*/
func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page","0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id : %s; page: %s; name: %s; message: %s", id, page, name, message)
	})
	router.Run(":8080")
}