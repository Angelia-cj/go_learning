package main

import (
	"Go+Gin+MongoDB+RESTful/crud 1.1/src/models"
	"Go+Gin+MongoDB+RESTful/crud 1.1/src/handlers"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux" //Go语言中的HTTP路由库
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
	"net/http"

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

func init() {

	//使用虚拟数据填充数据库
	var people []models.Person

	client,err := mongo.NewClient(options.Client().ApplyURI(CONNECTIONSTRING))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database(DBNAME)

	//将值从JSON文件加载到模型
	byteValues,err := ioutil.ReadFile("Person_data.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(byteValues,&people)

	//向DB数据库插入人的信息
	var ppl []interface{}
	for _,p := range people{
		ppl = append(ppl,p)
	}
	_,err = db.Collection(COLLNAME).InsertMany(context.Background(),ppl)
	if err != nil {
		log.Fatal(err)
	}
}

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/people",handlers.GetAllPersonEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}",handlers.GetPersonEndPoint).Methods("GET")
	router.HandleFunc("/people",handlers.CreatePersonEndPoint).Methods("POST")
	router.HandleFunc("/people",handlers.DeletePersonEndPoint).Methods("DELETE")
	router.HandleFunc("/people/{id}",handlers.UpdatePersonEndPoint).Methods("PUT")
	fmt.Println("Starting server on port 8000...")
	log.Fatal(http.ListenAndServe(":8000",router))
}