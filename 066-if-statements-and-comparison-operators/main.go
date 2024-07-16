package main

import (
	"fmt"
)

func main() {
	x := 10

	// If statements
	if x > 9 {
		fmt.Println("10 is greater than 9")
	}

	if x != 9 {
		fmt.Println("10 is not equal to 9")
	}

	if x == 10 {
		fmt.Println("10 is equal to 10")
	}

	// If-Else statement
	if x < 10 {
		fmt.Printf("%v is less than 10\n", x)
	} else {
		fmt.Printf("%v is not less than 10\n", x)
	}

	// If-Else If-Else
	if x >= 11 {
		fmt.Printf("%v is greater than or equal to 11\n", x)
	} else if x <= 9 {
		fmt.Printf("%v is less than or equal to 9\n", x)
	} else {
		fmt.Printf("%v is greater than 9 but less than 11\n", x)
	}

	// If statement initialization
	if z := 99; z >= 95 {
		fmt.Printf("%v is greater or equal to 95", z)
	}

}
