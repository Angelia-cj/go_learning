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
	示例代码网站：https://docs.mongodb.com/manual/tutorial/query-arrays/
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
			{"tags", bson.A{"blank", "red"}},
			{"dim_cm", bson.A{14, 21}},
		},
		bson.D{
			{"item", "notebook"},
			{"qty", 50},
			{"tags", bson.A{"red", "blank"}},
			{"dim_cm", bson.A{14, 21}},
		},
		bson.D{
			{"item", "paper"},
			{"qty", 100},
			{"tags", bson.A{"red", "blank", "plain"}},
			{"dim_cm", bson.A{14, 21}},
		},
		bson.D{
			{"item", "planner"},
			{"qty", 75},
			{"tags", bson.A{"blank", "red"}},
			{"dim_cm", bson.A{22.85, 30}},
		},
		bson.D{
			{"item", "postcard"},
			{"qty", 45},
			{"tags", bson.A{"blue"}},
			{"dim_cm", bson.A{10, 15.25}},
		},
	}

	result,err := collection.InsertMany(context.TODO(),docs)
	fmt.Println("Result: \n",result.InsertedIDs)

	//匹配一个数组
	cursor_1,err := collection.Find(
		context.TODO(),
		bson.D{{"tags",bson.A{"red","blank"}}})

	//相反，如果你希望找到一个包含元素“red”和“blank”的数组，而不考虑数组中的顺序或其他元素，使用$ all运算符
	cursor_2,err := collection.Find(
		context.TODO(),
		bson.D{
			{"tags",bson.D{{"$all",bson.A{"red","blank"}}}},
		})

	//查询一个元素得到它所在的数组
	cursor_3,err := collection.Find(
		context.TODO(),
		bson.D{
			{"tags","red"},
		})

	//以下操作查询数组dim_cm包含至少一个值大于25的元素的所有文档
	cursor_4,err := collection.Find(
		context.TODO(),
		bson.D{
			{"dim_cm",bson.D{
				{"$gt",25}, //"$gt"：大于
			}},
		})

	//指定复合条件从数组中查询
	cursor_5,err := collection.Find(
		context.TODO(),
		bson.D{
			{"dim_cm",bson.D{
				{"$gt",15},
				{"$lt",20},
			}},
		})

	//查询符合多个元素条件的数组
	cursor_6,err := collection.Find(
		context.TODO(),
		bson.D{
			{"dim_cm",bson.D{
				{"$elemMatch",bson.D{
					{"$gt",22},
					{"lt",30},
				}},
			}},
		})

	//按数组索引的位置查询元素
	cursor_7,err := collection.Find(
		context.TODO(),
		bson.D{
			{"dim_cm.1",bson.D{
				{"$gt",25},
			}},
		})

	//根据数组的长度查询一个数组
	cursor_8,err := collection.Find(
		context.TODO(),
		bson.D{
			{"tags",bson.D{
				{"$size",3},
			}},
		})
}
