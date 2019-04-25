package main

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	db			="ToDo"	//数据库
	collection	="ToDoList"	//集合
	host		="127.0.0.1:27017"	//主机地址
)



//构造一个TODO结构体
type (
	todoModel struct {
		Title		string		`bson:"title" json:"title"`
		Completed	int		    `bson:"completed" json:"completed"`
		CreatedAt	time.Time	`bson:"creatAt" json:"creatAt"`
	}
	transformedTodo struct {
		Id    		string 		`"bson":"_id" json:"id"`
		Title		string 		`"bson":"title" json:"title"`
		Completed	bool		`bson:"completed" json:"completed"`
		CreatedAt	time.Time	`bson:"creatAt" json:"creatAt"`
	}
)

//初始化
var globalS *mgo.Session
//连接数据库
func init()  {
	globalS,err := mgo.Dial(host)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	//创建路由，以及路由的初始化
	router := gin.Default()

	//创建一个路由组v1
	v1 := router.Group("api/v1/todos")
	{
		v1.POST("/", createTodo)
		v1.GET("/", feathAllTodo)
		v1.GET("/:id", fetchSingleTodo)
		v1.PUT("/:id", updateTodo)
		v1.DELETE("/:id", deleteTodo)
	}
	router.Run()
}
//createTodo路由方法
func createTodo(c *gin.Context)  {
	completed,_ := strconv.Atoi(c.PostForm("completed"))
	todo := todoModel{
		ID: 	bson.NewObjectId().Hex(),
		Title:	c.PostForm("title"),
		Completed: completed,
		CreatedAt:time.Now(),
	}
	ms := globalS.Copy() //每一次操作都copy一份 Session,避免每次创建Session,导致连接数量超过设置的最大值
	mc := ms.DB(db).C(collection)//获取文档对象 c := Session.DB(db).C(collection)
	defer ms.Close()
	mc.Insert(todo)
	//从gin的上下文获取post的数据，并保存数据到数据库，如果成功，返回对应的id
	c.JSON(http.StatusCreated,gin.H{
		"status": http.StatusCreated,
		"message":"Todo item created successfully!",
		"resourceId":todo.ID,
	})
}

//feathAllTodo路由方法
func feathAllTodo(c *gin.Context) {
	var todos []todoModel
	var _todos []transformedTodo

	ms := globalS.Copy()
	mc := ms.DB(db).C(collection)
	defer ms.Close()
	mc.Find(nil).All(&todos)

	if len(todos) <= 0 {
		c.JSON(http.StatusNotFound,gin.H{
			"ststus":http.StatusNotFound,
			"message":"No todo found!",
		})
		return
	}
	for _,item := range todos{
		completed := false
		if item.Completed == 1{
			completed = true
		}else {
			completed = false
		}
		_todos = append(_todos,transformedTodo{ID:item.ID,Title:item.Title,Completed:completed,CreatedAt:time.Now()})
	}
	c.JSON(http.StatusOK,gin.H{
		"status":http.StatusOK,
		"data":_todos})
}

//fetchSingleTodo路由方法
func fetchSingleTodo(c *gin.Context)  {
	var todo todoModel
	id := c.Param("id")
	ms := globalS.Copy()
	mc := ms.DB(db).C(collection)
	mc.FindId(id)

	if todo.ID == "" {
		c.JSON(http.StatusNotFound,gin.H{
			"status": http.StatusNotFound,
			"message":"No todo found!"})
		return
	}
	completed := false
	if todo.Completed == 1{
		completed = true
	}else {
		completed = false
	}
	_todo := transformedTodo{ID:todo.ID,Title:todo.Title,Completed:completed,CreatedAt:todo.CreatedAt}
	c.JSON(http.StatusOK,gin.H{
		"status":http.StatusOK,
		"data":_todo})
}

//updateTodo路由方法
func updateTodo(c *gin.Context)  {
	var todo todoModel
	id := c.Param("id")
	ms := globalS.Copy()
	mc := ms.DB(db).C(collection)
	defer ms.Close()
	mc.FindId(id).One(&todo)

	if todo.ID == "" {
		c.JSON(http.StatusNotFound,gin.H{"status":http.StatusNotFound,"message":"No todo Found!"})
		return
	}
	todo.Title = c.Param("title")
	todo.Completed,_ = strconv.Atoi(c.PostForm("completed"))

	mc.UpdateId(id,todo)
	c.JSON(http.StatusOK,gin.H{"status":http.StatusOK,"message":"Todo update successfully!"})
}

//deleteTodo路由方法
func deleteTodo(c *gin.Context)  {
	var todo todoModel
	id := c.Param("id")
	ms := globalS.Copy()
	mc := ms.DB(db).C(collection)
	defer ms.Close()
	mc.FindId(id).One(&todo)
	if todo.ID == ""{
		c.JSON(http.StatusNotFound,gin.H{"status":http.StatusNotFound,"message":"No todo found!"})
		return
	}
	mc.RemoveId(id)
	c.JSON(http.StatusOK,gin.H{"status":http.StatusOK,"message":"Todo deleted successfully!!"})
}
