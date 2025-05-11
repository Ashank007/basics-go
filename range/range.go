package main

import "fmt"

func main () {

	// nums := []int {5,6,7}

	// for i := 0 ; i<len(nums) ; i++ {
	//    fmt.Println(nums[i])
	// }

	// sum := 0 
	// for i,num := range nums {
	// 	fmt.Println(i,num)
	//    sum = sum + num
	// }
	// fmt.Println(sum)

	// m := map[string]string {"fname":"john","lname":"loa"}
	//
	// for k,val := range (m) {
	//   fmt.Println(k,val)
	// }

	for i,c := range("golang") {
    fmt.Println(i,c)
	}
}
