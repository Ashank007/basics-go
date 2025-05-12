package main

import "fmt"

// type order struct {
// 	id        string
// 	amount    float32
// 	status    string
// 	createdAt time.Time
// }
//
// func newOrder(id string,amount float32,status string) *order {
//
// 	myOrder := order{
// 		id:     id,
// 		amount: amount,
// 		status: status,
// 	}
// 	return &myOrder
// }

// func (o *order) changeStatus(status string) {
// 	o.status = status
// }
//
// func (o order) getAmount() float32 {
// 	return o.amount
// }

func main() {

	language := struct {
		name   string
		isGood bool
	}{"golang", true}

	fmt.Println(language)

	// myOrder := newOrder("1",400.00,"Done")
	//
	// fmt.Println(myOrder.amount)

	// myOrder2 := order{
	// 	id:        "2",
	// 	amount:    1000.0,
	// 	status:    "Pending",
	// 	createdAt: time.Now(),
	// }
	//
	// fmt.Println("Order 2 Struct ", myOrder2)

}
