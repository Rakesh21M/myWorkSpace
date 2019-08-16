package main

import "fmt"
import "sync"

//Semaphore
var wg sync.WaitGroup

func func1() {
	for i := 0; i <= 10; i++ {
		fmt.Println("This is Raki")
	}
	wg.Done()
}

func func2() {
	for i := 0; i <= 100; i++ {
		fmt.Println("This is cha")
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go func1()
	go func2()
	wg.Wait()
}
