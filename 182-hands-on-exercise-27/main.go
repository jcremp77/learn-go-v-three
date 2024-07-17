package main

import (
	"fmt"

	"math/rand"
)

func init() {
	fmt.Printf("\n")
	fmt.Printf("..... initializing package ......\n")
}

func main() {
	for i := 0; i < 42; i++ {
		x := rand.Intn(5)
		fmt.Printf("\n")
		fmt.Printf("ITERATION #%v\n", i)
		fmt.Printf("Value of x: %v\n", x)
		if x == 4 {
			fmt.Printf("\nvalue of 4 detected breaking loop\n")
			break // break here
		}
	}
	fmt.Printf("\n")
	fmt.Println("Exiting program")
}
