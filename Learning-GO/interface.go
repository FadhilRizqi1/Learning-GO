package main

import "fmt"

type nama interface {
	tampilkanNama() string
}

func tampilkanNama(orang nama) {
	fmt.Println(orang.tampilkanNama())
}

type Orang struct {
	namaDepan, namaBelakang string
	umur int
}

func (orang Orang) tampilkanNama() string {
	return "Nama: " + orang.namaDepan + " " + orang.namaBelakang + ", Umur: " + fmt.Sprint(orang.umur)
}

type Teman struct {
	Nama string
}

func (orang Teman) tampilkanNama() string {
	return "Nama Teman: " + orang.Nama
}

func main() {
	var Fadhil Orang

	Fadhil.namaDepan = "Fadhil"
	Fadhil.namaBelakang = "Rizqi"
	Fadhil.umur = 19

	tampilkanNama(Fadhil)

	Thoriq := Teman{
		Nama: "Thoriq",
	}
	tampilkanNama(Thoriq)
}
