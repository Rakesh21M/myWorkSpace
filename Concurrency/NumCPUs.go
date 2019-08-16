package main

import "fmt"
import "time"
import "sync"
import "runtime"

var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func PrintData1() {
	fmt.Println("Raki")
	wg.Done()
}

func PrintData2() {
	fmt.Println("Cha")
	wg.Done()
}
func main() {
	wg.Add(2)
	go PrintData1()
	go PrintData2()
	time.Sleep(5 * time.Second)
	wg.Wait()
}
