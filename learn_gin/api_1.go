package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	参数路径
	Handler：处理器
*/
func main() {

	router := gin.Default()

	//这个处理程序将匹配/user/john，但不匹配/user/或/user
	router.GET("/web/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK,"hello %s",name)
	})//先运行，然后在浏览器中输入  127.0.0.1:8080/web/api_1

	/*
		但是，这个将匹配/user/john/和/user/john/send
		如果没有其他路由匹配/user/john，它将重定向到/user/john/
	*/
	router.GET("/web/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + "is" + action

		c.String(http.StatusOK,message)
	})//先运行，然后在浏览器中输入  127.0.0.1:8080/web/api_1/

	router.Run(":8080")
}
