package main

import "fmt"

func main() {

	nilai := 0

	fmt.Scan(&nilai)

	if nilai >= 90 {
		fmt.Println("A")
	} else if nilai >= 80 {
		fmt.Println("B")
	} else if nilai >= 70 {
		fmt.Println("C")
	} else if nilai >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}

	fmt.Println("Nilai Anda:", nilai)

}