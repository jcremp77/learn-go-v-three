package main

import (
	"fmt"
)

func init() {
	fmt.Printf("\n")
	fmt.Printf("......initializing......")
}

func main() {
	// create a for loop using only a condition
	// 'i' decrements from 10 to 0
	// printing of 'i' inclusive of 0 value
	for i := 10; i >= 0; i-- {
		fmt.Printf("%v\n", i)
	}

}
