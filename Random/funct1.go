package main

import "fmt"

func main() {
	var1 := func() string {
		return "Hello World"
	}
	fmt.Println(var1())

	func() {
		fmt.Println("This is my function")
	}()
}
