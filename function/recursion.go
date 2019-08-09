package main

import "fmt"

func main() {
	var num int
	//var result int
	fmt.Println("Enter the number\n")
	fmt.Scan(&num)
	result := fibonacci(num)
	fmt.Println(result)
}

func fibonacci(n int) int {
	if n == 1 {
		return 0
	} else if n == 2 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}
