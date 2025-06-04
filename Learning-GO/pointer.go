package main

import "fmt"

type Adress struct {
	kota, provinsi, negara string
}

func main() {
	alamat1 := Adress{"Palembang", "Sumatera Selatan", "Indonesia"}
	alamat2 := alamat1

	alamat2.kota = "Jakarta"

	fmt.Println(alamat2)
	fmt.Println(alamat1)
}