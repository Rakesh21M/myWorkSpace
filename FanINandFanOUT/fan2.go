package main

import (
	"fmt"
	"sync"
)

func main() {
	ch2 := input(2, 3)
	c2 := sq(ch2)
	c3 := sq(ch2)
	cn2 := merge(c2, c3)

	for res := range cn2 {
		fmt.Println("The result ", res)
	}
}

func input(nums ...int) chan int {
	ch1 := make(chan int)
	go func() {
		for _, n := range nums {
			ch1 <- n
		}
		close(ch1)
	}()
	return ch1
}

func sq(ch3 chan int) chan int {
	c1 := make(chan int)
	go func() {
		for n := range ch3 {
			c1 <- n * n
		}
		close(c1)
	}()
	return c1
}
func merge(c ...chan int) chan int {
	cn1 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(c))
	for _, n := range c {
		go func(cs chan int) {
			for n1 := range cs {
				cn1 <- n1
			}
			wg.Done()
		}(n)
	}

	go func() {
		wg.Wait()
		close(cn1)
	}()
	return cn1
}
