package datastructs

import (
	"fmt"
)

// LearningArrays -
// array in go is value not reference
// - declaring array, len, array literal, copying arrays, array of pointers, passing arrays as parameters reference

func test(arr *[5]int) {

	fmt.Println(*arr)
	(*arr)[0] = 99
}

func LearningArrays() {

	var arr1 = [5]int{1, 2, 3, 4, 5}

	p := &arr1

	test(p)

	fmt.Println(arr1)
}
