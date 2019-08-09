package main

import "fmt"

func MyFunc(a int) (int, int) {
	b := 25
	return a, b
}

func main() {
	fmt.Println(MyFunc(10))
}
