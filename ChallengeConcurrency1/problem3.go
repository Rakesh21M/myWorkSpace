// creating a channel for factorial problem
package main

import (
	"fmt"
)

func main() {
	f := make(chan int)
	go func() {
		f <- 4
		close(f)
	}()
	out := factorial(f)
	fmt.Println(out)
}

func factorial(f chan int) int {
	var val int = 1
	//res := make(chan int)
	for n := range f {
		if n == 1 {
		} else {
			for i := 1; i <= n; i++ {
				val = val * i
			}
		}
	}
	return val
}
