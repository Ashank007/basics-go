package main

import (
	"fmt"
)

// func processNum(numChan chan int) {
//
// 	for num := range numChan {
// 		fmt.Println("Processing a Number", num)
// 		time.Sleep(time.Second * 1)
// 	}
// }

// func sum(result chan int,num1 int,num2 int) {
// 	numResult := num1 + num2
// 	result <- numResult
// }

// func task(done chan bool) {
// 	defer func () { done <- true } ()
//   fmt.Println("Processing....")
// }

// func emailSender(emailChan chan string, done chan bool) {
// 	defer func() { done <- true }()
// 	for email := range emailChan {
// 		fmt.Println("Sending Email to", email)
// 		time.Sleep(time.Second * 1)
// 	}
// }

func main() {

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "golang"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("Received Data from Chan1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("Received Data from Chan2", chan2Val)
		}
	}

	// emailChan := make(chan string, 100)
	// doneChan := make(chan bool)
	// go emailSender(emailChan, doneChan)
	//
	// for i := 0; i < 10; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }
	//
	// fmt.Println("Done Sending")
	//
	// close(emailChan)
	//
	// <-doneChan

	// done := make(chan bool)
	//
	// go task(done)
	//
	// <- done

	// result := make(chan int)
	//
	// go sum(result,4,5)
	//
	// res := <- result // blocking
	//
	// fmt.Println(res)

	// numChan := make(chan int)
	//
	// go processNum(numChan)
	//
	// for {
	// 	numChan <- rand.Intn(100)
	// }
	//
	// messageChan := make(chan string)
	//
	// messageChan <- "Ping"
	//
	// msg := <- messageChan
	//
	// fmt.Println(msg)

}
