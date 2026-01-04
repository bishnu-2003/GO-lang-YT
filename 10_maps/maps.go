package main

import "fmt"

//maps -> hash,objrct, dict

func main() {
	// creating map

	m := make(map[string]string)

	//setting an element
	m["name"] = "go lang"

	//get an element

	fmt.Println(m["name"])
}
	
