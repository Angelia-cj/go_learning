package main

import "github.com/gin-gonic/gin"

/*
	请求方的method必须是post，Content-Type或者表单中要设成multipart/form-data。
	服务器端也要使用post方式接收
*/

func main() {
	router := gin.Default()
	router.POST("/upload", func(c *gin.Context) {
		file ,_ := c.FormFile("myfile")//获取文件
		filename := file.Filename
		size := file.Size
		header := file.Header
		c.JSON(200,gin.H{
			"filename":filename,
			"size":size,
			"header":header,
		})
	})
	router.Run()
}