package main

import "fmt"

func main() {
	type raki int
	var a raki = 21
	fmt.Println(a)
	fmt.Printf("%T", a)
}
