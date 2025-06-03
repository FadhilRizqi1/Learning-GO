package main

import "fmt"

type Identitas struct {
	namaDepan, namaBelakang string
	umur int
}

func (orang Identitas) tampilkanNama() {
	fmt.Println("Nama Depan:", orang.namaDepan)
	fmt.Println("Nama Belakang:", orang.namaBelakang)
	fmt.Println("Umur:", orang.umur)
}

func (orang Identitas) perkenalan(nama string) {
	fmt.Println("Halo, nama saya", nama, "umur saya", orang.umur, "tahun.")
}

func main() {
	var orang Identitas

	orang.namaDepan = "Fadhil"
	orang.namaBelakang = "Rizqi"
	orang.umur = 19

	orang.tampilkanNama()

	orang.perkenalan("Fadhil Rizqi")
}
