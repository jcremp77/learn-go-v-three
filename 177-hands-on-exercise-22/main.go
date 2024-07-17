package main

import (
	"fmt"

	"math/rand"
)

func init() {
	fmt.Printf("\n")
	fmt.Println("++++ Initializing ++++n")
	fmt.Printf("\n")
}

func main() {
	// create 2 pseudo-random integers
	// inclusive of 0 and exclusive of 10
	// assign variables 'x' and 'y'
	x := rand.Intn(10)
	y := rand.Intn(10)
	fmt.Printf("value of x: %v\n", x)
	fmt.Printf("value of y: %v\n", y)
	fmt.Printf("\n")
	switch {
	case x < 4 && y < 4:
		fmt.Printf("%v and %v are both less than 4\n", x, y)
	case x > 6 && y > 6:
		fmt.Printf("%v and %v are both greater than 6\n", x, y)
	case x >= 4 && x <= 6:
		fmt.Printf("%v is greater than or equal to 4 and %v is less than or equal to 6\n", x, x)
	case y != 5:
		fmt.Printf("%v is not equal to 5\n", y)
	default:
		fmt.Printf("none of the previous caes were met\n")
	}
}
