package main


//映射
type User struct {
	ID		objectid.ObjectID	"_id,omitempty"
	Name 	string				`bson:"dbname",json:"name"`
	Phone	string
}

//连接数据库
ctx,cancle := context.WithTimeout( )