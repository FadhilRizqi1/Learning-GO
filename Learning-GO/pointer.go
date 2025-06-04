package main

import "fmt"

type Adress struct {
	kota, provinsi, negara string
}

func main() {
	alamat1 := Adress{"Palembang", "Sumatera Selatan", "Indonesia"}
	alamat2 := &alamat1
	alamat3 := &alamat1

    alamat2.kota = "Jakarta"
	
	*alamat2 = Adress{"Bandung", "Jawa Barat", "Indonesia"}

	fmt.Println(alamat2)
	fmt.Println(alamat1)
	fmt.Println(alamat3)
}