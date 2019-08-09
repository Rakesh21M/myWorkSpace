package main

import "fmt"

func namedReturn() (n int) {
	n = 45
	return
}

func main() {
	fmt.Println(namedReturn())
}
