package main

import "fmt"

func main() {

	// membuat nama tipe data
	type NoKTP string

	var noKtpAl NoKTP = "12121212121212100"

	var contoh string = "2112121212121"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(noKtpAl)
	fmt.Println(contohKtp)
}
