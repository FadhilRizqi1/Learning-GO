package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagi tidak boleh nol")
	}
	hasil := nilai / pembagi
	return hasil, nil
}

func main() {
	var nilai, pembagi int
	
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nilai)
	fmt.Print("Masukkan pembagi: ")
	fmt.Scan(&pembagi)

	pembagian, err := pembagian(nilai, pembagi)

	if  err != nil {
		fmt.Println("Terjadi kesalahan:", err)
	} else {
		fmt.Printf("Hasil pembagian:", pembagian)
	}

}