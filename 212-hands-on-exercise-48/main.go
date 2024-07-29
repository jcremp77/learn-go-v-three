package main

import (
	"fmt"
)

func main() {
	x1 := make([]string, 0)
	x2 := make([]string, 0)

	x1 = append(x1, "James", "Bond", "Shaken, not stirred")
	x2 = append(x2, "Miss", "Moneypenny", "I'm 008")

	xxs := [][]string{x1, x2}
	fmt.Println("Multi-dimensional Slice ", xxs)

	for i, v := range xxs {
		fmt.Println("=======================================")
		fmt.Println(i, v)
		fmt.Println("=======================================")
		for a, b := range v {
			fmt.Println(a, b)
		}
	}

}
