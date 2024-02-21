package main

import "fmt"

func getHello(name string) string {

	if name == "" {
		return "hello cuy"
	} else {
		return "hello " + name
	}
}

func main() {
	result := getHello("al")
	fmt.Println(result)

	fmt.Println(getHello(""))
}
