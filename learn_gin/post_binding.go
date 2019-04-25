package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if username == "abc" && password == "123"{
			c.JSON(200,gin.H{"message":"welcome"})
		}else {
			c.JSON(401,gin.H{"message": "wrong username or password"})
		}
	})
	router.Run()
}
