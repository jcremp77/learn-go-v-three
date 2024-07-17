package main

import (
	"fmt"
	"math/rand"
)

var x int = 40

func main() {
	// first statement to run
	fmt.Println("This is the first statement to run")
	// second statement to run
	fmt.Println("This is the second statement to run")

	z := 2 * rand.Intn(x)

	switch {
	case z < 50:
		fmt.Printf("%v is less than 50\n", z)
		fallthrough
	case z > 25:
		fmt.Printf("%v is greater than 25\n", z)
		fallthrough
	case z != 39:
		fmt.Printf("%v is not equal to 39\n", z)
		fallthrough
	case z >= 10:
		fmt.Printf("%v is greater than or equal 10\n", z)
	}

	switch z {
	case 40:
		fmt.Println("x is 40")
		fallthrough
	case 41:
		fmt.Println("printed because of fallthrough: z is 41")
		fallthrough
	case 42:
		fmt.Println("printed because of fallthrough: z is 42")
		fallthrough
	case 43:
		fmt.Println("printed because of fallthrough: z is 43")
		fallthrough
	default:
		fmt.Printf("printed due to the default case for z; z equals %v", z)
	}
}
