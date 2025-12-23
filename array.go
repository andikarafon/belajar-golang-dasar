package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Andika"
	names[1] = "Rafon"
	names[2] = "Sinuhaji"
	// names[3] = "Titis" // Jika dibuat sampai ini, maka akan error saat di runing

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
}
