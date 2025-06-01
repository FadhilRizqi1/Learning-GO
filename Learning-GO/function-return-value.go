package main

import "fmt"

	func panggilNama(name string) string {
		if name == "" {
			return "Nama tidak boleh kosong"
		} else {
			return "Nama saya adalah " + name
		}
	}

	func main() {
		var nama string

		fmt.Scan (&nama)
		fmt.Println(panggilNama(nama))
	}