package main

import (
	"fmt"
)

func main() {
	flavors := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)"}

	fmt.Printf("Slice \"flavors\"\n")
	fmt.Printf("------------------------------\n")
	for _, v := range flavors {

		fmt.Printf("%q\n", v)
	}

	fmt.Printf("\n")

	flavors = flavors[2:]

	fmt.Printf("New Slice \"flavors\"\n")
	fmt.Printf("------------------------------\n")
	for _, v := range flavors {
		fmt.Printf("%q\n", v)
	}
}
