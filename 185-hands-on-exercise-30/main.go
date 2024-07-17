package main

import (
	"fmt"
)

func init() {
	fmt.Printf("...initializing\n\n")
}

func main() {
	// a for range loop to print each value
	// and the index position of each value
	xi := []int{42, 43, 44, 45, 46, 47}
	for i, v := range xi {
		fmt.Printf("index=%v, value=%v\n", i, v)
	}
}
