// Used wait group to pass values from multiple function
//to channel

package main

import (
	"fmt"
	"sync"
)

var wg1 sync.WaitGroup

//var count1 int
//ch:= make(chan int)

func main() {
	ch1 := make(chan int)
	wg1.Add(2)
	go func() {
		for i := 0; i < 4; i++ {
			ch1 <- i
		}
		wg1.Done()
	}()
	go func() {
		for i := 10; i < 13; i++ {
			ch1 <- i
		}
		wg1.Done()
	}()
	go func() {
		wg1.Wait()
		close(ch1)

	}()
	for n := range ch1 {
		fmt.Println("From channel 1", n)
	}
}
