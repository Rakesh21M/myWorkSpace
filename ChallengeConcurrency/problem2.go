package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
	}()
	fmt.Println(<-ch)
}
