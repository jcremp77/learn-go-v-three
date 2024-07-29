package main

import (
	"fmt"
)

func main() {
	m := make(map[string][]string)
	m["apple"] = []string{`red`, `round`}
	m["bannana"] = []string{`yellow`, `curve`}
	m["grape"] = []string{`purple`, `round`}
	m["lemon"] = []string{`yellow`, `round`}

	for k, v := range m {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
		fmt.Println()
	}

	fmt.Println("--------record deleted--------")
	delete(m, "lemon")
	fmt.Println("--------record deleted--------")
	fmt.Println()

	for k, v := range m {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
		fmt.Println()
	}
}
