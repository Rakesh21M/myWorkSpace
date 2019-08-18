// Sub go routine function will not wait for fk things to return, it will return imdtly

package main

import (
	"fmt"
	//"time"
)

func main() {
	r1 := FuncReturn3()
	r2 := FuncReturn4()
	//time.Sleep(5*time.Second)
	fmt.Println(r1, r2)
}

func FuncReturn3() string {
	//time.Sleep(5*time.Second)
	var str string
	go func() {
		str = "Rakesh"
	}()
	return str
}

func FuncReturn4() string {
	//time.Sleep(5*time.Second)
	var str1 string
	go func() {
		str1 = "charan"
	}()
	return str1
}
