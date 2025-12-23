package main

import "fmt"

func main() {
	/*
		//deklraasi variabel dengan jenisnya (sesuai datanya, cth ini String)
		name := "Andika Rafon"
		fmt.Println(name)

		//di deklarasi yang kedua, variabel yang name tidak boleh pakai :=
		name = "Andika Rafon Sinuhaji"
		fmt.Println(name)
	*/

	var (
		firstName  = "Andika"
		middleNama = "Rafon"
		lastName   = "Sinuhaji"
		//catatan : di Golang, jika ada variabel yang tdk digunakan maka error
	)

	fmt.Println(firstName)
	fmt.Println(middleNama)
	fmt.Println(lastName)

}
