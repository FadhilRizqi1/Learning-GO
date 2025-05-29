package main

import "fmt"
func main() {

	var nama = [...]string{
		"Ahmad",
		"Fadhil",
		"Rizqi",
	}

	var sliceNama = nama[1:3]
	println(sliceNama[0])

	var kendaraan = [...]string{
		"Mobil",
		"Motor",
		"Sepeda",
		"Truk",
		"Bus",
	}

	var sliceKendaraan = kendaraan[:3]

	fmt.Println(sliceKendaraan)
	fmt.Println("Jumlah kendaraan: ", len(sliceKendaraan))
	fmt.Println("Jumlah kapasitas: ", cap(sliceKendaraan))

	var sliceKendaraan2 = make([]string, 2, 5)
	sliceKendaraan2[0] = "Mobil"
	sliceKendaraan2[1] = "Motor"

	fmt.Println(sliceKendaraan2)

	var sliceKendaraan3 = append(sliceKendaraan2, "Sepeda", "Truk", "Bus")

	fmt.Println(sliceKendaraan3)

	fmt.Println(sliceKendaraan2)
}
