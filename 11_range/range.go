package main

// iterating over data structures

func main(){
	nums := []int{10, 20, 30, 40, 50}

	for index, value := range nums {
		println("Index:", index, "Value:", value)
	}
}