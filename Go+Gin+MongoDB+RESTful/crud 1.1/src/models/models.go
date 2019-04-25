package models

/*
	定义人和联系信息结构体
*/

type Person struct {
	ID          string `json:"id,omitempty"`
	Firstname   string `json:"firstname,omitempty"`
	Lastname    string `json:"lastname,omitempty"`
	Contactinfo `"json:"contactinfo,omitempty"`
}

type Contactinfo struct {
	City		string		`json:"city,omitempty"`
	Zipcode		string		`json:"zipcode,omitempty"`
	Phone 		string		`json:"phone,omitempty"`
}