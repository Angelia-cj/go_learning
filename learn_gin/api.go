package main

import "github.com/gin-gonic/gin"

/*
	API的例子：
		使用 GET, POST, PUT, PATCH, DELETE 和 OPTIONS
*/

func main() {

	/*
		使用gin的Default方法创新一个Handler路由，日志记录器和恢复(无崩溃)中间件
	*/
	router := gin.Default()

	router.GET("/someGet",getting)
	router.POST("/somePost",posting)
	router.PUT("/somePut",posting)
	router.DELETE("/someDelete",deleting)
	router.PATCH("/somePatch",patching)
	router.HEAD("/someHead",head)
	router.OPTIONS("/someOptions",options)

	/*
		使用默认的端口8080，除非自定义一个端口号
	*/
	router.Run()
	// router.Run(":3000") 一个指定的端口号
}
