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
	users := []interface{}{
		User{Name: "UserName_0", Phone: "123"},
		User{Name: "UserName_1", Phone: "456"},
		User{Name: "UserName_2", Phone: "789"},
	}

	db.Collection(COLLNAME).InsertMany(c,users)

	c.JSON(http.StatusCreated,gin.H{
		"status":	http.StatusCreated,
		"message":	"Todo item created successfully!",})
}

//GetOne：获的所有的todos
func GetOne(c *gin.Context)  {

	db.Collection(COLLNAME).FindOne(c,bson.M{"phone":"456"})

	c.JSON(http.StatusCreated,gin.H{
		"status":	http.StatusCreated,
		"message":	"Todo item created successfully!",})
}

//GetMany：获的一个的todo
func GetMany(c *gin.Context)  {
	db.Collection(COLLNAME).Find(c,bson.M{"phone": primitive.Regex{Pattern: "456", Options: ""}})

	c.JSON(http.StatusCreated,gin.H{
		"status":	http.StatusCreated,
		"message":	"Todo item created successfully!",})
}

//UpdateOne：更新一个todo
func UpdateOne(c *gin.Context)  {
	db.Collection(COLLNAME).UpdateOne(c,bson.M{"phone":"123"},
	bson.M{"$set":bson.M{"dbname":"UserName_changed"}})
}

//DeleteOne：删除一个todo
func DeleteOne(c *gin.Context)  {
	db.Collection(COLLNAME).DeleteOne(c,bson.M{"phone": primitive.Regex{Pattern: "456", Options: ""}})
}



















