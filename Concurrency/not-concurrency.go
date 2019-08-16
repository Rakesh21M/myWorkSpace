package main

import "fmt"

func PrintAll1() {
	for i := 0; i <= 10; i++ {
		fmt.Println("Rakesh")
	}
}

func PrintAll2() {
	for i := 0; i <= 10; i++ {
		fmt.Println("Charan")
	}
}

func main() {
	PrintAll1()
	PrintAll2()
}
