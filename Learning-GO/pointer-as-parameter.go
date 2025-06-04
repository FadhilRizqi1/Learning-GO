package main

import "fmt"

type Adress struct {
	kota, provinsi, negara string
}

func changeAdress(alamat *Adress) {
	alamat.negara = "Indonesia"
}

func main() {
	var alamat1 Adress = Adress{"Palembang", "Sumatera Selatan", " "}
	fmt.Println(alamat1)

	var alamat2 *Adress = &alamat1
	changeAdress(alamat2)
	fmt.Println(alamat1)
	
}