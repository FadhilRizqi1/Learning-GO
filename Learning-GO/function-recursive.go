package main

import "fmt"

func factorial(n int) int {
	result := 1
	for i := n; i > 0; i-- {
		result *= i
	}
	return result
}

func main() {
	var n int
	fmt.Println("Masukkan angka untuk menghitung faktorial:")
	fmt.Scan(&n)
	loop := factorial(n)

	fmt.Println("Hasil faktorial dari", n, "adalah:", loop)
}