package main

import (
	"fmt"
)

var (
	street string = "9300 Montana"
)

const (
	city string = "El Paso"
)

func main() {
	apt := 9001

	fmt.Printf("The address is: %v %v, %v", street, apt, city)
}
