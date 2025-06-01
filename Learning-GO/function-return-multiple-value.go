package main

import "fmt"

func returnMultiple(firstName string, lastName string) (string, string) {
	return "Halo", firstName + " " + lastName
}

func main() {

	var firstName, lastName string

	fmt.Println("Masukkan nama depan anda:")
	fmt.Scan(&firstName)

	fmt.Println("Masukkan nama belakang anda:")
	fmt.Scan(&lastName)

	fmt.Println(returnMultiple(firstName, lastName))

}
