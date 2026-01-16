package main

import "fmt"

// PaymentGateway interface
type PaymentGateway interface {
	Pay(amount float64) string
}

// Razorpay implementation
type Razorpay struct{}

func (r Razorpay) Pay(amount float64) string {
	return fmt.Sprintf("Paid ₹%.2f using Razorpay", amount)
}

// Stripe implementation
type Stripe struct{}

func (s Stripe) Pay(amount float64) string {
	return fmt.Sprintf("Paid ₹%.2f using Stripe", amount)
}

// PayPal implementation
type PayPal struct{}

func (p PayPal) Pay(amount float64) string {
	return fmt.Sprintf("Paid ₹%.2f using PayPal", amount)
}

// Payment Service (uses interface)
type PaymentService struct {
	gateway PaymentGateway
}

func (ps PaymentService) MakePayment(amount float64) {
	fmt.Println(ps.gateway.Pay(amount))
}

func main() {

	service := PaymentService{gateway: Razorpay{}}
	service.MakePayment(500)

	service = PaymentService{gateway: Stripe{}}
	service.MakePayment(1000)

	service = PaymentService{gateway: PayPal{}}
	service.MakePayment(750)
}
