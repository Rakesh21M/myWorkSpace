package main

import (
	"fmt"
)

func main() {
	c1 := make(chan int)
	out := Sender(c1)
	for n := range out {
		fmt.Println(n)
	}
}

func Sender(c1 chan int) chan int {
	go func() {
		for i := 1; i <= 100; i++ {
			c1 <- i
		}
		close(c1)
	}()

	return c1
}
