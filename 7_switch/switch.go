package main

import "fmt"

// simple switch
func main(){
i := 3

switch i {
case 1:
	fmt.Println("one")
case 2:
	fmt.Println("two")
case 3:
	fmt.Println("three")
default:
	fmt.Println("other")
}
}
