package main

import "fmt"

func main() {
	
	var (
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

}