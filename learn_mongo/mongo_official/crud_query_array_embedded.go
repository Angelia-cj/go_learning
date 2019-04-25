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
	示例代码网站：https://docs.mongodb.com/manual/tutorial/query-array-of-documents/
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
			{"instock", bson.A{
				bson.D{
					{"warehouse", "A"},
					{"qty", 5},
				},
				bson.D{
					{"warehouse", "C"},
					{"qty", 15},
				},
			}},
		},
		bson.D{
			{"item", "notebook"},
			{"instock", bson.A{
				bson.D{
					{"warehouse", "C"},
					{"qty", 5},
				},
			}},
		},
		bson.D{
			{"item", "paper"},
			{"instock", bson.A{
				bson.D{
					{"warehouse", "A"},
					{"qty", 60},
				},
				bson.D{
					{"warehouse", "B"},
					{"qty", 15},
				},
			}},
		},
		bson.D{
			{"item", "planner"},
			{"instock", bson.A{
				bson.D{
					{"warehouse", "A"},
					{"qty", 40},
				},
				bson.D{
					{"warehouse", "B"},
					{"qty", 5},
				},
			}},
		},
		bson.D{
			{"item", "postcard"},
			{"instock", bson.A{
				bson.D{
					{"warehouse", "B"},
					{"qty", 15},
				},
				bson.D{
					{"warehouse", "C"},
					{"qty", 35},
				},
			}},
		},
	}

	result, err := collection.InsertMany(context.Background(), docs)
	fmt.Println("Result: \n",result.InsertedIDs)

	//查询嵌套在数组中的文档
	cursor_1,err := collection.Find(
		context.Background(),
		bson.D{
			{"instock",bson.D{
				{"warehouse","A"},
				{"qty",5},
			}},
		})
	//整个嵌入/嵌套文档上的等价匹配需要指定文档的完全匹配，包括字段顺序.
	//下面这个查询文档与任何的集合都不匹配
	cursor_2,err := collection.Find(
		context.Background(),
		bson.D{
			{"instock",bson.D{
				{"qty",5},
				{"warehouse","A"},
			}},
		})//查询条件的顺序不对，故查不到任何数据

	//在文档数组中的字段上指定查询条件
	cursor_3,err := collection.Find(
		context.Background(),
		bson.D{
			{"instock.qty",bson.D{
				{"$lte",20},
			}},
		})

	//使用数组索引查询嵌入文档中的字段
	cursor_4,err := collection.Find(
		context.Background(),
		bson.D{
			{"instock.0.qty",bson.D{
				{"$lte",20},
			}},
		})

	//单个嵌套文档在嵌套字段上遇到多个查询条件
	cursor_5,err := collection.Find(
		context.Background(),
		bson.D{
			{"instock",bson.D{
				{"$elemMatch",bson.D{
					{"qty",5},
					{"warehouse","A"},
				}},
			}},
		})

	//查询instock数组至少有一个嵌入文档的文档，该文档包含大于10且小于或等于20的字段数量
	cursor_6,err := collection.Find(
		context.Background(),
		bson.D{
			{"instock",bson.D{
				{"$elemMatch",bson.D{
					{"qty",bson.D{
						{"$gt",10},
						{"$lte",20},
					}},
				}},
			}},
		})

	//如果数组字段上的复合查询条件不使用$ elemMatch运算符
	// 则查询将选择其数组包含满足条件的任何元素组合的文档
	cursor_7,err := collection.Find(
		context.Background(),
		bson.D{
			{"instock.qty",bson.D{
				{"$gt",10},
				{"$lte",20},
			}},
		})

	//instock数组至少有一个包含字段数量等于5的嵌入文档和
	// 至少一个包含字段等于A的嵌入文档（但不一定是相同的嵌入文档）
	cursor_8,err := collection.Find(
		context.Background(),
		bson.D{
			{"instock",5},
			{"instock.warehouse","A"},
		})
}
