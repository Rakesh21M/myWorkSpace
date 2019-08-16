package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("Raki")
	}()

	go func() {
		fmt.Println("Cha")
	}()

	time.Sleep(2 * time.Millisecond)
}
