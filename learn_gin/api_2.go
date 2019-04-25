package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	查询字符串参数
*/

func main() {
	router := gin.Default()

	/*
		使用现有的底层请求对象解析查询字符串参数。
		请求响应一个url匹配:/welcome?firstname=Jane&lastname=Doe
	*/
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname","Guest")
		lastname := c.Query("lastname")

		c.String(http.StatusOK,"hello %s %s",firstname,lastname)
	})
	router.Run(":8080")
	//先运行，然后在浏览器中输入http://127.0.0.1:8080/welcome?firstname=Jane&lastname=Doe
}