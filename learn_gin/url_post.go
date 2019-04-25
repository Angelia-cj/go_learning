package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default() //服务器端必须使用router.POST方式来接收，
	// 不能使用router.GET方式接收，因为router.GET只能接收GET方式传递的数据
	router.POST("/post", func(c *gin.Context) {

		name := c.PostForm("name")
		age  := c.PostForm("age")
		sex  := c.DefaultQuery("sex","male")
		addr := c.Query("addr")
		hobby := c.DefaultQuery("hobby","basketball")

		c.JSON(200,gin.H{
			"name":name,
			"age":age,
			"sex":sex,
			"addr":addr,
			"hobby":hobby,
		})
	})
	router.Run()
}