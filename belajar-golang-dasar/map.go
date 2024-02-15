package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "al",
		"address": "subang",
	}

	person["title"] = "programmer"

	fmt.Println(person)
	fmt.Println(len(person))
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	delete(person, "title")
	fmt.Println(person)
	fmt.Println(len(person))

}
