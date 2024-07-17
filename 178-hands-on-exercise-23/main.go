package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Printf("\n")
	fmt.Printf(".....initializing.....")
	fmt.Printf("\n")
}

func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("%v\n", i)
	}

	for a := 0; a < 100; a++ {
		fmt.Printf("\n")
		fmt.Printf("=================================================================")
		fmt.Printf("\n")
		fmt.Printf("\n")
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
		fmt.Printf("\n")
		fmt.Printf("this is the %v iteration in this loop\n", a)
	}
}
