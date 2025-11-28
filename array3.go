package main

import "fmt"

func main() {
	var values = [3]int{
		90, 80, 70,
	}

	fmt.Println(values)
	fmt.Println(len(values)) //mendapatkan panjang dari array nya
	values[0] = 100          //mengganti nilai dari index ke 0
	fmt.Println(values)

	//tidak ada operasi data hapus di array untuk Golang
}
