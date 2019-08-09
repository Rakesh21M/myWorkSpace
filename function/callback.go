package main

import "fmt"

func cbfunc(sf []int, sample func()) {
	sample()
}

func main() {
	sf := []int{1, 2, 3, 4}
	cbfunc(sf, func() {
		fmt.Println("raki")
	})
}
