package main

import "fmt"

type Human struct {
	name string
	age int
	weight int
}

type Student struct {
	Human  //匿名字段，那么默认Student就包含了Human的所有字段
	speciality string
}

func main(){
	//初始化一个学生
	mark := Student{Human{"Mark",25,120},"Computer Science"}

	//访问相应的字段
	fmt.Println("His name is ",mark.name)
	fmt.Println("His age is ",mark.age)
	fmt.Println("His weight is ",mark.weight)
	fmt.Println("His speciality is ",mark.speciality)

	//修改对应的备注信息
	mark.speciality = "AI"
	fmt.Println("Mark changed his speciality")
	fmt.Println("His speciality is ",mark.speciality)

	//修改他的年龄
	fmt.Println("Mark become old")
	mark.age = 36
	fmt.Println("His age is ",mark.age)

	/*mark.age = 36
	fmt.Println("Mark chenged his age")
	fmt.Println("His age is ",mark.age)*/


	//修改他的体重信息
	fmt.Println("Mark is not ......")
	mark.weight += 60
	fmt.Println("His weight is ",mark.weight)
	//mark.weight = 136
	//fmt.Println("Mark changed his weight")
	//fmt.Println("His weight is ",mark.weight)
}