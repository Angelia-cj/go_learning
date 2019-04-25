package dao

import (
	"Go+Gin+MongoDB+RESTful/crud 1.1/src/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-go-driver-0.0.17/bson"

	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

/*
	数据库操作
*/

//数据库连接
const CONNECTIONSTRING  = "mongodb://localhost:27017"

//创建数据库
const DBNAME  = "users"

//创建集合
const COLLNAME  = "people"

//初始化一个数据库对象
var db  *mongo.Database

//通过Connect与数据库建立一个连接
func init()  {
	client,err := mongo.NewClient(options.Client().ApplyURI(CONNECTIONSTRING))//与MongoDB数据库建立连接
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//Collection类型可用于访问数据库
	db = client.Database(DBNAME)
}

//通过切片插入多条数据
func InsertManyValues(people []models.Person)  {//定义一个底层数组people
	var ppl [] interface{} //定义一个新的空slice ppl
	for _,p := range people{ //在people中，从下标为0开始，找到值为p的位置结束
		ppl = append(ppl,p) //然后把ppl和扩展的p 切片，重新放到ppl中
	}
	
	_,err := db.Collection(COLLNAME).InsertMany(context.Background(),ppl)
	if err != nil {
		log.Fatal(err)
	}
}

//插入一条数据，从person结构体中
func InsertOneValue(person models.Person)  {
	fmt.Println(person)
	_,err := db.Collection(COLLNAME).InsertOne(context.Background(),person)
	if err != nil {
		log.Fatal(err)
	}
}

//从DB数据库中获取所有的人
func GetAllPeople() []models.Person  {
	cur,err := db.Collection(COLLNAME).Find(context.Background(),nil,nil)
	if err != nil {
		log.Fatal(err)
	}
	var  elements []models.Person
	var  elem models.Person

	//从光标获取下一个结果
	for cur.Next(context.Background()) {
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		elements = append(elements,elem)
	}
	if err := cur.Err();err != nil{
		log.Fatal(err)
	}
	cur.Close(context.Background())
	return elements
}

//删除一个已经存在的person
func DeletePerson(person models.Person) {
	_, err := db.Collection(COLLNAME).DeleteOne(context.Background(), person, nil)
	if err != nil {
		log.Fatal(err)
	}
}

//更新一个已经存在的person
func UpdatePerson(person models.Person,personID string)  {
		doc := db.Collection(COLLNAME).FindOneAndUpdate(
		context.Background(),
		bson.NewDocument(
			bson.EC.String("id", personID),
		),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$set",
				bson.EC.String("firstname", person.Firstname),
				bson.EC.String("lastname", person.Lastname),
				bson.EC.String("contactinfo.city", person.City),
				bson.EC.String("contactinfo.phone", person.Phone),
				bson.EC.String("contactinfo.zipcode", person.Zipcode)),
		),
		nil)
	fmt.Println(doc)
}