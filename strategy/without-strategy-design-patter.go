package main

import "fmt"

// PaymentContext handles payment processing directly, without strategies
type PaymentContext2 struct {
	Method string
}

// ExecutePayment processes the payment based on the method
func (p *PaymentContext2) ExecutePayment(amount int) {
	if p.Method == "credit_card" {
		fmt.Printf("Paid %d using Credit Card.\n", amount)
	} else if p.Method == "paypal" {
		fmt.Printf("Paid %d using PayPal.\n", amount)
	} else if p.Method == "bitcoin" {
		fmt.Printf("Paid %d using Bitcoin.\n", amount)
	} else {
		fmt.Println("Invalid payment method.")
	}
}

func main() {
	// Payment using Credit Card
	context := &PaymentContext2{Method: "credit_card"}
	context.ExecutePayment(100)

	// Payment using PayPal
	context.Method = "paypal"
	context.ExecutePayment(200)

	// Payment using Bitcoin
	context.Method = "bitcoin"
	context.ExecutePayment(300)
}
