package main

/*
We learned about bitwise operations in the last video. Now let's learn about "iota" and use it in a
program
● to learn about iota
○ golang spec
○ effective go
● modify this program to use iota
○ https://go.dev/play/p/jZUmqIhqaIC
○ (note how iota only needs to be mentioned at the top of a block of code)
*/

import (
	"fmt"
)

const (
	_ = 1 << iota
	a //1<<1 == 10
	b //1<<2 == 100
	c //1<<3 == 1000
	d //1<<4 == 10000
	e //1<<5 == 100000
	f //1<<6 == 1000000
)

func main() {
	fmt.Printf("%d \t %b\n", 1, 1)
	fmt.Printf("%d \t %b\n", a, a)
	fmt.Printf("%d \t %b\n", b, b)
	fmt.Printf("%d \t %b\n", c, c)
	fmt.Printf("%d \t %b\n", d, d)
	fmt.Printf("%d \t %b\n", e, e)
	fmt.Printf("%d \t %b\n", f, f)
}
