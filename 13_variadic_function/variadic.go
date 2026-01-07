package main

import "fmt" 

func sum(nums ...int) int {
	total := 0

	for i := range nums {
		total += nums[i]
	}

	return total
}

func main(){
	result := sum(1, 2, 3, 4, 5)
	fmt.Println("The sum is:", result)
}