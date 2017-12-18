package errors

import "fmt"

// defer
// - executes in LIFO(stack) order
func TestFunc() {

	defer func() {
		r := recover()
		fmt.Println("recovered from panic => ", r)
	}()

	panic("No more memory")
	// str := recover()
	fmt.Println("str")
}
