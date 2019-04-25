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
	要求：以请求的ip为主键，存储 char ，path ；同一个ip同一个字符只记录3条记录，保存到mongo、
	实现功能：
		1.接收到前端的数据
		2.将前端的数据存到数据库中
		3.将接收成功的信息返回给前端
*/

/*const (
	CONNECTIONSTRING = "mongodb://127.0.0.1:27017"
	DBNAME 			 = "vaptcha"
	COLLNAME         = "22"
)*/
var db_1 *mongo.Database

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
	db_1 = client.Database("vaptcha")//访问集合
}

/*type Track struct {
	//Id 	   objectid.ObjectID	`_id`
	Id 		string	  `json:"ip"`	//	form:"ip"
	Char 	string	  `json:"char"`	// form:"char"
	Path 	string	  `json:"path"`	// form:"path"
}*/

func main() {
	router := gin.Default()
	router.POST("/post", insertOne_1)
	router.Run()
}

func insertOne_1(c *gin.Context)  {
	//var track Track
	/*var track = Track{
	}*/
	//id := c.PostForm("id")
	ip := c.PostForm("ip")
	char := c.PostForm("char")
	path := c.PostForm("path")

	db_1.Collection("22").InsertOne(c,bson.D{
		//{"id", id},
		{"ip", ip},
		{"char", char},
		{"path", path},
	})
	//print("123")
	c.JSON(200,gin.H{
		//"id":id,
		"ip":ip,
		"char":char,
		"path":path,
		"status": "Document storage succeeded",
	})
}
