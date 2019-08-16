package main

import "fmt"
import "time"

func FirstFunc() {
	//time.Sleep(3 * time.Second)
	fmt.Println("Rakesh")
}

func SecondFunc() {
	time.Sleep(3 * time.Second)
	fmt.Println("Charan")
}

func main() {
	go FirstFunc()
	go SecondFunc()
	time.Sleep(20 * time.Second)
}
