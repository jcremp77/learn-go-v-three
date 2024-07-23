package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []float64{3, 1, 4, 2}

	fmt.Println(medianOne(n))
	fmt.Println("after medianOne", n)

	y := []float64{3, 1, 4, 2}
	fmt.Println(medianTwo(y))
	fmt.Println("after medianTwo", y)
}

// uses the same underlying array
// everything is pass by value in go
// the value is being passed into the func
// and then assigned to x
func medianOne(x []float64) float64 {
	sort.Float64s(x)        // 1, 2, 3, 4
	i := len(x) / 2         // i := 4 / 2 = 2
	fmt.Println(len(x))     // value of 4
	fmt.Println(len(x) % 2) // value of 0
	if len(x)%2 == 1 {      // if 4%2 == 1
		return x[i/2]
	}
	return (x[i-1] + x[i]) / 2 // (x[2-1] + x[2]) / 2
} // (2 + 3) / 2
// 5 / 2 = 2.5

func medianTwo(x []float64) float64 {
	// allocate a new underlying array
	n := make([]float64, len(x))
	copy(n, x)
	fmt.Println(n) // [3, 1, 4, 2]

	sort.Float64s(n)
	i := len(n) / 2    // len(n) == 4; 4 / 2 = 2
	fmt.Println(i)     // i := 2
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
