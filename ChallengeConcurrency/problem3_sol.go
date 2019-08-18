package main

import (
	"fmt"
)

func main() {
	out := fibo(5)
	for n := range out {
		fmt.Println(n)
	}
}

func fibo(v int) chan int {
	ch := make(chan int)
	var total int
	for i := 1; i <= v; i++ {
		total = total + i
	}
	ch <- total
	return ch
}
