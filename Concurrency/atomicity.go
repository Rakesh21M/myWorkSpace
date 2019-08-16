package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg4 sync.WaitGroup
var count int64

func incre(s string) {
	for i := 1; i <= 100; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		atomic.AddInt64(&count, 1)
		fmt.Println(s, "---", count)
	}
	wg4.Done()

}
func main() {
	fmt.Println("Hello World")
	wg4.Add(2)
	incre("R")
	incre("C")
	wg4.Wait()
	fmt.Println("The final counter", count)
}
