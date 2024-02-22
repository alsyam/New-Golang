package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// method
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello ", name, "my name is", customer.Name)
}

func main() {
	var al Customer
	al.Name = "syam"
	al.Address = "indo"
	al.Age = 24

	fmt.Println(al)
	fmt.Println(al.Name)
	fmt.Println(al.Address)
	fmt.Println(al.Age)

	joko := Customer{
		Name:    "Joko",
		Address: "indo",
		Age:     76,
	}

	fmt.Println(joko)

	tono := Customer{"Tono", "indo", 99}
	fmt.Println(tono)

	// method

	tono.sayHello("al")

}
