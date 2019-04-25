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
	示例代码网址：https://docs.mongodb.com/manual/tutorial/update-documents/
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
			{"item", "canvas"},
			{"qty", 100},
			{"size", bson.D{
				{"h", 28},
				{"w", 35.5},
				{"uom", "cm"},
			}},
			{"status", "A"},
		},
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
			{"item", "mat"},
			{"qty", 85},
			{"size", bson.D{
				{"h", 27.9},
				{"w", 35.5},
				{"uom", "cm"},
			}},
			{"status", "A"},
		},
		bson.D{
			{"item", "mousepad"},
			{"qty", 25},
			{"size", bson.D{
				{"h", 19},
				{"w", 22.85},
				{"uom", "in"},
			}},
			{"status", "P"},
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
		bson.D{
			{"item", "sketchbook"},
			{"qty", 80},
			{"size", bson.D{
				{"h", 14},
				{"w", 21},
				{"uom", "cm"},
			}},
			{"status", "A"},
		},
		bson.D{
			{"item", "sketch pad"},
			{"qty", 95},
			{"size", bson.D{
				{"h", 22.85},
				{"w", 30.5},
				{"uom", "cm"},
			}},
			{"status", "A"},
		},
	}

	result, err := collection.InsertMany(context.Background(), docs)
	fmt.Println("Result: ",result.InsertedIDs)

	//更新文档在集合中
	result_1,err := collection.UpdateOne(
		context.Background(),
		bson.D{	//需要更新的数组元素
			{"item","paper"},
		},
		bson.D{	//新的结果
			{"$set",bson.D{
				{"size.uom","cm"},
				{"status","p"},
			}},//$set：更新值
			{"$currentDate",bson.D{
				{"lastModified",true},
			}},//"$currentDate"：返回最后修改的时间，如果"lastModified"不存在，$currentDate"新创建一个
		},
	)

	//更新多个文档
	result_2,err := collection.UpdateMany(
		context.Background(),
		bson.D{
			{"qty",bson.D{
				{"$lt",50},
			}},
		},
		bson.D{
			{"$set",bson.D{
				{"size.uom","cm"},
				{"status","p"},
			}},
			{"$currentDate",bson.D{
				{"lastModified",true},
			}},
		})

	//替换整个文档
	result_3,err := collection.ReplaceOne(
		context.Background(),
		bson.D{
			{"item","paper"},
		},
		bson.D{
			{"item","paper"},
			{"instock",bson.A{
				bson.D{
					{"warehouse","A"},
					{"qty",40},
				},
			}},
		})

}