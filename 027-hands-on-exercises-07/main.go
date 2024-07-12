package main

import "fmt"

func main() {
	a, b, c := 747, 911, 90210

	fmt.Printf("%d \t %b \t\t %#X\n", a, a, a)
	fmt.Printf("%d \t %b \t\t %#X\n", b, b, b)
	fmt.Printf("%d \t %b \t %#X\n", c, c, c)
}

/*
write a program that uses print verbs to show the following numbers
● 747
● 911
● 90210
as
● decimal
● binary
● hexadecimal
https://go.dev/play/p/7cFkjvbpKuY
*/
