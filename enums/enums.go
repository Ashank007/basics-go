package main

import "fmt"

type OrderStatus string

const (
	Received  OrderStatus = "received"
	Confirmed             = "confirmed"
	Prepared              = "prepared"
	Delivered             = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing Order Status to", status)
}

func main() {
	changeOrderStatus(Received)
}
