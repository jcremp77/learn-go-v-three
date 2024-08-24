package main

import (
	"fmt"
)

type Circle struct {
	Radius float64
}

func (c *Circle) Scale(factor float64) {
	c.Radius *= factor
}

func main() {
	c := &Circle{Radius: 5}
	c.Scale(2)
	fmt.Println(c.Radius) // output: 10
}
