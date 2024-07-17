package main

import (
	"fmt"
)

func main() {
	// x = the number of times 40 will go into 83
	// y = the remainder of the 83 divided by 40
	x := 83 / 40
	y := 83 % 40
	fmt.Printf("40 goes into 83, %v times with a remainder of %v", x, y)

	// while 'i' is less than 100
	// get the remainder of 'i' divided by 2
	// if there is a remainder then continue in the loop
	// if there is a '0' remainder then print
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Printf("%v  is an even number\n", i)
		}
	}
}
