package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// address2 := address1 //copy value

	// address2.City = "Bogor"

	// fmt.Println(address1)
	// fmt.Println(address2)

	// poniter
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1 //copy value

	address2.City = "Bogor"

	fmt.Println(address1)
	fmt.Println(address2)
}
