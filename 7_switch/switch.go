package main

import (
	"fmt"
	"time"
)

// simple switch
/*func main(){
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
}*/

// multiple condition switch

func main(){
	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
	fmt.Println("its weekend")
	default:
		fmt.Println("its workday")	
	}
}
