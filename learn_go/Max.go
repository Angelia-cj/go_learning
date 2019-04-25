package main

import "fmt"

func max(a,b int) int {
	 if a > b{
	 	return a
	 }
	 return b
}

func main(){
	x := 0
	y := 2
	z := 3

	max_xy := max(x,y)
	max_xz := max(x,z)

	fmt.Printf("max(%d,%d) = %d \n",x,y,max_xy)
	fmt.Printf("max(%d,%d) = %d \n",x,z,max_xz)
	fmt.Printf("max(%d,%d) = %d \n",y,z,max(y,z))
}