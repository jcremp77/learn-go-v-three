package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("\nInitializing...\n")
}

func main() {
	// create a non-negative pseudo-random int, between 0 and 499
	// store the random int as variable 'x'
	x := rand.Intn(500)

	//x := 0
	fmt.Printf("x is a non-negative pseudo-random number with the value of: %v\n", x)

	// write switch statements to check the value of 'x'
	// print out the range into which the value of 'x' corresponds
	switch {
	case x <= 100:
		fmt.Printf("%v is between 0 and 100\n", x)
	case x >= 101 && x <= 200:
		fmt.Printf("%v is between 101 and 200\n", x)
	case x >= 201 && x <= 300:
		fmt.Printf("%v is between 201 and 300", x)
	default:
		fmt.Printf("%v is greater than 300\n", x)
	}
}
