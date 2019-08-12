package main

import "fmt"

type Person struct {
	First string
	Last  string
	age   int
}

type Animal struct {
	Person
	name  string
	isPet bool
}

func main() {
	ani1 := Animal{}
	ani1.Person.Last = "manjunatha"
	ani1.Person.First = "Raki"
	ani1.Person.age = 21
	ani1.name = "Human"
	ani1.isPet = false
	//per1.age = 21
	//per1.First = "Raki"
	//per1.Last = "manjunatha"

	fmt.Println(ani1)
}
