package main

import (
	"fmt"
)
import "os"
import "bufio"

func main() {
	var name string
	reader := bufio.NewReader(os.Stdin)
	name, _ = reader.ReadString('\n')
	fmt.Println("Hello,", name)

}
