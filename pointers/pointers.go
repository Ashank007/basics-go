package main

import "fmt"

func changeInNum(num int) {

	num = 5
	fmt.Println("In changeNum", num)

}

func changeNum(num *int) {
	*num = 10
}

func main() {

	num := 1

	changeInNum(num)

	changeNum(&num)

	fmt.Println("After changeNum in main", num)

}
