package main

import (
	//"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("filename.txt")
	if err != nil {
		//fmt.Println("Error Opening File!! ", err)
		//log.Println("Error Opening File!! ", err)
		//panic(err)
		log.Fatalln(err)
	}
}
