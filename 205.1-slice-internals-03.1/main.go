package main

import (
	"fmt"
	"sort"
)

func main() {
	y := []float64{3, 1, 4, 2}
	fmt.Println(medianTwo(y))
	fmt.Println("after medianTwo", y)
	// slice 'y' was never sorted as when the value of 'y'
	// was passed to the function 'medianTwo' the value of 'y'
	// was assigned to 'x', and 'x' was a pointer to the same
	// array that was underneath 'y'.
	// The 'Println' statement above is referencing the slice
	// of 'y' which was 'never sorted'
	fmt.Println("---------------------")
}

func medianTwo(x []float64) float64 { // {3, 1, 4, 2} passed to slice 'x'
	// allocate a new underlying array for slice 'n'
	n := make([]float64, len(x))
	// contents of array underneath 'x' copied to new array for slice 'n'
	copy(n, x)
	sort.Float64s(n) // sort performed on slice 'n' {3, 1, 4, 2}
	fmt.Println(n)
	i := len(n) / 2 // len(n) == 4; 4 / 2 = 2
	// i := 2
	if len(n)%2 == 1 { // 2%2 == 0; false
		return n[i/2]
		// when you divide
		// you get the whole number quotient
		// without the remainder modulus
		// https://go.dev/play/p/2r5WrelUEh7
	}
	return (n[i-1] + n[i]) / 2
	// (n[2-1] + n[2]) / 2
	// (1 + 4) / 2 = 2.5
}
