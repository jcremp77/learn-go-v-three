package main

import (
	"fmt"
)

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// from slice 'x' index 0 up to 2 (not including index 3)
	// from slice 'x' index 7 up to the end of slice 'x'
	// ... 'exploding' the slice into arguments
	x = append(x[0:3], x[7:]...)
	fmt.Println("slice x: ", x)
}
