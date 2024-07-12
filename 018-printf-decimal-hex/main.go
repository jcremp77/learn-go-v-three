package main

import (
	"fmt"
)

func main() {
	adams := 42
	fmt.Printf("42 as binary, %b \n", adams)
	fmt.Printf("42 as hexadecimal, %x \n", adams)

	// print these values as both binary and hexadecimal
	a, b, c, d, e, f := 100, 101, 102, 103, 104, 105
	fmt.Printf("%v \t %b \t %#X \n", a, a, a)
	fmt.Printf("%v \t %b \t %#X \n", b, b, b)
	fmt.Printf("%v \t %b \t %#X \n", c, c, c)
	fmt.Printf("%v \t %b \t %#X \n", d, d, d)
	fmt.Printf("%v \t %b \t %#X \n", e, e, e)
	fmt.Printf("%v \t %b \t %#X \n", f, f, f)
}
