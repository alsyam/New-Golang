package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1 //copy value

	address2.City = "Bogor"

	fmt.Println(address1) //ikut berubah
	fmt.Println(address2) //berubah menjadi bogor

	// address2 = &Address{"Jakarta", "DKI jakarta", "Indonesia"}
	fmt.Println(address1) //tidak ikut berubah
	fmt.Println(address2) //berubah menjadi jakarta

	*address2 = Address{"Jakarta", "DKI jakarta", "Indonesia"}
	fmt.Println(address1) //ikut berubah jadi address2
	fmt.Println(address2) //berubah menjadi jakarta
}
