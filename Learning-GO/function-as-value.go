package main

import "fmt"

func hitungSemua(numbers ...int) int {
	var total int

	for _, value := range numbers {
		total += value
	}

	return total
}

func main() {
	
	tambah := hitungSemua

	totalDari := tambah(10, 20, 30, 40, 50)
	fmt.Println(totalDari) 
}