package main

import "fmt"

func main(){

	sum := 0
	for index := 0;index < 10; index ++{
		sum += index
	}
	fmt.Println("sun is equal to",sum)
}
