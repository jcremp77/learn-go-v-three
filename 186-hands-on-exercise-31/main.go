package main

import (
	"fmt"
)

func init() {
	fmt.Printf("...initializing\n\n")
}

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for k, v := range m {
		fmt.Printf("key=%v\tvalue=%v\n", k, v)
	}
}
