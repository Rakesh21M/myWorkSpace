package main

import "fmt"

type shape struct {
	Side float64
}

func (s shape) area() float64 {
	a := s.Side * s.Side
	return a
}

func main() {
	s1 := shape{10}
	a1 := s1.area()
	fmt.Println("The area of square :", a1)
}
