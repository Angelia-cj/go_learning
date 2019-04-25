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
	示例代码网站：https://docs.mongodb.com/manual/tutorial/query-for-null-fields/
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
			{"_id", 1},
			{"item", nil},
		},
		bson.D{
			{"_id", 2},
		},
	}

	result, err := collection.InsertMany(context.Background(), docs)
	fmt.Println("Result: \n",result.InsertedIDs)


	//等价过滤器
	cursor_1,err := collection.Find(
		context.Background(),
		bson.D{
			{"item",nil},
		})

	//类型检查
	cursor_2,err := collection.Find(
		context.Background(),
		bson.D{
			{"item",bson.D{
				{"$type",10},
			}},
		})

	//是否存在检查
	cursor_3,err := collection.Find(
		context.Background(),
		bson.D{
			{"item",bson.D{
				{"$exists",false},
			}},
		})
}