package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Ahmad Fadhil Rizqi", "Fadhil") )
	fmt.Println(strings.Contains("Ahmad Fadhil Rizqi", "MAFALQI"))

	fmt.Println(strings.Count("Ahmad Fadhil Rizqi", "a"))
	fmt.Println(strings.Split("Ahmad Fadhil Rizqi", " "))
	fmt.Println(strings.ToLower("Ahmad Fadhil Rizqi"))
	fmt.Println(strings.ToUpper("Ahmad Fadhil Rizqi"))
	fmt.Println(strings.Trim("   Ahmad Fadhil Rizqi   ", " "))
	fmt.Println(strings.TrimSpace("   Ahmad Fadhil Rizqi   "))
	fmt.Println(strings.ReplaceAll("Ahmad Fadhil Rizqi", "Fadhil", "Fadhil Rizqi"))

}