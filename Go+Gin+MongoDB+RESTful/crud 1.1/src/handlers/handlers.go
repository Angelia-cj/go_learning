package handlers

import (
	"Go+Gin+MongoDB+RESTful/crud 1.1/src/dao"
	"Go+Gin+MongoDB+RESTful/crud 1.1/src/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var people []models.Person

//获得一个人的信息
func GetPersonEndPoint(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	payload := dao.GetAllPeople()
	for _,p := range payload{
		if p.ID == params["id"]{
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	json.NewEncoder(w).Encode("Person not found!")
}

//得到所有人的信息
func GetAllPersonEndPoint(w http.ResponseWriter, r *http.Request){
	payload := dao.GetAllPeople()
	json.NewEncoder(w).Encode(payload)
}

//创建一个人
func CreatePersonEndPoint(w http.ResponseWriter, r *http.Request){
	var person models.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	dao.InsertOneValue(person)
	json.NewEncoder(w).Encode(person)
}

//删除一个人
func DeletePersonEndPoint(w http.ResponseWriter, r *http.Request)  {
	var person models.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	dao.DeletePerson(person)
}

//更新一个人
func UpdatePersonEndPoint(w http.ResponseWriter, r *http.Request)  {
	personID := mux.Vars(r)["id"]
	var  person models.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	dao.UpdatePerson(person,personID)
}