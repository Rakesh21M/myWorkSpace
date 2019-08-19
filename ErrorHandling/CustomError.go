package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Finding the square of positive number")
	_, err := sqrt(-23)
	if err != nil {
		log.Fatal(err)
	}
}

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, errors.New("no negative numbers")
	}
	val := n * n
	return val, nil
}
