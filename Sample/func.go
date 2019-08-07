package main

import "fmt"
import "time"

func read_things() {
	fmt.Println("Read this line of code")
}
func main() {
	time.Sleep(4*time.Second)
	read_things()
	fmt.Println("Now main function")
}
