package main

import "fmt"

func parameterMassage(i int) {
	fmt.Println("Perulangan ke-", i)
}

func main() {

	for i := 1; i <= 10; i++ {
		parameterMassage(i)
	}
}
