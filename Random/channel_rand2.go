package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex
var count int

func main() {
	wg.Add(2)
	go func() {
		for i := 0; i < 4; i++ {
			mu.Lock()
			x := count
			x++
			count = x
			mu.Unlock()
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 3; i++ {
			mu.Lock()
			y := count
			y++
			count = y
			mu.Unlock()
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("The Count :", count)
}
