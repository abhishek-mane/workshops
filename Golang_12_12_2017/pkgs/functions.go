package pkgs

import "fmt"

// basic function
func add(x int, y int) int {
	return x + y
}

// func can return any number of values
func swap(x, y string) (string, string) {
	return y, x
}

// func can have named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // "Naked" returns, should be used only in short functions
}

// variadic functions
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func LearningFunctions() {

	f := func(a, b int) int {
		return a + b
	}

	funcAreValues(f)

}

func funcAreValues(f func(int, int) int) int {
	return f(1, 4)
}
