// Used semaphore to pass values from multiple function
//to channel

package main

import (
	"fmt"
)

//var count1 int
//ch:= make(chan int)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan bool)
	go func() {
		for i := 0; i < 4; i++ {
			ch1 <- i
		}
		ch2 <- true
	}()
	go func() {
		for i := 10; i < 13; i++ {
			ch1 <- i
		}
		ch2 <- true
	}()
	go func() {
		<-ch2
		<-ch2
		close(ch1)

	}()
	for n := range ch1 {
		fmt.Println("From channel 1", n)
	}
}
