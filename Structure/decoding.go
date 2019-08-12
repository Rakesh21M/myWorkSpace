package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person21 struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	var p1 Person21
	rdr := strings.NewReader(`{"First":"Rakesh","Last":"Manjunatha","Age":24}`)
	json.NewDecoder(rdr).Decode(&p1)
	fmt.Println(p1.First)
}
