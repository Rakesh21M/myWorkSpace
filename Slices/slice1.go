package main

import "fmt"

func main() {
	sl1 := make([]int, 2)
	sl1[0] = 2
	sl1[1] = 4
	sl1 = append(sl1, 5)
	sl1[2]++
	fmt.Println(sl1)
}
