package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	User	string		`form:"user" json:"user" binding:"required"`
	Password string		`form:"password" json:"password" binding:required`
}

func main() {
	router:= gin.Default()

	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json);err == nil{
			if json.User == "abc" && json.Password == "123"{
				c.JSON(http.StatusOK,gin.H{"status": "you are logged in"})
			}else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		}else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		if err := c.Bind(&form);err == nil{
			if form.User == "abc" && form.Password == "123"{
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			}else{
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		}else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
	router.Run()
}