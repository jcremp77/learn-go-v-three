package main

import (
	"fmt"
)

func main() {
	// for range loop
	// data structure - slice
	// a slice has an 'index' and 'value'; index starts at '0'
	xi := []int{42, 43, 44, 45, 46, 47}
	for i, v := range xi {
		fmt.Println("ranging over a slice", i, v)
	}

	//for range loop
	// data structures - map
	// a map is made up of 'key' 'value' pairs; they are NOT ordered
	// maps will return an 'unordered' set of k:v pairs
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	for k, v := range m {
		fmt.Println("ranging over a map", k, v)
	}
}
