package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 50; i++ {

		x := rand.Intn(10)
		if x%2 != 0 {
			fmt.Printf("\n")
			fmt.Printf("Iteration # %v\n", i)
			fmt.Printf("%v is an odd number\n", x)
			continue
		}
	}
	fmt.Printf("Exiting program")
}
