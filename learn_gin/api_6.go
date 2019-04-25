package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

/*
	上传文件

	文件名始终是可选的，不得由应用程序盲目使用：应删除路径信息，并应转换为服务器文件系统规则。
*/

func main() {
	router := gin.Default()
	//为多部分表单设置较低的内存限制（默认值为32 MiB）
	// router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		//单个文件
		file,_ := c.FormFile("file")
		log.Println(file.Filename)

		//将文件上传到特定的dst。
		// c.SaveUploadedFile（file，dst）
		c.String(http.StatusOK,fmt.Sprintf(" '%s' uploaded!",file.Filename))
	})
	router.Run(":8080")
}