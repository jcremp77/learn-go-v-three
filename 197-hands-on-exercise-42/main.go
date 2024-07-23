package main

import (
	"fmt"
)

func main() {
	// create a slice
	flavors := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)"}

	// for range loop over slice of strings
	// printed on seperate lines
	for _, v := range flavors {
		fmt.Printf("%v\n", v)
	}

	// print by index position
	fmt.Println()
	fmt.Println(flavors[1])
	fmt.Println()

	// statement-statement for loop to print based on index reference
	for i := 0; i < len(flavors); i++ {
		fmt.Println(flavors[i])
	}
}
