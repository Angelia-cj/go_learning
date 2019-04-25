package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

//创建一个Trainer结构
type Trainer struct {
	Name	string
	Age 	int
	City	string
}

func main() {
/*
	//创建一个新的客户端
	client ,err := mongo.NewClient(options.Client().ApplyURI("momgo://127.0.0.1:27017"))
	if err != nil {
	   log.Fatal(err)
	}
	//设置超时时间
	ctx,cancle := context.WithTimeout(context.Background(),20*time.Second)
	if err != nil {
	   log.Fatal(err)
	}
	defer cancle()
	//创建连接,通过ctx字符串创建连接
	err = client.Connect(ctx)
	if err != nil {
	   log.Fatal(err)
	}
	*/
	//通过mongo.Connect()建立连接
	client,err := mongo.Connect(context.TODO(),options.Client().ApplyURI("mongodb://loaclhost:27017"))
	if err != nil {
	   log.Fatal(err)
	}
	//检查连接
	err = client.Ping(context.TODO(),nil)
	if err != nil {
	   log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	//连接test数据库和trainers集合
	collection := client.Database("test").Collection("trainers")


	//断开连接
	err = client.Disconnect(context.TODO())
	if err != nil {
	   log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed")


	//创建一些新的Trainers结构体的数据
	joe := Trainer{"Joe",10,"Chongqing"}
	jim := Trainer{"Jim",15,"Beijing"}
	ane := Trainer{"Ane",14,"Shanghai"}

	//向数据库插入一条数据
	insertone,err := collection.InsertOne(context.TODO(),joe)
	if err != nil {
	   log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ",insertone.InsertedID)

	//向数据库插入多条数据
	trainers := []interface{}{jim,ane}
	insertmany,err := collection.InsertMany(context.TODO(),trainers)
	if err != nil {
	   log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertmany.InsertedIDs)

	//过滤一条数据
	filter := bson.D{{"mame","Jim"}}

	//需要更新的数据
	update := bson.D{
		{"$inc",bson.D{
			{"age",1},
		}},
	}
	//更新一条数据
	updateResult,err := collection.UpdateOne(context.TODO(),filter,update)
	if err != nil {
	   log.Fatal(err)
	}
	fmt.Println("Matched %v documents and updated %v documents.\n",updateResult.MatchedCount,updateResult.ModifiedCount)


	//创建一个可以解码的结果值
	var result Trainer
	//查找一条数据，需要过滤文件（filter document）
	err = collection.FindOne(context.TODO(),filter).Decode(&result)
	if err != nil {
	   log.Fatal(err)
	}
	fmt.Println("Found a single document: %+v\n",result)

	//将这些options传递给Find方法
	findOptions := options.Find()
	findOptions.SetLimit(2)

	//建立一个数组，可以存储解码后的文件
	var results []*Trainer

	//通过nil值作为过滤器来筛选符合集合中的所有文档
	cur,err := collection.Find(context.TODO(),nil,findOptions)
	if err != nil {
	   log.Fatal(err)
	}

	//查找多个文档会返回一个游标
	//通过光标迭代可以让我们一次解码一个文档
	for cur.Next(context.TODO()){
		//创建一个单个文件就可以被解码的值
		var elem Trainer
		err := cur.Decode(&elem)
		if err != nil {
		   log.Fatal(err)
		}
		results = append(results,&elem)
	}
	if err := cur.Err(); err != nil {
	   log.Fatal(err)
	}

	//一旦结束关闭光标
	cur.Close(context.TODO())

	fmt.Println("Found multiple documents (array of pointers): %+v\n", results)

	//删除文件
	deleteResult,err := collection.DeleteMany(context.TODO(),nil)
	if err != nil {
	   log.Fatal(err)
	}
	fmt.Println("Deleted %v documents in the trainers collection\n",deleteResult.DeletedCount)
}
