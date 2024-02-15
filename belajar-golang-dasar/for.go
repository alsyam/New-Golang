package main

import (
	"fmt"
)

func main() {

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke-", counter)

	}

	slice := []string{"muhammmad", "al", "syam"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("index", i, "=", value)
		// fmt.Println(value)
	}

	person := make(map[string]string)
	person["name"] = "al"
	person["title"] = "web developer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}

}
