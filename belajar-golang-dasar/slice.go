package main

import "fmt"

func main() {
	var months = [...]string{
		"January",
		"February",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// months[5] = "diubuah"
	// fmt.Println(slice1)

	var slice2 = months[2:7]
	fmt.Println(slice2)

	var slice3 = append(slice2, "al")
	fmt.Println(slice3)
	slice3[1] = "dede"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)
	fmt.Println(cap(months))

	newSlice := make([]string, 2, 5)

	newSlice[0] = "aku"
	newSlice[1] = "kamu"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
