package main

import (
	"fmt"
)

// && and
func main() {
	// && AND
	if x := 1; x > 0 && x <= 1 {
		fmt.Println("x is greater than 0 and less than or equal to 1")
	}

	// || OR
	if x := 1; x < 5 || x != 0 {
		fmt.Println("x is either less than 5 or is not equal to 0")
	}

	// ! NOT
	if x := 9; x != 10 {
		fmt.Println("x does not equal 10")
	}
}
