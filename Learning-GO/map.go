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

	fmt.Println("Buku 1:", buku["buku1"]["judul"], "oleh", buku["buku1"]["pengarang"])
	fmt.Println("Buku 2:", buku["buku2"]["judul"], "oleh", buku["buku2"]["pengarang"])

}