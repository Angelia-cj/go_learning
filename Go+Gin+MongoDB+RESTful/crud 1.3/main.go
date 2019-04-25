package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

/*
	语言：Go
	数据库：MongoDB
	数据库驱动：mongo-go-driver
	框架：gin-gonic
*/


/*
	数据库操作
	1、创建连接
	2、创建数据库
	3、创建集合
	4、初始化一个对象
*/
const (
	CONNECTIONSTRING = "mongodb://127.0.0.1:27017"
	DBNAME 			 = "users"
	COLLNAME         = "people"
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

	v1 := router.Group("/api/v1/todos")
	{
		v1.POST("/",InsertMany)
		v1.GET("/",GetOne)
		v1.GET("/:id",GetMany)
		v1.PUT("/:id",UpdateOne)
		v1.DELETE("/:id",DeleteOne)
	}
	router.Run()
}


type User struct {
	ID    primitive.ObjectID "_id,omitempty"
	Name  string            `bson:"dbname",json:"jsonname"`
	Phone string
}

//InsertMany：增加一个新的todo
func InsertMany(c *gin.Context)  {
	//user := User{Name:"UserName", Phone:"18738485362",}
	userColl := db.Collection(COLLNAME)
	u1 := User{Name: "UserName_0", Phone: "123"}
	u2 := User{Name: "UserName_1", Phone: "456"}
	u3 := User{Name: "UserName_2", Phone: "789"}

	users := []interface{}{u1,u2,u3}
	reslut ,err := userColl.InsertMany(c,users)
	if err != nil{
		c.JSON(http.StatusNotFound,gin.H{
			"status":http.StatusNotFound,
			"message":"Data Not Found!"})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"status":http.StatusOK,
			"result": reslut.InsertedIDs})
	}
}

//GetOne：获的一个的todos
func GetOne(c *gin.Context)  {
	userColl := db.Collection(COLLNAME)
	result := userColl.FindOne(c,bson.M{"phone":"456"})
	var user User
	c.Bind(&user)
	if err := result.Decode(&user);err != nil  {
		c.JSON(http.StatusNotFound,gin.H{"status":http.StatusNotFound,"message":"Data not Found!"})
	}else {
		c.JSON(200,gin.H{
			"status":http.StatusOK,
			"result":	user})
	}

}

//GetMany：获的一个的todo
func GetMany(c *gin.Context)  {
	userColl := db.Collection(COLLNAME)
	cur,err := userColl.Find(c,bson.M{"phone":primitive.Regex{Pattern:"456",Options:""}})
	if err != nil {
		c.JSON(http.StatusNotFound,gin.H{"status":http.StatusNotFound,"message":"Data not Found!"})
	}
	defer cur.Close(c)
	for cur.Next(c) {
		var user User
		if err := cur.Decode(&user);err != nil {
			c.JSON(http.StatusUnauthorized,gin.H{"status":http.StatusUnauthorized,"message":"Data not Found!"})
		}
		c.JSON(http.StatusOK,gin.H{
			"status":	http.StatusOK,
			"result":	user})
	}
}

//UpdateOne：更新一个todo
func UpdateOne(c *gin.Context)  {
	userColl := db.Collection(COLLNAME)
	if result,err := userColl.UpdateOne(
		c,bson.M{"phone":"123"},
		bson.M{
			"$set":bson.M{"dbname":"UserName_changjia"}});err == nil{
		c.JSON(http.StatusOK,gin.H{
			"status":	http.StatusOK,
			"result":	result})
	}else {
		c.JSON(http.StatusNotFound,gin.H{"status":http.StatusNotFound,"message":"Data not Found!"})
	}
}

//DeleteOne：删除一个todo
func DeleteOne(c *gin.Context)  {
	userColl := db.Collection(COLLNAME)
	if result,err := userColl.DeleteOne(c,bson.M{"phone":"123"});err == nil{
		c.JSON(http.StatusOK,gin.H{
			"status":	http.StatusOK,
			"result":	result})
	}else {
		c.JSON(http.StatusNotFound,gin.H{"status":http.StatusNotFound,"message":"Data not Found!"})
	}
}
