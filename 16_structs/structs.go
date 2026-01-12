package main

import "fmt"

type order struct {
	id        string
	amount    float32
	status    string
	//createdAt time.Time // nanosecond precision

}

func main() {

	order := order{
		id:     "12345",
		amount: 99.99,
		status: "received",
	}

	fmt.Println("order struct", order)
}