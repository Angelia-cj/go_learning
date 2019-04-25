package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
	表单和Body参数
	多部分/ Urlencoded形式
*/
const (
	CONNECTIONSTRING = "mongodb://127.0.0.1:27017"
	DBNAME 			 = "vaptcha"
	COLLNAME         = "33"
)
var db *mongo.Database

//打开数据库连接
func init()  {
	client,err := mongo.NewClient(options.Client().ApplyURI(CONNECTIONSTRING))//打开数据库
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())//通过context.Background()获取一个空的全局变量context
	if err != nil {
		log.Fatal(err)
	}
	db = client.Database(DBNAME)//访问集合
}

func main() {
	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		// 获取post过来的message内容
		// 获取的所有参数内容的类型都是 string
		message := c.PostForm("message")
		// 如果不存在，使用第二个当做默认内容
		nick := c.DefaultPostForm("nick","anonymous")

		c.JSON(200,gin.H{
			"status": "posted",
			"message": message,
			"nick":  nick,
		})
	})

	db.Collection(COLLNAME).InsertOne(context.TODO(),bson.M{"message": "message","nick": "nick"})


	router.Run(":8080")
}