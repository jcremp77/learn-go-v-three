package main

// create a program that uses iota to calculate the size of each measurement of bytes

import "fmt"

type ByteSize int

const (
	_ ByteSize = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
	EB
)

func main() {
	fmt.Printf("%d \t\t\t %b\n", KB, KB)
	fmt.Printf("%d \t\t %b\n", MB, MB)
	fmt.Printf("%d \t\t %b\n", GB, GB)
	fmt.Printf("%d \t\t %b\n", TB, TB)
	fmt.Printf("%d \t %b\n", PB, PB)
	fmt.Printf("%d \t %b\n", EB, EB)
}
