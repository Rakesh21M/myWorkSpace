// if we list all the natural number below 10 that are multiplies by 3 or 5, we get 3, 5 ,6, 9 the sum of
//these multiplies is 23. find the sum ofall the multiplies of 3or 5 below  1000
package main

import "fmt"

func main() {
	var sum int = 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum = sum + i
		}
	}
	fmt.Println("The total sum is ", sum)
}
