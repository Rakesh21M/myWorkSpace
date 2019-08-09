// One Variadic parameter that finds the greater in list
package main

import "fmt"

func Largest(vf ...int) int {
	large := vf[0]
	for _, value := range vf {
		if value > large {
			large = value
		}
	}
	return large
}

func main() {
	sl := []int{2, 4, 6, 9, 1, 150}
	n := Largest(sl...)
	fmt.Println(n)

}
