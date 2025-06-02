package main

import "fmt"

func printWithFiler(name string, filter func(string) string) {
	name = filter(name)
	fmt.Println("Halo " + filter(name))
}

func spamFilter(name string) string {
	if name == "Gacor" {
		return "Rungkad"
	} else {
		return name
	}

}

func main() {
	var name string

	fmt.Println("Masukkan nama anda:")
	fmt.Scan(&name)

	printWithFiler(name, spamFilter)	

}
