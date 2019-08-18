package main

import (
	"fmt"
)

func main() {
	ch1 := FirstOne("Raki")
	ch2 := SecondOne(ch1)
	fmt.Println("The total sum :", <-ch2)
}

func FirstOne(s string) chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			out <- i
			fmt.Println(s, "---", i)
		}
		close(out)
	}()
	return out
}

func SecondOne(c2 chan int) chan int {
	out := make(chan int)
	var sum int
	go func() {
		for n := range c2 {
			sum += n
		}
		out <- sum
	}()
	return out
}
