package main

import "fmt"

func getFullName() (string, string) {
	return "al", "syam"
}

func main() {
	firstName, _ := getFullName()
	fmt.Println(firstName)
	// fmt.Println(lastName)
}
