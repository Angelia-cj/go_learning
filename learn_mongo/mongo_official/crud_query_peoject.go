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
	示例代码网站：https://docs.mongodb.com/manual/tutorial/project-fields-from-query-results/
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
			{"status", "A"},
			{"size", bson.D{
				{"h", 14},
				{"w", 21},
				{"uom", "cm"},
			}},
			{"instock", bson.A{
				bson.D{
					{"warehouse", "A"},
					{"qty", 5},
				},
			}},
		},
		bson.D{
			{"item", "notebook"},
			{"status", "A"},
			{"size", bson.D{
				{"h", 8.5},
				{"w", 11},
				{"uom", "in"},
			}},
			{"instock", bson.A{
				bson.D{
					{"warehouse", "EC"},
					{"qty", 5},
				},
			}},
		},
		bson.D{
			{"item", "paper"},
			{"status", "D"},
			{"size", bson.D{
				{"h", 8.5},
				{"w", 11},
				{"uom", "in"},
			}},
			{"instock", bson.A{
				bson.D{
					{"warehouse", "A"},
					{"qty", 60},
				},
			}},
		},
		bson.D{
			{"item", "planner"},
			{"status", "D"},
			{"size", bson.D{
				{"h", 22.85},
				{"w", 30},
				{"uom", "cm"},
			}},
			{"instock", bson.A{
				bson.D{
					{"warehouse", "A"},
					{"qty", 40},
				},
			}},
		},
		bson.D{
			{"item", "postcard"},
			{"status", "A"},
			{"size", bson.D{
				{"h", 10},
				{"w", 15.25},
				{"uom", "cm"},
			}},
			{"instock", bson.A{
				bson.D{
					{"warehouse", "B"},
					{"qty", 15},
				},
				bson.D{
					{"warehouse", "EC"},
					{"qty", 35},
				},
			}},
		},
	}

	result,err := collection.InsertMany(context.TODO(),docs)
	fmt.Println("Result: \n",result.InsertedIDs)

	//返回所有符合字符断的文档
	cursor_1,err := collection.Find(
		context.Background(),
		bson.D{
			{"stattus","A"},
	})
	//等价于：SELECT * from inventory WHERE status = "A"

	//返回指定字段和仅_id字段
	projection_1 := bson.D{
		{"item",1},
		{"status",1},
	}
	cursor_2,err := collection.Find(
		context.Background(),
		bson.D{
			{"status","A"},
		},
		options.Find().SetProjection(projection_1),
	)

	//等价于：SELECT _id, item, status from inventory WHERE status = "A"

	//抑制_id字段
	projection_2 := bson.D{
		{"item",1},
		{"status",1},
		{"_id",0},
	}
	cursor_3,err := collection.Find(
		context.Background(),
		bson.D{
			{"status","A"},
		},
		options.Find().SetProjection(projection_2),
	)
	//等价于：SELECT item, status from inventory WHERE status = "A"

	//返回所有但排除的字段
	projection_3 := bson.D{
		{"status",0},
		{"instock",0},
	}
	cursor_4,err := collection.Find(
		context.Background(),
		bson.D{
			{"status","A"},
		},
		options.Find().SetProjection(projection_3),
	)

	//返回嵌入式文档中的特定字段
	projection_4 := bson.D{
		{"item",1},
		{"status",1},
		{"size.uom",1},
	}
	cursor_5,err := collection.Find(
		context.Background(),
		bson.D{
			{"status","A"},
		},
		options.Find().SetProjection(projection_4),
	)

	//抑制嵌入式文档中的特定字段
	projection_5 := bson.D{
		{"size.uom",0},
	}
	cursor_6,err := collection.Find(
		context.Background(),
		bson.D{
			{"status","A"},
		},
		options.Find().SetProjection(projection_5),
	)

	//对数组上的嵌入式文档做一个投影
	projection_6 := bson.D{
		{"item",1},
		{"status",1},
		{"instock.qty",1},
	}
	cursor_7,err := collection.Find(
		context.Background(),
		bson.D{
			{"status","A"},
		},
		options.Find().SetProjection(projection_6),
	)

	//返回数组中的特定数组元素做一个投影
	projection_7 := bson.D{
		{"item",1},
		{"status",1},
		{"instock",bson.D{
			{"$slice",-1},
		}},
	}
	cursor_8,err := collection.Find(
		context.Background(),
		bson.D{
			{"status","A"},
		},
		options.Find().SetProjection(projection_7),
	)




















}
