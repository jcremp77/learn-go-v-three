package main

import (
	"fmt"
	"strings"
)

func main() {
	x := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("The sum is", x)

	fmt.Println()

	// Example usage
	result1 := Greet("nobody")
	fmt.Println(result1) // Output: nobody []

	s := []string{"Joe", "Anna", "Eileen"}
	result2 := Greet("Hello:", s...)
	fmt.Println(result2) // Output: Hello: [Joe Anna Eileen]
}

// func (r receiver) identifier(p parameters) (return(s)) {code}

func sum(ii ...int) int {
	fmt.Println(ii)
	fmt.Printf("%T\n", ii)

	n := 0
	for _, v := range ii {
		n += v
	}
	return n
}

// Greet takes a prefix and a variadic number of names, and returns a formatted string
func Greet(prefix string, who ...string) string {
	// Join the 'who' slice into a single string with ", " separator
	names := strings.Join(who, ", ")
	// Return the concatenated result
	return fmt.Sprintf("%s %s", prefix, names)
}
