package main

import (
	"fmt"
)

type Person2 struct {
	Name string
	Age  int
}

type Animal1 struct {
	Person2
	Atype   string
	Eatable bool
}

func (p Animal1) printAll() {
	fmt.Println(p.Name)
	fmt.Println(p.Age)
	fmt.Println(p.Atype)
	fmt.Println(p.Eatable)
}

func main() {
	p1 := Animal1{
		Person2: Person2{
			Name: "Raki",
			Age:  23,
		},
		Atype:   "Human",
		Eatable: false,
	}
	p1.printAll()
}
