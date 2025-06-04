package main

import "fmt"

func random() interface{} {
	return 1
}

func main() {
	var result interface{} = random()
	var resultInt = result.(int)
	fmt.Println("Hasil konversi:", resultInt)
}