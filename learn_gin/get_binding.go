package main

import "github.com/gin-gonic/gin"

type Login_1 struct {
	User1 	    `form:"user" json:"user" binding:"required"`
	Password1 	`form:"password" json:"password" binding:"required"`
}
func main() {
	router := gin.Default()

	router.GET("/login", func(c *gin.Context) {
		var json Login_1
		if err := c.ShouldBindQuery(&json); err == nil{
			if json.User1 == "abc" && json.Password1 == "123"{
				c.JSON(200,gin.H{"message":"welcome"})
			}else {
				c.JSON(401,gin.H{"message":"wrong username or password"})
			}
		}else{
			c.JSON(400,gin.H{"error":err.Error()})
		}
	})
	router.Run()
}

