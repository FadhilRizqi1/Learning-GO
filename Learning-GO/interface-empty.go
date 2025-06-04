package main

import "fmt"


func any(angka int) interface{} {
	if angka == 1{
	 return	1
	} else if angka == 2 {
		return true
	} else {
		return fmt.Sprintf("Angka lebih dari 2")
   }
}

func main() {
	var contoh interface{} = any(2)
	fmt.Println("Ini adalah contoh interface kosong:", contoh)

	var angka int
	fmt.Println("Masukkan angka")
	fmt.Scanln(&angka)

	result := any(angka)
	fmt.Println(result)
}