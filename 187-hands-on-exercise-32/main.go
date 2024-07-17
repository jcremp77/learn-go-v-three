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

	age := m["James"]
	fmt.Println(age)

	age1 := m["Q"]
	fmt.Println("The age of Q is:", age1)

	if k, ok := m["Q"]; !ok {
		fmt.Printf("There is no Q here...\n")
		fmt.Println("zero value of an int", k)
	}
}
