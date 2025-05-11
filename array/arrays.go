package main

import "fmt"

func main () {

	var nums [4]int 

	nums[0] = 1

	fmt.Println(nums[0])
	fmt.Println(len(nums))
	fmt.Println(nums)

	var vals[4]bool

	vals[2] = true

	fmt.Println(vals)

	var names[3] string
	names[0] = "Golang"
	fmt.Println(names)

	new_nums := [3]int {4,5,9}

	fmt.Println(new_nums)

}
