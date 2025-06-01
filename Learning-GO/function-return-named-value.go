package main

import "fmt"

func returnNamedValue (firstName string, middleName string, lastName string) (greeting string, fullName string, message string) {
	greeting = "Halo"
	fullName = firstName + " " + middleName + " " + lastName
	message = "Selamat datang"

	return
}

func main() {

	var firstName, middleName, lastName string

	fmt.Println("Masukkan nama depan anda:")
	fmt.Scan(&firstName)

	fmt.Println("Masukkan nama tengah anda:")
	fmt.Scan(&middleName)

	fmt.Println("Masukkan nama belakang anda:")
	fmt.Scan(&lastName)

	greeting, fullName, message := returnNamedValue(firstName, middleName, lastName)
	fmt.Println(greeting, fullName, message)
}
