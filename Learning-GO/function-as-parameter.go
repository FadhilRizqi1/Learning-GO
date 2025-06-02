package main

import "fmt"

type Filter = func(name string) string 

func printWithFilter(name string, filter Filter) {
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

	printWithFilter(name, spamFilter)	

}
