package main

import "fmt"

func logging() {
	fmt.Println("selesai")
}

func runApp() {
	defer logging()
	fmt.Println("run app")

}

func main() {
	runApp()
}
