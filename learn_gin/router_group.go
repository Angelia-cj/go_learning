package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/demo", func(c *gin.Context) {
			c.JSON(200,gin.H{"data":"/v1/demo"})
		})
		v1.GET("/test", func(c *gin.Context) {
			c.JSON(200,gin.H{"data":"/v1/test"})
		})
	}
	v2 := router.Group("/v2")
	{
		v2.GET("/demo", func(c *gin.Context) {
			c.JSON(200,gin.H{"data":"/v1/demo"})
		})
		v2.GET("/test", func(c *gin.Context) {
			c.JSON(200,gin.H{"data":"/v1/test"})
		})
	}
	router.Run()
}