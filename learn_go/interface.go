package main

import (
	"fmt"
)

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human
	school string
	loan float32
}

type Employee struct {
	Human
	company string
	money float32
}

//Human实现SayHi方法
func (h Human) SayHi()  {
	fmt.Printf("Hi ,I an %s you can call me on %s\n",h.name,h.phone)
}

//Human实现Sing方法
func (h Human) Sing(lyrics string)  {
	fmt.Println("la la la al ....",lyrics)
}

//Employee重载Human的SayHi方法
func (e Employee) SayHi()  {
	fmt.Printf("Hi , I am %s ,I work an  %s .Call me on %s\n",e.name,e.company,e.phone)
}

//Interface Men被Human，Student和Employee实现
//因为这三个类型都实现了这两个方法
type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main()  {
	mark := Student{Human{"Mark",24,"333-333-222"},"MIT",0.00}
	paul := Student{Human{"Paul",45,"333-555-444"},"Harvard",100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	//定义Men类型的变量i
	var i Men

	//i能存储Student
	i = mark
	fmt.Println("This is Mark , a student :")
	i.SayHi()
	i.Sing("November rain")

	//i也能存储Employee
	i = tom
	fmt.Println("This is Tom , an Employee :")
	i.SayHi()
	i.Sing("Born to be wild")

	//定义slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men,3)

	//这三个都是不同类型的元素，但是他们实现了interface同一个接口
	x[0],x[1],x[2] = paul,sam,mark

	for _,value := range x{
		value.SayHi()
	}
}













