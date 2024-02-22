package main

import (
	"fmt"
)

type Filter func(string) string

func sayHelloNameWithFilter(name string, filter Filter) {
	fmt.Println("Hallo " + filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloNameWithFilter("al", spamFilter)

	filter := spamFilter
	sayHelloNameWithFilter("Anjing", filter)
}
