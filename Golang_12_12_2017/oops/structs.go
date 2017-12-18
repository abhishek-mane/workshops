package oops

import "fmt"

// Rectangle
/*
	-> structs - declare, initialize, use, access
	-> methods - value receiver, pointer receiver
				You can only declare a method with a receiver whose type is defined in the
				same package as the method. You cannot declare a method with a receiver whose type
				is defined in another package (which includes the built-in types such as int)
	-> structs are values
*/
type Rectangle struct {
	L int
	W int
}

func (r Rectangle) showValueRecvMethod() {
	fmt.Println("Inside showValueRecvMethod()")
}

func (r *Rectangle) showPointerRecvMethod() {
	fmt.Println("Inside showPointerRecvMethod()")
}

func LeaningStructs() {

	r := Rectangle{
		4,
		6,
	}
	p := &r

	r.showValueRecvMethod()
	r.showPointerRecvMethod()

	p.showPointerRecvMethod()
	(*p).showValueRecvMethod()

}
