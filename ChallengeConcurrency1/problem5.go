//Square of a number using channel

package main

import (
	"fmt"
)

func main() {
	ch2 := inputChannel(2, 3, 4)
	c2 := squareChannel(ch2)
	for num := range c2 {
		fmt.Println("The numbers are :", num)
	}

}

func inputChannel(nums ...int) <-chan int {
	ch1 := make(chan int)
	go func() {
		for _, n := range nums {
			ch1 <- n
		}
		close(ch1)
	}()
	return ch1
}

func squareChannel(ch3 <-chan int) chan int {
	c1 := make(chan int)
	go func() {
		for n1 := range ch3 {
			c1 <- n1 * n1
		}
		close(c1)
	}()
	return c1
}
