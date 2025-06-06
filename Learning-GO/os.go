package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args
	fmt.Println("Argument : ")
	fmt.Println(arg)

	name, error := os.Hostname()
	if error != nil {
		fmt.Println("Error : ", error)
	} else {
		fmt.Println("Hostname : ", name)
	}
}