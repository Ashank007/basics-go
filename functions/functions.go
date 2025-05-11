package main

import "fmt"

func add (a,b int) int {
	return a + b
}

// func getLanguages () (string,string,string) {
// 	return "golang","javascript","c"
// }

// func processIt() func(a int) int {
//   return func(a int) int {
//    return 4
// 	}
// }

func sum(nums ...int) int {
	total := 0
	for _,num := range(nums) {
    total = total + num
	}
	return total
}

func main () {

	result := add(2,3)

	ans := sum(1,2,3,4,5,6,7,8,9)

	fmt.Println(result)

	fmt.Println(ans)

}
