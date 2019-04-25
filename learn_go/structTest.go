package main

import "fmt"

type person struct {
	name string
	age int
}

//比较两个人的年龄。返回年龄大的那个人，并且返回年龄差
//struct也是传值的

func Older(p1,p2 person)(person,int) {
	if p1.age > p2.age{   //比较p1和p2这两个人的年龄
		return p1,p1.age - p2.age
	}
	return p2,p2.age-p1.age
}


func main(){

	var tom person    //声明人Tom
	tom.name,tom.age = "Tom",20   //赋值初始化

	bob := person{age:26, name:"Bob"} //两个字段都写清楚的初始化

	paul := person{"Paul",45}  //按照struct定义顺序初始化值

	tb_Older,tb_diff := Older(tom,bob)
	tp_Older,tp_diff := Older(tom, paul)
	bp_Older,bp_diff := Older(bob,paul)

	fmt.Printf("Of %s and %s,%s is Older by %d years\n",tom.name,bob.name,tb_Older.name,tb_diff)
	fmt.Printf("Of %s and %s,%s is Older by %d years\n",tom.name,paul.name,tp_Older.name,tp_diff)
	fmt.Printf("Of %s and %s,%s is Older by %d years\n",bob.name,paul.name,bp_Older.name,bp_diff)
}
