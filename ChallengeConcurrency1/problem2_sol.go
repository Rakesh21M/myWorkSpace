// use range to solve that problem
//if we use range make sure that channel is closed.

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
		close(ch)
	}()
	for n := range ch {
		fmt.Println(n)
	}
}
