package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	// message := recover()
	// fmt.Println("terjadi panic", message)
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("error")
	}

}

/*
hay guys
*/
func main() {
	runApp(true)
	fmt.Println("al")
}
