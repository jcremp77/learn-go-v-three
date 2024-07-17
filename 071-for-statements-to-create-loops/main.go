package main

import (
	"fmt"
)

func main() {
	y := 5
	// for loop runs until i = 9
	// sum = 0,  i = 0 , i becomes 1 , sum + i, sum = 1
	// sum = 1,  i = 1 , i becomes 2 , sum + i, sum = 3
	// sum = 3,  i = 2 , i becomes 3 , sum + i, sum = 6
	// sum = 6,  i = 3 , i becomes 4 , sum + i, sum = 10
	// sum = 10, i = 4 , i becomes 5 , sum + i, sum = 15
	// sum = 15, i = 5 , i becomes 6 , sum + i, sum = 21
	// sum = 21, i = 6 , i becomes 7 , sum + i, sum = 28
	// sum = 28, i = 7 , i becomes 8 , sum + i, sum = 36
	// sum = 36, i = 8 , i becomes 9 , sum + i, sum = 45
	// end of for loop...
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}

	fmt.Printf("\n")
	fmt.Printf("First FOR loop\n")

	for i := 0; i < 5; i++ {
		fmt.Printf("%v of 5\n", i)
	}

	fmt.Printf("\n")
	fmt.Printf("Second FOR loop\n")

	for y < 10 {
		fmt.Printf("y is %v\n", y)
		y++
	}

	fmt.Printf("\n")
	fmt.Printf("Third FOR loop\n")

	// break takes you out of a loop
	for {
		fmt.Printf("y is %v\n", y)
		if y > 20 {
			break
		}
		y++
	}

	// continue takes it to the next iteration
	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("counting even numbers: ", i)
	}
}
