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
		4.限制同一字符只能存储3次
*/

var db_2 *mongo.Database

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
	db_2 = client.Database("vaptcha")//访问集合

}

func main() {
	router := gin.Default()
	router.POST("/post", insertOne_2)
	router.Run()
}

func insertOne_2(c *gin.Context)  {

	id := c.PostForm("ip")
	char := c.PostForm("char")
	path := c.PostForm("path")
	coll := db_2.Collection("11")

	//判断需要存入数据库的值是否存在，不存在，err为空；存在，err不为空
	/*resp,err := coll.Find(c,bson.D{
		{"char",bson.D{
			{"$exists",true}},
	}})*/

	num,err:=db_2.Collection("11").CountDocuments(c,bson.M{"char":char})
	if (num >=0 && num <4) && err == nil{
		coll.InsertOne(c,bson.D{
			{"ip", id},
			{"char", char},
			{"path", path},
		})
		c.JSON(200,gin.H{
			/*"ip":     id,
			"char":   char,
			"path":   path,*/
			"status": "Document storage succeeded",
		})
	}else{
		c.JSON(400,gin.H{
			"status": "More than three documents",
		})
	}
}
