package main

import "fmt"

func main() {
	y := 11
	fmt.Println("Hello Rakesh")
	{
		x := 21
		fmt.Println("This is the inner block")
		fmt.Println(x)
		fmt.Println(y)
	}
	//fmt.Println(x)
}
