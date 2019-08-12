package main

import "fmt"

func main() {
	Mycountry := make(map[string]string)
	if Mycountry == nil {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
