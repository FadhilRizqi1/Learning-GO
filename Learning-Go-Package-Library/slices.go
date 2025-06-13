package main

import (
	"fmt"
	"slices"
)

func main() {
	name := []string{"Ahmad", "Fadhil", "Rizqi", "MAFALQI", "Falqi"}
	value := []int{100, 90, 80, 70, 60}


	fmt.Println(slices.Max(value))
	fmt.Println(slices.Min(name))
	fmt.Println(slices.Contains(name, "Fadhil"))
	fmt.Println(slices.Index(name, "Alfin"))
}