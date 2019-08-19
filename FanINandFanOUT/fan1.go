package main

import (
	"fmt"
)

func main() {
	c := fanout(boring("Raki"), boring("Cha"))
	for i := 1; i <= 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You both are boring, i'm leaving bye")
}

func boring(str1 string) <-chan string {
	c1 := make(chan string)
	go func() {
		c1 <- str1
	}()
	return c1
}

func fanout(c2, c3 <-chan string) <-chan string {
	c4 := make(chan string)
	//str2 := "San"
	go func() {
		c4 <- <-c2
	}()
	go func() {
		c4 <- <-c3
	}()
	return c4
}
