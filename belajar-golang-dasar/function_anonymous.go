package main

import (
	"fmt"
)

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You Are Blocked", name)
	} else {
		fmt.Println("Wellcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "Anjing"
	}

	registerUser("al", blacklist)

	// anonymous
	registerUser("Anjing", func(name string) bool {
		return name == "Anjing"
	})
}
