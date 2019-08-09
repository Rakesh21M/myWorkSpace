package main

import "fmt"

func main() {
	myVariable := func() string {
		return "raki"
	}
	fmt.Println(myVariable())
}
