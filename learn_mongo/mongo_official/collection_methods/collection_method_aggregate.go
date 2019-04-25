package collection_methods

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func main() {

	//创建服务器客户端
	client,err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
	   log.Fatal(err)
	}
	//创建数据库
	db := client.Database("dbname")
	//创建集合
	coll := db.Collection("collection")

	//db.collection.aggregate(pipeline, options)
	//计算集合或视图中数据的聚合值

}