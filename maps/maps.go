package main

import (
	"fmt"
	"maps"
)

func main () {

	// m := make(map[string]string)
	//
	// m["name"] = "golang"
	// m["area"] = "backend"
	//
	// fmt.Println(m["name"])
	// fmt.Println(m["area"])
	// fmt.Println(m["font"])

	// m := make(map[string]int)
	// m["age"] = 30
	// fmt.Println(m["age"])
	// fmt.Println(m["phone"])
	// fmt.Println(len(m))
	// delete(m,"age")
	// clear(m)
	// fmt.Println(m["age"])
	
	// m := map[string]int{"price":40,"phones":3}
	// // fmt.Println(m["price"])
	// // fmt.Println(m["phones"])
	// v,ok := m["phones"]
	// fmt.Println(v)
	// if ok {
	//   fmt.Println("All Ok")
	// } else{
	//   fmt.Println("Not Ok")
	// }

	 m1 := map[string]int{"price":40,"phones":3}
	 m2 := map[string]int{"price":40,"phones":3}
	 fmt.Println(maps.Equal(m1,m2))


}
