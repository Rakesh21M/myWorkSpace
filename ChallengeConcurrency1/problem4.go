// Factorial using channel

package main

import (
	"fmt"
)

func main() {
	var n int = 4
	ch1 := factorial1(n)
	fmt.Println(<-ch1)

}
func factorial1(n int) <-chan int {
	var total int = 1
	ch := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			total = total * i
		}
		ch <- total
	}()
	return ch
}
