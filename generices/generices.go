package main

import "fmt"

// func printSlice[T int | string | bool](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func printStringSlice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

type Stack [T any] struct {
	elements []T
}

func main() {
	stack := Stack[string]{elements: []string{"golang"}}
	fmt.Println(stack)
	// nums := []int {1,2,3}
	// names := []string {"golang","typescript"}
	// vals := [] bool {true,false,true}
	// printSlice(names)
	// printSlice(nums)
	// printSlice(vals)


}
