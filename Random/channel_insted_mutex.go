package main

import (
	"fmt"
	"time"
)

//var wg1 sync.WaitGroup
//var count1 int
//ch:= make(chan int)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		for i := 0; i < 4; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	go func() {
		for i := 10; i < 13; i++ {
			ch2 <- i
		}
		close(ch2)
	}()
	for n := range ch1 {
		fmt.Println("From channel 1", n)
	}
	for m := range ch2 {
		fmt.Println("From channel 2", m)
	}
}
