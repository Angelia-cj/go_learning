package main

import "fmt"

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human //匿名字段
	school string
}

type Employee struct {
	Human
	company string
}

//在human上面定义了一个method
func (h *Human) SaHai() {
	fmt.Printf("Hi,I am %s you can call me on %s\n",h.name,h.phone)
}

func main()  {
	mark := Student{Human{"Mark",25,"222-333-444"},"MIT"}
	jim := Employee{Human{"Jim" ,24,"333-222-555"},"Golsng Inc"}

	mark.SaHai()
	jim.SaHai()

}