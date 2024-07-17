package main

import (
	"fmt"

	"math/rand"
)

func main() {

	// Statement Statement idiom func Intn(n int) int will choose
	// a number between 0 - 19 and assign to 'x'
	for i := 0; i < 100; i++ {
		fmt.Printf("BEG ITERATION # %v\n", i)
		x := rand.Intn(5)
		if x == 3 {
			fmt.Printf("x is: %v\n", x)
		}
	}
}
