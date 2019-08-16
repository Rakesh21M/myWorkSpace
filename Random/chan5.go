// go not wait for any fk things

package main

import "fmt"

func main() {
	go func() {
		fmt.Println("Rakesh")
	}()
	fmt.Println("Charan")
}
