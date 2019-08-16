package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg2 sync.WaitGroup
var counter1 int
var mu sync.Mutex

func main() {
	wg2.Add(2)
	go increment1("Raki")
	go increment1("Cha")
	wg2.Wait()
	fmt.Println("Final Counter :", counter1)
	fmt.Println("Success")
}

func increment1(s string) {
	for i := 1; i <= 100; i++ {
		mu.Lock()
		x := counter1
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter1 = x
		mu.Unlock()
		fmt.Println(s, "---", counter1)
	}
	wg2.Done()
}
