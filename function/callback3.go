package main

import "fmt"

func sample1(cbfunc func(int) bool) {
	n := 21
	if cbfunc(n) {
		fmt.Println("Its working")
	} else {
		fmt.Println("Not Working")
	}
}

func main() {
	sample1(func(n int) bool {
		if n < 21 {
			return true
		} else {
			return false
		}
	})
}
