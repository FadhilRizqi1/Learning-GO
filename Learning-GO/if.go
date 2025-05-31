package main

import "fmt"

func main() {

	nilai := 0

	fmt.Scan(&nilai)

	if nilai >= 90 {
		fmt.Println("Nilai A")
	} else if nilai >= 80 {
		fmt.Println("Nilai B")
	} else if nilai >= 70 {
		fmt.Println("Nilai C")
	} else if nilai >= 60 {
		fmt.Println("Nilai D")
	} else {
		fmt.Println("Nilai E")
	}


}