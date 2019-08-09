package main

import "fmt"

func average(a ...float64) float64 {
	var sum float64
	for _, value := range a {
		sum = sum + value
	}
	return sum
}
func main() {
	number := average(10, 20, 30)
	fmt.Printf("%T\n", number)
	fmt.Println(number)
}
