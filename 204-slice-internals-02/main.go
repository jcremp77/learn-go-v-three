package main

import (
	"fmt"
)

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	b := make([]int, 6)
	// copy contents the contents of the array
	// that the slice 'a' is pointing to into
	// a new array which slice 'b' now points to
	copy(b, a)
	fmt.Println("slice 'b'", b)

	fmt.Println("--------------------------")
	a[0] = 7
	fmt.Println("slice 'a'", a)
}
