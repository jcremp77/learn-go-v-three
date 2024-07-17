package main

import (
	"fmt"

	"math/rand"
)

func main() {
	// x := 20

	/*	Statement Statement idiom
			func Intn(n int) int
			will choose a number between 0 - 19
			as x above has been assigned the
			value of 20

		if z := 2 * rand.Intn(x); z >= x {
			fmt.Printf("z is %v and that is GREATER THAN OR EQUAL to x which is %v\n", z, x)
		} else {
			fmt.Printf("z is %v and that is LESS THAN x which is %v\n", z, x)
		}
	*/

	// comma ok
	var timeZone = map[string]int{
		"UTC": 0 * 60 * 60,
		"EST": -5 * 60 * 60,
		"CST": -6 * 60 * 60,
		"MST": -7 * 60 * 60,
		"PST": -8 * 60 * 60,
	}

	offset := timeZone["EST"]
	fmt.Println(offset)
	
		var seconds int
		var ok bool
		seconds, ok = timeZone[tz]

	func offset(tz string) int {
		if seconds, ok := timeZone[tz]; ok {
			return seconds
		}
		log.Println("unknown time zone:", tz)
		return 0
	}
	
}
