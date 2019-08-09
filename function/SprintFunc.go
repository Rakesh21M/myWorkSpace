package main

import "fmt"

func Sprintthings() string {
	s := "namedThings"
	return s
}

func main() {
	st := fmt.Sprintln(Sprintthings())
	fmt.Printf("%T\n", st)
	fmt.Println(st)
}
