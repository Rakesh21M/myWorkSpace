package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int

func main() {
	wg.Add(2)
	go increment("Raki")
	go increment("Cha")
	wg.Wait()
	fmt.Println("Success")
}

func increment(s string) {
	for i := 1; i <= 100; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		fmt.Println(s, "---", counter)
	}
	fmt.Println("Final Counter :", counter)
	wg.Done()
}
