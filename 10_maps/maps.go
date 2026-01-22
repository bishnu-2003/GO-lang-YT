package main

import "fmt"

//maps -> hash,objrct, dict

func main() {
	// creating map

	m := make(map[string]string)

	//setting an element
	m["name"] = "go lang"
	m["version"] = "1.19"

	//get an element

	fmt.Println(m["name"], m["version"])
}

    
