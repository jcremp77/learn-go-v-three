package main

import (
	"fmt"

	"math/rand"
)

func init() {
	fmt.Printf("\n")
	fmt.Printf("..... initializing package ......")
	fmt.Printf("\n")
}

func main() {
	for i := 0; i < 42; i++ {
		x := rand.Intn(5)
		fmt.Printf("\n")
		fmt.Println("======================================================")
		fmt.Printf("\n")
		switch {
		case x >= 0 && x <= 4:
			fmt.Printf("variable 'x' is a non-negative pseudo-random integer\n")
			fmt.Printf("and it is inclusive of the value 0 and exclusive of 5\n")
			fmt.Printf("and the value of x is: %v\n", x)
		default:
			fmt.Printf("...")
		}
		fmt.Printf("\n")
		fmt.Printf("iteration # %v", i)
		fmt.Printf("\n")
	}
}
