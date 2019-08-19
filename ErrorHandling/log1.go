package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf) // We are setting output, so log.Println puts output here!!
}

func main() {
	_, err := os.Open("Filename.txt")
	if err != nil {
		log.Println("Error Opening file", err)
	}
}
