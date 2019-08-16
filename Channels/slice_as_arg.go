package main

import (
	"fmt"
)

func main() {
	s := incrementor()
	sumA := sum(s)
	fmt.Println("The Total sum :", sumA)
}

func incrementor() []int {
	Slice := make([]int, 0)
	for i := 1; i <= 10; i++ {
		Slice = append(Slice, i)
	}
	return Slice
}

func sum(slc []int) int {
	var X int
	for i := range slc {
		X = X + i
	}
	return X
}
