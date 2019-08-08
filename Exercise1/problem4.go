package main

import "fmt"

func main() {
	var n1, n2, result int
	//var result float32
	fmt.Println("Enter the first number ")
	fmt.Scan(&n1)
	fmt.Println("Enter the second number ")
	fmt.Scan(&n2)
	result = n1 % n2
	fmt.Println("The result is ", result)

}
