package main

import (
	"fmt"
)

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println("slice 'x'")
	fmt.Println("====================")
	for i, v := range x {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}
	fmt.Println("====================")
	fmt.Printf("slice 'x' Type: %T\n", x)
}
