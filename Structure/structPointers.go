package main

import "fmt"

type Image struct {
	size string
	name string
}

func main() {
	i1 := &Image{"720x1024", "Raki.doc"}
	fmt.Println(i1.size, "\t", i1.name)
	//fmt.Println(i1.size,"\t",i1.name)

}
