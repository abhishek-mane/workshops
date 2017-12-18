package interfaces

import "fmt"

type duration int

func (d *duration) pretty() string {
	return fmt.Sprintf("Duration: %d", *d)
}

/*
LearningMethodSets -
*/
func LearningMethodSets() {

	fmt.Println(duration(45).pretty())

}
