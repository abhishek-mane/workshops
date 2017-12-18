package datastructs

import (
	"fmt"
)

// LearningSlices -
// 				Slices are tiny objects that abstract and manipulate an underlying array
/*
	-> creating slice using make function
	-> using literal
	-> []string{99: ""}
	-> empty vs nil slices
	-> slicing - creating new slice from existing slice, using indexes [1:2:3]
	-> append
	-> slice as parameter to the function
*/
func LearningSlices() {

	var slice = []int{1, 2, 3, 4, 5, 6}

	s2 := slice[1:4]

	s2 = append(s2, 99)

	fmt.Println(s2)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))

	slice[1] = 56

	fmt.Println(slice)
	fmt.Println(s2)
}
