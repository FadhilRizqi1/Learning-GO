package main

import "fmt"

func deferFunction() {
	fmt.Println("Function telah dieksekusi")
}

func main() {
	defer deferFunction()
	fmt.Println("Program sedang berjalan")

	var n int

	fmt.Println("Masukkan nilai n:")
	fmt.Scan(&n)

	hasilbagi := 10/n
	fmt.Println("Hasil bagi:", hasilbagi)

	fmt.Println("Program selesai")
}