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


	slice := []string{"apel", "jeruk", "mangga", "pisang"}

	for j := 0; j < len(slice); j++ {
		fmt.Println("Buah:", slice[j])
	}

	for _, value := range slice {
		fmt.Printf("Buah: %s\n", value)
	}

	for key, value := range map[string]string{
		"nama": "Nama",
		"usia": "19",
		"kota": "Palembang",
	} {
		fmt.Println(key, ":", value)
	}

}
