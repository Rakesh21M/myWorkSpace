package main

import "fmt"

func average(sf []float64) float64 {
	var sum float64
	for _, value := range sf {
		sum = sum + value
	}
	return sum
}

func main() {
	data := []float64{10, 20, 30, 40}
	total := average(data)
	fmt.Println(total)
}
