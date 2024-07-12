package main

import "fmt"

var (
	big   int8  = 127
	small uint8 = 255
)

func main() {
	fmt.Printf("%v is the largest value of type: %T\n", big, big)
	fmt.Printf("%v is the largest value of type: %T\n", small, small)
}

/*
write a program that declares two variables
● one variable to store a VALUE of TYPE int8
○ assign to it the largest number possible, then print it
● one variable to store a VALUE of TYPE uint8
○ assign to it the largest number possible, then print it
https://go.dev/play/p/-q6V8dbGkgz
*/
