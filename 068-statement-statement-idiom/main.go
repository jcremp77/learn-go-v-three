package main

import (
	"fmt"

	"math/rand"
)

func main() {
	x := 20

	/*
		func Intn(n int) int
		will choose a number between 0 - 19
		as x above has been assigned the
		value of 20
	*/
	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v and that is GREATER THAN OR EQUAL to x which is %v\n", z, x)
	} else {
		fmt.Printf("z is %v and that is LESS THAN x which is %v\n", z, x)
	}
}
