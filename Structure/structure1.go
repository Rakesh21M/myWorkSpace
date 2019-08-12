package main

import "fmt"

type Person1 struct {
	First string
	Last  string
	Age   int
}

func (p Person1) printNames() {
	fmt.Println(p.First, "\t", p.Last)
}

func main() {
	per1 := Person1{"Raki", "Manjunatha", 23}
	per2 := Person1{
		First: "Charan",
		Last:  "Manjunatha",
		Age:   19,
	}
	per3 := Person1{}
	per3.First = "Sanjay"
	per3.Last = "Babu"
	per3.Age = 23
	per1.printNames()
	per2.printNames()
	per3.printNames()

}
