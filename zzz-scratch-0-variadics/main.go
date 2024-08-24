package main

import (
	"fmt"
	"strings"
)

// Greet takes a prefix and a variadic number of names, and returns a formatted string
func Greet(prefix string, who ...string) string {
	// Join the 'who' slice into a single string with ", " separator
	names := strings.Join(who, ", ")
	// Return the concatenated result
	return fmt.Sprintf("%s %s", prefix, names)
}

func main() {
	// Example usage
	result1 := Greet("nobody")
	fmt.Println(result1) // Output: nobody

	s := []string{"Joe", "Anna", "Eileen"}
	result2 := Greet("Hello:", s...)
	fmt.Println(result2) // Output: Hello: Joe, Anna, Eileen
}
