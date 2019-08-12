package main

import "fmt"

type Country struct {
	capital string
	name    string
}

type State struct {
	Country
	capital string
	name    string
}

func (s State) PrintAll() {
	fmt.Println(s.name)
	fmt.Println(s.capital)
}

func (c Country) PrintAll() {
	fmt.Println(c.name)
	fmt.Println(c.capital)
}

func main() {
	c1 := Country{"New Delhi", "India"}
	s1 := State{
		Country: Country{"New Delhi 1", "India 1"},
		capital: "Bangalore",
		name:    "Karnataka",
	}
	s1.PrintAll()
	c1.PrintAll()

}
