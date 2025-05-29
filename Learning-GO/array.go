package main

import "fmt"

func main() {

	var nama [5]string

	nama[0] = "Fadhil"
	nama[1] = "Rizqi"
	nama[2] = "Ahmad"
	nama[3] = "Falzi"
	nama[4] = "MAFALQI"

	fmt.Println(nama[4])

	var kendaraan = [5]string{
		"Mobil",
		"Motor",
		"Sepeda",
		"Truk",
		"Bus",
	}
	
fmt.Println(kendaraan)

fmt.Println("Jumlah kendaraan: ", len(kendaraan))
fmt.Println("Jumlah huruf: ", len(nama[4]))



}