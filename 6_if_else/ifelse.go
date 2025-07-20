package main

import "fmt"

func main() {

	/*age := 18 

	if age >= 18 {
	fmt.Println("you are an adult")
	}else {
			fmt.Println("person not an adult")
		}*/
	

		// we can declare a variable inside if construct


		if age := 20 ; age >= 18 {
	fmt.Println("person is an adult", age)
}else if age >= 12 {
	fmt.Println("person is teenager", age)
}

}

