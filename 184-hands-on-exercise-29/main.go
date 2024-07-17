package main

import (
	"fmt"
)

func init() {
	fmt.Printf("\n")
	fmt.Printf("...initializing")
	fmt.Printf("\n")
	fmt.Printf("\nthis program will execute an outer and inner loop")
	fmt.Printf("both the outer and inner will execute 5 times...\n")
}

func main() {
	fmt.Printf("\n")
	for i := 0; i < 5; i++ {
		fmt.Printf("outer loop # %v\n", i)
		for z := 0; z < 5; z++ {
			fmt.Printf("	inner loop # %v\n", z)
		}
	}
}
