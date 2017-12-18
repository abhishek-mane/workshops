package mathpkg

type alertCounter int

func New(value int) alertCounter {
	return alertCounter(value)
}

func Main() {

	/*
			This program compiles and runs, but why? The New function is returning a
		value of the unexported type alertCounter, yet main is able to accept that value and
		create a variable of the unexported type.
			This is possible for two reasons. First, identifiers are exported or unexported, not
		values. Second, the short variable declaration operator is capable of inferring the type
		and creating a variable of the unexported type. You can never explicitly create a variable
		of an unexported type, but the short variable declaration operator can.
	*/
	// counter := mathpkg.New(10)
	// fmt.Println("Counter = ", counter)

}
