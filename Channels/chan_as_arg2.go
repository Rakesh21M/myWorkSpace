package main

import (
	"fmt"
)

func main() {
	ch := incre()
	out := sumP(ch)
	fmt.Println("The total sum :", out)
}

func incre() chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func sumP(ch chan int) int {
	var add int
	for n := range ch {
		add += n
	}
	return add
}
