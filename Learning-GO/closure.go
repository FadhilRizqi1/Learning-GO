package main

import "fmt"

func main() {
nama := "Fadhil"
n := 0

increment := func() {
	nama = "Fadhil"
	nama := "Rizqi"

	n ++
	
	fmt.Println("Nama:", nama)

}
	increment()
	increment()
	
	fmt.Println("Nilai n:", n)
	fmt.Println("Nama:", nama)
}