package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)


func main() {

	//首先，mongo-go-driver驱动程序的基本用法是先创建客户端并通过字符串建立连接
	//为此，请调用NewClient和Connect函数
	client,err := mongo.NewClient(options.Client().ApplyURI("mongo://127.0.0.1:27017"))//创建客户端
	if err != nil {
		log.Fatal(err)
	}
	//按照参数设置的超时时间结束，cancle方法，是context.WithTimeout()对整个goroutine结束的操作
	ctx,cancle := context.WithTimeout(context.Background(),20*time.Second)
	defer cancle()
	err = client.Connect(ctx)//通过ctx字符串创建连接
	if err != nil {
		log.Fatal(err)
	}
	/*
		这将创建一个新客户端并开始通过localhost监视MongoDB服务器。
		Database和Collection类型可用于访问数据库

		使用client客户端创建数据库的集合
	*/
 	collection := client.Database("test").Collection("abc")

   // 使用集合来查询数据库或插入文档
   res,err := collection.InsertOne(context.Background(),bson.M{"hello":"world"})
	if err != nil {
		log.Fatal(err)
	}
   id := res.InsertedID

	//下面的一些方法需要返回一个下标
	cur,err := collection.Find(context.Background(),bson.D{ })
	if err != nil {
	   log.Fatal(err)
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()){
		raw := rawType{}
		if err := cur.Decode(&raw); err != nil {
		   log.Fatal(err)
		}
		if err := cur.Err();err != nil {
			log.Fatal(err)
		}
	}

}