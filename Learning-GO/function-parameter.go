package main

import "fmt"

func parameterMassage(i int) {
	fmt.Println("Perulangan ke-", i)
}

func main() {

	for i := 1; i <= 10; i++ {
		parameterMassage(i)
	}

	for j := 1; j <= 5; j++ {
		parameterMassage(j * 2) 
	}
}
