package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)
/*
	示例代码网站：https://docs.mongodb.com/manual/tutorial/query-embedded-documents/
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

    //创建一些需要插入的数据
	docs := []interface{}{
		bson.D{
			//{"_id","1"},
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
			//{"_id","2"},
			{"item", "notebook"},
			{"qty", 50},
			{"size", bson.D{
				{"h", 8.5},
				{"w", 11},
				{"uom", "in"},
			}},
			{"status", "A"},
		},
		bson.D{
			//{"_id","3"},
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
			//{"_id","4"},
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
			//{"_id","5"},
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
	//向数据库插入操作
	result,err := collection.InsertMany(context.TODO(),docs)
	fmt.Println("Result: \n",result.InsertedIDs)

	cursor_1,err := collection.Find(
		context.TODO(),
		bson.D{})
	//等价于：SELECT * FROM inventory

	//查询一条数据
	cursor_2,err := collection.Find(
		context.TODO(),
		bson.D{{"status","D"}})
	//等价于：SELECT * FROM inventory WHERE status = "D"

	//查询多条数据,使用"$in"运算符，同一字段执行相等性检查
	cursor_3 ,err := collection.Find(
		context.TODO(),
		bson.D{{"status",bson.D{{"$in",bson.A{"A","D"}}}}})
	//等价于：SELECT * FROM inventory WHERE status in ("A", "D")

	//AND 条件
	cursor_4,err := collection.Find(
		context.TODO(),
		bson.D{{"status","A"},
			{"qty",bson.D{{"$lt",30}}}})
	//等价于：SELECT * FROM inventory WHERE status = "A" AND qty < 30

	//OR条件  "$or"：或者运算符，"$lt"：小于符号
	cursor_5,err := collection.Find(
		context.TODO(),
		bson.D{
			{"$or",
				bson.A{
					bson.D{{"status","A"}},
					bson.D{{"qty",bson.D{{"$lt",30}}}},
				}},
		})
	//等价于：SELECT * FROM inventory WHERE status = "A" OR qty < 30

	//指定AND以及OR条件
	cursor_5,err := collection.Find(
		context.TODO(),
		bson.D{
			{"status","A"},
			{"$or",bson.A{
				bson.D{{"qty",bson.D{{"$lt",30}}}},
				bson.D{{"item",primitive.Regex{Pattern:"^p",Options:""}}},
			}},
		})
	//等价于：SELECT * FROM inventory WHERE status = "A" AND ( qty < 30 OR item LIKE "p%")

	//在嵌套字段上指定等价匹配
	cursor_6,err := collection.Find(
		context.TODO(),
		bson.D{{"size.uom","in"}})

	//使用查询运算符指定匹配
	cursor_7,err := collection.Find(
		context.TODO(),
		bson.D{
			{"size.h",bson.D{
				{"$lt",15},
			}},
		})
	//指定AND条件查询
	cursor_8,err := collection.Find(
		context.TODO(),
		bson.D{
			{"size.h",bson.D{
				{"$lt",15},
			}},
			{"size.uom","in"},
			{"status","D"},
		})








}
