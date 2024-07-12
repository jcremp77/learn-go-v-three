package main

/*
write a program that uses the following:
● var for zero value
● short declaration operator
● multiple initializations
● var when specificity is required
● blank identifier
*/

import "fmt"

var x int

func main() {
	fmt.Println("1. variable x has a value of:", x)

	i, h := 5, 6
	fmt.Println("2. variables i and h are short declarations valued as:", i, h)

	_ = 5 // this is a blank identifier

	var y string = "hello"
	fmt.Println("3. y is a string variable with a value of:", y)

	var z float64 = 42.42
	fmt.Printf("4. variable z has a value of %v and a type of %T", z, z)
}
