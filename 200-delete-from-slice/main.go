package main

import (
	"fmt"
)

func main() {
	flavors := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)"}
	fmt.Printf("Elements of Slice \"flavors\"\n")
	fmt.Printf("------------------------------\n")
	for _, v := range flavors {
		fmt.Printf("%q\n", v)
	}

	fmt.Printf("\n")

	newflavors := flavors[:2]
	fmt.Printf("Elements of Slice \"flavors\"\n")
	fmt.Printf("------------------------------\n")
	for _, v := range newflavors {
		fmt.Printf("%q\n", v)
	}
}
