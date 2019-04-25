package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

/*
	示例代码网址：https://docs.mongodb.com/manual/tutorial/remove-documents/
*/

func main() {

	//创建客户端
	client ,err := mongo.NewClient(options.Client().ApplyURI("mongodb:127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}
	//ctx是最大的响应时间，达到规定的时间，由cancle取消
	ctx ,cancle := context.WithTimeout(context.TODO(),10*time.Second)
	defer cancle()

	//建立连接
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//cloth：布，canvas：油画
	collection := client.Database("test").Collection("cloth")

	docs := []interface{}{
		bson.D{
			{"item", "journal"},
			{"qty", 25},
			{"size", bson.D{
				{"h", 14},
				{"w", 21},
				{"uom", "cm"},
			}},
			{"status", "A"},
		},
		bson.D{
			{"item", "notebook"},
			{"qty", 50},
			{"size", bson.D{
				{"h", 8.5},
				{"w", 11},
				{"uom", "in"},
			}},
			{"status", "P"},
		},
		bson.D{
			{"item", "paper"},
			{"qty", 100},
			{"size", bson.D{
				{"h", 8.5},
				{"w", 11},
				{"uom", "in"},
			}},
			{"status", "D"},
		},
		bson.D{
			{"item", "planner"},
			{"qty", 75},
			{"size", bson.D{
				{"h", 22.85},
				{"w", 30},
				{"uom", "cm"},
			}},
			{"status", "D"},
		},
		bson.D{
			{"item", "postcard"},
			{"qty", 45},
			{"size", bson.D{
				{"h", 10},
				{"w", 15.25},
				{"uom", "cm"},
			}},
			{"status", "A"},
		},
	}

	result, err := collection.InsertMany(context.Background(), docs)
	fmt.Println("Result: ",result)

	//删除所有文档,bson.D{ }空值
	result_1,err := collection.DeleteMany(context.Background(),bson.D{ })

	//删除符合某一条件的所有文档
	result_2,err := collection.DeleteMany(
		context.Background(),
		bson.D{
			{"status","A"},
		})

	//删除符合某一条件的一个文档
	result_3,err := collection.DeleteOne(
		context.Background(),
		bson.D{
			{"status","D"},
		})
}
