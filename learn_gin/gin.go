package main

import "github.com/gin-gonic/gin" //导入gin包

/*
	简单的gin框架
*/

func main(){

	//创建一个路由
	router := gin.Default()	//使用gin的Default方法创建一个路由handler
	router.GET("/ping", func(c *gin.Context) {  //通过HTTP方法绑定路由规则和路由参数
							c.JSON(200,gin.H{   //通过JSON方法来返回参数
								"message":"pong",
							})
	})
	router.Run() //启用路由的Run方法监听路由：0.0.0.0:8080。默认绑定8080
}

//先运行，然后在浏览器中输入  127.0.0.1:8080/ping