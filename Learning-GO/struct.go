package main

import "fmt"

func main() {

	type Identitas struct {
		namaDepan, namaBelakang string
		umur int
	}

	var orang Identitas

	orang.namaDepan = "Fadhil"
	orang.namaBelakang = "Rizqi"
	orang.umur = 19

	fmt.Println(orang)

	fmt.Println("Nama Depan:", orang.namaDepan)
	fmt.Println("Nama Belakang:", orang.namaBelakang)
	fmt.Println("Umur:", orang.umur)

	Fadhil := Identitas{
		namaDepan: "Fadhil",
		namaBelakang: "Rizqi",
		umur: 19,
	}
	fmt.Println(Fadhil)

	Rizqi := Identitas{"Fadhil", "Rizqi", 19}
	fmt.Println(Rizqi)

}