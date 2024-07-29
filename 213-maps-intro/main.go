package main

import (
	"fmt"
)

func main() {
	m1 := make(map[string]string)
	m1["building"] = "90880"
	m1["street"] = "Ironwood Dr"
	m1["city"] = "Corona"
	m1["state"] = "CA"
	m1["zipcode"] = "99501"
	fmt.Println("============================================================")
	fmt.Println("print only the street name: ", m1["street"])
	fmt.Println("============================================================")
	fmt.Printf("complete address is: %v %v, %v %v, %v\n", m1["building"], m1["street"], m1["city"], m1["state"], m1["zipcode"])
	fmt.Println("============================================================")
	fmt.Println("the length of the map is: ", len(m1))
	fmt.Println("============================================================")

	/*fmt.Println(m1)
	m1["route"] = 66
	i := m1["root"]
	fmt.Println(m1)
	fmt.Println(i)
	m1["root"] = 100
	i = m1["root"]
	fmt.Println(i)*/
}
