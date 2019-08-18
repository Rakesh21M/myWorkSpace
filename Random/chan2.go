package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := prging()
	fmt.Println(<-c1)

}
func prging() chan int {
	out := make(chan int)
	//fmt.Println("This is samp",out)
	go func() {
		time.Sleep(5 * time.Second)
		out <- 20
	}()
	//close(out)
	return out
}
