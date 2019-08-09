package main

import "fmt"

func sample1() {
	fmt.Println("This is sample1 function")
}
func sample2(cbcall func()) {
	cbcall()
	fmt.Println("This is sample2 function")

}
func sample3() {
	fmt.Println("This is sample3 function")
}

func main() {
	sample2(func() {
		fmt.Println("This is callback function")
	})
}
