package oops

import (
	"fmt"
)

// Person -
type Person struct {
	Name   string
	Gender string
}

// SayHello -
func (p Person) SayHello() {
	fmt.Println("Hello , I'm ", p.Name)
}

// Doctor -
type Doctor struct {
	Person // Embedded type
	Degree string
}

// SayHello -
func (d Doctor) SayHello() {
	fmt.Println("Hello, I'm Dr.", d.Name, ", ", d.Degree)
}

// Diagnose -
func (d Doctor) Diagnose() {
	fmt.Println("It's Lupus.")
}

// LearningComposition - OR Type Embedding
// 		- Go allows you to take existing types and both extend and change their behavior
//		- inner type promotions
//		- embedding, overriding, overriden methods, vars
//		- inner type never loses its identity
//		- if inner type implements any interface, outer type bydefault implements that interface
//		- What if the outer type doesn’t want to use the inner type’s implementation of the interface, define own implementation
func LearningComposition() {

	d := Doctor{
		Person{
			"Max",
			"Male",
		},
		"MBBS",
	}

	fmt.Println(d.Name)
	fmt.Println(d.Gender)
	fmt.Println(d.Degree)
	d.SayHello()
	d.Person.SayHello()
	d.Diagnose()
}
