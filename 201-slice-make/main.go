package main

import "fmt"

func main() {
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Printf("Length of slice: %v\n", len(x))
	fmt.Printf("Capacity of slice: %v\n\n", cap(x))
	x[0] = 42
	x[9] = 999

	fmt.Println(x)
	fmt.Printf("Length of slice: %v\n", len(x))
	fmt.Printf("Capacity of slice: %v\n\n", cap(x))

	x = append(x, 3423)

	fmt.Println(x)
	fmt.Printf("Length of slice: %v\n", len(x))
	fmt.Printf("Capacity of slice: %v\n\n", cap(x))

	x = append(x, 3425)

	fmt.Println(x)
	fmt.Printf("Length of slice: %v\n", len(x))
	fmt.Printf("Capacity of slice: %v\n\n", cap(x))
}
