package main

import "fmt"

func main() {
	type NoKTP string

	var ktpAndika NoKTP = "111111"

	var contoh string = "222222"

	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpAndika)
	fmt.Println(contohKtp)
}
