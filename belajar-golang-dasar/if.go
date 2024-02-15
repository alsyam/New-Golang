package main

import "fmt"

func main() {
	name := "salman"

	if name == "al" {
		fmt.Println("Hello al")
	} else if name == "salman" {
		fmt.Println("Hello salman")
	} else {
		fmt.Println("Hello syam")
	}

	if length := len(name); length > 5 {
		fmt.Println("panjang")
	} else {
		fmt.Println("cuy")
	}
}
