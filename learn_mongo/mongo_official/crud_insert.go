package  main

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
	示例代码网站：https://docs.mongodb.com/manual/tutorial/insert-documents/
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


	//插入一条数据
	result,err := collection.InsertOne(
		context.TODO(),
		bson.D{
			{"item","canvas"},//商品名称
			{"qty",100},//商品重量
			{"tags",bson.A{"cotton"}},//商品标签，cotton：棉
			{"size",bson.D{//商品尺寸
				{"h",28},//长
				{"w",35.5},//宽
				{"uom","cm"},//面积
			}},
		})
	fmt.Println("Data：%v\n",result.InsertedID)

	//查找一个数据
	cursor,err := collection.Find(
		context.TODO(),bson.D{{"item","canvas"}})
	fmt.Println("Find：%v\n",cursor)




}
