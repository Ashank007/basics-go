package main

import (
	"fmt"
	"slices"
)

func main () {

	//  var nums []int
	// fmt.Println(nums)
	// fmt.Println(len(nums))

	// var nums = make([]bool,0,4)
	// fmt.Println(cap(nums))
	// fmt.Println(nums)	

	// var nums = make([]int,5)
	// nums = append(nums, 2);
	// var nums2 = make([]int,len(nums))
	// copy(nums2,nums)
	// fmt.Println(nums2)

	// var nums = []int {1,2,3}
	// fmt.Println(nums[:3])
	//
	var nums1 = []int {1,2,3}
	var nums2 = []int {1,2,3}

	fmt.Println(slices.Equal(nums1,nums2))

	var nums = [][]int{{2,3},{4,5}}
	fmt.Println(nums)
}
