package helper

import "fmt"

func SayName(name string) {
	fmt.Println("Hello", name)
}

// var unexported = "Halo, saya dari package helper" // unexported variable

var Exported = "Halo, saya dari package helper" // exported variable
