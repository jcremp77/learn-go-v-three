package main

import (
	"fmt"
)

func main() {
	xi := []int{42, 43, 44}
	fmt.Println(xi)
	fmt.Println()
	// append to existing slice
	// existing variable and not a new declaration
	xi = append(xi, 45, 46)
	fmt.Println(xi)
}
