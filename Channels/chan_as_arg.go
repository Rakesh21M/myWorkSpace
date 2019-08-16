package main

import (
	"fmt"
	"time"
)

func main() {
	s := incrementor1()
	sumA := sum1(s)
	time.Sleep(5 * time.Second)
	fmt.Println("The Total sum :", sumA)
}

func incrementor1() chan int {
	c := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			c <- i
		}
	}()
	return c
}

func sum1(c chan int) int {
	var X int
	go func() {
		for i := range c {
			X = X + i
		}
	}()
	return X
}
