package main

import "fmt"

func main() {

	var (
	nilai int
	akreditasi string
	)

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