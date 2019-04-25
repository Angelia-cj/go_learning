package main
import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)
/*
	以请求的ip为主键，存储 char ，path ；
	同一个ip同一个字符只记录3条记录，保存到mongo
*/

var db_3 *mongo.Database

//打开数据库连接
func init()  {
	client,err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))//打开数据库
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())//通过context.Background()获取一个空的全局变量context
	if err != nil {
		log.Fatal(err)
	}
	db_3 = client.Database("vaptcha")//访问集合
}

func main() {
	router := gin.Default()
	router.POST("/post", insertOne_3)
	router.Run()
}

func insertOne_3(c *gin.Context)  {
	id := c.PostForm("ip")
	char := c.PostForm("char")
	path := c.PostForm("path")

	db_3.Collection("11").InsertOne(c,bson.D{
		{"ip",id},
		{"char",char},
		{"path",path},
	})

	c.JSON(200,gin.H{
		"message":"Document storage succeeded",
	})
}

