package main

import (
	"fmt"
	//"math/rand"
	"time"
)

func main() {
	c := make(chan int)
	go func() {
		//time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		for i := 1; i <= 10; i++ {
			c <- i
		}
	}()

	go func() {
		//time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		for {
			fmt.Println(<-c)
		}
	}()

	time.Sleep(3 * time.Second)
}
