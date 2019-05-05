package main

import "fmt"

var(
	FirstName,SecondName,ThridName string
	i int
	f float32
	Input1 = "5.2 / 100 / Golang"//用户自定义变量，便于之后对这个字符串的处理
	format = "%f / %d / %s"
)

func main() {
	fmt.Printf("Please enter your full name: ")
	fmt.Scanln(&FirstName,&SecondName)//Scanln 扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行
	// fmt.Scanf("%s %s", &firstName, &lastName)    //Scanf与其类似，除了 Scanf 的第一个参数用作格式字符串，用来决定如何读取


	fmt.Printf("Hi %s %s!\n",FirstName,SecondName)
	//Sscan 和以 Sscan 开头的函数则是从字符串读取，除此之外，与 Scanf 相同。
	// 如果这些函数读取到的结果与您预想的不同，您可以检查成功读入数据的个数和返回的错误
	fmt.Sscanf(Input1,format, &f,&i,&ThridName)

	fmt.Println("From the Input we read: ",f,i,ThridName)

}
