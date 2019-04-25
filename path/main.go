package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
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
const (
	CONNECTIONSTRING = "mongodb://127.0.0.1:27017"
	DBNAME 			 = "vaptcha"
	COLLNAME         = "22"
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

type Track struct {
	//Id 	   objectid.ObjectID	`_id`
	Id 		string	  `form:"ip"	json:"ip"`	//	form:"ip"
	Char 	string	  `form:"char"	json:"char"`	// form:"char"
	Path 	string	  `form:"path" 	json:"path"`	// form:"path"
}

func main() {
	router := gin.Default()
	router.POST("/post", insertOne)
	router.Run()
}

func insertOne(c *gin.Context)  {
	/*ip := track.Id
	char := track.Char
	path := track.Path*/
	/*var track = Track{
	}*/
	var track  Track
	num,err := db.Collection("22").CountDocuments(c,bson.M{"char":track.Char})
	if (num >= 0 && num <4) && err == nil{
		c.Bind(&track)
		c.JSON(200,gin.H{
			"ip":track.Id,
			"char":track.Char,
			"path":track.Path,
			"status": "Document storage succeeded",
		})

		db.Collection(COLLNAME).InsertOne(c,bson.D{
			{"ip",track.Id},
			{"char",track.Char},
			{"path",track.Path},
		})
	}
	//{"ip", ip},
		//{"char", char},
		//{"path", path},
	//if err := c.ShouldBindJSON(&track);err == nil {
	//	c.JSON(200,gin.H{
	//		/*"ip":track.Id,
	//		"char":track.Char,
	//		"path":track.Path,*/
	//		"status": "Document storage succeeded",
	//	})
	//}
}