package main

import "fmt"

func main() {
    i := 10               // Declare an integer variable 'i'
    var p *int            // Declare a pointer 'p' to an integer
    p = &i                // Assign the memory address of 'i' to 'p'

    fmt.Println("Value of i:", i)       // Output: Value of i: 10
    fmt.Println("Address of i:", &i)    // Output: Address of i: 0x... (some memory address)
    fmt.Println("Value of p:", p)       // Output: Value of p: 0x... (same memory address as &i)
    fmt.Println("Value at address *p:", *p) // Output: Value at address *p: 10

    *p = 20 // Change the value at the memory address 'p' points to
    fmt.Println("New value of i:", i) // Output: New value of i: 20 (original 'i' is modified)
}
