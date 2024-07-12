package main

/* write a program that uses the following:
● for a VARIABLE storing a VALUE of TYPE
○ string
○ int
○ float64
● use print verbs to show
○ value
○ type
https://go.dev/play/p/
*/

import "fmt"

func main() {
	s, i, f := "meow", 42, 42.42
	fmt.Printf("%v \t is of type: %T\n", s, s)
	fmt.Printf("%v \t is of type: %T\n", i, i)
	fmt.Printf("%v \t is of type: %T\n", f, f)
}
