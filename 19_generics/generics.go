package main

import "fmt"

// Constraint
type Number interface {
	int | float64
}

// Generic function
func Sum[T Number](a, b T) T {
	return a + b
}

// Generic struct
type Box[T any] struct {
	Value T
}

// Generic helper
func Contains[T comparable](arr []T, val T) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func main() {

	// Generic function
	fmt.Println(Sum(10, 20))
	fmt.Println(Sum(10.5, 20.5))

	// Generic struct
	intBox := Box[int]{Value: 100}
	strBox := Box[string]{Value: "Go"}

	fmt.Println(intBox.Value)
	fmt.Println(strBox.Value)

	// Generic slice helper
	fmt.Println(Contains([]int{1, 2, 3}, 2))
	fmt.Println(Contains([]string{"a", "b"}, "c"))
}
