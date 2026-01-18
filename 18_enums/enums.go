package main

import "fmt"

/*
	ENUM 1: Integer enum using iota
*/
type Status int

const (
	Pending Status = iota
	Approved
	Rejected
)

/*
	ENUM 2: Enum with custom integer values
*/
type Role int

const (
	Admin Role = 1
	User  Role = 2
	Guest Role = 3
)

/*
	ENUM 3: Enum with String() method
*/
type OrderStatus int

const (
	Created OrderStatus = iota
	Shipped
	Delivered
	Cancelled
)

func (o OrderStatus) String() string {
	return [...]string{
		"CREATED",
		"SHIPPED",
		"DELIVERED",
		"CANCELLED",
	}[o]
}

/*
	ENUM 4: String-based enum (Best for APIs)
*/
type PaymentMethod string

const (
	CreditCard PaymentMethod = "CREDIT_CARD"
	UPI        PaymentMethod = "UPI"
	NetBanking PaymentMethod = "NET_BANKING"
)

/*
	ENUM Validation
*/
func IsValidStatus(s Status) bool {
	switch s {
	case Pending, Approved, Rejected:
		return true
	default:
		return false
	}
}

func main() {

	// Integer enum
	var s Status = Approved
	fmt.Println("Status:", s)
	fmt.Println("Valid Status?", IsValidStatus(s))

	// Custom value enum
	var r Role = Admin
	fmt.Println("Role:", r)

	// Enum with String method
	var o OrderStatus = Shipped
	fmt.Println("Order Status:", o)

	// String enum
	var p PaymentMethod = UPI
	fmt.Println("Payment Method:", p)
}
