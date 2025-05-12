package main

import "fmt"

type paymenter interface {
	pay(amount float32)
	refund (amount float32,account string)
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	//razorpayPaymentGw := razorpay{}
	// razorpayPaymentGw.pay(amount)
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	// logic to make payment
	fmt.Println("Making Payment using RazorPay", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	// Logic to make payment
	fmt.Println("Making Payment using Stripe", amount)
}

type fakepayment struct{}

func (p fakepayment) pay(amount float32) {
	// Logic to make payment
	fmt.Println("Making Payment using Fake Gateway",amount)
}

type paypal struct {} 

func (p paypal) pay(amount float32) {
	// Logic to make payment
	fmt.Println("Making Payment using Paypal",amount)
}

func (p paypal) refund (amount float32,account string) {
} 

func main() {
	//razorpayPaymentGw := razorpay{}
	//fakePayementGw := fakepayment {}
	paypalPaymentGw := paypal {}
	// stripePaymentGw := stripe{}
	newPayment := payment{
   gateway: paypalPaymentGw,
	}

	newPayment.makePayment(1000)

}
