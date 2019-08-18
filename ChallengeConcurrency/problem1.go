package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	ch <- 23 // code blocks here, i am putting on to the channel, nothing to receive
	fmt.Println(<-ch)
}
