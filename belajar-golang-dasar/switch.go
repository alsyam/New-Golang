package main

import "fmt"

func main() {
	name := "syam"

	switch name {
	case "al":
		fmt.Println("halo al")
	case "syam":
		fmt.Println("halo syam")
	default:
		fmt.Println("hai guys")
	}

	switch length := len(name); length > 3 {
	case true:
		fmt.Println("nama panjang")
	case false:
		fmt.Println("nama pendek")
	}

	length := len(name)

	switch {
	case length > 10:
		fmt.Println("panjang")
	case length < 3:
		fmt.Println("pendek")
	default:
		fmt.Println("benar")
	}

}
