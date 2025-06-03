package main

import "fmt"

func recoverFunction() {
	massage := recover()
	if massage != nil {
		fmt.Println("Terjadi kesalahan:", massage)
	}
	fmt.Println("Function telah dieksekusi")
}

func endApp(error bool) {
	if error == true {
		panic ("Terjadi kesalahan, program  dihentikan")
	} 
}

func main() {
	defer recoverFunction()
	fmt.Println("Program sedang berjalan")

	var n int

	fmt.Println("Masukkan nilai n:")
	fmt.Scan(&n)

	if n == 0 {
		endApp(true)
	}

	hasilbagi := 10 / n
	fmt.Println("Hasil bagi:", hasilbagi)

	fmt.Println("Program selesai")
}