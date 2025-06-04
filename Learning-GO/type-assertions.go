package main

import "fmt"

func random() interface{} {
	return true
}

func main() {
	var result interface{} = random()
	// var resultInt = result.(int)
	// fmt.Println("Hasil konversi:", resultInt)

	switch value := result.(type) {
	case int:
		fmt.Println("Hasil konversi:", value, "adalah int")
	case string:
		fmt.Println("Hasil konversi:", value, "adalah string")
	case bool:
		fmt.Println("Hasil konversi:", value, "adalah bool")
	default:
		fmt.Println("Tipe data tidak dikenali")
	}
}	