package main

import "fmt"

func main() {
	orang := map[string]string{
		"nama":    "Fadhil",
		"alamat":  "Palembang",
		"pekerjaan": "Programmer/IT Support",
		"umur":    "19",
	}

	fmt.Println(orang)
	
	fmt.Println("Nama:", orang["nama"])
	fmt.Println("Alamat:", orang["alamat"])
	fmt.Println("Pekerjaan:", orang["pekerjaan"])
	fmt.Println("Umur:", orang["umur"])

	buku := map[string]map[string]string{
		"buku1": {
			"judul": "Belajar Go",
			"pengarang": "Fadhil",
		},
		"buku2": {
			"judul": "Belajar Laravel",
			"pengarang": "Fadhil",
		},
	}
}