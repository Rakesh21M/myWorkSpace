package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- 24
	}()
	fmt.Println(<-ch)
}
