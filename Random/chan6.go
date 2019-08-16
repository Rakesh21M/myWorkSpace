package main

import (
	"fmt"
)

func main() {
	ch := RandFunc()
	//fmt.Printf("The type of channel is %T",ch)
	for n := range ch {
		fmt.Println(n)
	}
}

func RandFunc() chan int {
	out := make(chan int)
	go func() {
		out <- 2019
		close(out)
	}()
	//fmt.Println("Waiting")
	//close(out)
	return out
}
