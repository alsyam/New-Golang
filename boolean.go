package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absensi = 80

	var lulusUjian = nilaiAkhir > 80
	var lulusAbsensi = absensi > 80
	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var lulus = lulusAbsensi && lulusUjian

	fmt.Println(lulus)
}
