package main

import (
	"fmt"
)

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	x1 := append(x, 52)
	fmt.Println("slice x1: ", x1)

	x2 := append(x1, 53, 54, 55)
	fmt.Println("slice x2: ", x2)

	y := []int{56, 57, 58, 59, 60}
	x3 := append(x2, y...)
	fmt.Println("slice x3: ", x3)
}
