package main

import (
	"fmt"
)

func main() {
	m := make(map[string][]string)
	m["apple"] = []string{`red`, `round`}
	m["bannana"] = []string{`yellow`, `curve`}
	m["grape"] = []string{`purple`, `round`}

	fmt.Printf("%#v\n", m)     // prints the map
	fmt.Println(m["apple"][1]) // prints round
	fmt.Println()

	for k, v := range m {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
		fmt.Println()
	}
}
