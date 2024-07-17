package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// create a non-negative pseudo-random int; rand.Intn()
	// store the random int as variable 'x'
	x := rand.Intn(250)
	//x := 0
	fmt.Printf("x is a non-negative pseudo-random number with the value of: %v\n", x)

	// write sequential 'if' statements to check the value of 'x'
	// print out the range into which the value of 'x' corresponds
	if x <= 100 {
		fmt.Printf("%v is between 0 and 100\n", x)
	} else if x >= 101 && x < 200 {
		fmt.Printf("%v is between 101 and 200\n", x)
	} else {
		fmt.Printf("%v is between 201 and 250\n", x)
	}

	// print out the value of 'y'; rand.Intn()
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
}
