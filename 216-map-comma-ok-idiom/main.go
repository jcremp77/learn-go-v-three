package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	m1["fred"] = 46
	m1["barney"] = 44
	m1["Wilma"] = 39

	if v, ok := m1["bam bam"]; !ok {
		fmt.Println("key doesn't exist")
	} else {
		fmt.Println("the value prints", v)
	}
	if v, ok := m1["fred"]; !ok {
		fmt.Println("key doesn't exist")
	} else {
		fmt.Println("the value prints", v)
	}
}
