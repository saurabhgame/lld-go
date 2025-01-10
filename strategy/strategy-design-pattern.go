package main

import "fmt"

// PaymentStrategy defines the interface for payment strategies
type PaymentStrategy interface {
	Pay(amount int)
}

//-----------------------------------------------------------------------------------------------------------------------------

// CreditCardPayment represents a payment method using a credit card
type CreditCardPayment struct {
	CardNumber string
}

func (c *CreditCardPayment) Pay(amount int) {
	fmt.Printf("Paid %d using Credit Card: %s\n", amount, c.CardNumber)
}

type PayPalPayment struct {
	Email string
}

func (p *PayPalPayment) Pay(amount int) {
	fmt.Printf("Paid %d using PayPal: %s\n", amount, p.Email)
}

type BitcoinPayment struct{}

func (b *BitcoinPayment) Pay(amount int) {
	fmt.Printf("Paid %d using Bitcoin.\n", amount)
}

//-----------------------------------------------------------------------------------------------------------------------------

type PaymentContext struct {
	Strategy PaymentStrategy
}

func (p *PaymentContext) ExecutePayment(amount int) {
	p.Strategy.Pay(amount)
}

//-----------------------------------------------------------------------------------------------------------------------------

func main() {

	// Pay using Credit Card
	creditCard := &CreditCardPayment{CardNumber: "1234-5678-9876-5432"}
	context := &PaymentContext{Strategy: creditCard}
	context.ExecutePayment(100)

	// Pay using PayPal
	payPal := &PayPalPayment{Email: "user@example.com"}
	context.Strategy = payPal
	context.ExecutePayment(200)

	// Pay using Bitcoin
	bitcoin := &BitcoinPayment{}
	context.Strategy = bitcoin
	context.ExecutePayment(300)

}
