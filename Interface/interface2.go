package main

import "fmt"

type Company struct {
	Manager
	SeniorEmployee
	JuniorEmployee
	Fresher
}

type Manager interface {
	Salary() float64
}

type SeniorEmployee interface {
	Salary() float64
}

type JuniorEmployee interface {
	Salary() float64
}

type Fresher interface {
	Salary() float64
}

func (c Company) Salary() float64 {
	return 25000.00
}

func PrintSalary(n Manager) {
	fmt.Println(n.Salary())
}

func main() {
	var c Company
	PrintSalary(c)
}
