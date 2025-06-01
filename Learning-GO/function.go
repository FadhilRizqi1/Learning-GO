package main

import "fmt"

var i int

func loopMassage() {

	fmt.Println("Perulangan ke-", i)
}

func main() {

	for i = 1; i <= 10; i++ {
		loopMassage()
	}

}