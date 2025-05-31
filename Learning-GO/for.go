package main

import "fmt"

func main() {

	var (
		perulangan int = 0
		i int
	)

	for perulangan <= 10 {
			fmt.Println ("Perulangan ke-", perulangan)
			perulangan++
	}

	for i = 0; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Println("Bilangan genap:", i)
		} else {
			fmt.Println("Bilangan ganjil:", i)
		}
	}

}