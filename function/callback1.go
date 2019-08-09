package main

import "fmt"

func sample(cbfunc func()) {
	cbfunc()
}

func main() {
	sample(func() {
		fmt.Println("This is callback function")
	})
	fmt.Println("This works")
}
