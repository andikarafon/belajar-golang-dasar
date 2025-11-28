//tipe data slice adalah potongan dari array

package main

import "fmt"

func main() {
	names := [...]string{"Andika", "Rafon", "Sinuhaji", "Titis", "Wulandari", "Ulun"}
	slice := names[4:6]

	fmt.Println(slice[0])
	fmt.Println(slice[1])

	slice2 := names[:3] // 0 1 2
	fmt.Println(slice2[0])
	fmt.Println(slice2[1])
	fmt.Println(slice2[2])
	// fmt.Println(slice2[3])

}
