package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := NewMap("lala")
	if data == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(data)
		fmt.Println(data["name"])
	}

}
