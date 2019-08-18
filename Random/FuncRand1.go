// What is stopping main function from exit its return
package main

import (
	"fmt"
	"time"
)

func main() {
	r1 := FuncReturn1()
	r2 := FuncReturn2()
	fmt.Println(r1, r2)
}

func FuncReturn1() string {
	time.Sleep(5 * time.Second)
	return "Raki"
}

func FuncReturn2() string {
	time.Sleep(5 * time.Second)
	return "cha"
}
