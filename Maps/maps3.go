package main

import "fmt"

func main() {
	state := make(map[string]string)
	state["Karnataka"] = "Bangalore"
	state["Tamil Nadu"] = "Chennai"
	state["Andra Pradesh"] = "Hyderabad"
	delete(state, "Andra Pradesh")
	for _, value := range state {
		fmt.Println(value)
	}
	if val, error := state["Andra Pradesh"]; error {
		fmt.Println("Ok val  present", val)
	} else {
		fmt.Println("Ok val not present", val)
	}

}
