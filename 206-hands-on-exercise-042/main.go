package main

import (
	"fmt"
)

func main() {
	a := [5]int{2, 4, 3}
	a[3] = 5
	a[4] = 9

	fmt.Println("array 'a'")
	fmt.Println("====================")
	for i, v := range a {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}
	fmt.Println("====================")
	fmt.Printf("array 'a' Type: %T\n", a)
}
