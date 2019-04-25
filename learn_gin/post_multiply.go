package main

import "github.com/gin-gonic/gin"

/*

*/

func main() {
	router := gin.Default()
	router.POST("/upload", func(c *gin.Context) {
		file_1,_ := c.FormFile("file_1")
		file_2,_ := c.FormFile("file_2")
		file_3,_ := c.FormFile("file_3")

		c.IndentedJSON(200,gin.H{
			"file_1":file_1.Filename,
			"file_2":file_2.Filename,
			"file_3":file_3.Filename,
		})
	})
	router.Run()
}