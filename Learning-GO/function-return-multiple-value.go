package main

import "fmt"

func returnMultiple(firstName string, middleName string, lastName string) (string, string, string) {
	return "Halo", firstName + " " + middleName + " " + lastName , "Selamat datang"
} 

func main() {
	
	var firstName, middleName, lastName string

	fmt.Println("Masukkan nama depan anda:")
	fmt.Scan(&firstName)

	fmt.Println("Masukkan nama tengah anda:")
	fmt.Scan(&middleName)

	fmt.Println("Masukkan nama belakang anda:")
	fmt.Scan(&lastName)

	fmt.Println(returnMultiple(firstName, middleName, lastName))

}
