package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 10
	var c = a * b
	fmt.Println(c)

	// augmneteh assignment
	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)
	i += 5 // i = i + 5
	fmt.Println(i)

	// unary operator

	i++ // i = i + 1
	fmt.Println(i)
	i++ // i = i + 1
	fmt.Println(i)

}
