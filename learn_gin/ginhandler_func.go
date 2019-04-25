package main

import "github.com/gin-gonic/gin"

type User struct{
	User	    string 		`form:"user"  json:"user"   binding:"required"`
	Password	string		`form:"password"  json:"password"  binding:"required"`
}

//声明一个gin.HandlerFunc
func LoginCheck(c *gin.Context){
	var json User
	var err error

	if err := c.ShouldBindQuery(&json);err == nil{ //尝试从get请求中获取参数
		if json.User == "abc" && json.Password == "123"{
			c.JSON(200,gin.H{"message":"login in success"})
		}else {
			c.JSON(200,gin.H{"message":"login in failed"})
		}
	}else if err := c.ShouldBind(&json);err ==nil{//尝试从post请求中获取参数
		if json.User == "abc" && json.Password == "123"{
			c.JSON(200,gin.H{"message":"login in success"})
		}else{
			c.JSON(200,gin.H{"message":"login in failed"})
		}
	}
	if err != nil{  //解析请求中的参数失败
		c.JSON(400,gin.H{"message":err.Error()})
	}
}

func main() {
	router := gin.Default()
	router.GET("login",LoginCheck)
	router.POST("login",LoginCheck)
	router.Run()
}
