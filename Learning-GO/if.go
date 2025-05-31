package main

import "fmt"

func main() {

	var (
	nilai int
	akreditasi string
	nama string
	length int
	)

	fmt.Print("Masukkan nama: ")
	fmt.Scan(&nama)

	if length = len(nama); length <= 3 {
		fmt.Println("Nama harus lebih dari 3 karakter")
		return
	}
	

	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nilai)

	if nilai >= 90 {
		akreditasi = "A"
	} else if nilai >= 80 {
		akreditasi = "B"
	} else if nilai >= 70 {
		akreditasi = "C"
	} else if nilai >= 60 {
		akreditasi = "D"
	} else {
		akreditasi = "E"
	}

	fmt.Println("Nilai Anda:", akreditasi)



}