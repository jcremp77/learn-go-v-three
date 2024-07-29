package main

import (
	"fmt"
)

func main() {
	m1 := make(map[string]string)
	// Following the letter(s) that identify the service type, the three-digit numeric
	// portion identifies the tire's "Section Width" (cross section) in millimeters.
	m1["section width"] = "245"
	// Following the three digits identifying the tire's Section Width in millimeters
	// is a two-digit number that identifies the tire's profile or aspect ratio. This
	// is the “height” of the sidewall, expressed as a percentage of the section width.
	m1["sidewall aspect ratio"] = "75"
	// A letter (R in this case) that identifies the tire's internal construction
	// follows the two digits used to identify the aspect ratio.
	// 'R' identifies that the tire has a Radial construction.
	// 'D' identifies that the tire has a bias-ply (Diagonal) construction.
	// 'RF' identifies that the tire has a self-supporting Run Flat construction.
	m1["internal construction"] = "R"
	// The 16 indicates the tire and wheel diameter designed to be matched together.
	m1["tire and wheel diameter"] = "16"
	// The 91S represents the tire's Service Description. A Service Description
	// identifies the tire's Load Index (91) and Speed Rating (S).
	m1["service description"] = "91s"

	fmt.Println("-------------------------------------------------")
	fmt.Println("tire size details")
	fmt.Println("-------------------------------------------------")
	for k, v := range m1 {
		fmt.Printf("key: %v: %v\n\n", k, v)
	}
	fmt.Printf("tire size: %v%v%v%v%v\n", m1["section width"], m1["sidewall aspect ration"], m1["internal construction"], m1["tire and wheel diameter"], m1["service description"])

	xi := []int{42, 43, 44}
	fmt.Println()
	fmt.Println("-------------------------------------------------")
	fmt.Println("slice 'xi'")
	fmt.Println("-------------------------------------------------")
	fmt.Println("index and values of slice 'xi'")
	for i, v := range xi {
		fmt.Println(i, v)
	}
	fmt.Println("-------------------------------------------------")
	fmt.Println("only values of slice 'xi' #no index position")
	for _, v := range xi {
		fmt.Println(v)
	}
}
