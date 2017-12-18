package interfaces

// Animal -
type Animal interface {
	Type() string
	Swim() string
}

// Dog -
type Dog struct {
	Name  string
	Breed string
}

func (d *Dog) Swim() string {
	return "Paddle"
}

func (d *Dog) Type() string {
	return "Doggie"
}

/*
LearningInterfaces1 -
	- if you implement an interface using a pointer receiver, then only pointers of that type implement the interface.
	- If you implement an interface using a value receiver, then both values and pointers of that type implement the interface.
*/
func LearningInterfaces1() {

	// d := Dog{
	// 	Name: "Max",
	// }

	// f := Frog{
	// 	"Charlie",
	// 	"Green",
	// }

	// var i Animal = d

	// fmt.Println(i.Swim())
}
