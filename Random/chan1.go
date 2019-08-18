//fk flow of the program

package main

import (
	"fmt"
	"time"
)

func main() {
	n := prgingDesk1()
	prgingDesk2()
	fmt.Println(n)
}
func prgingDesk1() int {
	go func() {
		fmt.Println("Raki")
	}()

	return 21

}

func prgingDesk2() {
	go func() {
		fmt.Println("cha")
	}()
	time.Sleep(5 * time.Second)

}
