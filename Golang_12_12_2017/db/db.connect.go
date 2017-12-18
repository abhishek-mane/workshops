package db

import (
	"fmt"
)

func init() {
	fmt.Println("Another init function")
}

func ConnectDB() {
	fmt.Println("Connecting to DB & returning connection object")
}
