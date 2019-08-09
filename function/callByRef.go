package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Println("This is main")
	fmt.Println("Enter two number")
	fmt.Scan(&n1, &n2)
	swap(&n1, &n2)
	fmt.Println(n1, n2)
}
func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
