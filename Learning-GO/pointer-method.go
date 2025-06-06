package main

import "fmt"

type Person struct {
	Name string
}

func (orang *Person) Perkenalan() {
	fmt.Println("Halo, nama saya", orang.Name)
}

func main() {
	var orang Person

	fmt.Println("Masukkan nama anda:")
	fmt.Scan(&orang.Name)

	orang.Perkenalan()
}