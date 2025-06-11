package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args
	fmt.Println("Argument : ")
	fmt.Println(arg)

	name, err := os.Hostname()
	if err != nil {
		fmt.Println("Error : ", err.Error())
	} else {
		fmt.Println("Hostname : ", name)
	}
}