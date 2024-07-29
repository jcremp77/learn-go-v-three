package main

import (
	"fmt"
)

func main() {
	x := make([]string, 1)
	fmt.Println("=====================================================================")
	fmt.Printf("Slice 'x'\tType: %T\t", x)
	x[0] = "<intentionally blank>"
	x = append(x, "Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado",
		"Connecticut", "Delaware", "Florida", "Georgia", "Hawaii", "Idaho",
		"Illinois", "Indiana", "Iowa", "Kansas", "Kentucky", "Louisiana",
		"Maine", "Maryland", "Massachusetts", "Michigan", "Minnesota", "Mississippi",
		"Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey",
		"New Mexico", "New York", "North Carolina", "North Dakota", "Ohio",
		"Oklahoma", "Oregon", "Pennsylvania", "Rhode Island", "South Carolina",
		"South Dakota", "Tennessee", "Texas", "Utah", "Vermont", "Virginia",
		"Washington", "West Virginia", "Wisconsin", "Wyoming")

	fmt.Printf("\tLength: %v", len(x))
	fmt.Printf("\tCapacity: %v\n", cap(x))
	fmt.Println("=====================================================================")
	for i := 0; i < len(x); i++ {
		fmt.Printf("Index: %v, Value: %v\n", i, x[i])
	}
}
