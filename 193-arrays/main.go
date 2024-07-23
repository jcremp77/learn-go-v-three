package main

import (
	"fmt"
)

func init() {
	fmt.Printf("...initializing\n\n")
}

func main() {
	// array literal
	a := [3]int{42, 43, 44}
	fmt.Printf("%#v \t\t\t\t\t %T\n", a, a)

	b := [...]string{"Hello", "Gophers!"}
	fmt.Printf("%#v \t\t\t\t %T\n", b, b)

	var c [2]int
	c[0] = 7
	c[1] = 8
	fmt.Printf("%#v \t\t\t\t\t\t %T\n", c, c)

	// declare a variable of type [7]int
	var ni [7]int
	// assign a value to index position zero
	ni[0] = 42
	fmt.Printf("%#v \t\t\t\t %T\n", ni, ni)

	// array literal
	ni2 := [4]int{55, 56, 57, 58}
	// another way to write it
	// ni2 := [...]int{55, 56, 57, 58}
	fmt.Printf("%#v \t\t\t\t\t %T\n", ni2, ni2)

	// array literal
	// have compiler count elements
	ns := [...]string{"chocolate", "vanilla", "strawberry"}
	fmt.Printf("%#v \t %T\n", ns, ns)

	// use the builtin len function
	// https://pkg.go.dev/builtin#len
	fmt.Printf("\n")
	fmt.Printf("length of array 'ni' is: %v\n", len(ni))
	fmt.Printf("length of array 'ni2' is: %v\n", len(ni2))
	fmt.Printf("length of array 'ns' is: %v\n", len(ns))
	// print just the length of the array
	fmt.Println()
	fmt.Println(len(ni))
	fmt.Println(len(ni2))
	fmt.Println(len(ns))
}
