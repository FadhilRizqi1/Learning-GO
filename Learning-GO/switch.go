package main

import "fmt"

func main() {
	
	var (
		nilai int
		akreditasi string
		nama string
		length int
	)

	fmt.Println("Masukkan Nama:")
	fmt.Scan(&nama)

	switch length = len(nama); length  <= 3 {
	case true:
		fmt.Println("Nama harus lebih dari 3 karakter")
		return
	case false:
		fmt.Println("Nama valid")
	}
	
	fmt.Println("Masukkan Nilai:")
	fmt.Scan(&nilai)

	switch nilai {
		case 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100:
			akreditasi = "A"
		case 80, 81, 82, 83, 84, 85, 86, 87, 88, 89:
			akreditasi = "B"
		case 70, 71, 72, 73, 74, 75, 76, 77, 78, 79:
			akreditasi = "C"
		case 60, 61, 62, 63, 64, 65, 66, 67, 68, 69:
			akreditasi = "D"
		default:
			akreditasi = "E"
	}

	fmt.Println("Nilai Anda:", akreditasi)

}