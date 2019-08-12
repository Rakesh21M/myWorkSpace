package main

import (
	"fmt"
)

type Person3 struct {
	Name string
	Age  int
}

type Animal2 struct {
	Person3
	Name    string
	Atype   string
	Eatable bool
}

func (p Animal2) printAll() {
	fmt.Println(p.Name)
	fmt.Println(p.Age)
	fmt.Println(p.Atype)
	fmt.Println(p.Eatable)
}

func main() {
	p1 := Animal2{
		Person3: Person3{
			Name: "Raki",
			Age:  23,
		},
		Name:    "Charan",
		Atype:   "Human",
		Eatable: false,
	}
	p1.printAll()
}
