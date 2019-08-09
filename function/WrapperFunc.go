package main

import "fmt"

func WrapperFunc() func() string {
	return func() string {
		return "Rakesh"
	}
}

func main() {
	myVariable := WrapperFunc()
	fmt.Println(myVariable())
}
