package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type User struct {
	Name string
	Age  int
	City string
}

func main()  {
	//设置客户端选项
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")

	//建立MongoDB数据库的连接
	client,err := mongo.Connect(context.TODO(),clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	//检测连接是否成功
	err = client.Ping(context.TODO(),nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	//给集合找一个处理
	collection := client.Database("test").Collection("users")

	//添加一些虚拟数据到数据库中
	joe := User{"Joe",15,"重庆"}
	ane := User{"Ane",18,"北京"}
	jim := User{"Jim",16,"上海"}

	//向数据库中插入一条数据
	insertResult,err := collection.InsertOne(context.TODO(),joe)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	//向数据库中插入多条数据
	insertManyResult,err := collection.InsertMany(context.TODO(),ane.jim)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	//更新一条数据
	filter := bson.D{{"name","joe"}}
	update := bson.D{
		{"$inc",bson.D{
			{"age",12},
		}},
	}
	updateResult,err := collection.UpdateOne(context.TODO(),filter,update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n",updateResult.MatchedCount,updateResult.ModifiedCount)

	//查找一条数据
	var result User

	err = collection.FindOne(context.TODO(),filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n",result)

	findOptions := options.Find()
	findOptions.SetLimit(i:2)

	var results []*User

	//查找多个文档将返回一个下标
	cur,err := collection.Find(context.TODO(),bson.D{{}},findOptions)
	if err != nil {
		log.Fatal(err)
	}

	//遍历下标
	for cur.Next(context.TODO()){
		var elem User
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results,&elem)
	}
	if err := cur.Err();err != nil{
		log.Fatal(err)
	}
	// 完成后关闭光标
	cur.Close(context.TODO())

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)

	// 删除集合中的所有文档
	deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

	// 一旦不再需要，关闭连接
	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection to MongoDB closed.")
	}


	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/",)
		v1.GET("/",)
		v1.PUT("/",)
		v1.DELETE("/",deleteResult)
	}
	router.Run()
}


